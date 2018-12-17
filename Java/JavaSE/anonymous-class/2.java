// "static void main" must be defined in a public class.
public class Main {
	
    public static void main(String... args) {

        System.out.println(args.length);
        
        if(args.length==0 && new Main(){{ main("a"); }} !=null ){
            
            System.out.println("A");
        }else{
            
            System.out.println("B");
        }
        
        System.out.println("end");
    }
}