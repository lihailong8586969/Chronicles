#include<stdio.h>
#include<sys/types.h>
#include<sys/socket.h>
#include<arpa/inet.h>
#include<string.h>
#include<unistd.h>
#include<errno.h>
#include<fcntl.h>
#include<time.h>
#include<pthread.h>

#define PORT 6666 //端口号
#define SIZE 1024 //定义的数组大小
#define FLAGS O_APPEND | O_CREAT | O_RDWR
#define LINE \n
int user[201]={0};
int *c_link;

int creat_socket()
{
	int listen_sock = socket(AF_INET,SOCK_STREAM,0);
	if(listen_sock==-1)
		{
			perror("socket");
			return -1;
		}
	struct sockaddr_in addr;
	memset(&addr,0,sizeof(addr));
	addr.sin_family = AF_INET;
	addr.sin_port = htons(PORT);
	addr.sin_addr.s_addr=htonl(INADDR_ANY);
	int reb = bind(listen_sock,(struct sockaddr*)&addr,sizeof(addr));
	if(reb==-1)
		{
			perror("bind");
			return -1;
		}
	reb = listen(listen_sock,5);
	if(reb==-1)
		{
			perror("listen");
			return -1;
		}
	return listen_sock;
}

int wait_listen(int listen_sock)
{
	struct sockaddr_in c_addr;
	int stl = sizeof(c_addr);
	int cilent_sock = accept(listen_sock,(struct sockaddr*)&c_addr,&stl);
	if(cilent_sock==-1)
		{
			perror("accept");
			return -1;
		}
	return cilent_sock;
}

void w_flog(char buf[],int re)
{
	int i = 0;
	time_t t_temp;
	char *tt;
	char w_time[30];
	int count=0;
	int file = open("/opt/Cwork/flog.txt",FLAGS,S_IRWXU | S_IRWXO | S_IRWXG);
        if(file==-1)
        {
                perror("open");
        }
	while(1)
	{	
		if(buf[i] == '\0'){buf[i] = '\n'; break;}
		else i++;
	}
	time(&t_temp);
	tt = asctime(gmtime(&t_temp));
	while(*tt)
	{
		w_time[count]=*(tt++);
		printf("%c",w_time[count]);
		count++;
	}
	write(file,asctime(gmtime(&t_temp)),count);
	re = write(file,buf,i+1);
	if(re==-1)
	{
		perror("writen");
		close(file);
	}
}

void hand(int *t_sock)
{	
	char buf[SIZE];
	int i , re , reval;
	int cilent_sock = *t_sock;
while(1)
{
	i = 1;
	memset(buf,0,SIZE);
	re = read(cilent_sock,buf,SIZE-1);
	if(re==-1)
	{
		perror("readn");
		close(cilent_sock);
		break;
	}
	if(re==0)
	{
		continue;
	}
	w_flog(buf,re);
	if(strncmp(buf,"end",3)==0)break;
	while(user[i] > 1)
	{
		if(user[i] != cilent_sock)
		write(user[i],buf,re);
		i++;
	}
}
	close(cilent_sock);
	pthread_exit(&reval);
}

int main()
{
	int listen_sock = creat_socket();
	int cilent_sock,t_dech;
	pthread_t t_id;
	c_link = user;
	user[200] = 1;
while(1)
{
	if((*c_link) == 1) break;
	c_link++;
	*c_link = wait_listen(listen_sock);
	printf("link%d\n",*c_link);
	if((*c_link) != -1)
	{
		pthread_create(&t_id,NULL,(void *)(&hand),(void *) c_link);
		t_dech = pthread_detach(t_id);
		if(t_dech != 0)
		{
			perror("detach");
			pthread_cancel(t_id);
			break;
		}
	}
}
	for(c_link = user;(*c_link) != 1 ;c_link++)
		close(*c_link);
	close(listen_sock);
	return 0;
}	