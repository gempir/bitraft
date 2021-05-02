(this.webpackJsonpweb=this.webpackJsonpweb||[]).push([[0],{37:function(n,e,t){"use strict";t.r(e);var r,i,a,o=t(4),c=t(0),s=t.n(c),d=t(21),b=t.n(d),l=t(5),h=t(6),u=t(1),g=l.b.div(r||(r=Object(o.a)(["\n    display: inline-block;\n    background: var(--bg-bright);\n    border: 1px solid var(--bg-brighter);\n    padding: 0.5rem;\n    margin: 1rem;\n    margin-right: 0;\n    margin-bottom: 0;\n"])));function j(n){var e=n.activeChannels,t=n.joinedChannels;return Object(u.jsxs)(u.Fragment,{children:[Object(u.jsxs)(g,{children:["Joined Channels: ",Object(u.jsx)("strong",{children:t})]}),Object(u.jsxs)(g,{children:["Active Channels: ",Object(u.jsx)("strong",{children:e})]})]})}function p(){var n=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{src:"https://static-cdn.jtvnw.net/user-default-pictures-uv/cdd517fe-def4-11e9-948e-784f43822e80-profile_image-70x70.png"},e=n.src;return Object(u.jsx)("img",{src:e,width:"30",height:"30",alt:"profile"})}var m=l.b.div(i||(i=Object(o.a)(["\n    display: flex;\n"])));function x(n){var e=n.records;return Object(u.jsx)(m,{children:e.map((function(n){return Object(u.jsx)(O,{record:n},n.title)}))})}var f=l.b.div(a||(a=Object(o.a)(["\n    background: var(--bg-bright);\n    border: 1px solid var(--bg-brighter);\n    margin: 1rem;\n    margin-right: 0;\n    padding: 1rem;\n\n    h2 {\n        color: var(--text);\n        margin: 0;\n        margin-bottom: 1rem;\n        padding: 0;\n    }\n\n    ol {\n        color: white;\n        font-size: 1.5rem;\n        font-weight: bold;\n        width: 500px;\n        padding: 0;\n        margin: 0;\n        margin-right: 15px;\n        background: var(--lightBackground);\n        border: 1px solid var(--lightBorder);\n        border-radius: 3px;\n\n        li {\n            display: flex;\n            align-items: center;\n            margin-bottom: 0.25rem;\n\n            img {\n                margin-right: 10px;\n            }\n\n            .value {\n                text-align: right;\n                flex: 1 1 auto;\n            }\n        }\n    }\n"])));function O(n){var e=n.record;return Object(u.jsxs)(f,{children:[Object(u.jsx)("h2",{children:e.title}),Object(u.jsx)("ol",{children:e.scores.map((function(n){return Object(u.jsxs)("li",{children:[Object(u.jsx)(p,{src:n.user.profilePicture}),Object(u.jsx)("span",{children:n.user.displayName}),Object(u.jsx)("span",{className:"value",children:n.score})]},n.user.id)}))})]})}var v,k,w,S,C=t(14),y=t(2),T=t(25),N=function n(e,t){Object(T.a)(this,n),function n(){var r=new WebSocket("".concat(e.replace("https://","wss://").replace("http://","ws://"),"/api/ws"));r.onmessage=function(n){t(JSON.parse(n.data))},r.onclose=function(e){console.log("Socket is closed. Reconnect will be attempted in 1 second.",e.reason),setTimeout(n,1e3)},r.onerror=function(n){console.error("Socket encountered error: ",n,"Closing socket"),r.close()}}()},B=t(15),U={state:{apiBaseUrl:"https://spamchamp-api.gempir.com",twitchClientId:"il5oqz0li81219qvw2hs7umntr493e",baseUrl:"https://spamchamp.gempir.com",scToken:window.localStorage.getItem("scToken")},setState:function(n){},setScToken:function(n){}},J=Object(c.createContext)(U),P=J.Provider,z=function(n){var e=n.children,t=Object(c.useState)(Object(B.a)({},U.state)),r=Object(h.a)(t,2),i=r[0],a=r[1];return Object(u.jsx)(P,{value:{state:i,setState:a,setScToken:function(n){null===n?window.localStorage.removeItem("scToken"):window.localStorage.setItem("scToken",n),a(Object(B.a)(Object(B.a)({},i),{},{scToken:n}))}},children:e})},I=l.b.a(v||(v=Object(o.a)(["\n    position: absolute;\n    display: block;\n    color: white;\n    top: 1rem;\n    right: 1rem;\n    padding: 1rem 2rem;\n    text-decoration: none;\n    font-weight: bold;\n    border-radius: 3px;\n    background: var(--twitch);\n\n    &:hover {\n        background: var(--twitch-dark);\n    }\n"])));function R(){var n=Object(c.useContext)(J).state;return n.scToken?Object(u.jsx)(F,{}):Object(u.jsx)(I,{href:_(n.twitchClientId,n.apiBaseUrl).toString(),children:"Login"})}var E=l.b.div(k||(k=Object(o.a)(["\n    position: absolute;\n    top: 1rem;\n    right: 1rem;\n    display: flex;\n\n    a {\n        text-decoration: none;\n    }\n"]))),A=l.b.a(w||(w=Object(o.a)(["\n    display: block;\n    color: white;\n    width: 150px;\n    padding: 1rem 2rem;\n    text-decoration: none;\n    font-weight: bold;\n    border-radius: 3px;\n    background: var(--twitch);\n    cursor: pointer;\n\n    &:hover {\n        background: var(--twitch-dark);\n    }\n"]))),L=l.b.div(S||(S=Object(o.a)(["\n    display: block;\n    color: white;\n    margin-right: 1rem;\n    padding: 1rem 2rem;\n    text-decoration: none;\n    font-weight: bold;\n    border-radius: 3px;\n    background: var(--theme2-dark);\n    cursor: pointer;\n\n    &.dashboard {\n        background: var(--theme);\n\n        &:hover {\n        background: var(--theme-bright);\n    }\n    }\n\n    &:hover {\n        background: var(--theme2);\n    }\n"])));function F(){var n=Object(c.useContext)(J).state,e=Object(c.useState)("Logged In"),t=Object(h.a)(e,2),r=t[0],i=t[1];return Object(u.jsxs)(E,{children:[Object(u.jsx)(C.b,{to:"/",children:Object(u.jsx)(L,{children:"Home"})}),Object(u.jsx)(C.b,{to:"/dashboard",children:Object(u.jsx)(L,{className:"dashboard",children:"Dashboard"})}),Object(u.jsx)(A,{href:_(n.twitchClientId,n.apiBaseUrl).toString(),onMouseEnter:function(){return i("force login")},onMouseLeave:function(){return i("Logged In")},children:r})]})}function _(n,e){var t=new URL("https://id.twitch.tv/oauth2/authorize");return t.searchParams.set("client_id",n),t.searchParams.set("redirect_uri",e+"/api/callback"),t.searchParams.set("response_type","code"),t.searchParams.set("claims",JSON.stringify({userinfo:{picture:null,preferred_username:null}})),t.searchParams.set("scope","channel:read:redemptions channel:manage:redemptions"),t}var M,X=t(39);function q(n,e){return 403===e.status&&n(null),e}var D,H=l.b.div(M||(M=Object(o.a)(["\n    position: absolute;\n    top: 1rem;\n    left: 1rem;\n    display: block;\n    color: white;\n    padding: 1rem 2rem;\n    text-decoration: none;\n    font-weight: bold;\n    border-radius: 3px;\n    background: var(--danger-dark);\n    cursor: pointer;\n\n    &:hover {\n        background: var(--danger);\n    }\n"])));function W(n){var e=n.setUserConfig;return Object(u.jsx)(H,{onClick:function(){return e(null)},children:"Reset"})}function G(){var n=Object(c.useState)(["userConfig"]),e=Object(h.a)(n,2),t=e[0],r=e[1],i=function(n){var e=Object(c.useContext)(J).state,t=e.scToken,r=e.apiBaseUrl,i=Object(c.useContext)(J).setScToken,a=Object(c.useState)(void 0),o=Object(h.a)(a,2),s=o[0],d=o[1],b=Object(c.useState)(0),l=Object(h.a)(b,2),u=l[0],g=l[1],j=function(){t&&fetch(r+"/api/userConfig",{headers:{Authorization:"Bearer "+t}}).then((function(n){return q(i,n)})).then((function(n){return n.json()})).then((function(n){return d(n)})).catch()};return Object(c.useEffect)(j,[t,r,i]),Object(X.a)((function(){u&&s&&t?fetch(r+"/api/userConfig",{headers:{Authorization:"Bearer "+t},method:"POST",body:JSON.stringify(s)}).then((function(n){return q(i,n)})).then(n):u&&null===s&&t&&fetch(r+"/api/userConfig",{headers:{Authorization:"Bearer "+t},method:"DELETE"}).then((function(n){return q(i,n)})).then(j)}),500,[u]),[s,function(n){d(n),g(u+1)}]}((function(){var n=t.slice();n.push("saved"),r(n),setTimeout((function(){r(["userConfig"])}),500)})),a=Object(h.a)(i,2),o=a[0],s=a[1];return Object(u.jsxs)(V,{children:[o&&Object(u.jsx)(W,{setUserConfig:s}),o&&Object(u.jsx)("div",{className:t.join(" "),children:Object(u.jsxs)("div",{className:"redemption",children:[Object(u.jsx)("img",{src:"/images/bttv.png",alt:"bttv"}),Object(u.jsxs)("label",{className:"switch",children:[Object(u.jsx)("input",{type:"checkbox",checked:o.Redemptions.Bttv.Active,onChange:function(n){var e=JSON.parse(JSON.stringify(o));e.Redemptions.Bttv.Active=n.target.checked,s(e)}}),Object(u.jsx)("span",{className:"slider round"})]}),Object(u.jsxs)("div",{className:"redemption-title",children:[Object(u.jsx)("span",{children:"Channel Points Reward Name"}),Object(u.jsx)("input",{type:"text",value:o.Redemptions.Bttv.Title,spellCheck:!1,onChange:function(n){var e=JSON.parse(JSON.stringify(o));e.Redemptions.Bttv.Title=n.target.value,s(e)}})]}),Object(u.jsxs)("span",{className:"hint",children:["make sure ",Object(u.jsx)("strong",{children:"gempbot"})," is bttv editor"]})]})})]})}var K,Q,V=l.b.div(D||(D=Object(o.a)(['\n    margin-top: 5rem;\n    margin-left: 1rem;\n    margin-right: 1rem;\n\n    .userConfig {\n        padding-bottom: 2rem;\n        background: var(--bg);\n        transition: background-color ease-in-out 0.2s;\n\n        &.saved {\n            background: var(--theme);\n        }\n    }\n\n    .redemption {\n        display: flex;\n        align-items: center;\n        background: var(--bg-bright);\n        border: 1px solid var(--bg-brighter);\n        padding: 0.5rem;\n\n        img {\n            max-height: 3rem;\n            margin-left: 1rem;\n            margin-right: 2rem;\n        }\n\n        .redemption-title {\n            position: relative;\n\n            span {\n                position: absolute;\n                top: -13px;\n                left: 18px;\n                font-size: 11px;\n            }\n\n            input {\n                margin: 0;\n                padding: 0;\n                margin-left: 1rem;\n                font-size: 1rem;\n                background: var(--bg);\n                border: 1px solid var(--bg-bright);\n                padding: 5px;\n                color: white;\n\n                &:focus {\n                    outline: none;\n                    border: 1px solid var(--theme2);\n                }\n            }\n        }\n        \n        .hint {\n            margin-left: 1rem;\n\n            strong {\n                color: var(--theme-bright);\n            }\n        }\n    }\n\n    /* The switch - the box around the slider */\n    .switch {\n        position: relative;\n        display: inline-block;\n        width: 60px;\n        height: 34px;\n\n        input {\n            opacity: 0;\n            width: 0;\n            height: 0;\n        }\n    }\n    /* The slider */\n    .slider {\n        position: absolute;\n        cursor: pointer;\n        top: 0;\n        left: 0;\n        right: 0;\n        bottom: 0;\n        background-color: #ccc;\n        -webkit-transition: .4s;\n        transition: .4s;\n\n        &:before {\n            position: absolute;\n            content: "";\n            height: 26px;\n            width: 26px;\n            left: 4px;\n            bottom: 4px;\n            background-color: white;\n            -webkit-transition: .4s;\n            transition: .4s;\n        }\n    }\n\n    input:checked + .slider {\n        background-color: var(--theme-bright);\n    }\n\n    input:focus + .slider {\n     box-shadow: 0 0 1px var(--theme-bright);\n    }\n\n    input:checked + .slider:before {\n        -webkit-transform: translateX(26px);\n        -ms-transform: translateX(26px);\n        transform: translateX(26px);\n    }\n\n    .slider.round {\n        border-radius: 34px;\n\n        &:before {\n            border-radius: 50%;\n        }\n    }\n']))),Y=l.b.main(K||(K=Object(o.a)(["\n\n"])));function Z(){var n=Object(c.useContext)(J),e=n.state,t=n.setScToken,r=Object(c.useState)(0),i=Object(h.a)(r,2),a=i[0],o=i[1],s=Object(c.useState)(0),d=Object(h.a)(s,2),b=d[0],l=d[1],g=Object(c.useState)([]),p=Object(h.a)(g,2),m=p[0],f=p[1],O=window.location.search;return Object(c.useEffect)((function(){var n=new URLSearchParams(O).get("scToken");n&&(window.history.replaceState({},document.title,"/dashboard"),t(n))}),[O,t]),Object(c.useEffect)((function(){new N(e.apiBaseUrl,(function(n){o(n.joinedChannels),l(n.activeChannels),f(n.records)}))}),[e.apiBaseUrl]),Object(u.jsx)(Y,{children:Object(u.jsxs)(C.a,{children:[Object(u.jsx)(R,{}),Object(u.jsxs)(y.c,{children:[Object(u.jsx)(y.a,{path:"/dashboard",children:Object(u.jsx)(G,{})}),Object(u.jsxs)(y.a,{path:"/",children:[Object(u.jsx)(j,{activeChannels:b,joinedChannels:a}),Object(u.jsx)(x,{records:m})]})]})]})})}var $=Object(l.a)(Q||(Q=Object(o.a)(["\n    body {\n        --bg: #0e0e10;\n        --bg-bright: #18181b;\n        --bg-brighter: #3d4146;\n        --bg-dark: #121416;\n        --theme: #009148;\n        --theme-bright: #00a552;\n        --theme2: #2980b9;\n        --theme2-bright: #3498db;\n        --theme2-dark: #24618a;\n        --text: #F5F5F5;\n        --text-dark: #616161;\n        --twitch: #6441a5;\n        --twitch-dark: #4c317e;\n        --danger: #e74c3c;\n        --danger-dark: #c0392b;\n\n        background: var(--bg);\n        margin: 0;\n        padding: 0;\n        color: var(--text);\n        margin: 0;\n        font-family: Helvetica, Arial, sans-serif;\n        height: 100%;\n        width: 100%;\n\n        * {\n            box-sizing: border-box;\n        }\n    }\n"])));b.a.render(Object(u.jsx)(c.StrictMode,{children:Object(u.jsx)(z,{children:Object(u.jsxs)(s.a.Fragment,{children:[Object(u.jsx)($,{}),Object(u.jsx)(Z,{})]})})}),document.getElementById("root"))}},[[37,1,2]]]);
//# sourceMappingURL=main.358cc77a.chunk.js.map