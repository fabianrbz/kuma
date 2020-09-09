package manager_test

import (
	"context"
	"time"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	core_manager "github.com/kumahq/kuma/pkg/core/resources/manager"
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
	core_store "github.com/kumahq/kuma/pkg/core/resources/store"
	core_metrics "github.com/kumahq/kuma/pkg/metrics"
	"github.com/kumahq/kuma/pkg/plugins/resources/memory"
	test_metrics "github.com/kumahq/kuma/pkg/test/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type countingResourcesManager struct {
	store       core_store.ResourceStore
	getQueries  int
	listQueries int
}

func (c *countingResourcesManager) Get(ctx context.Context, res core_model.Resource, fn ...core_store.GetOptionsFunc) error {
	c.getQueries++
	return c.store.Get(ctx, res, fn...)
}

func (c *countingResourcesManager) List(ctx context.Context, list core_model.ResourceList, fn ...core_store.ListOptionsFunc) error {
	c.listQueries++
	return c.store.List(ctx, list, fn...)
}

var _ core_manager.ReadOnlyResourceManager = &countingResourcesManager{}

var _ = Describe("Cached Resource Manager", func() {

	var store core_store.ResourceStore
	var cachedManager core_manager.ReadOnlyResourceManager
	var countingManager *countingResourcesManager
	var res *core_mesh.DataplaneResource
	var metrics core_metrics.Metrics
	expiration := 100 * time.Millisecond

	BeforeEach(func() {
		// given
		store = memory.NewStore()
		countingManager = &countingResourcesManager{
			store: store,
		}
		m, err := core_metrics.NewMetrics()
		metrics = m
		Expect(err).ToNot(HaveOccurred())
		cachedManager, err = core_manager.NewCachedManager(countingManager, expiration, metrics)
		Expect(err).ToNot(HaveOccurred())

		// and created resources
		res = &core_mesh.DataplaneResource{
			Spec: mesh_proto.Dataplane{
				Networking: &mesh_proto.Dataplane_Networking{
					Address: "127.0.0.1",
					Inbound: []*mesh_proto.Dataplane_Networking_Inbound{
						{
							Port:        80,
							ServicePort: 8080,
						},
					},
				},
			},
		}
		err = store.Create(context.Background(), res, core_store.CreateByKey("dp-1", "default"))
		Expect(err).ToNot(HaveOccurred())
	})

	It("should cache Get() queries", func() {
		// when fetched resources multiple times
		fetch := func() core_mesh.DataplaneResource {
			fetched := core_mesh.DataplaneResource{}
			err := cachedManager.Get(context.Background(), &fetched, core_store.GetByKey("dp-1", "default"))
			Expect(err).ToNot(HaveOccurred())
			return fetched
		}

		for i := 0; i < 100; i++ {
			fetch()
		}

		// then real manager should be called only once
		Expect(fetch().Spec).To(Equal(res.Spec))
		Expect(countingManager.getQueries).To(Equal(1))

		// when
		time.Sleep(expiration)

		// then
		Expect(fetch().Spec).To(Equal(res.Spec))
		Expect(countingManager.getQueries).To(Equal(2))

		// and metrics are published
		Expect(test_metrics.FindMetric(metrics, "store_cache", "operation", "get", "result", "hit").Counter.GetValue()).To(Equal(100.0))
		Expect(test_metrics.FindMetric(metrics, "store_cache", "operation", "get", "result", "miss").Counter.GetValue()).To(Equal(2.0))
	})

	It("should cache List() queries", func() {
		// when fetched resources multiple times
		fetch := func() core_mesh.DataplaneResourceList {
			fetched := core_mesh.DataplaneResourceList{}
			err := cachedManager.List(context.Background(), &fetched, core_store.ListByMesh("default"))
			Expect(err).ToNot(HaveOccurred())
			return fetched
		}

		for i := 0; i < 100; i++ {
			fetch()
		}

		// then real manager should be called only once
		list := fetch()
		Expect(list.Items).To(HaveLen(1))
		Expect(list.Items[0].GetSpec()).To(Equal(&res.Spec))
		Expect(countingManager.listQueries).To(Equal(1))

		// when
		time.Sleep(expiration)

		// then
		list = fetch()
		Expect(list.Items).To(HaveLen(1))
		Expect(list.Items[0].GetSpec()).To(Equal(&res.Spec))
		Expect(countingManager.listQueries).To(Equal(2))

		// and metrics are published
		Expect(test_metrics.FindMetric(metrics, "store_cache", "operation", "list", "result", "miss").Counter.GetValue()).To(Equal(2.0))
		Expect(test_metrics.FindMetric(metrics, "store_cache", "operation", "list", "result", "hit").Counter.GetValue()).To(Equal(100.0))
	})
})
