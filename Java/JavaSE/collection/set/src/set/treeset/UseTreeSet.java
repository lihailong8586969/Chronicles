package set.treeset;

import java.util.Iterator;
import java.util.Set;
import java.util.TreeSet;

public class UseTreeSet {




    public static void main( String[] args ){


        Set<Integer> integerSet = new TreeSet<Integer>() ;

        integerSet.add( 10 ) ;
        integerSet.add( 1 ) ;
        integerSet.add( 5 ) ;
        integerSet.add( 7 ) ;
        integerSet.add( 3 ) ;

        // 先使用for循环 循环TreeSet
        System.out.println( "以 for : 循环的方式循环 TreeSet" );
        for( int e : integerSet ){

            System.out.println( "当前 TreeSet 的元素为 : " + e ) ;
        }


        // 以迭代器的方式 循环TreeSet
        System.out.println( "以迭代器的迭代方式循环 TreeSet" );
        Iterator iterator = integerSet.iterator() ;
        while( iterator.hasNext() ){

            System.out.println( "当前 TreeSet 的元素为 : " + iterator.next() );
        }


    }

}
