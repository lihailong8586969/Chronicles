package pck02.food.fruit;

public class Apple extends Fruit {

    private String name = "" ;


    public Apple( String name ){

        this.name = name ;
    }

    public void introduce(){

        System.out.println( "我是苹果 , 我的名字是 : " + this.name );
    }

}
