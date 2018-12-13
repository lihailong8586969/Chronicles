
// 加载 CSS 模块
import loadCSS from './loadCSS.js' ;



import Vue from 'vue' ;

import MyHead from './components/MyHead.vue' ;

import MyBody from './components/MyBody.vue' ;

import MyFoot from './components/MyFoot.vue' ;


new Vue({

  el: '#myFirstVueApp' ,

  data : {

    msg : [ "你负责头部，对不？" , "你负责身体，对不？" , "你负责脚部，对不？" ]

  } ,

  components : {

    MyHead ,

    MyBody ,

    MyFoot

  } ,

  methods : {

    myAlert : function( e ){

      alert( event ) ;
    }


  }



}) ;


