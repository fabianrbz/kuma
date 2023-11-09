import{A as q,a as F}from"./AccordionList-45f507ad.js";import{d as j,R as x,a as N,o as t,c as o,p as y,f as a,F as m,D as E,t as f,e as h,w as i,v as S,b as _,a5 as Q,B as W,C as X,_ as H,q as U,a8 as Y,H as A,I as K}from"./index-9c7b51f6.js";import{_ as z}from"./CodeBlock.vue_vue_type_style_index_0_lang-da007c9a.js";import{P as J}from"./PolicyTypeTag-637d362a.js";import{T as G}from"./TagList-b2ff20af.js";import{t as M}from"./toYaml-4e00099e.js";const O=b=>(W("data-v-d33ad9ee"),b=b(),X(),b),Z={class:"policies-list","data-testid":"builtin-gateway-dataplane-policies"},V={class:"mesh-gateway-policy-list"},ee=O(()=>y("h3",{class:"mb-2"},`
        Gateway policies
      `,-1)),te={key:0},se=O(()=>y("h3",{class:"mt-6 mb-2"},`
        Listeners
      `,-1)),ne=O(()=>y("b",null,"Host",-1)),oe=O(()=>y("h4",{class:"mt-2 mb-2"},`
                Routes
              `,-1)),ae={class:"dataplane-policy-header"},ie=O(()=>y("b",null,"Route",-1)),le=O(()=>y("b",null,"Service",-1)),ce={key:0,class:"badge-list"},re={class:"mt-1"},pe=j({__name:"BuiltinGatewayPolicies",props:{gatewayDataplane:{},policyTypesByName:{}},setup(b){const k=b,C=x(()=>B(k.gatewayDataplane)),L=x(()=>T(k.gatewayDataplane.policies));function B(v){const d=[],n=v.listeners??[];for(const e of n)for(const l of e.hosts)for(const s of l.routes){const c=[];for(const r of s.destinations){const g=T(r.policies),u={routeName:s.route,route:{name:"policy-detail-view",params:{mesh:v.gateway.mesh,policyPath:"meshgatewayroutes",policy:s.route}},service:r.tags["kuma.io/service"],policies:g};c.push(u)}d.push({protocol:e.protocol,port:e.port,hostName:l.hostName,routeEntries:c})}return d}function T(v){if(v===void 0)return[];const d=[];for(const n of Object.values(v)){const e=k.policyTypesByName[n.type];d.push({type:n.type,name:n.name,route:{name:"policy-detail-view",params:{mesh:n.mesh,policyPath:e.path,policy:n.name}}})}return d}return(v,d)=>{const n=N("router-link"),e=N("KBadge"),l=N("RouterLink");return t(),o("div",Z,[y("div",V,[ee,a(),L.value.length>0?(t(),o("ul",te,[(t(!0),o(m,null,E(L.value,(s,c)=>(t(),o("li",{key:c},[y("span",null,f(s.type),1),a(`:

          `),h(n,{to:s.route},{default:i(()=>[a(f(s.name),1)]),_:2},1032,["to"])]))),128))])):S("",!0),a(),se,a(),y("div",null,[(t(!0),o(m,null,E(C.value,(s,c)=>(t(),o("div",{key:c},[y("div",null,[y("div",null,[ne,a(": "+f(s.hostName)+":"+f(s.port)+" ("+f(s.protocol)+`)
            `,1)]),a(),s.routeEntries.length>0?(t(),o(m,{key:0},[oe,a(),h(F,{"initially-open":[],"multiple-open":""},{default:i(()=>[(t(!0),o(m,null,E(s.routeEntries,(r,g)=>(t(),_(q,{key:g},Q({"accordion-header":i(()=>[y("div",ae,[y("div",null,[y("div",null,[ie,a(": "),h(n,{to:r.route},{default:i(()=>[a(f(r.routeName),1)]),_:2},1032,["to"])]),a(),y("div",null,[le,a(": "+f(r.service),1)])]),a(),r.policies.length>0?(t(),o("div",ce,[(t(!0),o(m,null,E(r.policies,(u,p)=>(t(),_(e,{key:`${c}-${p}`},{default:i(()=>[a(f(u.type),1)]),_:2},1024))),128))])):S("",!0)])]),_:2},[r.policies.length>0?{name:"accordion-content",fn:i(()=>[y("ul",re,[(t(!0),o(m,null,E(r.policies,(u,p)=>(t(),o("li",{key:`${c}-${p}`},[a(f(u.type)+`:

                        `,1),h(l,{to:u.route},{default:i(()=>[a(f(u.name),1)]),_:2},1032,["to"])]))),128))])]),key:"0"}:void 0]),1024))),128))]),_:2},1024)],64)):S("",!0)])]))),128))])])])}}});const ue=H(pe,[["__scopeId","data-v-d33ad9ee"]]),ye={class:"policy-type-heading"},de={class:"policy-list"},_e={key:0},me=j({__name:"PolicyTypeEntryList",props:{id:{type:String,required:!1,default:"entry-list"},policyTypeEntries:{type:Object,required:!0}},setup(b){const k=[{label:"From",key:"sourceTags"},{label:"To",key:"destinationTags"},{label:"On",key:"name"},{label:"Conf",key:"config"},{label:"Origin policies",key:"origins"}],C=b;function L({headerKey:B}){return{class:`cell-${B}`}}return(B,T)=>{const v=N("router-link");return t(),_(F,{"initially-open":[],"multiple-open":""},{default:i(()=>[(t(!0),o(m,null,E(C.policyTypeEntries,(d,n)=>(t(),_(q,{key:n},{"accordion-header":i(()=>[y("h3",ye,[h(J,{"policy-type":d.type},{default:i(()=>[a(f(d.type)+" ("+f(d.connections.length)+`)
          `,1)]),_:2},1032,["policy-type"])])]),"accordion-content":i(()=>[y("div",de,[h(U(Y),{class:"policy-type-table",fetcher:()=>({data:d.connections,total:d.connections.length}),headers:k,"cell-attrs":L,"disable-pagination":"","is-clickable":""},{sourceTags:i(({rowValue:e})=>[e.length>0?(t(),_(G,{key:0,class:"tag-list",tags:e},null,8,["tags"])):(t(),o(m,{key:1},[a(`
                —
              `)],64))]),destinationTags:i(({rowValue:e})=>[e.length>0?(t(),_(G,{key:0,class:"tag-list",tags:e},null,8,["tags"])):(t(),o(m,{key:1},[a(`
                —
              `)],64))]),name:i(({rowValue:e})=>[e!==null?(t(),o(m,{key:0},[a(f(e),1)],64)):(t(),o(m,{key:1},[a(`
                —
              `)],64))]),origins:i(({rowValue:e})=>[e.length>0?(t(),o("ul",_e,[(t(!0),o(m,null,E(e,(l,s)=>(t(),o("li",{key:`${n}-${s}`},[h(v,{to:l.route},{default:i(()=>[a(f(l.name),1)]),_:2},1032,["to"])]))),128))])):(t(),o(m,{key:1},[a(`
                —
              `)],64))]),config:i(({rowValue:e,rowKey:l})=>[e!==null?(t(),_(z,{key:0,id:`${C.id}-${n}-${l}-code-block`,code:e,language:"yaml","show-copy-button":!1},null,8,["id","code"])):(t(),o(m,{key:1},[a(`
                —
              `)],64))]),_:2},1032,["fetcher"])])]),_:2},1024))),128))]),_:1})}}});const he=H(me,[["__scopeId","data-v-9a1971d5"]]),ge={class:"policy-type-heading"},fe={class:"policy-list"},ke={key:1,class:"tag-list-wrapper"},ve={key:0},be={key:1},Te={key:0},$e={key:0},Pe=j({__name:"RuleEntryList",props:{id:{type:String,required:!1,default:"entry-list"},ruleEntries:{type:Object,required:!0}},setup(b){const k=[{label:"Type",key:"type"},{label:"Addresses",key:"addresses"},{label:"Conf",key:"config"},{label:"Origin policies",key:"origins"}],C=b;function L({headerKey:B}){return{class:`cell-${B}`}}return(B,T)=>{const v=N("router-link");return t(),_(F,{"initially-open":[],"multiple-open":""},{default:i(()=>[(t(!0),o(m,null,E(C.ruleEntries,(d,n)=>(t(),_(q,{key:n},{"accordion-header":i(()=>[y("h3",ge,[h(J,{"policy-type":d.type},{default:i(()=>[a(f(d.type)+" ("+f(d.connections.length)+`)
          `,1)]),_:2},1032,["policy-type"])])]),"accordion-content":i(()=>[y("div",fe,[h(U(Y),{class:"policy-type-table",fetcher:()=>({data:d.connections,total:d.connections.length}),headers:k,"cell-attrs":L,"disable-pagination":"","is-clickable":""},{type:i(({rowValue:e})=>[e.sourceTags.length===0&&e.destinationTags.length===0?(t(),o(m,{key:0},[a(`
                —
              `)],64)):(t(),o("div",ke,[e.sourceTags.length>0?(t(),o("div",ve,[a(`
                  From

                  `),h(G,{class:"tag-list",tags:e.sourceTags},null,8,["tags"])])):S("",!0),a(),e.destinationTags.length>0?(t(),o("div",be,[a(`
                  To

                  `),h(G,{class:"tag-list",tags:e.destinationTags},null,8,["tags"])])):S("",!0)]))]),addresses:i(({rowValue:e})=>[e.length>0?(t(),o("ul",Te,[(t(!0),o(m,null,E(e,(l,s)=>(t(),o("li",{key:`${n}-${s}`},f(l),1))),128))])):(t(),o(m,{key:1},[a(`
                —
              `)],64))]),origins:i(({rowValue:e})=>[e.length>0?(t(),o("ul",$e,[(t(!0),o(m,null,E(e,(l,s)=>(t(),o("li",{key:`${n}-${s}`},[h(v,{to:l.route},{default:i(()=>[a(f(l.name),1)]),_:2},1032,["to"])]))),128))])):(t(),o(m,{key:1},[a(`
                —
              `)],64))]),config:i(({rowValue:e,rowKey:l})=>[e!==null?(t(),_(z,{key:0,id:`${C.id}-${n}-${l}-code-block`,code:e,language:"yaml","show-copy-button":!1},null,8,["id","code"])):(t(),o(m,{key:1},[a(`
                —
              `)],64))]),_:2},1032,["fetcher"])])]),_:2},1024))),128))]),_:1})}}});const we=H(Pe,[["__scopeId","data-v-3e59037c"]]),Be={"data-testid":"standard-dataplane-policies"},Re=y("h2",{class:"visually-hidden"},`
      Policies
    `,-1),Ce={key:0,class:"mt-2"},Le=y("h2",{class:"mb-2"},`
        Rules
      `,-1),Ne=j({__name:"StandardDataplanePolicies",props:{sidecarDataplanes:{},rules:{},policyTypesByName:{}},setup(b){const k=b,C=x(()=>B(k.sidecarDataplanes)),L=x(()=>v(k.rules));function B(n){const e=new Map;for(const s of n){const{type:c,service:r}=s,g=typeof r=="string"&&r!==""?[{label:"kuma.io/service",value:r}]:[],u=c==="inbound"||c==="outbound"?s.name:null;for(const[p,$]of Object.entries(s.matchedPolicies)){e.has(p)||e.set(p,{type:p,connections:[]});const R=e.get(p),P=k.policyTypesByName[p];for(const I of $){const w=T(I,P,s,g,u);R.connections.push(...w)}}}const l=Array.from(e.values());return l.sort((s,c)=>s.type.localeCompare(c.type)),l}function T(n,e,l,s,c){const r=n.conf&&Object.keys(n.conf).length>0?M(n.conf):null,u=[{name:n.name,route:{name:"policy-detail-view",params:{mesh:n.mesh,policyPath:e.path,policy:n.name}}}],p=[];if(l.type==="inbound"&&Array.isArray(n.sources))for(const{match:$}of n.sources){const P={sourceTags:[{label:"kuma.io/service",value:$["kuma.io/service"]}],destinationTags:s,name:c,config:r,origins:u};p.push(P)}else{const R={sourceTags:[],destinationTags:s,name:c,config:r,origins:u};p.push(R)}return p}function v(n){const e=new Map;for(const s of n){e.has(s.policyType)||e.set(s.policyType,{type:s.policyType,connections:[]});const c=e.get(s.policyType),r=k.policyTypesByName[s.policyType],g=d(s,r);c.connections.push(...g)}const l=Array.from(e.values());return l.sort((s,c)=>s.type.localeCompare(c.type)),l}function d(n,e){const{type:l,service:s,subset:c,conf:r}=n,g=c?Object.entries(c):[];let u,p;l==="ClientSubset"?g.length>0?u=g.map(([w,D])=>({label:w,value:D})):u=[{label:"kuma.io/service",value:"*"}]:u=[],l==="DestinationSubset"?g.length>0?p=g.map(([w,D])=>({label:w,value:D})):typeof s=="string"&&s!==""?p=[{label:"kuma.io/service",value:s}]:p=[{label:"kuma.io/service",value:"*"}]:l==="ClientSubset"&&typeof s=="string"&&s!==""?p=[{label:"kuma.io/service",value:s}]:p=[];const $=n.addresses??[],R=r&&Object.keys(r).length>0?M(r):null,P=[];for(const w of n.origins)P.push({name:w.name,route:{name:"policy-detail-view",params:{mesh:w.mesh,policyPath:e.path,policy:w.name}}});return[{type:{sourceTags:u,destinationTags:p},addresses:$,config:R,origins:P}]}return(n,e)=>(t(),o("div",Be,[Re,a(),h(he,{id:"policies","policy-type-entries":C.value,"data-testid":"policy-list"},null,8,["policy-type-entries"]),a(),L.value.length>0?(t(),o("div",Ce,[Le,a(),h(we,{id:"rules","rule-entries":L.value,"data-testid":"rule-list"},null,8,["rule-entries"])])):S("",!0)]))}}),xe=j({__name:"DataPlanePoliciesView",props:{data:{}},setup(b){const k=b;return(C,L)=>{const B=N("RouteTitle"),T=N("DataSource"),v=N("KCard"),d=N("AppView"),n=N("RouteView");return t(),_(n,{name:"data-plane-policies-view",params:{mesh:"",dataPlane:""}},{default:i(({route:e,t:l})=>[h(d,null,{title:i(()=>[y("h2",null,[h(B,{title:l("data-planes.routes.item.navigation.data-plane-policies-view"),render:!0},null,8,["title"])])]),default:i(()=>[a(),h(v,null,{body:i(()=>{var s,c,r;return[((r=(c=(s=k.data.dataplane)==null?void 0:s.networking)==null?void 0:c.gateway)==null?void 0:r.type)==="BUILTIN"?(t(),_(T,{key:0,src:"/*/policy-types"},{default:i(({data:g,error:u})=>[h(T,{src:`/meshes/${e.params.mesh}/dataplanes/${e.params.dataPlane}/gateway-dataplane-policies`},{default:i(({data:p,error:$})=>[u?(t(),_(A,{key:0,error:u},null,8,["error"])):$?(t(),_(A,{key:1,error:$},null,8,["error"])):p===void 0||g===void 0?(t(),_(K,{key:2})):(t(),_(ue,{key:3,"policy-types-by-name":g.policies.reduce((R,P)=>Object.assign(R,{[P.name]:P}),{}),"gateway-dataplane":p},null,8,["policy-types-by-name","gateway-dataplane"]))]),_:2},1032,["src"])]),_:2},1024)):(t(),_(T,{key:1,src:"/*/policy-types"},{default:i(({data:g,error:u})=>[h(T,{src:`/meshes/${e.params.mesh}/dataplanes/${e.params.dataPlane}/sidecar-dataplane-policies`},{default:i(({data:p,error:$})=>[h(T,{src:`/meshes/${e.params.mesh}/dataplanes/${e.params.dataPlane}/rules`},{default:i(({data:R,error:P})=>[u?(t(),_(A,{key:0,error:u},null,8,["error"])):$?(t(),_(A,{key:1,error:$},null,8,["error"])):P?(t(),_(A,{key:2,error:P},null,8,["error"])):g===void 0||p===void 0||R===void 0?(t(),_(K,{key:3})):(t(),_(Ne,{key:4,"policy-types-by-name":g.policies.reduce((I,w)=>Object.assign(I,{[w.name]:w}),{}),"sidecar-dataplanes":p.items,rules:R.items},null,8,["policy-types-by-name","sidecar-dataplanes","rules"]))]),_:2},1032,["src"])]),_:2},1032,["src"])]),_:2},1024))]}),_:2},1024)]),_:2},1024)]),_:1})}}});export{xe as default};
