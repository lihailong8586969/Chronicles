

// 1、导入 CSS
import loadCSS from './assets/js/loadCSS.js' ;

// 2、导入 Vue
import Vue from 'vue' ;


// 3、、导入路由
import router from './assets/js/route.js' ;
import VueRouter from 'vue-router' ;
Vue.use( VueRouter ) ;


// 4、导入 element-ui
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI) ;



import App from './components/app/App.vue' ;


new Vue({

  router ,

  components : {

    "app" : App ,
  }

}).$mount( "#Root" ) ;




