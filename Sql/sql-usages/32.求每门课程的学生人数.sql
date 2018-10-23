



SELECT
  A.`id`,
  A.`course_id`,
  B.`course_name`,
  COUNT(`student_id`) '学生人数'
FROM `sql`.`score` A LEFT JOIN `sql`.`course` B ON A.`course_id` = B.`id`
GROUP BY `course_id`;










