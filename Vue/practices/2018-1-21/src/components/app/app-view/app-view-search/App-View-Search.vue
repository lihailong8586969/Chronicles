<template>
<div id="App-View-Search" >

  <div style="padding: 100px 20px;text-align: left;" >我要搜索的内容是 : {{ ( wd === "" ) ? "暂无搜索内容" : wd }} </div>

</div>
</template>

<script>
export default {

  name: 'App-View-Search',

  data () {

    return {

      "wd" : this.$router.currentRoute.query.wd
    }
  } ,

  beforeRouteUpdate ( to , from , next ) {

    // 可以打印 to , from , next 查看都有什么属性
    this.wd = to.query.wd ;

    next() ;
  }


}
</script>

<style lang="scss">

</style>
