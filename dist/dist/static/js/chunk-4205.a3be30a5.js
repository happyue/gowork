(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-4205"],{"1Xwi":function(e,t,n){},"E7+A":function(e,t,n){"use strict";var l=n("Pk8f");n.n(l).a},JsRX:function(e,t,n){"use strict";n.r(t);var l=n("GQeE"),a=n.n(l),i=[{id:0,event:"Event-0",timeLine:50},{id:1,event:"Event-1",timeLine:100,children:[{id:2,event:"Event-2",timeLine:10},{id:3,event:"Event-3",timeLine:90,children:[{id:4,event:"Event-4",timeLine:5},{id:5,event:"Event-5",timeLine:10},{id:6,event:"Event-6",timeLine:75,children:[{id:7,event:"Event-7",timeLine:50,children:[{id:71,event:"Event-7-1",timeLine:25},{id:72,event:"Event-7-2",timeLine:5},{id:73,event:"Event-7-3",timeLine:20}]},{id:8,event:"Event-8",timeLine:25}]}]}]}],o={name:"TreeTableDemo",components:{treeTable:n("itRl").a},data:function(){return{defaultExpandAll:!1,showCheckbox:!0,key:1,columns:[{label:"Checkbox",checkbox:!0},{label:"",key:"id",expand:!0},{label:"Event",key:"event",width:200,align:"left"},{label:"Scope",key:"scope"},{label:"Operation",key:"operation"}],data:i}},watch:{showCheckbox:function(e){e?this.columns.unshift({label:"Checkbox",checkbox:!0}):this.columns.shift(),this.reset()}},methods:{reset:function(){++this.key},click:function(e){console.log(e);var t=e.row,n=a()(t).map(function(e){return"<p>"+e+": "+t[e]+"</p>"}).join("");this.$notify({title:"Success",dangerouslyUseHTMLString:!0,message:n,type:"success"})}}},s=(n("E7+A"),n("KHd+")),c=Object(s.a)(o,function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"app-container"},[n("div",{staticStyle:{"margin-bottom":"20px"}},[n("el-button",{staticClass:"option-item",attrs:{type:"primary",size:"small"}},[n("a",{attrs:{href:"https://github.com/PanJiaChen/vue-element-admin/tree/master/src/components/TreeTable",target:"_blank"}},[e._v("Documentation")])]),e._v(" "),n("div",{staticClass:"option-item"},[n("el-tag",[e._v("Expand All")]),e._v(" "),n("el-switch",{attrs:{"active-color":"#13ce66","inactive-color":"#ff4949"},on:{change:e.reset},model:{value:e.defaultExpandAll,callback:function(t){e.defaultExpandAll=t},expression:"defaultExpandAll"}})],1),e._v(" "),n("div",{staticClass:"option-item"},[n("el-tag",[e._v("Show Checkbox")]),e._v(" "),n("el-switch",{attrs:{"active-color":"#13ce66","inactive-color":"#ff4949"},model:{value:e.showCheckbox,callback:function(t){e.showCheckbox=t},expression:"showCheckbox"}})],1)],1),e._v(" "),n("tree-table",{key:e.key,attrs:{"default-expand-all":e.defaultExpandAll,data:e.data,columns:e.columns,border:""},scopedSlots:e._u([{key:"scope",fn:function(t){var l=t.scope;return[n("el-tag",[e._v("level: "+e._s(l.row._level))]),e._v(" "),n("el-tag",[e._v("expand: "+e._s(l.row._expand))]),e._v(" "),n("el-tag",[e._v("select: "+e._s(l.row._select))])]}},{key:"operation",fn:function(t){var l=t.scope;return[n("el-button",{attrs:{type:"primary",size:""},on:{click:function(t){e.click(l)}}},[e._v("Click")])]}}])})],1)},[],!1,null,"4ffa88db",null);c.options.__file="index.vue";t.default=c.exports},Pk8f:function(e,t,n){},itRl:function(e,t,n){"use strict";var l=n("4d7F"),a=n.n(l),i=n("Kw5r");function o(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},n=t.parent,l=void 0===n?null:n,a=t.preIndex,s=void 0!==a&&a,c=t.level,r=void 0===c?0:c,d=t.expand,u=void 0!==d&&d,h=t.children,p=void 0===h?"children":h,v=t.show,f=void 0===v||v,_=t.select,w=void 0!==_&&_;e.forEach(function(e,t){var n=(s?s+"-"+t:t)+"";i.default.set(e,"_id",n),i.default.set(e,"_level",r),i.default.set(e,"_expand",u),i.default.set(e,"_parent",l),i.default.set(e,"_show",f),i.default.set(e,"_select",w),e[p]&&e[p].length>0&&o(e[p],{parent:e,level:r+1,expand:u,preIndex:n,children:p,status:status,select:w})})}var s={name:"TreeTable",props:{data:{type:Array,required:!0,default:function(){return[]}},columns:{type:Array,default:function(){return[]}},defaultExpandAll:{type:Boolean,default:!1},defaultChildren:{type:String,default:"children"},indent:{type:Number,default:50}},data:function(){return{guard:1}},computed:{children:function(){return this.defaultChildren},tableData:function(){var e=this.data;return 0===this.data.length?[]:(o(e,{expand:this.defaultExpandAll,children:this.defaultChildren}),function e(t){var n=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"children",l=[];return t.forEach(function(t,a){if(i.default.set(t,"_index",a),l.push(t),t[n]&&t[n].length>0){var o=e(t[n],n);l=l.concat(o)}}),l}(e,this.defaultChildren))}},methods:{addBrother:function(e,t){e._parent?e._parent.children.push(t):this.data.push(t)},addChild:function(e,t){e.children||this.$set(e,"children",[]),e.children.push(t)},delete:function(e){var t=e._index,n=e._parent;n?n.children.splice(t,1):this.data.splice(t,1)},getData:function(){return this.tableData},showRow:function(e){var t=e.row,n=t._parent,l=!n||n._expand&&n._show;return t._show=l,l?"animation:treeTableShow 1s;-webkit-animation:treeTableShow 1s;":"display:none;"},showSperadIcon:function(e){return e[this.children]&&e[this.children].length>0},toggleExpanded:function(e){var t=this.tableData[e],n=!t._expand;t._expand=n},handleCheckAllChange:function(e){this.selcetRecursion(e,e._select,this.defaultChildren),this.isIndeterminate=e._select},selcetRecursion:function(e,t){var n=this,l=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"children";t&&(this.$set(e,"_expand",!0),this.$set(e,"_show",!0));var a=e[l];a&&a.length>0&&a.map(function(e){e._select=t,n.selcetRecursion(e,t,l)})},updateTreeNode:function(e){var t=this;return new a.a(function(n){var l=e._id,a=e._parent,i=l.split("-").slice(-1)[0];a?(a.children.splice(i,1,e),n(t.data)):(t.data.splice(i,1,e),n(t.data))})}}},c=(n("x7Y3"),n("KHd+")),r=Object(c.a)(s,function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-table",e._g(e._b({attrs:{data:e.tableData,"row-style":e.showRow}},"el-table",e.$attrs,!1),e.$listeners),[e._t("selection"),e._v(" "),e._t("pre-column"),e._v(" "),e._l(e.columns,function(t){return n("el-table-column",{key:t.key,attrs:{label:t.label,width:t.width,align:t.align||"center","header-align":t.headerAlign},scopedSlots:e._u([{key:"default",fn:function(l){return[e._t(t.key,[t.expand?[n("span",{style:{"padding-left":+l.row._level*e.indent+"px"}}),e._v(" "),n("span",{directives:[{name:"show",rawName:"v-show",value:e.showSperadIcon(l.row),expression:"showSperadIcon(scope.row)"}],staticClass:"tree-ctrl",on:{click:function(t){e.toggleExpanded(l.$index)}}},[l.row._expand?n("i",{staticClass:"el-icon-minus"}):n("i",{staticClass:"el-icon-plus"})])]:e._e(),e._v(" "),t.checkbox?[l.row[e.defaultChildren]&&l.row[e.defaultChildren].length>0?n("el-checkbox",{style:{"padding-left":+l.row._level*e.indent+"px"},attrs:{indeterminate:l.row._select},on:{change:function(t){e.handleCheckAllChange(l.row)}},model:{value:l.row._select,callback:function(t){e.$set(l.row,"_select",t)},expression:"scope.row._select"}}):n("el-checkbox",{style:{"padding-left":+l.row._level*e.indent+"px"},on:{change:function(t){e.handleCheckAllChange(l.row)}},model:{value:l.row._select,callback:function(t){e.$set(l.row,"_select",t)},expression:"scope.row._select"}})]:e._e(),e._v("\n        "+e._s(l.row[t.key])+"\n      ")],{scope:l})]}}])})})],2)},[],!1,null,null,null);r.options.__file="index.vue";t.a=r.exports},x7Y3:function(e,t,n){"use strict";var l=n("1Xwi");n.n(l).a}}]);