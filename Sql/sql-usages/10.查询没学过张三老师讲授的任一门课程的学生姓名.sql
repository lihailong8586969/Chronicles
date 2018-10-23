


# 笔记 : count(field) 计算时 ，未匹配的field(field显示为NULL的)不会记录在内

SELECT


  D.`student_id`,

  COUNT(C.`course_id`) learn_zhangsan_course_count
FROM (

       SELECT B.`id` course_id
       FROM (

              SELECT *
              FROM `sql`.`teacher`
              WHERE `teacher_name` = '张三'

            ) A LEFT JOIN `sql`.`course` B ON A.`id` = B.`teacher_id`

     ) C RIGHT JOIN `sql`.`score` D ON C.`course_id` = D.`course_id`
GROUP BY D.`student_id`
HAVING learn_zhangsan_course_count < 1;


