package set.hashset;

import java.util.HashSet;
import java.util.Iterator;
import java.util.Set;

public class UseHashSet {


    public static void main( String[] args ){

        // Set<类型>
        // 虽然有时候类型可以省略,如 Set stringHashSet = new HashSet() 等 ,
        // 但是不推荐省略 , 写上可以提供代码的可读性
        Set<String> stringHashSet = new HashSet<String>() ;

        stringHashSet.add( "one" ) ;
        stringHashSet.add( "two" ) ;
        stringHashSet.add( "three" ) ;
        stringHashSet.add( "four" ) ;
        stringHashSet.add( "five" ) ;
        stringHashSet.add( "six" ) ;

        System.out.println( "当前 HashSet 的大小为 ：" + stringHashSet.size() ) ;


        // 返回一个 HashSet 的迭代器
        Iterator iterator = stringHashSet.iterator() ;
        System.out.println( "以迭代器的迭代方式循环 HashSet" );
        while( iterator.hasNext() ){

            System.out.println( "当前 HashSet 的元素为 : " + iterator.next() );
        }

        System.out.println( "以 for : 循环的方式循环 HashSet" );
        for( String s : stringHashSet ){

            System.out.println( "当前 HashSet 的元素为 : " + s );
        }



        // 定义一个临时的 HashSet 命名为 stringHashSet1
        Set<String> stringHashSet1 = new HashSet<>() ;
        stringHashSet1.add( "seven" ) ;
        stringHashSet1.add( "eight" ) ;
        stringHashSet1.add( "nine" ) ;
        stringHashSet.addAll( stringHashSet1 ) ;
        System.out.println( "临时的 HashSet( stringHashSet1 ) 的大小为 ：" + stringHashSet1.size() ) ;


        // 将 stringHashSet1 全部添加到 stringHashSet 中
        stringHashSet.addAll( stringHashSet1 ) ;
        System.out.println( "当前 HashSet 的大小为 ：" + stringHashSet.size() ) ;


        // 将临时的 HashSet stringHashSet1 清空
        stringHashSet1.clear();
        System.out.println( "临时的 HashSet( stringHashSet1 ) 的大小为 ：" + stringHashSet1.size() ) ;


    }

}
