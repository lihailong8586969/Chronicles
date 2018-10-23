



SELECT
  A.`student_id`,

  B.`s_name`,

  AVG(`score`) avg_score

FROM `sql`.`score` A
  LEFT JOIN `sql`.`student` B ON A.`student_id` = B.`id`

WHERE `score` < 60
GROUP BY `student_id`
HAVING count(`course_id`) >= 2



