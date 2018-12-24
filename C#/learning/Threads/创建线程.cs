// https://www.cnblogs.com/gangle/p/9285094.html

// 但是在子控件中不能创建线程( 我想应该就是这样 )


// 方法一：使用Thread类
public static void Main(string[] args)
{
     ThreadStart threadStart = new ThreadStart(Calculate);//通过ThreadStart委托告诉子线程执行什么方法　　
     Thread thread = new Thread(threadStart);
     thread.Start();//启动新线程
}

public static void Calculate()
{
     Console.Write("执行成功");
     Console.ReadKey();
}


// 方法三：使用ThreadPool.QueueworkItem

WaitCallback w = new WaitCallback(Calculate);
//下面启动四个线程,计算四个直径下的圆周长
ThreadPool.QueueUserWorkItem(w, 1.0);
ThreadPool.QueueUserWorkItem(w, 2.0);
ThreadPool.QueueUserWorkItem(w, 3.0);
ThreadPool.QueueUserWorkItem(w, 4.0);

public static void Calculate(double Diameter)
{
   return Diameter * Math.PI;
}