import{d as k,i as y,aw as f,o as s,c as t,a as n,Q as g,u as a,w as l,f as o,e as r,z as c,y as d,F as v,A as w,b as B,q as b,j as E,B as m,N as x,O as I,J as N}from"./index-5f1fbf13.js";const h=e=>(x("data-v-4154f15d"),e=e(),I(),e),S={class:"error-block"},C=h(()=>o("p",null,"An error has occurred while trying to load this data.",-1)),V={class:"error-block-details"},q=h(()=>o("summary",null,"Details",-1)),A={key:0},z={key:0,class:"badge-list"},D=k({__name:"ErrorBlock",props:{error:{type:[Error,null],required:!1,default:null}},setup(e){const i=e,u=y(()=>i.error instanceof f?i.error.causes:[]);return(F,j)=>(s(),t("div",S,[n(a(b),{"cta-is-hidden":""},g({title:l(()=>[n(a(B),{class:"mb-3",icon:"warning",color:"var(--black-500)","secondary-color":"var(--yellow-300)",size:"42"}),r(),C]),_:2},[e.error!==null||a(u).length>0?{name:"message",fn:l(()=>[o("details",V,[q,r(),e.error!==null?(s(),t("p",A,c(e.error.message),1)):d("",!0),r(),o("ul",null,[(s(!0),t(v,null,w(a(u),(_,p)=>(s(),t("li",{key:p},[o("b",null,[o("code",null,c(_.field),1)]),r(": "+c(_.message),1)]))),128))])])]),key:"0"}:void 0]),1024),r(),e.error instanceof a(f)?(s(),t("div",z,[e.error.code?(s(),E(a(m),{key:0,appearance:"warning"},{default:l(()=>[r(c(e.error.code),1)]),_:1})):d("",!0),r(),n(a(m),{appearance:"warning"},{default:l(()=>[r(c(e.error.statusCode),1)]),_:1})])):d("",!0)]))}});const L=N(D,[["__scopeId","data-v-4154f15d"]]);export{L as E};