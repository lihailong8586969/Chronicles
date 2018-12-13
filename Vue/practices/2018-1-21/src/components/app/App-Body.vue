<template>
<div id="App-Body" >
  <div class="container" >

    <transition name="slide"  >
      <router-view :class="{ 'def_displayInlineBlock' : def_displayInlineBlock }" ></router-view>
    </transition>

  </div>
</div>
</template>

<script>
export default {

  name: 'App-Body',

  data () {

    return {

      "def_displayInlineBlock" : false
    }
  } ,

  components : {


  } ,


}
</script>

<style type="text/scss" lang="scss" scoped >


  #App-Body{


    margin: 20px auto;

    .container{

      width: 100%;
      height: 100% ;
      max-width: 1000px;
      margin: 0 auto;
      min-height: 300px;

      overflow-x: hidden;
    }
  }





  @keyframes slide-enter-active {

    0%   { margin-right: -200px; opacity: 0; }
    100% { position:absolute;margin-right:0;opacity: 1;  }
  }
  .slide-enter-active{

    animation: slide-enter-active 0.5s  ;
    animation-timing-function: ease-out;
    animation-fill-mode: forwards;
  }


  .slide-enter{

    position:absolute;margin-right:0;opacity: 1;
  }




  @keyframes slide-leave-active {

    0%   { position:absolute; margin-left: 0; opacity: 1; }
    100% { margin-left:-200px;opacity: 0;  }
  }
  .slide-leave-active{

    animation: slide-leave-active 0.5s  ;
    animation-timing-function: ease-out;
    animation-fill-mode: forwards;
  }
  .slide-leave-to{

    position:absolute;margin-left:-200px;opacity: 0;
  }


</style>
