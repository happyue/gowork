(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-4675"],{"8PlU":function(e,t,a){},WXpr:function(e,t,a){"use strict";a.r(t);var n=a("gDS+"),i=a.n(n),s=a("P2sY"),r=a.n(s),o=a("t3Un");a("jUE0");var l="@@wavesContext";function u(e,t){function a(a){var n=r()({},t.value),i=r()({ele:e,type:"hit",color:"rgba(0, 0, 0, 0.15)"},n),s=i.ele;if(s){s.style.position="relative",s.style.overflow="hidden";var o=s.getBoundingClientRect(),l=s.querySelector(".waves-ripple");switch(l?l.className="waves-ripple":((l=document.createElement("span")).className="waves-ripple",l.style.height=l.style.width=Math.max(o.width,o.height)+"px",s.appendChild(l)),i.type){case"center":l.style.top=o.height/2-l.offsetHeight/2+"px",l.style.left=o.width/2-l.offsetWidth/2+"px";break;default:l.style.top=(a.pageY-o.top-l.offsetHeight/2-document.documentElement.scrollTop||document.body.scrollTop)+"px",l.style.left=(a.pageX-o.left-l.offsetWidth/2-document.documentElement.scrollLeft||document.body.scrollLeft)+"px"}return l.style.backgroundColor=i.color,l.className="waves-ripple z-active",!1}}return e[l]?e[l].removeHandle=a:e[l]={removeHandle:a},a}var d={bind:function(e,t){e.addEventListener("click",u(e,t),!1)},update:function(e,t){e.removeEventListener("click",e[l].removeHandle,!1),e.addEventListener("click",u(e,t),!1)},unbind:function(e){e.removeEventListener("click",e[l].removeHandle,!1),e[l]=null,delete e[l]}},c=function(e){e.directive("waves",d)};window.Vue&&(window.waves=d,Vue.use(c)),d.install=c;var p=d;Math.easeInOutQuad=function(e,t,a,n){return(e/=n/2)<1?a/2*e*e+t:-a/2*(--e*(e-2)-1)+t};var m=window.requestAnimationFrame||window.webkitRequestAnimationFrame||window.mozRequestAnimationFrame||function(e){window.setTimeout(e,1e3/60)};function f(e,t,a){var n=document.documentElement.scrollTop||document.body.parentNode.scrollTop||document.body.scrollTop,i=e-n,s=0;t=void 0===t?500:t;!function e(){s+=20,function(e){document.documentElement.scrollTop=e,document.body.parentNode.scrollTop=e,document.body.scrollTop=e}(Math.easeInOutQuad(s,n,i,t)),s<t?m(e):a&&"function"==typeof a&&a()}()}var h={name:"Pagination",props:{total:{required:!0,type:Number},page:{type:Number,default:1},limit:{type:Number,default:20},pageSizes:{type:Array,default:function(){return[10,20,30,50]}},layout:{type:String,default:"total, sizes, prev, pager, next, jumper"},background:{type:Boolean,default:!0},autoScroll:{type:Boolean,default:!0},hidden:{type:Boolean,default:!1}},computed:{currentPage:{get:function(){return this.page},set:function(e){this.$emit("update:page",e)}},pageSize:{get:function(){return this.limit},set:function(e){this.$emit("update:limit",e)}}},methods:{handleSizeChange:function(e){this.$emit("pagination",{page:this.currentPage,limit:e}),this.autoScroll&&f(0,800)},handleCurrentChange:function(e){this.$emit("pagination",{page:e,limit:this.pageSize}),this.autoScroll&&f(0,800)}}},v=(a("a5W9"),a("KHd+")),b=Object(v.a)(h,function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"pagination-container",class:{hidden:e.hidden}},[a("el-pagination",e._b({attrs:{background:e.background,"current-page":e.currentPage,"page-size":e.pageSize,layout:e.layout,"page-sizes":e.pageSizes,total:e.total},on:{"update:currentPage":function(t){e.currentPage=t},"update:pageSize":function(t){e.pageSize=t},"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}},"el-pagination",e.$attrs,!1))],1)},[],!1,null,"0565df52",null);b.options.__file="index.vue";var g={name:"SSHList",components:{Pagination:b.exports},directives:{waves:p},filters:{statusFilter:function(e){return{published:"success",draft:"info",deleted:"danger"}[e]}},data:function(){return{tableKey:0,list:null,total:0,listLoading:!0,listQuery:{page:1,limit:20,sort:"+id"},importanceOptions:[1,2,3],sortOptions:[{label:"ID Ascending",key:"+id"},{label:"ID Descending",key:"-id"}],tpyeOptions:["password","private key"],showReviewer:!1,temp:{id:void 0,servername:"",serverhost:"",serverusername:"",serverpassword:"",serverport:22,updateTime:new Date,servertype:"password"},dialogFormVisible:!1,dialogStatus:"",textMap:{update:"编辑",create:"添加"},dialogPvVisible:!1,pvData:[],rules:{serverusername:[{required:!0,message:"server username is required",trigger:"blur"}],serverhost:[{required:!0,message:"server host is required",trigger:"blur"}],serverport:[{required:!0,message:"server port is required",trigger:"blur"}]},downloadLoading:!1}},created:function(){this.getList()},methods:{getList:function(){var e=this;this.listLoading=!0,function(e,t,a){var n={page:e,limit:t,sort:a};return Object(o.a)({url:"/getSSHList",method:"post",data:n})}(this.listQuery.page,this.listQuery.limit,this.listQuery.sort).then(function(t){e.list=t.data.items,e.total=t.data.total,setTimeout(function(){e.listLoading=!1},100)})},handleFilter:function(){this.listQuery.page=1,this.getList()},handleStartSSH:function(e){this.$emit("handleStartSSH",e)},sortChange:function(e){console.log(e);var t=e.prop,a=e.order;"id"===t&&this.sortByID(a)},sortByID:function(e){this.listQuery.sort="ascending"===e?"+id":"-id",this.handleFilter()},resetTemp:function(){this.temp={id:void 0,servername:"",serverhost:"",serverusername:"",serverpassword:"",serverport:22,updateTime:new Date,servertype:"password"}},handleCreate:function(){var e=this;this.resetTemp(),this.dialogStatus="create",this.dialogFormVisible=!0,this.$nextTick(function(){e.$refs.dataForm.clearValidate()})},createData:function(){var e=this;this.$refs.dataForm.validate(function(t){t&&function(e){return Object(o.a)({url:"/sshAddMachine",method:"post",data:e})}(e.temp).then(function(){e.dialogFormVisible=!1,setTimeout(function(){e.getList()},200)})})},handleUpdate:function(e){var t=this,a=r()({},e);this.temp.id=a.id,this.temp.servername=a.name,this.temp.serverhost=a.host,this.temp.serverusername=a.user,this.temp.serverpassword=a.password,this.temp.serverport=a.port,this.temp.servertype=a.type,this.dialogStatus="update",this.dialogFormVisible=!0,this.$nextTick(function(){t.$refs.dataForm.clearValidate()})},updateData:function(){var e=this;this.$refs.dataForm.validate(function(t){t&&function(e){return Object(o.a)({url:"/sshUpdateMachine",method:"post",data:e})}(e.temp).then(function(){e.dialogFormVisible=!1,setTimeout(function(){e.getList()},200)})})},handleDelete:function(e){var t=this,a=r()({},e);this.temp.id=a.id,function(e){return Object(o.a)({url:"/sshDeleMachine",method:"post",data:e})}(this.temp).then(function(){setTimeout(function(){t.getList()},200)})}}},w=Object(v.a)(g,function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"app-container"},[a("div",{staticClass:"filter-container"},[a("el-button",{staticClass:"filter-item",staticStyle:{"margin-left":"10px"},attrs:{type:"primary",icon:"el-icon-edit"},on:{click:e.handleCreate}},[e._v(e._s(e.$t("table.add")))])],1),e._v(" "),a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],key:e.tableKey,staticStyle:{width:"100%"},attrs:{data:e.list,border:"",fit:"","highlight-current-row":""},on:{"sort-change":e.sortChange}},[a("el-table-column",{attrs:{prop:"id",label:"id",sortable:"custom",align:"center",width:"65"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.id))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"服务器名称",width:"180px",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.name))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"IP","min-width":"150px"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.host))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"端口",width:"80px",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.port))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"用户名",width:"150px",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.user))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"类型",width:"150px"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.type))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"更新时间",align:"center",width:"150"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(e._f("parseTime")(t.row.updatedAt,"{y}-{m}-{d} {h}:{i}")))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:e.$t("table.actions"),align:"center",width:"250","class-name":"small-padding fixed-width"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-button",{attrs:{type:"primary",size:"mini"},on:{click:function(a){e.handleUpdate(t.row)}}},[e._v(e._s(e.$t("table.edit")))]),e._v(" "),"published"!=t.row.status?a("el-button",{attrs:{size:"mini",type:"success"},on:{click:function(a){e.handleStartSSH(t.row)}}},[e._v("\n          "+e._s(e.$t("table.publish"))+"\n        ")]):e._e(),e._v(" "),"deleted"!=t.row.status?a("el-button",{attrs:{size:"mini",type:"danger"},on:{click:function(a){e.handleDelete(t.row)}}},[e._v("\n          "+e._s(e.$t("table.delete"))+"\n        ")]):e._e()]}}])})],1),e._v(" "),a("pagination",{directives:[{name:"show",rawName:"v-show",value:e.total>0,expression:"total>0"}],attrs:{total:e.total,page:e.listQuery.page,limit:e.listQuery.limit},on:{"update:page":function(t){e.$set(e.listQuery,"page",t)},"update:limit":function(t){e.$set(e.listQuery,"limit",t)},pagination:e.getList}}),e._v(" "),a("el-dialog",{attrs:{title:e.textMap[e.dialogStatus],visible:e.dialogFormVisible},on:{"update:visible":function(t){e.dialogFormVisible=t}}},[a("el-form",{ref:"dataForm",staticStyle:{width:"400px","margin-left":"50px"},attrs:{rules:e.rules,model:e.temp,"label-position":"left","label-width":"120px"}},[a("el-form-item",{attrs:{label:"服务器别名"}},[a("el-input",{model:{value:e.temp.servername,callback:function(t){e.$set(e.temp,"servername",t)},expression:"temp.servername"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"服务器 host",prop:"serverhost"}},[a("el-input",{model:{value:e.temp.serverhost,callback:function(t){e.$set(e.temp,"serverhost",t)},expression:"temp.serverhost"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"服务器端口",prop:"serverport"}},[a("el-input",{model:{value:e.temp.serverport,callback:function(t){e.$set(e.temp,"serverport",e._n(t))},expression:"temp.serverport"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"服务器用户名",prop:"serverusername"}},[a("el-input",{model:{value:e.temp.serverusername,callback:function(t){e.$set(e.temp,"serverusername",t)},expression:"temp.serverusername"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"服务器密码"}},[a("el-input",{model:{value:e.temp.serverpassword,callback:function(t){e.$set(e.temp,"serverpassword",t)},expression:"temp.serverpassword"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"权限类型"}},[a("el-select",{staticClass:"filter-item",attrs:{placeholder:"Please select"},model:{value:e.temp.servertype,callback:function(t){e.$set(e.temp,"servertype",t)},expression:"temp.servertype"}},e._l(e.tpyeOptions,function(e){return a("el-option",{key:e,attrs:{label:e,value:e}})}))],1)],1),e._v(" "),a("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(t){e.dialogFormVisible=!1}}},[e._v(e._s(e.$t("table.cancel")))]),e._v(" "),a("el-button",{attrs:{type:"primary"},on:{click:function(t){"create"===e.dialogStatus?e.createData():e.updateData()}}},[e._v(e._s(e.$t("table.confirm")))])],1)],1)],1)},[],!1,null,null,null);w.options.__file="complexTable.vue";var y=w.exports,_=a("Wb/s"),S=a("xVS5"),k=a("J66h"),T=a("ldua"),x=a("akkt"),$=(a("ie3l"),a("k2hz"),a("X4fA")),L={foreground:"#ffffff",background:"#1b212f",cursor:"#ffffff",selection:"rgba(255, 255, 255, 0.3)",black:"#000000",brightBlack:"#808080",red:"#ce2f2b",brightRed:"#f44a47",green:"#00b976",brightGreen:"#05d289",yellow:"#e0d500",brightYellow:"#f4f628",magenta:"#bd37bc",brightMagenta:"#d86cd8",blue:"#1d6fca",brightBlue:"#358bed",cyan:"#00a8cf",brightCyan:"#19b8dd",white:"#e5e5e5",brightWhite:"#ffffff"},z={name:"SSH",components:{ComplexTable:y},data:function(){return{editableTabsValue:"SSH",editableTabs:[],tabIndex:0,ws:null,term:null}},methods:{handleStartSSH:function(e){var t=this,a=++this.tabIndex+""+e.host;this.editableTabs.push({title:e.name+":"+e.host,name:a}),this.editableTabsValue=a,this.$nextTick(function(){_.Terminal.applyAddon(S),_.Terminal.applyAddon(T),_.Terminal.applyAddon(x),t.term=new _.Terminal({rows:30,fontSize:18,cursorBlink:!0,cursorStyle:"bar",bellStyle:"sound",theme:L}),t.term.open(document.getElementById("terminal")),t.term.webLinksInit(t.doLink),window.addEventListener("resize",t.onWindowResize),t.term.fit();var a=Object($.a)(),n="ws://127.0.0.1:8023/api/newSsh/"+(e.id||0)+"?cols="+t.term.cols+"&rows="+t.term.rows+"&Token="+a;t.ws=new WebSocket(n),t.ws.onerror=function(){t.$message.error("服务器连接错误，请检测配置！")},t.ws.onclose=function(){t.term.setOption("cursorBlink",!1),t.$message("服务器连接断开！")},function(e,t,a,n){e.socket=t;var s=null,r=function(t){n&&n>0?s?s+=t.data:(s=t.data,setTimeout(function(){e.write(s)},n)):e.write(t.data)},o=function(e){t.send(i()({type:"cmd",cmd:k.Base64.encode(e)}))};t.onmessage=r,a&&e.on("data",o);var l=setInterval(function(){t.send(i()({type:"heartbeat",data:""}))},2e4);t.addEventListener("close",function(){t.removeEventListener("message",r),e.off("data",o),delete e.socket,clearInterval(l)})}(t.term,t.ws,!0,-1),function(e,t){var a=function(e){t.send(i()({type:"resize",rows:e.rows,cols:e.cols}))};e.on("resize",a),t.addEventListener("close",function(){e.off("resize",a)})}(t.term,t.ws)})},removeTab:function(e){window.removeEventListener("resize",this.onWindowResize),this.ws&&this.ws.close(),this.term&&this.term.dispose();var t=this.editableTabs,a=this.editableTabsValue;a===e&&t.forEach(function(n,i){if(n.name===e){var s=t[i+1]||t[i-1];a=s?s.name:"SSH"}}),this.editableTabsValue=a,this.editableTabs=t.filter(function(t){return t.name!==e})},onWindowResize:function(){this.term.fit()},doLink:function(e,t){"click"===e.type&&window.open(t)}}},C=Object(v.a)(z,function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"ssh-container"},[a("el-tabs",{attrs:{type:"card"},on:{"tab-remove":e.removeTab},model:{value:e.editableTabsValue,callback:function(t){e.editableTabsValue=t},expression:"editableTabsValue"}},[a("el-tab-pane",{attrs:{label:"SSH 列表",name:"SSH"}},[a("keep-alive",[a("complex-table",{on:{handleStartSSH:e.handleStartSSH}})],1)],1),e._v(" "),e._l(e.editableTabs,function(e){return a("el-tab-pane",{key:e.name,attrs:{label:e.title,name:e.name,closable:""}},[a("div",{attrs:{id:"terminal"}})])})],2)],1)},[],!1,null,null,null);C.options.__file="sshTable.vue";t.default=C.exports},a5W9:function(e,t,a){"use strict";var n=a("8PlU");a.n(n).a},jUE0:function(e,t,a){}}]);