// 1、导入 CSS
import './assets/js/loadCSS.js' ;

// 2、导入 Vue
import Vue from 'vue' ;

// 3、、导入路由
import router from './router/index.js' ;

// 4 引入 axios
import Axios from 'axios' ;
Vue.prototype.$http = Axios ;


// 5、引入 vuex
// import store from './vuex/store.js' ;
import store from './vuex/store3.js' ;


// 引入 App.vue 组件
import App from './components/app/App.vue' ;


new Vue({

  // 路由
  router ,

  store ,

  // 组件
  components : {

    "app" : App ,
  }

}).$mount( "#Root" ) ;




