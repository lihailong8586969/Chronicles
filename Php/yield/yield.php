<?php

// 普通
function createRange($number){
    $data = [];
    for($i=0;$i<$number;$i++){
        $data[] = time();
    }
    return $data;
}




// 荷官( 赌场中发牌的人 )
function dealer($peopleNum){

    // 现在赌场中有 $peopleNum 个人在赌牌
    // 荷官就需要给这 $peopleNum 个人发牌 , 具体怎么发牌取决于下面的 foreach
    for ($i = 1; $i <= $peopleNum; $i++) {

        yield time();
    }
}

// 现在有1个荷官,需要给10个人发牌
$dealer = dealer(10); // 这里调用上面我们创建的函数

// 荷官已准备好发牌
foreach ($dealer as $card ) {

    // 要求每1秒发1次
    sleep(1); //这里停顿1秒

    // 荷官发出牌
    echo $card . '<br />';
}


