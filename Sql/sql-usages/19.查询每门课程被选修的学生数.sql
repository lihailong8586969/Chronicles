



SELECT
  A.`id`,
  A.`course_name`,
  COUNT(A.`id`) 'count'
FROM `sql`.`course` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`course_id`
GROUP BY A.`id`;




