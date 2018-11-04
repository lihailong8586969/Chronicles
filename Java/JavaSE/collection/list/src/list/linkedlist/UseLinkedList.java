package list.linkedlist;

import java.util.Iterator;
import java.util.LinkedList;
import java.util.List;

public class UseLinkedList {


    public static void main( String[] args ){

        List<String> stringList = new LinkedList<String>() ;
        stringList.add( "第3个元素" ) ;
        stringList.add( "第4个元素" ) ;
        stringList.add( "第1个元素" ) ;
        stringList.add( "第2个元素" ) ;
        stringList.add( "第5个元素" ) ;

        // 以迭代器的方式 循环LinkedList
        System.out.println( "以迭代器的迭代方式循环 LinkedList" );
        Iterator iterator = stringList.iterator() ;
        while( iterator.hasNext() ){

            System.out.println( "当前 LinkedList 的元素为 : " + iterator.next() );
        }

        // 以for循环的方式 循环LinkedList
        System.out.println( "以for循环的方式循环 LinkedList" );
        for( String e : stringList ){

            System.out.println( "当前 LinkedList 的元素为 : " + e );
        }
    }

}
