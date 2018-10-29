#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

int print_message_function (void *ptr);

int main()
{

    // 被等待线程的返回值
    void *retval; // retval是无类型的指针，可以被任意指针赋值

    // 定义2个线程
    pthread_t thread1, thread2;

    char *message1 = "thread1";
    char *message2 = "thread2";

    // 线程1和线程2的创建状态
    int thread1_create_status = pthread_create(&thread1, NULL, (void *)&print_message_function, (void *) message1);
    int thread2_create_status = pthread_create(&thread2, NULL, (void *)&print_message_function, (void *) message2);

    // 线程创建成功，返回0,失败返回失败号
    if (thread1_create_status != 0) {

        printf("线程1创建失败\n");
    }
    if (thread2_create_status != 0) {
        printf("线程2创建失败\n");
    }



    // 主线程以阻塞的方式等待线程thread1的结束
    if(0 != pthread_join(thread1, &retval)){

        printf("the main thread pthread_join thread1 with a error\n");
    }
    printf("thread1 return value(retval) is %d\n", (int)retval);
    printf("thread1 end\n");

    // 主线程以阻塞的方式等待线程thread2的结束
    if(0 != pthread_join(thread2, &retval)){

        printf("the main thread pthread_join thread2 with a error\n");
    }
    printf("thread2 return value(retval) is %d\n", (int)retval);
    printf("thread2 end\n");

    return 0;
}

int print_message_function( void *ptr ) {

    int i = 0;

    for (i; i<100000; i++) {

        printf("%s:%d\n", (char *)ptr, i);
    }

    return 98;
}