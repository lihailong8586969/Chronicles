package pck03;

import org.apache.commons.lang3.builder.ToStringBuilder;

public class Test {

    private String name = "Test类" ;

    public String toString() {

        return new ToStringBuilder(this).
                append("name", this.name ).
                toString();
    }


    public static void main( String[] args ){

        System.out.println( new Test().toString() ) ;

        Lvsi lvsi = new Lvsi<>() ;

        lvsi.setT( "bhcd" );

        System.out.println( lvsi.getT() );

        lvsi.setT( 876 );
        System.out.println( lvsi.getT() );


        lvsi.setT( 4345.232 );
        System.out.println( lvsi.getT() );

        System.out.println( lvsi.toString() );


    }

}
