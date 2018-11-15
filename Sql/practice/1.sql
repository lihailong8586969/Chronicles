# 1
Select * from a where a.HR_ID not in (select distinct HR_ID from b)

# 2
Select

a.*

from a

left join b

on a.HR_ID = b.HR_ID

Where b.HR_ID is null

# 3
Select

distinct

a.*

from a

left join b

on a.HR_ID = b.HR_ID

Where b.HR_ID is null  


# 4
Select b.HR_ID, count(distinct Promotion_ID)

From b

Group by 1

Order by 2 desc

Limit 1

# 5
Select b.HR_ID, count(distinct Promotion_ID) as promo_cnt

From b

Group by 1

Order by 2 desc

# 6
Select distinct c.HR_ID

From c

Where c.promo_cnt = (select max(promo_cnt) from c)

# 7
select * from table_name limit 3,1;                # 跳过前3条数据，从数据库中第4条开始查询，取一条数据，即第4条数据
select * from table_name limit 3 offset 1;      # 从数据库中的第2条数据开始查询3条数据，即第2条到第4条

# 8
select department_ id ,AVG(salary) from employees group by department_ id ;

# 9
select * from orders as e 
where (select count(distinct(e1.orderNo)) from orders as e1
 where  e1.city = e.city and e1.orderNo > e.orderNo) < 3 ;
# where (select count(distinct(e1.orderNo)) ???