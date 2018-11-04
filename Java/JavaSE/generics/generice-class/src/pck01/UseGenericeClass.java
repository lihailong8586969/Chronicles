package pck01;

import java.util.Iterator;

public class UseGenericeClass<T> {



    // 泛型不能直接运算
    // public T add( T t1 , T t2 ){

        // System.out.print( "t1= "+ t1 +" , t2="+ t2 +" , t1+t2 = " );

        // return t1 + t2 ;
    // }


    public static void main( String[] args ){


        // 构造一个专门放狗的房子
        House<Dog> dogHouse = new House<Dog>() ;

        // 放进去
        Dog dog1 = new Dog() ;
        Dog dog2 = new Dog() ;
        dogHouse.put( dog1 );
        dogHouse.put( dog2 );

        Iterator<Dog> iterator = dogHouse.getArrayList().iterator() ;


        for( Dog dog : dogHouse.getArrayList() ){

            System.out.println( dog.getIntroduceInfo() );
        }


        while( iterator.hasNext() ){

            System.out.println( iterator.next().getIntroduceInfo() );
        }

    }

}





