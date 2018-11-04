package pck02.table;

import java.util.ArrayList;
import java.util.List;


// 泛型类 : 桌子
public class Table<Plate> {


    // 盘子列表
    private List<Plate> plateList = new ArrayList<>() ;



    public void put( Plate t ){

        this.plateList.add( t ) ;
    }


    // 获取当前的盘子列表
    public List<Plate> getPlateList(){

        return this.plateList ;
    }


}
