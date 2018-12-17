// "static void main" must be defined in a public class.

public class Main {
    
    public static int times=0;
    
    public static void main(String... args) {
        
        new Main(){{ start(); }};
        System.out.println("end");
    }
    
    public static void start(){
        
        if(times>=3){
            
            return;
        }
        
        ++times;
        System.out.println(times);
        System.out.println("end2");
    }
}