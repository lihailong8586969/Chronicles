

SELECT

  A.id,
  A.`student_id`,

  B.`s_name`,
  B.`sex`

FROM `score` A
  LEFT JOIN `student` B ON A.`student_id` = B.`id`
WHERE `student_id` != 1
GROUP BY `student_id`
HAVING count(`course_id`) = (

  SELECT COUNT(*)
  FROM `score`
  WHERE `student_id` = 1

);

