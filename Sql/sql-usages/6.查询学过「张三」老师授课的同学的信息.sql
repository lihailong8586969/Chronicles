# mysql 如果未配置 ，不能使用 GROUP BY
SET @@GLOBAL.sql_mode = '';
SET sql_mode = 'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

# SELECT @@sql_mode;
# SELECT @@GLOBAL.sql_mode;



SELECT
  D.`id`,

  C.`course_id`,
  C.`course_name`,
  C.`teacher_id`,
  C.`teacher_name`,

  D.`student_id`,
  E.s_name,
  E.sex,
  D.`score`

FROM (

       SELECT
         A.`id` `course_id`,
         A.`course_name`,
         A.`teacher_id`,
         B.`teacher_name`
       FROM `course` A LEFT JOIN `teacher` B ON A.`teacher_id` = B.`id`
       WHERE B.`teacher_name` LIKE '张三%'

     ) C LEFT JOIN `score` D ON C.`course_id` = D.`course_id`
  LEFT JOIN `student` E ON D.student_id = E.id;