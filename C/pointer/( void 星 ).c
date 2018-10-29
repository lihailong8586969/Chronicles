
// 表示无类型的指针
( void * ) 


int *pr ;

// 表示无类型的指针
( void * )p ;
void *p ; // 指针 p 是无类型的

// 错误

double d＝3.14;
double *dptr=&d;
int *iptr=dptr;   //错误，double和int占用不同的字节，编译报错。


// 正确 : 无类型指针可以接受任何类型的指针。
double d=3.14;
double *dptr=&d;
void *vptr=dptr  //正确，无类型指针可以接受任何类型的指针。

