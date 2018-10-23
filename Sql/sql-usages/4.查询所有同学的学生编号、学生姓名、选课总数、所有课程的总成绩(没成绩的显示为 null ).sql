
--SELECT
--  A.id,
--  A.`student_id`,
--  B.s_name,
--
--  A.`course_count`,
--  A.score_sum
--FROM (
--
--       SELECT
--         `id`,
--         `student_id`,
--         COUNT(`student_id`) course_count,
--         SUM(`score`)        score_sum
--       FROM `sql`.`score`
--       GROUP BY `student_id`
--
--     ) A LEFT JOIN `sql`.`student` B ON A.`student_id` = B.`id`;


# mysql 如果未配置 ，不能使用 GROUP BY
SET @@GLOBAL.sql_mode = '';
SET sql_mode = 'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

# SELECT @@sql_mode;
# SELECT @@GLOBAL.sql_mode;



# 答案不唯一 ，因为题目没说清都要获取哪些数据 ，反正只要完成 " 查有成绩的学生信息 " 即可

SELECT
  A.`id`,
  A.`s_name`,
  count(`course_id`), -- count 计算某个字段的总记录数
  SUM(`score`) `score_sum` -- sum 就算所有记录在某个字段上的值的总和
FROM `sql`.`student` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`student_id` GROUP BY `id`;



SELECT
  A.id,
  A.s_name
FROM `sql`.`student` A
  RIGHT JOIN `sql`.`score` B ON A.`id` = B.`student_id` GROUP BY `id`;
