
import Vue from 'vue' ;
import Vuex from 'vuex' ;
Vue.use( Vuex ) ;




let state = {

  "name" : "吕思"
} ;

let mutations = {

  changeName ( state, payload ){

    state.name = payload.name ;
  }

} ;



let getters = {

  getName ( state ){

    return state.name ;
  }

} ;

let actions = {

  changeName ( context , payload ){

    context.commit( "changeName" , payload ) ;
  } ,

  changeName2 ( { dispatch , commit } , payload ){

    commit( "changeName" , payload ) ;
  }

} ;

const store = new Vuex.Store({

  state ,

  getters ,

  mutations ,

  actions

}) ;


export default store ;



