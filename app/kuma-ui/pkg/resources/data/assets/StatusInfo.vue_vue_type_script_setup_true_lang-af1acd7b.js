import{_ as l}from"./EmptyBlock.vue_vue_type_script_setup_true_lang-7207a998.js";import{E as n}from"./ErrorBlock-c00ab7a0.js";import{i as s}from"./RouteView.vue_vue_type_script_setup_true_lang-1e381d15.js";import{d as i,o as r,e as a,a as o,n as f}from"./index-452eba2d.js";const m={key:3},p=i({__name:"StatusInfo",props:{isLoading:{type:Boolean,default:!1},hasError:{type:Boolean,default:!1},isEmpty:{type:Boolean,default:!1},error:{type:[Error,null],required:!1,default:null}},setup(e){return(t,u)=>(r(),a("div",null,[e.isLoading?(r(),o(s,{key:0})):e.hasError||e.error!==null?(r(),o(n,{key:1,error:e.error},null,8,["error"])):e.isEmpty?(r(),o(l,{key:2})):(r(),a("div",m,[f(t.$slots,"default")]))]))}});export{p as _};