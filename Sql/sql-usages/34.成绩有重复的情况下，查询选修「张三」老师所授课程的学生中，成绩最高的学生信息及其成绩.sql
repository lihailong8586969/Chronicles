





-- 这道题有问题 ，题目没有说明张三老师是不是只教授 1 门课 ,
-- 虽然给的表中是只有 1 门 , 但是题目没说
-- 姑且认为张三老师只教授 1 门课吧


-- 第 1 种办法 :
SELECT
  A.`id`,
  A.`teacher_name`,
  B.`course_name`,

  C.`student_id`,
  C.`score`,

  D.`s_name`

FROM `sql`.`teacher` A LEFT JOIN `sql`.`course` B ON A.`id` = B.`teacher_id`
  LEFT JOIN `sql`.`score` C ON B.`id` = C.`course_id`
  LEFT JOIN `sql`.`student` D ON C.`student_id` = D.`id`
WHERE A.`teacher_name` = '张三' AND C.`score` = (

  SELECT MAX(`score`)
  FROM `sql`.`score`
  WHERE `course_id` = (

    SELECT A.`id`
    FROM `sql`.`course` A LEFT JOIN `sql`.`teacher` B ON A.`teacher_id` = B.`id`
    WHERE `teacher_name` = '张三'

  )

)
ORDER BY C.`score` DESC;





-- 第 2 种办法

-- 定义变量 , 把最高分存变量里面

SET @max_score = NULL;
SELECT @max_score := (SELECT  MAX(`score`)
  FROM `sql`.`score`
  WHERE `course_id` = (

    SELECT A.`id`
    FROM `sql`.`course` A LEFT JOIN `sql`.`teacher` B ON A.`teacher_id` = B.`id`
    WHERE `teacher_name` = '张三'

)) ;


SELECT
  A.`id`,
  A.`teacher_name`,
  B.`course_name`,

  C.`student_id`,
  C.`score`,

  D.`s_name`

FROM `sql`.`teacher` A LEFT JOIN `sql`.`course` B ON A.`id` = B.`teacher_id`
  LEFT JOIN `sql`.`score` C ON B.`id` = C.`course_id`
  LEFT JOIN `sql`.`student` D ON C.`student_id` = D.`id`
WHERE A.`teacher_name` = '张三' AND C.`score` = @max_score
ORDER BY C.`score` DESC;




