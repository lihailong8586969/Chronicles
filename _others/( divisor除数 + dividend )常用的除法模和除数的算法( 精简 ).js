// 一下两种命名方法都 OK


// divisor 除数
// dividend 被除数
function isSelfDividing(num){
    
    var divisor = num;
    while( divisor != 0 ){
        
        var dividend = divisor%10;
        if( num % dividend != 0){
            
            return false;
        }
        
        divisor = Math.floor(divisor/10);
    }
    
    return true;
}

// divisor 除数
// remainder 余数
function isSelfDividing(num){
    
    var divisor = num;
    while( divisor != 0 ){
        
        var remainder = divisor%10;
        if( num % remainder != 0){
            
            return false;
        }
        
        divisor = Math.floor(divisor/10);
    }
    
    return true;
}