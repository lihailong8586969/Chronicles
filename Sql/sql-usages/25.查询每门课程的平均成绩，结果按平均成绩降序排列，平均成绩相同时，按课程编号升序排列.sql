





SELECT
  A.`id`,
  A.`course_name`,
  AVG(B.`score`) 'avg_score'
FROM `sql`.`course` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`course_id`
GROUP BY A.`id` ORDER BY `avg_score` DESC , A.`id` DESC ;






