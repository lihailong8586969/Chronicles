



SET @preScore := NULL; -- 设置变量 , 记录上一个分数
SET @curRank := 0;     -- 设置变量  ,记录当前的排名变量
SELECT * ,
CASE
WHEN @preScore = `score` THEN @curRank
WHEN @preScore := `score` THEN @curRank := @curRank + 1
END AS rank

FROM `sql`.`score`
ORDER BY `score` DESC;

