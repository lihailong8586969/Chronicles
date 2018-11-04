package pck01;

public class Dog {

    private int age ;
    private String name ;

    public Dog(){

    }

    public String getIntroduceInfo(){

        return "我是一只狗,我的名字是" + this.getName() + " , 今年 " + this.getAge() + "岁了" ;
    }

    public int getAge() {
        return age;
    }

    public void setAge(int age) {
        this.age = age;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}








