(function(t){function e(e){for(var n,i,o=e[0],c=e[1],l=e[2],p=0,d=[];p<o.length;p++)i=o[p],Object.prototype.hasOwnProperty.call(r,i)&&r[i]&&d.push(r[i][0]),r[i]=0;for(n in c)Object.prototype.hasOwnProperty.call(c,n)&&(t[n]=c[n]);u&&u(e);while(d.length)d.shift()();return s.push.apply(s,l||[]),a()}function a(){for(var t,e=0;e<s.length;e++){for(var a=s[e],n=!0,o=1;o<a.length;o++){var c=a[o];0!==r[c]&&(n=!1)}n&&(s.splice(e--,1),t=i(i.s=a[0]))}return t}var n={},r={app:0},s=[];function i(e){if(n[e])return n[e].exports;var a=n[e]={i:e,l:!1,exports:{}};return t[e].call(a.exports,a,a.exports,i),a.l=!0,a.exports}i.m=t,i.c=n,i.d=function(t,e,a){i.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:a})},i.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},i.t=function(t,e){if(1&e&&(t=i(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var a=Object.create(null);if(i.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var n in t)i.d(a,n,function(e){return t[e]}.bind(null,n));return a},i.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return i.d(e,"a",e),e},i.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},i.p="/";var o=window["webpackJsonp"]=window["webpackJsonp"]||[],c=o.push.bind(o);o.push=e,o=o.slice();for(var l=0;l<o.length;l++)e(o[l]);var u=c;s.push([0,"chunk-vendors"]),a()})({0:function(t,e,a){t.exports=a("56d7")},"034f":function(t,e,a){"use strict";var n=a("64a9"),r=a.n(n);r.a},"56d7":function(t,e,a){"use strict";a.r(e);a("28a5"),a("cadf"),a("551c"),a("f751"),a("097d");var n=a("2b0e"),r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"app"}},[a("div",{staticClass:"app-wrapper"},[a("div",{staticClass:"app-container"},[a("div",{staticClass:"sidebar"},[a("div",{staticClass:"sidebar-box"},t._l(t.testData,(function(e,n){return a("a",{key:n,staticClass:"list-item",attrs:{href:"#"+e["title"]}},[t._v(t._s(e["title"]))])})),0)]),a("div",{staticClass:"index-content"},[a("Index",{attrs:{data:t.testData}})],1)])])])},s=[],i=(a("7f7f"),a("ac6a"),function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"index"}},[a("div",{staticClass:"wrapper"},[a("div",{staticClass:"index-container"},t._l(t.data,(function(e,n){return a("el-card",{key:n,staticClass:"box-card",attrs:{id:e["title"]}},[a("div",{staticClass:"clearfix",staticStyle:{"font-size":"20px"},attrs:{slot:"header"},slot:"header"},[a("span",[a("strong",[t._v(t._s(e["title"]))])])]),a("div",{staticClass:"i-con"},[a("p",{staticClass:"i-item"},[a("span",{staticClass:"i-item-1"},[a("h3",[t._v("Url"),a("strong",[t._v(":")])])]),a("span",{staticClass:"i-item-2"},[t._v(t._s(e["url"]))])]),a("p",{staticClass:"i-item"},[a("span",{staticClass:"i-item-1"},[a("h3",[t._v("Header"),a("strong",[t._v(":")])])]),a("span",{staticClass:"i-item-2"},[t._v(t._s(e["header"]))])]),a("p",{staticClass:"i-item"},[a("span",{staticClass:"i-item-1"},[a("h3",[t._v("Method"),a("strong",[t._v(":")])])]),a("span",{staticClass:"i-item-2"},[t._v(t._s(e["method"]))])]),a("el-table",{staticStyle:{width:"100%","font-size":"18px"},attrs:{data:e["params"]}},[a("el-table-column",{attrs:{prop:"name",label:"name"}}),a("el-table-column",{attrs:{prop:"type",label:"type"}}),a("el-table-column",{attrs:{prop:"explain",label:"explain"}}),a("el-table-column",{attrs:{"test-container":"",prop:"remark",label:"remark"}}),a("el-table-column",{attrs:{prop:"other",label:"other"}})],1),a("div",{staticClass:"test-container"},[a("div",{staticClass:"test-area"},[t._l(e["params"],(function(e,n){return a("div",{key:n,staticClass:"t-area-item"},[a("Input",{attrs:{label:e["name"]},model:{value:e["value"],callback:function(a){t.$set(e,"value",a)},expression:"item['value']"}})],1)})),a("button",{staticClass:"t-btn",on:{click:function(e){return t.startTest(n)}}},[t._v("测试")])],2),a("pre",{staticClass:"t-res-area"},[t._v(t._s(e["result"]))])])],1)])})),1)])])}),o=[],c=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"input-wrapper"},[a("div",{staticClass:"input-container"},[a("input",{staticClass:"input-data",attrs:{type:"text",required:""},on:{input:function(e){return t.$emit("input",e.target.value)}}}),a("div",{staticClass:"underline"}),a("label",{staticClass:"input-label"},[t._v(t._s(t.label))])])])},l=[],u={name:"i-input",model:{prop:"value",event:"input"},props:{value:String,label:String}},p=u,d=(a("b327"),a("2877")),f=Object(d["a"])(p,c,l,!1,null,"f190f03e",null),h=f.exports,v={name:"index",components:{Input:h},props:{data:Array},methods:{startTest:function(t){var e=this,a={};this.data[t]["params"].forEach((function(t){a["".concat(t["name"])]=t["value"]})),this.$http({url:"http://".concat(this.$doc_server.apiServer,"/").concat(this.data[t]["url"]),method:this.data[t]["method"],header:{"Content-Type":"multipart/form-data"},params:a}).then((function(a){e.data[t]["result"]=a.data})).catch((function(t){return console.log(t)}))}}},m=v,b=(a("e37c"),Object(d["a"])(m,i,o,!1,null,"5ba28968",null)),_=b.exports,g={name:"app",components:{Index:_},data:function(){return{testData:[],code:""}},created:function(){var t=this.$createElement;t("div",{attrs:{c:!0}}),this.fetchTestData()},methods:{processParams:function(t){var e=[];return Array.isArray(t)&&t.forEach((function(t){var a=t.split(":");e.push({name:a[0]||"",type:a[1]||"",explain:a[2]||"",remark:a[3]||"",other:a[4]||"",value:""})})),e},filter:function(t){return console.log(t),""!==t},hasKey:function(t){return void 0!==t&&null!==t},fetchTestData:function(){var t=this;console.log(this.$doc_server),this.$http.get("http://".concat(this.$doc_server.addr,":").concat(this.$doc_server.port,"/doc/v1")).then((function(e){var a=e.data.Docs;console.log(a),null!==a&&Array.isArray(a)&&a.forEach((function(e){var a=e["Param"];a.forEach((function(e){var a={title:t.hasKey(e["@title"])?e["@title"][0]:"",url:t.hasKey(e["@url"])?e["@url"][0]:"",header:t.hasKey(e["@header"])?e["@header"][0]:"",method:t.hasKey(e["@method"])?e["@method"][0]:"",params:t.processParams(e["@param"]),result:""};t.filter(a["url"])&&t.testData.push(a)}))})),console.log(t.testData)})).catch((function(t){console.log(t)}))},startTest:function(t){var e=this,a={};this.testData[t]["params"].forEach((function(t){a["".concat(t["name"])]=t["value"]})),this.$http({url:"http://".concat(this.$doc_server.apiServer,"/").concat(this.testData[t]["url"]),method:this.testData[t]["method"],header:{"Content-Type":"multipart/form-data"},params:a}).then((function(a){e.testData[t]["result"]=a.data})).catch((function(t){return console.log(t)}))}}},y=g,C=(a("034f"),Object(d["a"])(y,r,s,!1,null,null,null)),x=C.exports,$=a("bc3a"),j=a.n($),w=a("a7fe"),O=a.n(w),S=a("5c96"),k=a.n(S),D=(a("0fae"),a("8058"));function P(t,e){document.cookie="".concat(t,"=").concat(e)}function T(t){for(var e=document.cookie.split("; "),a=e.length,n="",r=0;r<a;r++){var s=e[r].split("=");if(s[0]===t){n=s[1];break}}return n}n["default"].prototype.$doc_server=D,n["default"].use(O.a,j.a),n["default"].use(k.a),n["default"].config.productionTip=!1,P("name00","zjh"),P("name01","zjh"),P("name02","zjh"),console.log(T("name00")),console.log(T("name01")),console.log(T("name02")),new n["default"]({render:function(t){return t(x)}}).$mount("#app")},"64a9":function(t,e,a){},"68d2":function(t,e,a){},8058:function(t){t.exports=JSON.parse('{"addr":"192.168.0.110","apiServer":"192.168.0.110:86","files":{"franchisee":"/home/hangiangai/go/src/JibeiServer/dbaccess/franchisee.go","template":"/home/hangiangai/go/src/JibeiServer/dbaccess/template.go"},"port":"8888"}')},b327:function(t,e,a){"use strict";var n=a("f1e4"),r=a.n(n);r.a},e37c:function(t,e,a){"use strict";var n=a("68d2"),r=a.n(n);r.a},f1e4:function(t,e,a){}});
//# sourceMappingURL=app.a5766ab2.js.map