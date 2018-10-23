

SELECT DISTINCT `student_id`
FROM `score`
WHERE `course_id`
      IN (SELECT `course_id`
          FROM `score`
          WHERE `student_id` = 1) AND `student_id` != 1;



