






SET @curRank := 0; -- 设置当前的排名变量为 0
SELECT
  *,
  @curRank := @curRank + 1 AS `rank`
FROM (
  SELECT
    A.`student_id`,
    SUM(`score`) `score_sum`

    -- @curRank := @curRank + 1 AS `rank` 如果写在这里 , rank 排名就会错误 。
    -- 这还是和 SQL 的执行顺序有关
  FROM `sql`.`score` A
  GROUP BY `student_id`
  ORDER BY `score_sum` DESC
) B ;






