




SELECT
  *,
  CHAR_LENGTH(`course_id`)                          'length_all',
  CHAR_LENGTH(SUBSTRING_INDEX(`course_id`, ',', 1)) 'length_first'
FROM (

       SELECT
         GROUP_CONCAT(`id`)         'id',
         GROUP_CONCAT(`student_id`) 'student_id',
         GROUP_CONCAT(`course_id`)  'course_id',
         `score`
       FROM `sql`.`score`
       GROUP BY `score`
       ORDER BY `score` DESC
     ) A
WHERE CHAR_LENGTH(`course_id`) > CHAR_LENGTH(SUBSTRING_INDEX(`course_id`, ',', 1));




