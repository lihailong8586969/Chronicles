

<template>
  <div id="foot" >

    <h2>这是脚部 <span class="text"> {{ "传给我的消息是 : " + msg }} </span> </h2>

  </div>
</template>

<script>
  
  export default {
  name: 'Foot',
  data () {
    return {

    }
  } ,

  props : [ "msg" ] ,



}
</script>

<style type="text/scss" scoped lang="scss" >

  #foot{

    width: 100%;
    height: 70px;
    background: #444;


    h2{

      line-height: 70px;
      height: 100%;
      text-align: center;
      color:#fff;
    }


  }



</style>
