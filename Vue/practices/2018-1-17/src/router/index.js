
// 导入路由
import VueRouter from 'vue-router' ;


// 导入 app-view
import AppViewIndex from '../components/app/app-view/app-view-index/App-View-Index.vue' ;
import AppViewExplore from '../components/app/app-view/app-view-explore/App-View-Explore.vue' ;
import AppViewSearch from '../components/app/app-view/app-view-search/App-View-Search.vue' ;
import AppViewTopic from '../components/app/app-view/app-view-topic/App-View-Topic.vue' ;


const routes = [

  { name : "index" , path: '/index', alias : "/" , component: AppViewIndex } ,

  { name : "explore" , path: '/explore', component: AppViewExplore } ,

  { name : "search" , path: '/search', component: AppViewSearch } ,

  { name : "topic" , path: '/topic', component: AppViewTopic } ,

] ;


const router = new VueRouter({

  "routes" : routes
}) ;


// 守卫以后 , 需要自己手动跳转
router.beforeEach( ( to , from , next ) => {


  console.clear() ;
  console.log( "beforeEach ——> to : " ,  to );
  console.log( "beforeEach ——> from : " ,  from );
  console.log( "beforeEach ——> next : " ,  next );
  next();
}) ;



export default router ;






