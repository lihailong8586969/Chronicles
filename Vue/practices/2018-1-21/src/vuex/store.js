
import Vue from 'vue' ;
import Vuex from 'vuex' ;
Vue.use( Vuex ) ;


const store = new Vuex.Store({

  state: {

    count: 0 ,

    todos : [

      { "id" : 1 , "title" : "做作业" } ,
      { "id" : 2 , "title" : "看电视" } ,
      { "id" : 3 , "title" : "上楼" } ,
      { "id" : 4 , "title" : "下楼" } ,
      { "id" : 5 , "title" : "喝水" } ,
      { "id" : 6 , "title" : "吃饭" }

    ]

  } ,

  getters: {

    getComputedCount () {

      return store.state.todos.filter( todo => todo.id >= 3 ) ;
    }

  } ,

  mutations: {

    increment (state) {

      state.count ++ ;
    } ,

    desc ( state ){

      state.count -- ;
    }
  }

}) ;


export default store ;



