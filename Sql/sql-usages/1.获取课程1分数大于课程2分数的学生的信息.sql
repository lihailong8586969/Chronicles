SELECT
  A.`id`,
  A.`student_id`,
  # A.`course_id` `course_id_1`,
  A.`score` `course_id_1_score`,

  # B.`course_id` `course_id_2`,
  B.`score` `course_id_2_score`,
  `course_table_first`.`course_name`  `course_id_1_name`,
  `course_table_second`.`course_name` `course_id_2_name`,

  `student`.`s_name`
FROM (
    (SELECT *
     FROM `score`
     WHERE `course_id` = 1) A LEFT JOIN
    (SELECT
       `student_id`,
       `course_id`,
       `score`
     FROM `score`
     WHERE `course_id` = 2) B ON A.student_id = B.`student_id`
  ) LEFT JOIN `course` `course_table_first` ON A.`course_id` = `course_table_first`.`id`
  LEFT JOIN `course` `course_table_second` ON B.`course_id` = `course_table_second`.`id`
  LEFT JOIN `student` ON A.`student_id` = `student`.`id`
WHERE A.`score` > B.`score`