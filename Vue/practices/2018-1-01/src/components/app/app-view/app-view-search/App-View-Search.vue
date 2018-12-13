<template>
<div id="App-View-Search" >

  <div style="padding: 100px 20px;text-align: left;" >我要搜索的内容是 : {{ ( wd === "" ) ? "暂无搜索内容" : wd }} </div>

</div>
</template>

<script>

import router from '../../../../assets/js/route.js' ;

export default {

  name: 'App-View-Search',

  data () {

    return {

      "wd" : router.currentRoute.query.wd
    }
  } ,

  beforeRouteEnter (to, from, next) {

    next() ;
  }


}
</script>

<style lang="scss">

</style>
