import{_ as o,a as c,M as y}from"./MeshResources-46669578.js";import{u as P}from"./store-eba5eacb.js";import{d as w,c as a,y as x,h as k,g as z,u as t,a as g,b as u,f as s,e as r,F as I,o as l}from"./runtime-dom.esm-bundler-91b41870.js";import{_ as S}from"./_plugin-vue_export-helper-c27b6911.js";import"./vue-router-573afc44.js";import"./Instance-eddcbc9c.js";import"./kongponents.es-3df60cd6.js";import"./index-5f465f86.js";import"./vuex.esm-bundler-df5bd11e.js";import"./constants-31fdaf55.js";const D={class:"chart-container mt-16"},N=w({__name:"MainOverviewView",setup(O){const e=P(),i=a(()=>e.getters["config/getMulticlusterStatus"]),_=a(()=>e.getters.getServiceResourcesFetching),h=a(()=>e.getters.getZonesInsightsFetching),n=a(()=>e.getters.getMeshInsightsFetching),m=a(()=>e.getters.getChart("services")),p=a(()=>e.getters.getChart("dataplanes")),f=a(()=>e.getters.getChart("meshes")),v=a(()=>e.getters.getChart("zones")),C=a(()=>e.getters.getChart("zonesCPVersions")),M=a(()=>e.getters.getChart("kumaDPVersions")),V=a(()=>e.getters.getChart("envoyVersions"));x(()=>i.value,function(){d()}),d();function d(){e.dispatch("fetchMeshInsights"),e.dispatch("fetchServices"),e.dispatch("fetchZonesInsights",i.value),i.value&&e.dispatch("fetchTotalClusterCount")}return(Z,F)=>(l(),k(I,null,[z("div",D,[t(i)?(l(),g(o,{key:0,class:"chart chart-1/2 chart-offset-left-1/6",title:{singular:"Zone",plural:"Zones"},data:t(v).data,url:{name:"zones"},"is-loading":t(h)},null,8,["data","is-loading"])):u("",!0),s(),t(i)?(l(),g(c,{key:1,class:"chart chart-1/2 chart-offset-right-1/6",title:"ZONE CP",data:t(C).data,url:{name:"zones"},"is-loading":t(h)},null,8,["data","is-loading"])):u("",!0),s(),r(o,{class:"chart chart-1/3",title:{singular:"Mesh",plural:"Meshes"},data:t(f).data,"is-loading":t(n)},null,8,["data","is-loading"]),s(),r(o,{class:"chart chart-1/3",title:{singular:"Service",plural:"Services"},data:t(m).data,"is-loading":t(_),"save-chart":""},null,8,["data","is-loading"]),s(),r(o,{class:"chart chart-1/3",title:{singular:"DP Proxy",plural:"DP Proxies"},data:t(p).data,"is-loading":t(n)},null,8,["data","is-loading"]),s(),r(c,{class:"chart chart-1/2 chart-offset-left-1/6",title:"KUMA DP",data:t(M).data,"is-loading":t(n)},null,8,["data","is-loading"]),s(),r(c,{class:"chart chart-1/2 chart-offset-right-1/6",title:"ENVOY",data:t(V).data,"is-loading":t(n),"display-am-charts-logo":""},null,8,["data","is-loading"])]),s(),r(y,{class:"mt-8"})],64))}});const Y=S(N,[["__scopeId","data-v-d759b307"]]);export{Y as default};