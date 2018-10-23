


SELECT
  C.`s_name`,
  A.`course_id`,
   B.`course_name`,
  A.`score`
FROM `sql`.`score` A LEFT JOIN `sql`.`course` B ON A.`course_id` = B.`id`

  LEFT JOIN `sql`.`student` C ON A.`student_id` = C.`id`

WHERE A.`score` >= 70;





