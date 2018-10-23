



SELECT
  A.`id`,
  A.`course_name`,
  group_concat(B.`student_id`),
  group_concat(B.`score` ORDER BY B.`score` DESC ),

  C.`s_name`
FROM `sql`.`course` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`course_id`
  LEFT JOIN `sql`.`student` C ON B.`student_id` = C.`id` GROUP BY A.`id` ORDER BY B.`score` DESC ;







