
SELECT
  C.`student_id`,
  D.`s_name`,
  D.`sex`,

  C.`course_sum`
FROM (

       SELECT
         B.`student_id`,
         COUNT(B.`student_id`) course_sum

       FROM `course` A
         LEFT JOIN `score` B ON A.`id` = B.`course_id`
       GROUP BY B.`student_id`
       HAVING `course_sum` < 3
     ) C LEFT JOIN `student` D ON C.`student_id` = D.`id`;

