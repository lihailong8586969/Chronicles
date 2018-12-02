





(function(){








    // 配置


    // 工具


    // 注册器
    var eventRGST = require( "../../js.core/js.core.rgst/event.rgst" ) ;


    // 服务


    // 插件
    var $ = require( "../../js.core/js.core.plug/jquery.plug" ) ;


    // 组件


    // API 接口
    var uploadThumbAPI = require( "../../js.core/js.core.api.user/uploadThumb.api" ) ;


    // 传送器





    // 私有模块 : 上传头像模块
    module.exports = uploadThumbPrivateModule = {




        install : function (){

            this.startup() ;
        } ,



        // 开启
        startup : function (){


            // 绑定事件
            this.constructProcess.bindEvent() ;
        } ,


        // 关闭
        shutdown : function (){

            this.startup = function (){

                return 0;
            }
        } ,


        // 构造
        constructProcess : {


            // 绑定事件
            bindEvent : function(){


                $( ".autoId" ).each(function(){

                    if( $( this ).attr( "ls-module" ) === 'module.private.userHome' ){

                        eventRGST.addListener( "#" + $( this ).attr( "id" ) + " .cpnt_profile #inputFile" , "change" , uploadThumbPrivateModule.constructProcess.handleEvent ) ;

                        // 退出循环
                        return false ;
                    }
                }) ;
            } ,


            // 事件处理程序
            handleEvent : function(){

                uploadThumbPrivateModule.constructProcess.gatherData( arguments[1].get(0) , arguments[1] ) ;
            } ,


            // 获取数据
            gatherData: function( _this , $this ){

                var formData = new FormData() ;


                // 存放上传的数据
                var len = _this.files.length ;
                for( var i=0; i<=len-1; ++i ){

                    formData.append( _this.files[i].name , _this.files[i] ) ;
                }

                // 必须使用 formData.get 或者 formData.getAll() 才能获取
                // https://segmentfault.com/q/1010000012507163
                console.log(formData.get("img"));

                // 清除原始值
                $this.val( "" ) ;

                this.runBS( formData ) ;

            } ,


            // 开始业务
            runBS : function( paramTransmit ){

                uploadThumbAPI.runup( paramTransmit ) ;
            }

        }

    } ;

})() ;






















