// 必须加这个 false 赋值 ( https://blog.csdn.net/kingboy2008/article/details/6529362 )
System.Windows.Forms.Control.CheckForIllegalCrossThreadCalls = false;
Task task = new Task(() => {

     while (true){
          this.fly();
     }
});
task.Start();
