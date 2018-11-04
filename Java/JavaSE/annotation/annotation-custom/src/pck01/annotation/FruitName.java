package pck01.annotation;


import java.lang.annotation.*;




/**
 * 水果名称注解
 */

@Target(ElementType.FIELD)
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface FruitName {

    String value() default "水果的名字未知" ;

}




