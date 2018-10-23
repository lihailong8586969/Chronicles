




SELECT

  A.id,
  A.student_id,


  B.course_id course_1,
  B.score     course_1_score,

  C.course_id course_2,
  C.score     course_2_score,

  D.course_id course_3,
  D.score     course_3_score ,


  AVG(E.score) avg_score
FROM (

    (SELECT *
     FROM `sql`.`score`) A

    LEFT JOIN

    (SELECT *
     FROM `sql`.`score`
     WHERE `course_id` = 1) B ON A.`student_id` = B.`student_id`

    LEFT JOIN

    (SELECT *
     FROM `sql`.`score`
     WHERE `course_id` = 2) C ON A.`student_id` = C.`student_id`

    LEFT JOIN

    (SELECT *
     FROM `sql`.`score`
     WHERE `course_id` = 3) D ON A.`student_id` = D.`student_id`
) LEFT JOIN `sql`.`score` E ON A.`student_id` = E.`student_id`

  # 注意 , 这里如果取别名则sql报错 , 这和sql的执行顺序有关(GroupBy，Having，Where，Orderby) ,
  # 一个SQL语句往往会产生多个临时视图，那么这些关键字的执行顺序就非常重要了，
  # 因为你必须了解这个关键字是在对应视图形成前的字段进行操作还是对形成的临时视图进行操作，
  # 这个问题在使用了别名的视图尤其重要。
  # ) AS F LEFT JOIN `sql`.`score` E ON A.`student_id` = E.`student_id`

GROUP BY A.`student_id` ORDER BY avg_score DESC ;



