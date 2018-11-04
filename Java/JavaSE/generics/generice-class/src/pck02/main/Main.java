package pck02.main;

import pck02.food.IFood;
import pck02.food.fruit.Apple;
import pck02.food.fruit.Fruit;
import pck02.food.meat.Meat;
import pck02.food.meat.PigMeat;
import pck02.plate.Plate;
import pck02.table.Table;

public class Main {



    public static void main( String[] args ){

        Fruit apple01 = new Apple( "apple01" ) ;
        Fruit apple02 = new Apple( "apple02" ) ;

        Meat pigMeat01 = new PigMeat( "pigMeat01" ) ;
        Meat pigMeat02 = new PigMeat( "pigMeat02" ) ;

        Plate<Fruit> fruitPlate = new Plate<>() ;

        Plate<Meat>  meatPlate = new Plate<>() ;


        fruitPlate.put( apple01 ) ;
        fruitPlate.put( apple02 ) ;


        meatPlate.put( pigMeat01 );
        meatPlate.put( pigMeat02 );


        Table<Plate<? extends IFood>> plateTable = new Table<>() ;

        plateTable.put( fruitPlate ) ;
        plateTable.put( meatPlate );

        for( Plate<? extends IFood> plate : plateTable.getPlateList() ){



        }

    }


}




