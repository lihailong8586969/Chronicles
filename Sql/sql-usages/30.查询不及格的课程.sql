



SELECT
  A.`course_id`,
  A.`score`,
  B.`course_name`
FROM `sql`.`score` A LEFT JOIN `sql`.`course` B ON A.`course_id` = B.`id`
WHERE A.`score` < 60
GROUP BY A.`course_id`;






