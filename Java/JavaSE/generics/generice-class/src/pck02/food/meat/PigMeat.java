package pck02.food.meat;

public class PigMeat extends Meat {

    private String name = "" ;

    public PigMeat( String name ){

        this.name = name ;
    }

    public void introduce(){

        System.out.println( "我是猪肉 , 我的名字是 : " + this.name );
    }
}
