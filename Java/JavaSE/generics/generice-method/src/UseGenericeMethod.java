


public class UseGenericeMethod {



    // <T> 声明此方法为一个泛型方法
    public static <T> int getRandomNum ( T t ){

        System.out.println( "传入的参数T是 : " + t );

        return (int)( Math.random() * 1000 ) ;
    }



    public static <T,U,V,W,X> String getMsg( T t , U u , V v , W w , X x ){

        System.out.println( "传入的参数T是 : " + t );
        System.out.println( "传入的参数U是 : " + u );
        System.out.println( "传入的参数V是 : " + v );
        System.out.println( "传入的参数W是 : " + w );
        System.out.println( "传入的参数X是 : " + x );

        return "Hello World" ;
    }



    // 泛型运算 !!!
    public static <T extends Number> Number add( T x1 , T x2  ) throws Exception {

        if( x1.getClass() == Integer.class && x2.getClass() == Integer.class ){

            return x1.intValue() + x2.intValue() ;
        }
        else if( x1.getClass() == Float.class && x2.getClass() == Float.class ){

            return x1.floatValue() + x2.floatValue() ;
        }
        else if( x1.getClass() == Double.class && x2.getClass() == Double.class ){

            return x1.doubleValue() + x2.doubleValue() ;
        }
        else{

            throw new Exception( "参数类型必须同时是 int , float , double" ) ;
        }

    }

    public static void main( String[] args ){

        System.out.println( "获取到的随机数为 : " + UseGenericeMethod.getRandomNum( 2323 ) );

        System.out.println( "获取到的 Msg 为 : " + UseGenericeMethod.getMsg( 2323 , 3 , 2.3 , "12" , "we" ) );

        try{

            System.out.println( "x1+x2=" + UseGenericeMethod.add( 2.323 , 2.3 ) ) ;
        }catch ( Exception e ){

            System.out.println( e.getMessage() );
        }


    }

}




