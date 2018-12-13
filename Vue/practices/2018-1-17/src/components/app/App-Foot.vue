<template>
<div id="App-Foot" >
  <div class="container" >

    <div class="text" >版权归知乎公司所有</div>

  </div>
</div>
</template>

<script>
export default {

  name: 'App-Foot',

  data () {

    return {

    }
  } ,

  components : {


  }

}
</script>

<style type="text/scss" lang="scss" scoped >


  #App-Foot{

    border-top:1px solid #ddd;
    background: #FFF;
    width: 100%;


    .container{

      width: 100%;
      height: 100% ;
      max-width: 1000px;
      margin: 0 auto;

      .text{

        padding: 30px 0;
        font-size: 13px;
        color: #666;
      }

    }

  }


</style>
