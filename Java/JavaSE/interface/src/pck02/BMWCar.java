package pck02;

public class BMWCar implements Car {

    @Override
    public void start() {

        System.out.println( "宝马汽车开始启动" );
    }

    @Override
    public void run() {

        System.out.println( "宝马汽车正在行驶" );
    }
}
