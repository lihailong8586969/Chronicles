




SELECT
  A.`course_name`,
  B.`student_id`,
  B.`score`,
  C.`s_name`
FROM `sql`.`course` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`course_id`

  LEFT JOIN `sql`.`student` C ON B.`student_id` = C.`id`

WHERE A.`course_name` = '数学' AND B.`score` < 60;






