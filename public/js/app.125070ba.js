(function(t){function e(e){for(var n,i,o=e[0],l=e[1],c=e[2],d=0,p=[];d<o.length;d++)i=o[d],Object.prototype.hasOwnProperty.call(r,i)&&r[i]&&p.push(r[i][0]),r[i]=0;for(n in l)Object.prototype.hasOwnProperty.call(l,n)&&(t[n]=l[n]);u&&u(e);while(p.length)p.shift()();return s.push.apply(s,c||[]),a()}function a(){for(var t,e=0;e<s.length;e++){for(var a=s[e],n=!0,o=1;o<a.length;o++){var l=a[o];0!==r[l]&&(n=!1)}n&&(s.splice(e--,1),t=i(i.s=a[0]))}return t}var n={},r={app:0},s=[];function i(e){if(n[e])return n[e].exports;var a=n[e]={i:e,l:!1,exports:{}};return t[e].call(a.exports,a,a.exports,i),a.l=!0,a.exports}i.m=t,i.c=n,i.d=function(t,e,a){i.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:a})},i.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},i.t=function(t,e){if(1&e&&(t=i(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var a=Object.create(null);if(i.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var n in t)i.d(a,n,function(e){return t[e]}.bind(null,n));return a},i.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return i.d(e,"a",e),e},i.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},i.p="/";var o=window["webpackJsonp"]=window["webpackJsonp"]||[],l=o.push.bind(o);o.push=e,o=o.slice();for(var c=0;c<o.length;c++)e(o[c]);var u=l;s.push([0,"chunk-vendors"]),a()})({0:function(t,e,a){t.exports=a("56d7")},"034f":function(t,e,a){"use strict";var n=a("64a9"),r=a.n(n);r.a},"11a4":function(t,e,a){"use strict";var n=a("3978"),r=a.n(n);r.a},2662:function(t,e,a){"use strict";var n=a("8016"),r=a.n(n);r.a},3978:function(t,e,a){},"56d7":function(t,e,a){"use strict";a.r(e);a("cadf"),a("551c"),a("f751"),a("097d");var n=a("2b0e"),r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"app"}},[a("div",{staticClass:"app-wrapper"},[a("div",{staticClass:"app-container"},[a("div",{staticClass:"sidebar"},[a("div",{staticClass:"sidebar-box"},t._l(t.docs,(function(e,n){return a("div",{key:n,staticClass:"sidebar-area"},t._l(e.Content,(function(e,n){return a("div",{key:n,staticClass:"menu-item"},["DELETE"===e["Method"]?a("span",{staticClass:"mark",staticStyle:{"background-color":"#F56C6C"}}):t._e(),"PUT"===e["Method"]?a("span",{staticClass:"mark",staticStyle:{"background-color":"#E6A23C"}}):t._e(),"GET"===e["Method"]?a("span",{staticClass:"mark",staticStyle:{"background-color":"#409EFF"}}):t._e(),"POST"===e["Method"]?a("span",{staticClass:"mark",staticStyle:{"background-color":"#67C23A"}}):t._e(),""===e["Method"]?a("span",{staticClass:"mark",staticStyle:{"background-color":"#303133"}}):t._e(),a("a",{staticClass:"list-item",attrs:{href:"#"+e["Title"]}},[t._v(t._s(e["Title"]))])])})),0)})),0)]),a("div",{ref:"test01",staticClass:"index-content"},t._l(t.docs,(function(e,n){return a("div",{key:n},t._l(e.Content,(function(t,e){return a("Index",{key:e,ref:"item",refInFor:!0,staticStyle:{"margin-top":"8px"},attrs:{data:t}})})),1)})),0)])])])},s=[],i=(a("ac6a"),function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{ref:"card",attrs:{id:"index"}},[a("div",{staticClass:"wrapper"},[a("div",{staticClass:"index-container"},[a("el-card",{staticClass:"box-card",attrs:{id:t.data["Title"]}},[a("div",{staticClass:"clearfix",staticStyle:{"font-size":"20px"},attrs:{slot:"header"},slot:"header"},[a("span",[a("strong",[t._v(t._s(t.data["Title"]))])])]),a("div",{staticClass:"i-con"},[a("p",{staticClass:"i-item"},[a("span",{staticClass:"i-item-1"},[a("h3",[t._v("Url"),a("strong",[t._v(":")])])]),a("span",{staticClass:"i-item-2"},[t._v(t._s(t.data["Url"]))])]),a("p",{staticClass:"i-item"},[a("span",{staticClass:"i-item-1"},[a("h3",[t._v("Header"),a("strong",[t._v(":")])])]),a("span",{staticClass:"i-item-2"},[t._v(t._s(t.data["Header"]))])]),a("p",{staticClass:"i-item"},[a("span",{staticClass:"i-item-1"},[a("h3",[t._v("Method"),a("strong",[t._v(":")])])]),a("span",{staticClass:"i-item-2"},[t._v(t._s(t.data["Method"]))])]),a("el-table",{staticStyle:{width:"100%","font-size":"18px"},attrs:{data:t.data["Params"]}},[a("el-table-column",{attrs:{prop:"name",label:"name"}}),a("el-table-column",{attrs:{prop:"type",label:"type"}}),a("el-table-column",{attrs:{prop:"explain",label:"explain"}}),a("el-table-column",{attrs:{"test-container":"",prop:"remark",label:"remark"}}),a("el-table-column",{attrs:{prop:"other",label:"other"}})],1),a("div",{staticClass:"divider",style:{"background-color":t.showTestAreaState?"#90929833":""},on:{click:function(e){t.showTestAreaState=!t.showTestAreaState}}},[a("i",{staticClass:"el-icon-d-arrow-right"})]),a("div",{staticClass:"test-container",style:{height:t.showTestAreaState?this.$refs.testArea.scrollHeight+"px":""}},[a("div",{ref:"testArea",staticClass:"test-area"},[t._l(t.data["Params"],(function(e,n){return a("div",{key:n,staticClass:"t-area-item"},[a("label",["string"===e["type"]?a("Input",{attrs:{label:e["name"]},model:{value:e["value"],callback:function(a){t.$set(e,"value",a)},expression:"p['value']"}}):t._e()],1),a("label",["bool"===e["type"]?a("Input",{attrs:{label:e["name"]},model:{value:e["value"],callback:function(a){t.$set(e,"value",a)},expression:"p['value']"}}):t._e()],1),a("label",["int"===e["type"]?a("Input",{attrs:{label:e["name"]},model:{value:e["value"],callback:function(a){t.$set(e,"value",a)},expression:"p['value']"}}):t._e()],1),a("label",["float"===e["type"]?a("Input",{attrs:{label:e["name"]},model:{value:e["value"],callback:function(a){t.$set(e,"value",a)},expression:"p['value']"}}):t._e()],1),"file"===e["type"]?a("input",{ref:"uploadFile",refInFor:!0,attrs:{type:"file"},on:{change:t.handleFileChange}}):t._e()])})),a("button",{staticClass:"t-btn",on:{click:function(e){return t.startTest()}}},[t._v("测试")])],2),a("pre",{staticClass:"t-res-area"},[t._v(t._s(t.data["Result"]))])])],1)])],1)])])}),o=[],l=(a("7f7f"),function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"input-wrapper"},[a("div",{staticClass:"input-container"},[a("input",{staticClass:"input-data",attrs:{type:"text",required:""},on:{input:function(e){return t.$emit("input",e.target.value)}}}),a("div",{staticClass:"underline"}),a("label",{staticClass:"input-label"},[t._v(t._s(t.label))])])])}),c=[],u={name:"i-input",model:{prop:"value",event:"input"},props:{value:String,label:String}},d=u,p=(a("2662"),a("2877")),f=Object(p["a"])(d,l,c,!1,null,"2243f04c",null),h=f.exports,v={name:"index",components:{Input:h},props:{data:Object},data:function(){return{file:"",formData:{},showTestAreaState:!1}},mounted:function(){this.data["offsetTop"]=this.$refs.card.offsetTop,this.data["clientHeight"]=this.$refs.card.clientHeight},methods:{showTestArea:function(){console.log(this.$refs.testArea.scrollHeight)},handleFileChange:function(){var t=this.$refs.uploadFile[0];this.file=t.files[0];var e=new FormData;e.append("file",this.file),this.formData=e},startTest:function(){var t=this,e={};this.data["Params"].forEach((function(t){e["".concat(t["name"])]=t["value"]}));var a={url:"http://".concat(this.$doc_server.apiServer,"/").concat(this.data["Url"]),method:this.data["Method"],headers:{"Content-Type":this.data["Header"]},params:e};""!==this.file&&(a["data"]=this.formData),this.$http(a).then((function(e){console.log(e),t.data["Result"]=e.data})).catch((function(e){return t.data["Result"]=e}))}}},m=v,b=(a("11a4"),Object(p["a"])(m,i,o,!1,null,"3dc4bfd4",null)),C=b.exports,g={name:"app",components:{Index:C},data:function(){return{docs:[],sidebar:[]}},created:function(){this.fetchTestData()},mounted:function(){},updated:function(){},methods:{fetchUpdated:function(){var t=this;this.$server.fetchUpdated().then((function(e){console.log(e);for(var a=e.data.updated,n=function(e){var n=a[e];n.Content=t.toDoc(n.Content,null),t.docs.map((function(t){t.Name===n.Name&&(t["Content"]=n.Content)}))},r=0;r<a.length;r++)n(r)}))},toDoc:function(t,e){for(var a=0;a<t.length;a++){var n=t[a];n["Url"]?(n["Method"]=n["Method"].toUpperCase(),n["Params"]=this.handleParams(n["Params"]),n["Result"]="",n["Header"]||(n["Header"]="multipart/form-data"),e&&e(n)):t.splice(a,1)}return t},fetchTestData:function(){var t=this;this.$server.fetchDocs().then((function(e){var a=e.data.Docs;null!==a&&Array.isArray(a)&&(a.forEach((function(e){var a=[];e.Content=t.toDoc(e.Content,null),t.sidebar.push(a)})),t.docs=a),console.log(t.docs),setInterval(t.fetchUpdated,300)}))},handleParams:function(t){var e=[];if(Array.isArray(t))for(var a=0;a<t.length;a++){var n=t[a];e.push({name:n[0]||"",type:n[1]||"",explain:n[2]||"",remark:n[3]||"",other:n[4]||"",value:""})}return e}}},y=g,_=(a("034f"),Object(p["a"])(y,r,s,!1,null,null,null)),x=_.exports,k=a("bc3a"),w=a.n(k),S=a("a7fe"),T=a.n(S),$=a("5c96"),O=a.n($),j=(a("0fae"),a("8058")),P=a("d225"),A=a("b0b4"),M=w.a.create({baseURL:"http://".concat(j.addr,":").concat(j.port)}),D=function(){function t(){Object(P["a"])(this,t)}return Object(A["a"])(t,[{key:"fetchDocs",value:function(){return M({url:"/doc/v1",method:"get",headers:{"Content-Type":"application/x-www-form-urlencode"}})}},{key:"fetchUpdated",value:function(){return M({url:"/doc/v2/updated",method:"get",headers:{"Content-Type":"application/x-www-form-urlencode"}})}}]),t}();function E(t,e){document.cookie="".concat(t,"=").concat(e)}n["default"].prototype.$doc_server=j,n["default"].use(T.a,w.a),n["default"].use(O.a),n["default"].config.productionTip=!1,n["default"].prototype.$server=new D,E("name00","zjh"),E("name01","zjh"),E("name02","zjh"),new n["default"]({render:function(t){return t(x)}}).$mount("#app")},"64a9":function(t,e,a){},8016:function(t,e,a){},8058:function(t){t.exports=JSON.parse('{"addr":"192.168.0.110","apiServer":"192.168.0.110:86","files":{"franchisee":"/home/hangiangai/go/src/JibeiServer/dbaccess/franchisee.go","template":"/home/hangiangai/go/src/JibeiServer/dbaccess/template.go"},"port":"8888"}')}});
//# sourceMappingURL=app.125070ba.js.map