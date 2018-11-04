package pck01.annotation;


import java.lang.annotation.*;

/**
 * 水果颜色注解
 *
 */

@Target(ElementType.FIELD)
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface FruitColor {

    /**
     * 颜色枚举
     *
     */
    enum Color{ BULE , RED , GREEN };

    /**
     * 颜色属性
     */
    Color fruitColor() default Color.GREEN;

}




