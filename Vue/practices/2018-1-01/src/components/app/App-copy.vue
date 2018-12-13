<template>
<div id="App" >

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
        <input type="text" placeholder="搜索" />
        <router-link to="/search" >
          <button class="btn" >搜索</button>
        </router-link>
      </div>

    </div>
  </div>



  <div id="App-Body" >
    <div class="container" >

      <router-view></router-view>

    </div>
  </div>



  <div id="App-Foot" >
    <div class="container" >

      <div class="text" >版权归知乎公司所有</div>

    </div>
  </div>

</div>
</template>

<script>
export default {

  name: 'App',

  data () {

    return {

    }
  } ,

  components : {


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



  #App-Body{

    .container{

      width: 100%;
      height: 100% ;
      max-width: 1000px;
      margin: 0 auto;
      min-height: 300px;
    }
  }



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
