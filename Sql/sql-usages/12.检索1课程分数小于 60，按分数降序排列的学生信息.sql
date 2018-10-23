



SELECT

  B.`id`,
  B.`s_name`,
  A.`score`,
  A.`course_id`

FROM `sql`.`score` A

  LEFT JOIN `sql`.`student` B ON A.`student_id` = B.`id`

WHERE `score` < 60 AND `course_id` = 1
ORDER BY `score` DESC;



