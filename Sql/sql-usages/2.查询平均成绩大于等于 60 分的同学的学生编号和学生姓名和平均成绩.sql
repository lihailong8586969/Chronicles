
# mysql 如果未配置 ，不能使用 GROUP BY
set @@GLOBAL.sql_mode='';
set sql_mode ='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

# SELECT @@sql_mode;
# SELECT @@GLOBAL.sql_mode;



SELECT * , AVG(`score`) FROM `sql`.`score` GROUP BY `student_id` HAVING AVG(`score`) >= 60  ;



