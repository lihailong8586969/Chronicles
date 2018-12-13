

<template>
  <div id="head" >

    <h2>这是头部

      <span class="text"> {{ "传给我的消息是 : " + msg }} </span>

      <button @click="replay()" >我要回复</button>

    </h2>

  </div>
</template>

<script>

export default {
  name: 'Head',
  data () {
    return {

    }
  } ,

  props : [ "msg" ] ,


  methods: {

    replay : function(){

      this.$emit( "replay" , "对的，我就是负责头部" ) ;
    }

  }

}
</script>

<style type="text/scss"  lang="scss" >

  #head{

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

  .text{

    font-size: 13px;
  }

</style>
