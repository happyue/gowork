(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-6712"],{MSNs:function(e,t,i){"use strict";i.d(t,"d",function(){return n}),i.d(t,"c",function(){return o}),i.d(t,"a",function(){return d}),i.d(t,"b",function(){return u});var a=i("t3Un");function n(){return Object(a.a)({url:"/qiniu/upload/token",method:"get"})}function o(){return Object(a.a)({url:"/qiniu/getMediaInfo",method:"get"})}function d(e,t){var i={uid:e,name:t};return Object(a.a)({url:"/qiniu/addMediaInfo",method:"post",data:i})}function u(e,t){var i={uid:e,name:t};return Object(a.a)({url:"/qiniu/deleMediaInfo",method:"post",data:i})}},XPP3:function(e,t,i){"use strict";i.r(t);var a=i("4d7F"),n=i.n(a),o=i("MSNs"),d={data:function(){return{dataObj:{token:"",key:""},dialogImageUrl:"",dialogImageName:"",dialogVisible:!1,image_uri:[],fileList:[]}},created:function(){this.getMedia()},methods:{getMedia:function(){var e=this;Object(o.c)().then(function(t){e.fileList=t.data})},addMedia:function(e){var t=this;Object(o.a)(e.uid,e.uid+"_"+e.name),setTimeout(function(){t.getMedia()},300)},deleMedia:function(e){var t=this;Object(o.b)(e.uid,e.name),setTimeout(function(){t.getMedia()},200)},beforeUpload:function(e){var t=this;return new n.a(function(i,a){Object(o.d)().then(function(a){var n=e.uid+"_"+e.name,o=a.data;t._data.dataObj.token=o,t._data.dataObj.key=n,i(!0)}).catch(function(e){console.log(e),a(!1)})})},handleSuccess:function(e,t,i){this.addMedia(t)},handleRemove:function(e,t){this.deleMedia(e)},handlePreview:function(e){this.dialogImageUrl=e.url,this.dialogImageName=e.name,this.dialogVisible=!0},beforeRemove:function(e,t){return this.$confirm("确定移除 "+e.name+"？")}}},u=i("KHd+"),r=Object(u.a)(d,function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("el-card",{staticStyle:{"margin-top":"20px","margin-left":"20px","margin-right":"20px"}},[i("el-upload",{attrs:{data:e.dataObj,multiple:!0,"on-preview":e.handlePreview,"on-remove":e.handleRemove,"on-success":e.handleSuccess,"before-upload":e.beforeUpload,"before-remove":e.beforeRemove,"file-list":e.fileList,"list-type":"picture-card",action:"http://upload.qiniup.com/"}},[i("i",{staticClass:"el-icon-plus"})]),e._v(" "),i("el-dialog",{attrs:{visible:e.dialogVisible},on:{"update:visible":function(t){e.dialogVisible=t}}},[i("img",{attrs:{width:"100%",src:e.dialogImageUrl,alt:""}}),e._v(" "),i("code",[e._v("Name: "+e._s(e.dialogImageName)),i("br"),e._v("URL: "+e._s(e.dialogImageUrl))])])],1)},[],!1,null,null,null);r.options.__file="upload.vue";t.default=r.exports}}]);