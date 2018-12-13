<template>
<div id="App-Head" >
  <div class="container">

    <div class="logo">
      <router-link to="/" class="text" >知乎</router-link>
    </div>

    <div class="navigation" >
      <router-link to="/" class="nav" >首页</router-link>
      <router-link to="/explore" class="nav" >发现</router-link>
      <router-link to="/topic" class="nav" >话题</router-link>
    </div>

    <div class="search">
      <input type="text" v-model.trim="wd" @keyup.enter="beginSearch()" placeholder="搜索" />
      <button @click="beginSearch()" class="btn" >搜索</button>
    </div>

  </div>
</div>
</template>

<script>

import router from "../../router" ;

export default {

  name: 'App-Head',

  data () {

    return {

      "$router" : router ,

      "wd" : ( typeof router.currentRoute.query.wd === "undefined" ) ? "" : router.currentRoute.query.wd   ,

    }
  } ,

  components : {


  } ,

  methods: {

    beginSearch () {

      if( this.wd !== "" ){

        router.push( { name : "search" , path : "/search" , query : { "wd" : this.wd } } ) ;
      }

    }

  }

}
</script>

<style type="text/scss" lang="scss" scoped >


  #App-Head{

    $height : 60px;

    width: 100%;
    height: $height;
    line-height: $height;
    overflow: hidden;
    background: #FFF;
    box-shadow: 0 1px 3px 0 rgba(0,34,77,.1);
    border-bottom:1px solid #ddd;


    .container{

      width: 100%;
      height: 100% ;
      max-width: 1000px;
      margin: 0 auto;


      .logo{

        display: inline-block;
        height: 100%;
        margin-right: 27px;

        .text{

          cursor: pointer;
          display: inline-block;
          font-size: 25px;
        }

      }


      .navigation{

        display: inline-block;
        height: 100%;
        margin-right: 27px;

        .nav{

          padding-right: 17px;
          display: inline-block;
          font-size: 13px;
        }

      }


      .search{

        display: inline-block;
        height: 100% ;
        line-height: $height;

        input{

          height: 15px;
          font-size: 13px;
          vertical-align: middle;
          margin-top: -8px;
          width: 220px;
        }

        .btn{

          border-radius: 4px;
          padding: 10px 17px;
          margin-top: -8px;
          color: #FFF;
          background: #555;
          display: inline-block;
          vertical-align: middle;
          margin-left: 10px;

          &:hover{

            cursor: pointer;
            background: #000;
          }
        }
      }

    }

  }


</style>
