
INSERT INTO forum.community (community_id, community_name, introduction, create_time, update_time)
VALUES
(1, 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10'),
(2, 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00'),
(3, 'CS:GO', 'Rush B...', '2018-08-07 08:30:00', '2018-08-07 08:30:00'),
(4, 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');

SELECT * FROM "forum"."community";

SELECT * FROM "forum"."community" LIMIT 10 OFFSET 40;

SELECT count(*) FROM "forum"."community";



DELETE FROM forum.community
WHERE id IN (
  SELECT id FROM forum.community
  ORDER BY id
  LIMIT 100
);
