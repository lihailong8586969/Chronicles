


-- OR 和 IN 都是一样的( OR 比 IN 好一点 , OR 会使用索引 )


-- before optimize
SELECT * FROM A WHERE `id` IN (10,28) ;



-- 分析 :
-- IN 不会使用索引 , 所以不要用 IN


-- optimize 1
-- OR 条件如果较少 , 用 OR 或者 UNION
-- ① : 用 OR
SELECT * FROM A WHERE `id` = 10 OR `id` = 28 ;

-- ② : 用 UNION
SELECT * FROM A WHERE `id` = 10
UNION ALL
SELECT * FROM A WHERE `id` = 28 ;



-- optimize 2
-- OR 条件如果较多 , 且 A表 和 B表 有相关的联系 , 用 EXISTS 或者 JOIN 优化
-- ① : 用 EXISTS
SELECT * FROM A WHERE EXISTS (

  SELECT * FROM B WHERE A.id = B.id
) ;

-- ② : 用 JOIN
SELECT A.* FROM A LEFT JOIN B ON A.id = B.id ;





-- optimize 3
-- OR 条件如果很多 , 而且表之间没有任何关系 , 用 存储过程 + 临时表 优化
-- 步骤 :
-- 先创建临时表 , 把数据插入到临时表中 , 然后这两张表就会有关系了 , 再用 EXISTS 或者 JOIN 查询即可
-- 其实也可以不用 存储过程 , 直接拼 SQL 就行 , 不过 SQL 可能会很长
CREATE TABLE temp_table_01{

    `id` PRIMARY KEY AUTO_INCREMENT,
    `value` INT
}




