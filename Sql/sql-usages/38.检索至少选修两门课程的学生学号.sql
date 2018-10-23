








SELECT
  B.*,
  C.`s_name`,
  C.`birth`,
  C.`sex`
FROM (

       SELECT
         A.`student_id`,
         COUNT(A.`course_id`) 'selected_course_count'
       FROM `sql`.`score` A
       GROUP BY A.`student_id`
       HAVING COUNT(A.`course_id`) >= 2

     ) B LEFT JOIN `sql`.`student` C ON B.`student_id` = C.`id`;



