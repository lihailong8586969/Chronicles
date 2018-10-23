


SELECT
  A.*,
  B.`student_id`,

  COUNT(A.`id`) 'selected_course_count'
FROM `sql`.`student` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`student_id`
GROUP BY A.`id` HAVING COUNT(A.`id`) = 2 ;




