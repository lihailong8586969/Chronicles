






(SELECT
   A.`s_name`,
   A.`sex`,
   COUNT(A.`s_name`) 'count'
 FROM (SELECT *
       FROM `sql`.`student`
       WHERE `sex` = '男') A
 GROUP BY A.`s_name`
 HAVING (`count`) > 1)

UNION

(SELECT
   A.`s_name`,
   A.`sex`,
   COUNT(A.`s_name`) 'count'
 FROM (SELECT *
       FROM `sql`.`student`
       WHERE `sex` = '女') A
 GROUP BY A.`s_name`
 HAVING (`count`) > 1);





