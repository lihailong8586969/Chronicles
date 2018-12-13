<template>
<div id="App-View-Explore" >

  <div style="padding: 100px 20px;text-align: left;" >我要探索</div>


  <el-row >
    <el-col :span="2" :offset="0"><div class="grid-content bg-purple"></div></el-col>
    <el-col :span="12" :offset="6"><div class="grid-content bg-purple"></div></el-col>
  </el-row>

</div>
</template>

<script>
export default {

  name: 'App-View-Explore',

  data () {

    return {

    }
  } ,

  components : {

  }

}
</script>


<style type="text/scss" lang="scss" scoped >


  .el-row {
    margin-bottom: 20px;
    &:last-child {
       margin-bottom: 0;
     }
  }
  .el-col {
    border-radius: 4px;
  }
  .bg-purple-dark {
    background: #99a9bf;
  }
  .bg-purple {
    background: #d3dce6;
  }
  .bg-purple-light {
    background: #e5e9f2;
  }
  .grid-content {
    border-radius: 4px;
    min-height: 36px;
  }
  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }



</style>
