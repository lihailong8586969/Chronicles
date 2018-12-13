import Vue from 'vue'
import VueRouter from 'vue-router' ;
Vue.use( VueRouter ) ;

// 导入 app-view
import AppViewIndex from '../components/app/app-view/app-view-index/App-View-Index.vue' ;
import AppViewExplore from '../components/app/app-view/app-view-explore/App-View-Explore.vue' ;
import AppViewSearch from '../components/app/app-view/app-view-search/App-View-Search.vue' ;
import AppViewTopic from '../components/app/app-view/app-view-topic/App-View-Topic.vue' ;
import AppViewVuexPractice from '../components/app/app-view/app-view-vuexPractice/App-View-vuexPractice.vue' ;
import AppViewVueSSR from '../components/app/app-view/app-view-vueSSR/App-View-VueSSR.vue' ;


const routes = [

  { name : "index" , path: '/index', alias : "/" , component: AppViewIndex } ,

  { name : "explore" , path: '/explore', component: AppViewExplore } ,

  { name : "search" , path: '/search', component: AppViewSearch } ,

  { name : "topic" , path: '/topic', component: AppViewTopic } ,

  { name : "vuexPractice" , path: "/vuexPractice" , component : AppViewVuexPractice } ,

  { name : "vueSSR" , path: "/vueSSR" , component :AppViewVueSSR  }

] ;


const router = new VueRouter({

  "routes" : routes
}) ;


// 守卫以后 , 需要自己手动跳转
router.beforeEach( ( to , from , next ) => {

  next();
}) ;



export default router ;
