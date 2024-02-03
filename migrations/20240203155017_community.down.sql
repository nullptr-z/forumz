-- Add down migration script here
-- 删除触发器
DROP TRIGGER IF EXISTS update_community_modtime ON forum.community;

-- 删除函数
DROP FUNCTION IF EXISTS forum.update_modified_column();

-- 删除表
DROP TABLE IF EXISTS forum.community;
