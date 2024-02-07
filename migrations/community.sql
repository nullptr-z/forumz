
INSERT INTO forum.community (community_id, community_name, introduction, create_time, update_time)
VALUES
(1, 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10'),
(2, 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00'),
(3, 'CS:GO', 'Rush B...', '2018-08-07 08:30:00', '2018-08-07 08:30:00'),
(4, 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');

SELECT * FROM "forum"."community";

SELECT * FROM "forum"."community" LIMIT 10 OFFSET 40 ;

SELECT count(*) FROM "forum"."community";


DELETE FROM forum.community
WHERE id IN (
  SELECT id FROM forum.community
  ORDER BY id
  LIMIT 100
);
-- ORDER BY avg(x)

-- CASE; 程序设计时应该避免存 null，存对应类型的默认值最好
SELECT  (CASE WHEN id is NULL then 0 ELSE id END) FROM "forum"."community" LIMIT 20 ;
-- DISTINCT去重
SELECT  DISTINCT introduction,introduction FROM "forum"."community";
-- WHERE
-- BETWEEN . and . 之间
-- is NOTNULL不等于 null
-- or 或，and 并且
-- in 多选
-- 不等于;
-- < 、> 、<>符号可用于时间
-- like 模糊查询；%任意，_一个字符，\%转义百分号
SELECT * FROM "forum"."community" WHERE id <> 71 and id BETWEEN 70 and 80
and id in(75,78) and create_time > '2018-08-07'  and community_name like '%code';
-- Order
-- DESC大到小；ASC小到大; 排序条件优先从左到右
SELECT * FROM "forum"."community" ORDER by community_id DESC, id ASC;

-------- 函数，有很多，列举几个--------
-- lower 转小写
SELECT *, lower(community_name) FROM "forum"."community";
-- concat 拼接字符串
SELECT id,community_name, concat(id,community_name) FROM "forum"."community";
-- char_length 字符串长度
SELECT community_name, char_length(community_name) FROM "forum"."community";
-- substring(0,5) 截取字符串
SELECT community_name, substring(community_name,0,5) FROM "forum"."community";
-- trim去除左右空格， ltrim去除左边空格，rtrim去除右边空格
SELECT trim('       ab c       '),ltrim('       ab c       ');
-- abs取绝对值
SELECT abs(-1);
-- ceil 向上取整
SELECT ceil(3.14);
-- floor 向下取整
SELECT floor(3.14);
-- round 保留小数位数，四舍五入
SELECT round(3.1415926,4);
-- round 保留小数位数，截断
SELECT trunc(3.1415926,4);
-- 当前时间
SELECT CURRENT_TIME;
SELECT now();
-- 获取时间的月份
SELECT MONTH(create_time) FROM "forum"."community";
-- 获取时间的年
SELECT YEAR(create_time) FROM "forum"."community";
SELECT weekday(create_time) FROM "forum"."community";

-------- 组函数，有很多，列举几个--------
-- max 最大值
SELECT max(community_id) FROM "forum"."community";
-- count 不统计 null
SELECT count(comm) FROM "forum"."community";
-- sum(x), avg(x)

-------- Group by --------
SELECT community_name,community_id FROM "forum"."community" GROUP by community_name,community_id
HAVING community_name = 'Go';
-- HAVING 专门用于 group，类似 where
-- HAVING avg(x)

-- 把查询结果当成临时的中间表, 再通过关联字段把多个表连接起来
-------- 表链接 JOIN ON --------
SELECT * from   "forum"."community" JOIN
-- 中间表，别名叫 tmp
(SELECT * FROM "forum"."community" LIMIT 10) as tmp  ON forum.community.id = tmp.id
WHERE forum.community.community_name = 'GO';

-------- 子查询 --------
SELECT * FROM "forum"."community" WHERE community_id= (SELECT max(community_id) FROM "forum"."community");
-- 大于平均 ID的 ID
SELECT * FROM "forum"."community" WHERE id > (SELECT avg(id) FROM "forum"."community");
