
<template>

  <div>

    <div id="moduleA" >

      <div style="padding: 20px 10px;">

        现在的 count = {{ countA }}

      </div>

      <div style="padding: 20px 10px;">

        <button @click=" addA() " > count +1 </button>

      </div>

      <div style="padding: 20px 10px;">

        <button @click=" descA() " > count -1 </button>

      </div>

      <div style="padding: 20px 10px;">

        <button > 获取 count 的值为 {{ getCountA }} </button>

      </div>

    </div>



    <div id="moduleB" >

      <div style="padding: 20px 10px;">

        现在的 count = {{ countB }}

      </div>

      <div style="padding: 20px 10px;">

        <button @click=" addB() " > count +1 </button>

      </div>

      <div style="padding: 20px 10px;">

        <button @click=" descB() " > count -1 </button>

      </div>

      <div style="padding: 20px 10px;">

        <button > 获取 count 的值为 {{ getCountB }} </button>

      </div>

    </div>




  </div>

</template>

<script>

import { mapState , mapGetters , mapActions } from 'vuex' ;

export default {

  name: 'App-View-Topic',

  data () {

    return {


    }
  } ,

  computed : {

    ...mapState( "moduleA", {

      "countA" : "count"
    }) ,

    ...mapState( "moduleB", {

      "countB" : "count"
    }) ,

    ...mapGetters( "moduleA" , {

      "getCountA" : "getCount"
    }) ,

    ...mapGetters( "moduleB" , {

      "getCountB" : "getCount"
    }) ,

  },

  methods : {

    ...mapActions( "moduleA" , {

      "addA":"add" ,
      "descA":"desc"
    }) ,

    ...mapActions( "moduleB" , {

      "addB":"add" ,
      "descB":"desc"
    }) ,

  } ,

  components : {

  } ,

}
</script>

<style type="text/scss" lang="scss" scoped >


  #moduleA{

    border: 2px solid #333;
    background: #FFF;
    padding: 17px;
  }

  #moduleB{

    border: 2px solid green;
    margin-top:20px;
  }

  button{

    padding: 10px;
    margin: 10px;
  }

</style>
