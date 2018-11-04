package pck02.plate;


import java.util.ArrayList;
import java.util.List;

// 泛型类 : 盘子
public class Plate<Food> {

    private String name = "" ;


    public Plate(){


    }


    // 食物列表
    private List<Food> foodList = new ArrayList<>() ;


    // 再往当前的食物列表中放食物
    public void put( Food food ){

        this.foodList.add( food ) ;
    }

    // 获取当前的食物列表
    public List<Food> getFoodList(){

        return this.foodList ;
    }


}
