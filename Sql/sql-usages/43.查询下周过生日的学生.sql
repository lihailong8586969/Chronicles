



SET @@time_zone = "PRC";

SELECT
  *
FROM `sql`.`student`
WHERE WEEK(date_format(`birth`, '%Y-%m-%d')) = (WEEK(NOW()) + 1);




