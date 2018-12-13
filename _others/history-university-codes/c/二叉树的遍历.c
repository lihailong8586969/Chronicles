#include <stdio.h>
#include <stdlib.h>



/*
 * 根据先序遍历，输入数据，生成一棵二叉树，
 * 
 * 通过visit()函数，我们可以得到结点的层数或遍历序列 
 *
 */

typedef char TElemType ;
 
 
 
// 定义二叉树结点BiTNode , 和 结点指针BiTree 
typedef struct BiTNode{
	
	TElemType data ;	//数据单元 
	struct BiTNode *lchild , *rchild ;	// 左右孩子指针 
}BiTNode , *BiTree ;



// 访问二叉树的具体操作 
void visit( TElemType data , int level ){
	
	printf( "%c 位于 %d 层\n" , data , level ) ;
	
}



//
//
//
///*	层序遍历访问结点
// *	算法思想：先输出根结点，将根节点的左右孩子进入队列 。。。。。依次。。 
// *	只不过队列的实现是通过链表 
// */
// 
//// 链表结点的数据单元 
//typedef char QElemType ;
//
//// 链表结点 
//typedef struct QNode{
//	
//	BiTree data_address	;	// 结点中的数据在二叉树中的地址 	
//	QElemType data ;		// 结点中的数据单元 
//	struct QNode *next ;	// 指向下一个结点 
//
//}QNode , *QueuePtr ; 
//
// 
//
////  参数 T 为二叉树根节点的地址 
//void LevelOrderTraverse( BiTree T ){
//	
//	if( T == NULL ){
//		
//		printf( "当前树为空！！！\n" ) ;
//	}
//	else{
//		
//		QueuePtr p1 , p2 ;
//		QueuePtr p = (QueuePtr)malloc( sizeof( QNode ) ) ;
//		
//		p -> data = T -> data ;	// 把根节点的数据存到第一个链表结点中
//		p -> data_address = T ;
//		while( p ){
//			
//			printf( " %c ," , p -> data ) ;
//			
//			if( ( p -> data_address -> lchild ) ){
//				
//				p1 = (QueuePtr)malloc( sizeof( QNode ) ) ;
//				p1 -> data = p -> data_address -> lchild -> data ;
//				p1 -> data_address = p -> data_address -> lchild ;
//				
//				p1 -> next = NULL ;
//			}
//			
//			if( ( p -> data_address -> rchild ) ){
//				
//				
//			 	p2 = (QueuePtr)malloc( sizeof( QNode ) ) ;
//				p2 -> data = p -> data_address -> rchild -> data ;
//				p2 -> data_address = p -> data_address -> rchild ;
//				
//				p2 -> next  = NULL ;
//			}
//			
//			
//			
//			if( p -> data_address -> lchild  && !(p -> data_address -> rchild ) ){	// 如果当前结点只有左孩子
//			
//				printf( "hh") ;
//				QueuePtr temp = p ;
//				
//				while( temp -> next ){
//					
//					temp = temp -> next ;
//				} 
//				
//				temp -> next = p1 ;
//			}
//			
//			else if( p -> data_address -> rchild && !(p -> data_address -> lchild ) ){	//  如果当前结点只有右孩子 
//				
//				QueuePtr temp = p ;
//				    
//				while( temp -> next ){
//					
//					temp = temp -> next ;
//				} 
//				
//				temp -> next = p2 ;
//			}
//			
//			else if( !( p -> data_address -> rchild ) && !( p -> data_address -> rchild ) ){	// 一个孩子都没有 
//			
//				p -> next = NULL ;
//			}
//			
//			else{	// 当前结点有2个孩子 
//				
//				QueuePtr temp = p ;
//				
//				while( temp -> next ){
//					
//					temp = temp -> next ;
//				} 
//				
//				temp -> next = p1 ;
//				p1 -> next = p2 ;
//			}
//
//			p = p -> next ;
//	   }
//		
//	}
//	
//} 





//  前序遍历访问结点 
void PreOrderTraverse( BiTree T , int level ){
	
	if( T ){
		
		visit( T -> data , level ) ;
		PreOrderTraverse( T -> lchild , level + 1 ) ;
		PreOrderTraverse( T -> rchild , level + 1 ) ;
	}
}


//  中序遍历访问结点 
void InOrderTraverse( BiTree T , int level ){
	
	if( T ){
		
		InOrderTraverse( T -> lchild , level + 1 ) ;
		visit( T -> data , level ) ;
		InOrderTraverse( T -> rchild , level + 1 ) ;
	}
}


//  后序遍历访问结点 
void PostOrderTraverse( BiTree T , int level ){
	
	if( T ){
		
		PostOrderTraverse( T -> lchild , level + 1 ) ;
		PostOrderTraverse( T -> rchild , level + 1 ) ;
		visit( T -> data , level ) ;
	}
}






//  以先序遍历 ——> 构造一棵二叉树 
//  @param 指针 指向根结点 
void CreateBiTreePreOrder( BiTree *T ){
	
	TElemType data ;
	
	scanf( "%c" ,&data ) ;
	
	if( data == ' ' ){
		
		*T = NULL ;
	}	
	else{
		
		*T = ( BiTNode * )malloc( sizeof( BiTNode ) ) ; 
		( *T ) -> data = data ;
		
		CreateBiTreePreOrder( &( *T ) -> lchild ) ;
		CreateBiTreePreOrder( &( *T ) -> rchild ) ;
	}
	
} 


 




int main(){
	
	int level = 1 ;
	
	// 申请一个结点作为根节点，并获取到其地址返给结点指针 T 
	BiTree T = NULL ;
	
	printf("请输入结点(空格表示结点为空) : \n") ;
		
	CreateBiTreePreOrder( &T ) ;
	
	printf("\n先序遍历如下：\n") ;
	PreOrderTraverse( T , level ) ;
	
	printf("\n中序遍历如下：\n") ;
	InOrderTraverse( T , level ) ;
	
	printf("\n中序遍历如下：\n") ; 
	PostOrderTraverse( T , level ) ;
	
//	printf("\n层序遍历如下：\n") ;
//	LevelOrderTraverse( T ) ;

	return 0 ;
}



