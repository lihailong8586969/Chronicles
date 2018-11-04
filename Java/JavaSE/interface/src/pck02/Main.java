package pck02;

public class Main {




    public static void main( String[] args ){


        // 类   作为 数据引用类型
        // pck02.BMWCar bmwCar = new pck02.BMWCar() ;

        // 接口 作为 数据引用类型
        Car bmwCar = new BMWCar() ;


        bmwCar.run();
        bmwCar.start();
    }

}




