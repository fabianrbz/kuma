package webhooks

import (
	"context"
	managers_mesh "github.com/Kong/kuma/pkg/core/managers/apis/mesh"
	"github.com/Kong/kuma/pkg/core/resources/apis/mesh"
	"github.com/Kong/kuma/pkg/core/validators"
	"github.com/Kong/kuma/pkg/plugins/resources/k8s"
	"github.com/Kong/kuma/pkg/plugins/resources/k8s/native/api/v1alpha1"
	mesh_k8s "github.com/Kong/kuma/pkg/plugins/resources/k8s/native/api/v1alpha1"
	"k8s.io/api/admission/v1beta1"
	"net/http"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func NewMeshValidatorWebhook(validator managers_mesh.MeshValidator, converter k8s.Converter) AdmissionValidator {
	return &MeshValidator{
		validator: validator,
		converter: converter,
	}
}

type MeshValidator struct {
	validator managers_mesh.MeshValidator
	converter k8s.Converter
	decoder   *admission.Decoder
}

func (h *MeshValidator) InjectDecoder(d *admission.Decoder) error {
	h.decoder = d
	return nil
}

func (h *MeshValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	switch req.Operation {
	case v1beta1.Create:
		return h.ValidateCreate(ctx, req)
	case v1beta1.Update:
		return h.ValidateUpdate(ctx, req)
	}
	return admission.Allowed("")
}

func (h *MeshValidator) ValidateCreate(ctx context.Context, req admission.Request) admission.Response {
	coreRes := &mesh.MeshResource{}
	k8sRes := &v1alpha1.Mesh{}
	if err := h.decoder.Decode(req, k8sRes); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	if err := h.converter.ToCoreResource(k8sRes, coreRes); err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}
	if err := h.validator.ValidateCreate(ctx, req.Name, coreRes); err != nil {
		if kumaErr, ok := err.(*validators.ValidationError); ok {
			return convertValidationError(kumaErr, k8sRes)
		}
		return admission.Denied(err.Error())
	}
	return admission.Allowed("")
}

func (h *MeshValidator) ValidateUpdate(ctx context.Context, req admission.Request) admission.Response {
	coreRes := &mesh.MeshResource{}
	k8sRes := &v1alpha1.Mesh{}
	if err := h.decoder.DecodeRaw(req.Object, k8sRes); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	if err := h.converter.ToCoreResource(k8sRes, coreRes); err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	oldCoreRes := &mesh.MeshResource{}
	oldK8sRes := &v1alpha1.Mesh{}
	if err := h.decoder.DecodeRaw(req.OldObject, oldK8sRes); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	if err := h.converter.ToCoreResource(oldK8sRes, oldCoreRes); err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	if err := h.validator.ValidateUpdate(ctx, oldCoreRes, coreRes); err != nil {
		if kumaErr, ok := err.(*validators.ValidationError); ok {
			return convertValidationError(kumaErr, k8sRes)
		}
		return admission.Denied(err.Error())
	}
	return admission.Allowed("")
}

func (h *MeshValidator) Supports(req admission.Request) bool {
	gvk := mesh_k8s.GroupVersion.WithKind("Mesh")
	return req.Kind.Kind == gvk.Kind && req.Kind.Version == gvk.Version && req.Kind.Group == gvk.Group
}
