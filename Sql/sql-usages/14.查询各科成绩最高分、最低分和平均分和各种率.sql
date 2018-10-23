

-- 题目如下
--    查询各科成绩最高分、最低分和平均分：
--    以如下形式显示：课程 ID，课程 name，最高分，最低分，平均分，及格率，中等率，优良率，优秀率
--    及格为>=60，中等为：70-80，优良为：80-90，优秀为：>=90
--    要求输出课程号和选修人数，查询结果按人数降序排列，若人数相同，按课程号升序排列











SELECT C.* ,

  COUNT(`score`) `learned_student_count`,
  AVG(`score`) '平均分',
  MAX(`score`) '最高分',
  MIN(`score`) '最低分',
  ( ( COUNT( `score` >= 90 OR NULL ) / COUNT(`score`) ) ) '优秀率',
  ( ( COUNT( (`score` >= 80 AND `score` < 90) OR NULL ) / COUNT(`score`) ) * 100 ) '优良率',
  ( ( COUNT( (`score` >= 70 AND `score` < 80) OR NULL ) / COUNT(`score`) ) * 100 ) '中等率',
  ( ( COUNT( `score` >= 60 OR NULL ) / COUNT(`score`) ) * 100 ) '及格率'
FROM (

  SELECT

  B.`course_id`,
  A.`course_name`,
  A.`teacher_id`,

  B.`score`
FROM `sql`.`course` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`course_id`

) C GROUP BY `course_id` ORDER BY `learned_student_count` ASC , `course_id` ASC ;










--------- 以下这个只是实验 ---------
SELECT C.* ,

  AVG(`score`) '平均分',
  MAX(`score`) '最高分',
  MIN(`score`) '最低分',
  COUNT( `score` >= 60 OR NULL ) '及格的学生数量',
  COUNT( (`score` >= 70 AND `score` < 80) OR NULL ) '中等的学生数量',
  COUNT( (`score` >= 80 AND `score` < 90) OR NULL ) '优良的学生数量',
  COUNT( `score` >= 90 OR NULL ) '优秀的学生数量'

FROM (

  SELECT

  B.`course_id`,
  A.`course_name`,
  A.`teacher_id`,

  B.`score`
FROM `sql`.`course` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`course_id`

) C GROUP BY `course_id`;







