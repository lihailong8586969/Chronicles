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

