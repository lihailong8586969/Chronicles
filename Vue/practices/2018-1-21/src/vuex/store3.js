


import Vue from 'vue' ;
import Vuex from 'vuex' ;
Vue.use( Vuex ) ;


let moduleA = {

  namespaced : true ,

  state : {

    "count" : 0

  } ,

  getters : {

    getCount ( state ){

      return state.count ;
    }

  } ,

  mutations : {

    add ( state , payload ){

      state.count ++ ;
    } ,

    desc ( state , payload ){

      state.count -- ;
    }

  } ,

  actions : {

    add ( context , payload ){

      context.commit( "add" , payload ) ;
    } ,

    desc ( context , payload ){

      context.commit( "desc" , payload ) ;
    }

  }

} ;


let moduleB = {

  namespaced : true ,

  state : {

    "count" : 0

  } ,

  getters : {

    getCount ( state ){

      return state.count ;
    }

  } ,

  mutations : {

    add ( state , payload ){

      state.count ++ ;
    } ,

    desc ( state , payload ){

      state.count -- ;
    }

  } ,

  actions : {

    add ( context , payload ){

      context.commit( "add" , payload ) ;
    } ,

    desc ( context , payload ){

      context.commit( "desc" , payload ) ;
    }

  }

} ;



const store = new Vuex.Store({

  // 开启严格模式
  strict: process.env.NODE_ENV !== 'production' ,

  modules: {

    "moduleA" : moduleA ,
    "moduleB" : moduleB
  }

}) ;


export default store ;



