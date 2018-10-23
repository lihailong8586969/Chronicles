


# mysql 如果未配置 ，不能使用 GROUP BY
SET @@GLOBAL.sql_mode = '';
SET sql_mode = 'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

# SELECT @@sql_mode;
# SELECT @@GLOBAL.sql_mode;

SELECT *
FROM `sql`.`student`
  RIGHT JOIN `sql`.`score` ON `sql`.`student`.`id` = `sql`.`score`.`student_id`
GROUP BY `sql`.`student`.`id`;

