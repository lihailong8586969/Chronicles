




SELECT
  A.`id`,
  A.`s_name`,
  B.`course_id`,
  AVG(B.`score`) 'avg_score'

FROM `sql`.`student` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`student_id`
GROUP BY A.`id`
HAVING `avg_score` >= 85
ORDER BY `avg_score` DESC;





