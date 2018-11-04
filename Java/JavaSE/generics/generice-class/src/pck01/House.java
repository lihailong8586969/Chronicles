package pck01;

import java.util.ArrayList;
import java.util.List;

// 定义泛型类 : 房屋
public class House<T> {

    private List<T> arrayList = new ArrayList<T>() ;


    public void put( T t ){

        this.arrayList.add( t ) ;
    }



    public T get( int index ){

        return this.arrayList.get( index ) ;
    }

    public List<T> getArrayList(){

        return this.arrayList ;
    }

}







