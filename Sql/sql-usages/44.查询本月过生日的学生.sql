



SET @@time_zone = "PRC";

SELECT
  *
FROM `sql`.`student`
WHERE MONTH(`birth`) = MONTH(NOW());


