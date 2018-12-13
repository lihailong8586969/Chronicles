
<template>

  <div>

    <div style="padding: 20px 10px;">
      name 的值是 {{ name }}
    </div>

    <div style="padding: 20px 10px;">
      name 的值是 {{ this.$store.state.name }}
    </div>

    <div style="padding: 20px 10px;">
      <button @click=" changeName() " >改变name</button>
    </div>

    <div style="padding: 20px 10px;">
      <button @click=" recoverName() " >name复原</button>
    </div>

    <div style="padding: 20px 10px;">
      <button @click=" getName() " >获取name</button>
    </div>

  </div>

</template>

<script>

// 导入 mapMutations
import { mapState , mapMutations } from 'vuex' ;

export default {

  name: 'App-View-Topic',

  data () {

    return {


    }
  } ,

  computed : mapState({

    "name" : () => "不会"

  }) ,


  methods : {

    getName(){

      alert( this.$store.getters.getName ) ;
    } ,

    changeName(){

      this.$store.commit( 'changeName' , { "name" : 'JK' } ) ;
    } ,

    recoverName(){

      this.$store.dispatch( 'changeName' , { "name" : '吕思' } ) ;
    }

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

    border: 2px solid red;
    margin-top:20px;
  }

  button{

    padding: 10px;
    margin: 10px;
  }

</style>
