package map.hashmap;


import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;

public class UseHashMap {


    public static void main( String[] args ){

        Map<String,String> stringStringMap = new HashMap<String, String>() ;
        
        stringStringMap.put( "one"   , "第1个元素" ) ;
        stringStringMap.put( "two"   , "第2个元素" ) ;
        stringStringMap.put( "three" , "第3个元素" ) ;
        stringStringMap.put( "four"  , "第4个元素" ) ;

        // 以迭代器的方式 HashMap
        System.out.println( "以迭代器的迭代方式循环 HashMap" );
        Iterator iterator = stringStringMap.entrySet().iterator() ;
        while( iterator.hasNext() ){

            System.out.println( "当前 HashMap 的元素为 : " + iterator.next() );
        }

        // 以for循环的方式 循环HashMap的key
        System.out.println( "以for循环的方式循环 HashMap 的key" );
        for( String e : stringStringMap.keySet() ){

            System.out.println( "当前 HashMap 的 key 为 : " + e );
        }

        // 以for循环的方式 循环HashMap的value
        System.out.println( "以for循环的方式循环 HashMap的value" );
        for( String e : stringStringMap.values() ){

            System.out.println( "当前 HashMap 的 value 为 : " + e );
        }


        // 输出 [two=第2个元素, one=第1个元素, three=第3个元素, four=第4个元素]
        System.out.println( stringStringMap.entrySet() ) ;

    }


}
