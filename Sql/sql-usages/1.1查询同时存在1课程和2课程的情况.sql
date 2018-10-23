



# JOIN 即可 ( 全连接 )

SELECT
  A.id ,
  A.`student_id`,
  A.`course_id` `course_id_1` ,
  A.`score` `course_id_1_score` ,

  B.`course_id` `course_id_2` ,
  B.`score` `course_id_2_score`
FROM (

    (SELECT
       `id`,
       `student_id`,
       `course_id`,
       `score`
     FROM `sql`.`score`
     WHERE `course_id` = 1) A
    JOIN

    (SELECT
       `student_id`,
       `course_id`,
       `score`
     FROM `sql`.`score`
     WHERE `course_id` = 2) B ON A.`student_id` = B.`student_id`

);
