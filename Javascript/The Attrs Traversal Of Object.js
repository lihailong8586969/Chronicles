function sayKeyName(obj){      
  function sayName(obj){
    if (!(obj instanceof Array) && !(obj instanceof Object)) {
       throw new TypeError('obj 类型错误！');
    }
    
    if (obj instanceof Array) {
      for (var i = 0; i < obj.length; ++i)
       {
          console.log('对象属性名：' , i);
          if (obj[i] instanceof Object) {
             sayName(obj[i]);
          }
       }
    } else {
       for (var key in obj) 
        {
          console.log('对象属性名：' , key);
          if (obj[key] instanceof Object) {
             sayName(obj[key]);
          }
        }
    }
  }
  
  sayName(obj);
}

var obj = [
   {a:
     [
       {
         b:1,
         c:[
             {d:1},
             {e:1}
           ]
        },
        {f:1}
      ]
    }
];
sayKeyName(obj);
F12查看结果...