package list.arraylist;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

public class UseArrayList {


    public static void main( String[] args ){

        List<String> stringList = new ArrayList<String>() ;

        stringList.add( "第 1 个元素" ) ;
        stringList.add( "第 2 个元素" ) ;
        stringList.add( "第 3 个元素" ) ;
        stringList.add( "第 4 个元素" ) ;
        stringList.add( "第 5 个元素" ) ;


        // 以迭代器的方式 循环ArrayList
        System.out.println( "以迭代器的迭代方式循环 ArrayList" );
        Iterator iterator = stringList.iterator() ;
        while( iterator.hasNext() ){

            System.out.println( "当前 ArrayList 的元素为 : " + iterator.next() );
        }

        // 以for循环的方式 循环ArrayList
        System.out.println( "以for循环的方式循环 ArrayList" );
        for( String e : stringList ){

            System.out.println( "当前 ArrayList 的元素为 : " + e );
        }

    }


}
