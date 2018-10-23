





SELECT
  A.`student_id`,
  B.`s_name`,
  A.`course_id`,
  A.`score`
FROM `sql`.`score` A

  LEFT JOIN `sql`.`student` B ON A.`student_id` = B.`id`

WHERE A.`course_id` = 1 AND A.`score` >= 80





