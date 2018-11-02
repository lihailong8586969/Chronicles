
bool flock ( int handle, int operation [, int &wouldblock] );
flock() 操作的 handle 必须是一个已经打开的文件指针。operation 可以是以下值之一：

要取得共享锁定（读取程序），将 operation 设为 LOCK_SH（PHP 4.0.1 以前的版本设置为 1）
要取得独占锁定（写入程序），将 operation 设为 LOCK_EX（PHP 4.0.1 以前的版本中设置为 2）
要释放锁定（无论共享或独占），将 operation 设为 LOCK_UN（PHP 4.0.1 以前的版本中设置为 3）
如果你不希望 flock() 在锁定时堵塞，则给 operation 加上 LOCK_NB（PHP 4.0.1 以前的版本中设置为 4）


但在PHP中，flock似乎工作的不是那么好！在多并发情况下，似乎是经常独占资源，不即时释放，或者是根本不释放，造成死锁，从而使服务器的cpu占用很高，甚至有时候会让服务器彻底死掉。好像在很多linux/unix系统中，都会有这样的情况发生。所以使用flock之前，一定要慎重考虑。
那么就没有解决方案了吗？其实也不是这样的。如果flock()我们使用得当，完全可能解决死锁的问题。当然如果不考虑使用flock()函数，也同样会有很好的解决方案来解决我们的问题。经过我个人的搜集和总结，大致归纳了解决方案有如下几种。