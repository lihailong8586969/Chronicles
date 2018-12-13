

<template>
  <div id="myBody" >

    {{ "传给我的消息是 : " + msg }}

  </div>
</template>

<script>

export default {
  name: 'myBody',
  data () {
    return {


    }
  },

  props : [ "msg" ] ,

}
</script>

<style type="text/scss" scoped lang="scss" >

  #myBody{

    width: 100%;
    height: 400px;
    background: #FFF;


    h2{

      height: 100%;
      text-align: center;
      color:#fff;
    }


  }



</style>
