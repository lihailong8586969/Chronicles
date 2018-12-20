#include<stdio.h>

struct ss{
    
    char name[10];
    int age;
    char sex;
}std[3], *p = std;

int main(void) {
    
    (*std).age = 23;
    printf( "%d", std[0].age );
    
    (*std).age = 23;
    
    int a[10];
    
    *a = 2;
    printf( "%d", a[0] );
    
    return 0;
}
