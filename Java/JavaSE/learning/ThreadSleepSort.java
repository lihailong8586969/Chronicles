public class ArraySort implements Runnable{

	private int number;

	public ArraySort(int number){

		this.number = number;
	}

	public static void main(String[] args){

		int[] numbers = new int[]{102,338,62,9132,580,666};
		for(int number : numbers){

			new Thread(new ArraySort(number)).start();
		}
	}

	public void run(){

		try{

			Thread.sleep(this.number);
			System.out.println(this.number);
		}catch(InterruptedException e){

			e.printStackTrace();
		}
	}

}