import{u as h}from"./vue-router-573afc44.js";import{k as A}from"./store-eba5eacb.js";import{Q as f}from"./QueryParameter-70743f73.js";import{D as P}from"./DataPlaneList-19d3f5d0.js";import{d as S,r as i,c as T,y as k,a as x,u as I,o as b}from"./runtime-dom.esm-bundler-91b41870.js";import"./vuex.esm-bundler-df5bd11e.js";import"./constants-31fdaf55.js";import"./datadogLogEvents-4578cfa7.js";import"./kongponents.es-3df60cd6.js";import"./ContentWrapper-b0fa1a61.js";import"./_plugin-vue_export-helper-c27b6911.js";import"./DataOverview-fc404ed4.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang-4047971f.js";import"./ErrorBlock-03bc2ec2.js";import"./LoadingBlock.vue_vue_type_script_setup_true_lang-d3176fee.js";import"./StatusBadge-81464ebd.js";import"./TagList-91d1133a.js";import"./YamlView.vue_vue_type_script_setup_true_lang-ca12ecd4.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-6be6d47a.js";import"./_commonjsHelpers-87174ba5.js";import"./toYaml-4e00099e.js";const ee=S({__name:"DataPlaneListView",props:{selectedDppName:{type:String,required:!1,default:null},gatewayType:{type:String,required:!1,default:"true"},offset:{type:Number,required:!1,default:0},isGatewayView:{type:Boolean,required:!1,default:!1}},setup(d){const t=d,v=50,w={name:{description:"filter by name or parts of a name"},service:{description:"filter by “kuma.io/service” value"},tag:{description:"filter by tags (e.g. “tag: version:2”)"},zone:{description:"filter by “kuma.io/zone” value"}},g={protocol:{description:"filter by “kuma.io/protocol” value"}},_={},s=h(),o=i([]),p=i(!0),c=i(null),l=i(null),y=i(t.offset),D=T(()=>({...w,...t.isGatewayView?_:g}));k(()=>s.params.mesh,function(){s.name!=="data-plane-list-view"&&s.name!=="gateway-list-view"||u(0)});function E(){const a=f.get("filterFields"),r=a!==null?JSON.parse(a):{};u(t.offset,{...r,gateway:t.gatewayType})}E();async function u(a,r={}){y.value=a,f.set("offset",a>0?a:null),f.set("gatewayType",r.gateway==="true"?"all":r.gateway),p.value=!0;const m=s.params.mesh,n=F(r,v,a,t.isGatewayView);try{const{items:e,next:L}=await A.getAllDataplaneOverviewsFromMesh({mesh:m},n);Array.isArray(e)&&e.length>0?(o.value=e,l.value=L):(o.value=[],l.value=null)}catch(e){e instanceof Error?c.value=e:console.error(e),o.value=[],l.value=null}finally{p.value=!1}}function F(a,r,m,n){const e={...a,size:r,offset:m};return n&&(!("gateway"in e)||e.gateway==="false")?e.gateway="true":n||(e.gateway="false"),e}return(a,r)=>(b(),x(P,{"data-plane-overviews":o.value,"is-loading":p.value,error:c.value,"next-url":l.value,"page-offset":y.value,"selected-dpp-name":t.selectedDppName,"is-gateway-view":t.isGatewayView,"gateway-type":t.gatewayType,"dpp-filter-fields":I(D),onLoadData:u},null,8,["data-plane-overviews","is-loading","error","next-url","page-offset","selected-dpp-name","is-gateway-view","gateway-type","dpp-filter-fields"]))}});export{ee as default};