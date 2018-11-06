#include <stdio.h>

int main()
{
	
	float math,english,political,profession,all;
	int c;
	
	scanf("%d",&c);
	
	while(c--){

		all=0.0;

		scanf("%f%f%f%f",&math,&english,&political,&profession);

		all = math+english+political+profession;

		if(math>=85&&profession>=85&&english>=55&political>=55&&all>305){

			if(all>=370){

				printf("A\n");
			}else{

				printf("B\n");
			}
		}
		else{

			printf("C\n");
		}
	}
	
   return 0;
}