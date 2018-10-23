






-- 定义变量
SET @curRank := 0;


-- 变量不断 +1
SELECT * , @curRank := @curRank + 1 AS rank
FROM `sql`.`score`
ORDER BY `score` DESC;



