// 代码来自：https://github.com/GFeather/Chat/blob/master/cilent.c


#include<stdio.h>
#include<sys/types.h>
#include<sys/socket.h>
#include<unistd.h>
#include<string.h>
#include<arpa/inet.h>
#include<pthread.h>

#define PORT 6666 //端口号
#define SIZE 1024 //定义的数组大小

int c_read(int *cilent_sock)
{
	char in[SIZE] = {'\n'};
	int re,reval;
while(1)
{
	memset(in,'\n',SIZE);
	re = read(*cilent_sock,in,SIZE-1);
	if(re == -1)
	{
		perror("read");
		break;
	}
	else if(re ==0)
		continue;
	else
		printf("other:%s\n",in);
}
	pthread_exit(&reval);
}

int main()
{
	char out[SIZE];
	int cilent_sock = socket(AF_INET,SOCK_STREAM,0);
	pthread_t t_id;
	int re;
	void *reval;
	struct sockaddr_in ser_addr;
	memset(&ser_addr,0,sizeof(ser_addr));
	ser_addr.sin_family=AF_INET;
	ser_addr.sin_port=htons(PORT);
	ser_addr.sin_addr.s_addr=inet_addr("127.0.0.1");
	int con = connect(cilent_sock,(struct sockaddr*)&ser_addr,sizeof(ser_addr));
	if(con == -1)
		{
			perror("connect");
			close(cilent_sock);
			return 0;
		}
	else if(con == 0 ) printf("success\n");
	pthread_create(&t_id,NULL,(void *)(&c_read),(void *)(&cilent_sock));
while(1)
{	
		memset(out,'\0',SIZE);
		printf("--------------------what fuck you want to say(enter 'end' to end):\n");
		scanf("%s",out);
		printf("out:%s\n",out);
		re = write(cilent_sock,out,SIZE-1);
		if(re==-1)
		{
			perror("send");
			break;
		}
		if(strncmp(out,"end",3)==0)break;
}
	printf("over\n");
	close(cilent_sock);
	pthread_cancel(t_id);
	return 0;
}