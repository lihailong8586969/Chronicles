#include <stdio.h>
#include <stdlib.h>



/*
 * ��������������������ݣ�����һ�ö�������
 * 
 * ͨ��visit()���������ǿ��Եõ����Ĳ������������ 
 *
 */

typedef char TElemType ;
 
 
 
// ������������BiTNode , �� ���ָ��BiTree 
typedef struct BiTNode{
	
	TElemType data ;	//���ݵ�Ԫ 
	struct BiTNode *lchild , *rchild ;	// ���Һ���ָ�� 
}BiTNode , *BiTree ;



// ���ʶ������ľ������ 
void visit( TElemType data , int level ){
	
	printf( "%c λ�� %d ��\n" , data , level ) ;
	
}



//
//
//
///*	����������ʽ��
// *	�㷨˼�룺���������㣬�����ڵ�����Һ��ӽ������ �������������Ρ��� 
// *	ֻ�������е�ʵ����ͨ������ 
// */
// 
//// ����������ݵ�Ԫ 
//typedef char QElemType ;
//
//// ������ 
//typedef struct QNode{
//	
//	BiTree data_address	;	// ����е������ڶ������еĵ�ַ 	
//	QElemType data ;		// ����е����ݵ�Ԫ 
//	struct QNode *next ;	// ָ����һ����� 
//
//}QNode , *QueuePtr ; 
//
// 
//
////  ���� T Ϊ���������ڵ�ĵ�ַ 
//void LevelOrderTraverse( BiTree T ){
//	
//	if( T == NULL ){
//		
//		printf( "��ǰ��Ϊ�գ�����\n" ) ;
//	}
//	else{
//		
//		QueuePtr p1 , p2 ;
//		QueuePtr p = (QueuePtr)malloc( sizeof( QNode ) ) ;
//		
//		p -> data = T -> data ;	// �Ѹ��ڵ�����ݴ浽��һ����������
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
//			if( p -> data_address -> lchild  && !(p -> data_address -> rchild ) ){	// �����ǰ���ֻ������
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
//			else if( p -> data_address -> rchild && !(p -> data_address -> lchild ) ){	//  �����ǰ���ֻ���Һ��� 
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
//			else if( !( p -> data_address -> rchild ) && !( p -> data_address -> rchild ) ){	// һ�����Ӷ�û�� 
//			
//				p -> next = NULL ;
//			}
//			
//			else{	// ��ǰ�����2������ 
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





//  ǰ��������ʽ�� 
void PreOrderTraverse( BiTree T , int level ){
	
	if( T ){
		
		visit( T -> data , level ) ;
		PreOrderTraverse( T -> lchild , level + 1 ) ;
		PreOrderTraverse( T -> rchild , level + 1 ) ;
	}
}


//  ����������ʽ�� 
void InOrderTraverse( BiTree T , int level ){
	
	if( T ){
		
		InOrderTraverse( T -> lchild , level + 1 ) ;
		visit( T -> data , level ) ;
		InOrderTraverse( T -> rchild , level + 1 ) ;
	}
}


//  ����������ʽ�� 
void PostOrderTraverse( BiTree T , int level ){
	
	if( T ){
		
		PostOrderTraverse( T -> lchild , level + 1 ) ;
		PostOrderTraverse( T -> rchild , level + 1 ) ;
		visit( T -> data , level ) ;
	}
}






//  ��������� ����> ����һ�ö����� 
//  @param ָ�� ָ������ 
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
	
	// ����һ�������Ϊ���ڵ㣬����ȡ�����ַ�������ָ�� T 
	BiTree T = NULL ;
	
	printf("��������(�ո��ʾ���Ϊ��) : \n") ;
		
	CreateBiTreePreOrder( &T ) ;
	
	printf("\n����������£�\n") ;
	PreOrderTraverse( T , level ) ;
	
	printf("\n����������£�\n") ;
	InOrderTraverse( T , level ) ;
	
	printf("\n����������£�\n") ; 
	PostOrderTraverse( T , level ) ;
	
//	printf("\n����������£�\n") ;
//	LevelOrderTraverse( T ) ;

	return 0 ;
}



