package pck01.annotationProcessor;

import pck01.annotation.FruitColor;
import pck01.annotation.FruitName;
import pck01.annotation.FruitProvider;

import java.lang.reflect.Field;



public class AnnotationProcessor {


    /**
     * 处理注解
     * @param classObj
     */
    public static void processAnnotation(Class<?> classObj){

        String strFruitName=" 水果名称：";
        String strFruitColor=" 水果颜色：";
        String strFruitProvicer="供应商信息：";

        Field[] fields = classObj.getDeclaredFields();

        for(Field field :fields){

            if(field.isAnnotationPresent(FruitName.class)){

                FruitName fruitName = (FruitName) field.getAnnotation(FruitName.class);

                System.out.println( "水果名称：" + fruitName.value()  ) ;
            }
            else if(field.isAnnotationPresent(FruitColor.class)){

                FruitColor fruitColor= (FruitColor) field.getAnnotation(FruitColor.class);

                System.out.println( "水果颜色：" + fruitColor.fruitColor().toString()  ) ;
            }
            else if(field.isAnnotationPresent(FruitProvider.class)){

                FruitProvider fruitProvider= (FruitProvider) field.getAnnotation(FruitProvider.class);
                strFruitProvicer="供应商编号："+fruitProvider.id()+" 供应商名称："+fruitProvider.name()+" 供应商地址："+fruitProvider.address();
                System.out.println(strFruitProvicer);
            }
        }
    }

}
