
var arr = [];

for(i=0;i<=4;++i){

  arr[i] = function(){
    
    return i;
  };
}

for(j=0;j<=4;++j){
  
  console.log(arr[j]()); // 全部输出 5
}

for(i=0;i<=4;++i){
  
  console.log(arr[i]()); // 输出 0,1,2,3,4
}

