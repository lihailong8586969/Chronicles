





SELECT
  A.`id`,
  A.`teacher_name`,
  B.`course_name`,

  C.`student_id`,
  C.`score`,

  D.`s_name`

FROM `sql`.`teacher` A LEFT JOIN `sql`.`course` B ON A.`id` = B.`teacher_id`
  LEFT JOIN `sql`.`score` C ON B.`id` = C.`course_id`
  LEFT JOIN `sql`.`student` D ON C.`student_id` = D.`id`
WHERE A.`teacher_name` = '张三'
ORDER BY C.`score` DESC
LIMIT 1;





