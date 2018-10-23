


SELECT * FROM `sql`.`student` A WHERE EXISTS(

  SELECT * FROM `sql`.`score` B WHERE A.`id` = B.`student_id`

);






-- 第 2 种
-- 选出所有记录 : student_id 字段的值不同
SELECT DISTINCT
  `student_id` 'id',
  B.s_name,
  B.`birth`,
  B.`sex`
FROM `sql`.`score` A LEFT JOIN `sql`.`student` B ON A.`student_id` = B.`id`;



