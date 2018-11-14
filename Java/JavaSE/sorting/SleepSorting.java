// 睡觉排序法
public class SleepSort {

    public static void main(String[] args) {
    
        int[] ints = {1,4,7,3,8,9,2,6,5};
    
        SortThread[] sortThreads = new SortThread[ints.length];
    
        // 创建多个线程
        for (int i = 0; i < sortThreads.length; i++) {
    
            sortThreads[i] = new SortThread(ints[i]);
        }
    
        // 线程开始分别执行, 但是如果数组很长，也就是说这个数组遍历需要很长时间，
        // 如果最后1个恰好是最小的睡眠时间，但由于遍历时间太长，前面的已经睡醒了，
        // 前面已经输出了，最小的这个还没输出
        // 所以这个方法只是理论上的
        for (int i = 0; i < sortThreads.length; i++) {
        
            // 线程开始执行
            sortThreads[i].start();
        }

    }
}

class SortThread extends Thread{

    // 睡觉的时间
    int ms = 0;

    public SortThread(int ms){

        this.ms = ms;
    }

    // 如果睡眠时间很长怎么办？ 除以一个很大的数可以吗 ？ 比如 1000000
    // 不可以的，因为所有线程的睡眠时间相差太近会出现没有及时唤醒的问题 ！！！
    public void run(){
    
        try {
    
            // 睡觉
            sleep(ms*10+10);

        } catch (InterruptedException e) {

            e.printStackTrace();
        }

        // 打印出数字
        System.out.println(ms);
    }

}