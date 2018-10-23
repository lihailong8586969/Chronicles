

-- before optimize
SELECT * FROM A WHERE `name` IS NOT NULL ;
SELECT * FROM A WHERE `id` > 0 AND `id` IS NOT NULL ;

-- 分析 :
-- 设置字段 NOT NULL , 设置默认值 , 查的时候不要用 IS NOT NULL


-- optimize 1
SELECT * FROM A WHERE `id` > 0 ; // 设置了id字段不能为NULL且默认值为0 , 而不是可以为NULL



