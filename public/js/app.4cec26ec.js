(function(t){function e(e){for(var s,i,l=e[0],o=e[1],c=e[2],p=0,d=[];p<l.length;p++)i=l[p],Object.prototype.hasOwnProperty.call(r,i)&&r[i]&&d.push(r[i][0]),r[i]=0;for(s in o)Object.prototype.hasOwnProperty.call(o,s)&&(t[s]=o[s]);u&&u(e);while(d.length)d.shift()();return n.push.apply(n,c||[]),a()}function a(){for(var t,e=0;e<n.length;e++){for(var a=n[e],s=!0,l=1;l<a.length;l++){var o=a[l];0!==r[o]&&(s=!1)}s&&(n.splice(e--,1),t=i(i.s=a[0]))}return t}var s={},r={app:0},n=[];function i(e){if(s[e])return s[e].exports;var a=s[e]={i:e,l:!1,exports:{}};return t[e].call(a.exports,a,a.exports,i),a.l=!0,a.exports}i.m=t,i.c=s,i.d=function(t,e,a){i.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:a})},i.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},i.t=function(t,e){if(1&e&&(t=i(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var a=Object.create(null);if(i.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var s in t)i.d(a,s,function(e){return t[e]}.bind(null,s));return a},i.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return i.d(e,"a",e),e},i.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},i.p="/";var l=window["webpackJsonp"]=window["webpackJsonp"]||[],o=l.push.bind(l);l.push=e,l=l.slice();for(var c=0;c<l.length;c++)e(l[c]);var u=o;n.push([0,"chunk-vendors"]),a()})({0:function(t,e,a){t.exports=a("56d7")},"034f":function(t,e,a){"use strict";var s=a("64a9"),r=a.n(s);r.a},"064e":function(t,e,a){"use strict";var s=a("545c"),r=a.n(s);r.a},"545c":function(t,e,a){},"56d7":function(t,e,a){"use strict";a.r(e);a("cadf"),a("551c"),a("f751"),a("097d");var s=a("2b0e"),r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"app"}},[a("div",{staticClass:"app-wrapper"},[a("div",{staticClass:"app-container"},[a("div",{staticClass:"sidebar"},[a("div",{staticClass:"sidebar-box"},t._l(t.docs,(function(e,s){return a("div",{key:s,staticClass:"menu-item"},["DELETE"===e["method"]?a("span",{staticClass:"mark",staticStyle:{"background-color":"#F56C6C"}}):t._e(),"PUT"===e["method"]?a("span",{staticClass:"mark",staticStyle:{"background-color":"#E6A23C"}}):t._e(),"GET"===e["method"]?a("span",{staticClass:"mark",staticStyle:{"background-color":"#409EFF"}}):t._e(),"POST"===e["method"]?a("span",{staticClass:"mark",staticStyle:{"background-color":"#67C23A"}}):t._e(),""===e["method"]?a("span",{staticClass:"mark",staticStyle:{"background-color":"#303133"}}):t._e(),a("a",{staticClass:"list-item",attrs:{href:"#"+e["title"]}},[t._v(t._s(e["title"]))])])})),0)]),a("div",{staticClass:"index-content"},t._l(t.docs,(function(t,e){return a("Index",{key:e,staticStyle:{"margin-top":"8px"},attrs:{data:t}})})),1)])])])},n=[],i=(a("28a5"),a("ac6a"),function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"index"}},[a("div",{staticClass:"wrapper"},[a("div",{staticClass:"index-container"},[a("el-card",{staticClass:"box-card",attrs:{id:t.data["title"]}},[a("div",{staticClass:"clearfix",staticStyle:{"font-size":"20px"},attrs:{slot:"header"},slot:"header"},[a("span",[a("strong",[t._v(t._s(t.data["title"]))])])]),a("div",{staticClass:"i-con"},[a("p",{staticClass:"i-item"},[a("span",{staticClass:"i-item-1"},[a("h3",[t._v("Url"),a("strong",[t._v(":")])])]),a("span",{staticClass:"i-item-2"},[t._v(t._s(t.data["url"]))])]),a("p",{staticClass:"i-item"},[a("span",{staticClass:"i-item-1"},[a("h3",[t._v("Header"),a("strong",[t._v(":")])])]),a("span",{staticClass:"i-item-2"},[t._v(t._s(t.data["header"]))])]),a("p",{staticClass:"i-item"},[a("span",{staticClass:"i-item-1"},[a("h3",[t._v("Method"),a("strong",[t._v(":")])])]),a("span",{staticClass:"i-item-2"},[t._v(t._s(t.data["method"]))])]),a("el-table",{staticStyle:{width:"100%","font-size":"18px"},attrs:{data:t.data["params"]}},[a("el-table-column",{attrs:{prop:"name",label:"name"}}),a("el-table-column",{attrs:{prop:"type",label:"type"}}),a("el-table-column",{attrs:{prop:"explain",label:"explain"}}),a("el-table-column",{attrs:{"test-container":"",prop:"remark",label:"remark"}}),a("el-table-column",{attrs:{prop:"other",label:"other"}})],1),a("div",{staticClass:"divider",style:{"background-color":t.showTestAreaState?"#90929833":""},on:{click:function(e){t.showTestAreaState=!t.showTestAreaState}}},[a("i",{staticClass:"el-icon-d-arrow-right"})]),a("div",{staticClass:"test-container",style:{height:t.showTestAreaState?this.$refs.testArea.scrollHeight+"px":""}},[a("div",{ref:"testArea",staticClass:"test-area"},[t._l(t.data["params"],(function(e,s){return a("div",{key:s,staticClass:"t-area-item"},["string"===e["type"]?a("Input",{attrs:{label:e["name"]},model:{value:e["value"],callback:function(a){t.$set(e,"value",a)},expression:"p['value']"}}):t._e(),"bool"===e["type"]?a("Input",{attrs:{label:e["name"]},model:{value:e["value"],callback:function(a){t.$set(e,"value",a)},expression:"p['value']"}}):t._e(),"int"===e["type"]?a("Input",{attrs:{label:e["name"]},model:{value:e["value"],callback:function(a){t.$set(e,"value",a)},expression:"p['value']"}}):t._e(),"float"===e["type"]?a("Input",{attrs:{label:e["name"]},model:{value:e["value"],callback:function(a){t.$set(e,"value",a)},expression:"p['value']"}}):t._e(),"file"===e["type"]?a("input",{ref:"uploadFile",refInFor:!0,attrs:{type:"file"},on:{change:t.handleFileChange}}):t._e()],1)})),a("button",{staticClass:"t-btn",on:{click:function(e){return t.startTest()}}},[t._v("测试")])],2),a("pre",{staticClass:"t-res-area"},[t._v(t._s(t.data["result"]))])])],1)])],1)])])}),l=[],o=(a("7f7f"),function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"input-wrapper"},[a("div",{staticClass:"input-container"},[a("input",{staticClass:"input-data",attrs:{type:"text",required:""},on:{input:function(e){return t.$emit("input",e.target.value)}}}),a("div",{staticClass:"underline"}),a("label",{staticClass:"input-label"},[t._v(t._s(t.label))])])])}),c=[],u={name:"i-input",model:{prop:"value",event:"input"},props:{value:String,label:String}},p=u,d=(a("b327"),a("2877")),f=Object(d["a"])(p,o,c,!1,null,"f190f03e",null),h=f.exports,v={name:"index",components:{Input:h},props:{data:Object},data:function(){return{file:"",formData:{},showTestAreaState:!1}},methods:{showTestArea:function(){console.log(this.$refs.testArea.scrollHeight)},handleFileChange:function(){var t=this.$refs.uploadFile[0];this.file=t.files[0];var e=new FormData;e.append("file",this.file),this.formData=e},startTest:function(){var t=this,e={};this.data["params"].forEach((function(t){e["".concat(t["name"])]=t["value"]}));var a={url:"http://".concat(this.$doc_server.apiServer,"/").concat(this.data["url"]),method:this.data["method"],headers:{"Content-Type":this.data["header"]},params:e};""!==this.file&&(a["data"]=this.formData),this.$http(a).then((function(e){console.log(e),t.data["result"]=e.data})).catch((function(e){return t.data["result"]=e}))}}},m=v,b=(a("064e"),Object(d["a"])(m,i,l,!1,null,"3e4b67c8",null)),y=b.exports,_={name:"app",components:{Index:y},data:function(){return{testData:[],code:"",docs:[]}},created:function(){var t=this.$createElement;t("div",{attrs:{c:!0}}),this.fetchTestData()},methods:{processParams:function(t){var e=[];return Array.isArray(t)&&t.forEach((function(t){var a=t.split(":");e.push({name:a[0]||"",type:a[1]||"",explain:a[2]||"",remark:a[3]||"",other:a[4]||"",value:""})})),e},filter:function(t){return""!==t},hasKey:function(t){return!!t},fetchTestData:function(){var t=this;this.$http.get("http://".concat(this.$doc_server.addr,":").concat(this.$doc_server.port,"/doc/v1")).then((function(e){var a=e.data.Docs;null!==a&&Array.isArray(a)&&a.forEach((function(e){var a=e["Param"];a.forEach((function(e){var a={title:t.hasKey(e["@title"])?e["@title"][0]:"",url:t.hasKey(e["@url"])?e["@url"][0]:"",header:t.hasKey(e["@header"])?e["@header"][0]:"multipart/form-data",method:t.hasKey(e["@method"])?e["@method"][0].toUpperCase():"",params:t.processParams(e["@param"]),result:""};a["header"]||(a["header"]="multipart/form-data"),t.filter(a["url"])&&t.docs.push(a)}))}))})).catch((function(t){console.log(t)}))}}},g=_,C=(a("034f"),Object(d["a"])(g,r,n,!1,null,null,null)),x=C.exports,k=a("bc3a"),S=a.n(k),w=a("a7fe"),$=a.n(w),T=a("5c96"),O=a.n(T),j=(a("0fae"),a("8058"));function A(t,e){document.cookie="".concat(t,"=").concat(e)}s["default"].prototype.$doc_server=j,s["default"].use($.a,S.a),s["default"].use(O.a),s["default"].config.productionTip=!1,A("name00","zjh"),A("name01","zjh"),A("name02","zjh"),new s["default"]({render:function(t){return t(x)}}).$mount("#app")},"64a9":function(t,e,a){},8058:function(t){t.exports=JSON.parse('{"addr":"192.168.0.110","apiServer":"192.168.0.110:86","files":{"franchisee":"/home/hangiangai/go/src/JibeiServer/dbaccess/franchisee.go","template":"/home/hangiangai/go/src/JibeiServer/dbaccess/template.go"},"port":"8888"}')},b327:function(t,e,a){"use strict";var s=a("f1e4"),r=a.n(s);r.a},f1e4:function(t,e,a){}});
//# sourceMappingURL=app.4cec26ec.js.map