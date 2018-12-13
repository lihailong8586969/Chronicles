

<script>


import span from './span.vue' ;

export default {

  name: 'App-View-Topic',

  data () {

    return {

    }
  } ,

  components : {

  } ,


  render : function( createElement ){

    return createElement(

      "ul" ,

      {} ,

      [

        createElement( 'li', {

          'class' : {

            'sb' : true ,
          } ,

          attrs : {

            "data-id" : 345
          }

        } , [

          createElement( span )

        ] ) ,

      ]

    ) ;
  }

}
</script>

<style lang="scss">

</style>
