




SELECT IF(
           (SELECT MONTH(CURRENT_DATE) < date_format(`birth`, '%m'))
           OR
           ((SELECT MONTH(CURRENT_DATE) = date_format(`birth`, '%m')) AND
            (SELECT DAY(CURRENT_DATE) < date_format(`birth`, '%d'))),

           @age := (SELECT YEAR(CURRENT_DATE) - date_format(`birth`, '%Y')) - 1,
           @age := (SELECT YEAR(CURRENT_DATE) - date_format(`birth`, '%Y'))
       ) AS 'age'
FROM `sql`.`student`;







