<template>
<div id="App-View-Index" >

  <div class="mainColumn" >

    <div class="columnHead" >

      <span class="nav" >提问</span>
      <span class="nav" >回答</span>
      <span class="nav" >写文章</span>
      <span class="nav" >写想法</span>

    </div>


    <div class="columnBody" >

      <template v-if=" loaded " >

        <template v-if=" feeds.length > 0 " >

          <div class="feed" v-for=" feed in feeds " >

            <div class="feed-meta" >

              <span>来自话题 : </span>

              <span>{{ feed.topic }}</span>

            </div>

            <div class="feed-author">

              <img class="thumb" :src=" feed.thumb "  />
              <span class="text" >{{ feed.name }}</span>

            </div>

            <div class="feed-title">{{ feed.title }}</div>

            <div class="feed-answer" >{{ feed.answer }}</div>

          </div>

        </template>

        <template v-else >

          <div style="padding: 50px 20px ; border-radius: 2px; text-align: center; background: #FFF; ">暂无推荐</div>

        </template>

      </template>

    </div>


    <div class="columnFoot">

      <div class="loading" :class="{ 'loaded' : loaded }" >{{ loadedText }}</div>

    </div>

  </div>

  <div class="sideBar" ></div>

</div>
</template>

<script>

let $vm = {} ;


let myFinallyHandle = function(){

  console.log( "请求结束了 ..." );
} ;

export default {

  name: 'App-View-Index',

  data () {

    return {

      "loaded" : false ,

      "loadedText" : "正在为您推荐..." ,

      "feeds" : [  ] ,

    }
  } ,

  components : {

  } ,

  created : function(){

  } ,

  mounted : function(){

    $vm = this ;

    this.$http.get( 'http://localhost:9815/index.php' , { params : { "id" : "lvsi" } } ).then( function( response ) {

      $vm.loaded = true ;
      $vm.feeds = response.data ;

      myFinallyHandle() ;

    }).catch( function( error ) {

      $vm.loadedText = "加载失败了" ;
      console.log( "失败了 ：" , error );

      myFinallyHandle() ;

    }) ;

  }



} ;


</script>

<style type="text/scss" lang="scss" scoped >

  #App-View-Index{


    // 主列
    .mainColumn{

      width:694px;


      .columnHead{

        line-height: 58px;
        padding: 0 20px;
        background: #FFF;
        border-radius: 2px;
        margin-bottom: 10px;

        .nav{

          display: inline-block;
          margin-right: 20px;
        }

      }


      .columnBody{

        .feed{

          background: #FFF;
          margin-bottom: 10px;
          padding: 16px 20px;
          border-radius: 2px;

          .feed-meta{

            margin-bottom: 14px;
            font-size: 15px;
            color: #8590a6;
          }

          .feed-author{

            margin-bottom: 14px;

            .thumb{

              width: 25px;
              height:25px;
              vertical-align: middle;
              border-radius: 3px;
              margin-right: 10px;
            }

            .text{

              font-size: 14px;
              font-weight: bold;
            }

          }

          .feed-title{

            font-size: 18px;
            font-weight: 600;
            font-synthesis: style;
            line-height: 1.6;
            color: #1e1e1e;
            margin-bottom: 14px;
          }

          .feed-answer{

            font-size: 15px;
            color: #262626;
            margin-bottom: 14px;
          }

        }


      }


      .columnFoot{


        .loading{

          display: block;
          text-align: center;padding: 20px 0;
          font-size: 13px;
        }

        .loaded{

          display: none;
        }

      }

    }

    // 侧边栏
    .sideBar{



    }

  }

</style>
