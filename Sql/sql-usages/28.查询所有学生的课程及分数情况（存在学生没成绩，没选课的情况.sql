




SELECT
  A.`id`,
  A.`s_name`,
  B.`course_id`,
  B.`score`
FROM `sql`.`student` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`student_id`



