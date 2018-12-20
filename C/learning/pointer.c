#include<stdio.h>

int main(void) {
    
    int var1 =3, var2=4, *ptr;
    
    ptr = &var1;
    var2 = *ptr;
    
    printf("%d", var2); // 3
    
    return 0;
}


// -------------------------------------------

void sub(int x, int y, int *z){
    
    *z = y-x;
}

int main(void) {
    int a,b,c;
    
    sub(10,5,&a);
    sub(7,a,&b);
    sub(a,b,&c);
    
    printf("%d %d %d", a,b,c);
    
    return 0;
}