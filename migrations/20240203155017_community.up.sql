-- 在forum schema下创建表
CREATE TABLE IF NOT EXISTS forum.community (
    id SERIAL PRIMARY KEY,
    community_id INT NOT NULL,
    community_name VARCHAR(128) NOT NULL,
    introduction TEXT NOT NULL,
    create_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 在forum schema下创建函数
CREATE OR REPLACE FUNCTION forum.update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_time = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 在forum schema下创建触发器
CREATE TRIGGER update_community_modtime
BEFORE UPDATE ON forum.community
FOR EACH ROW
EXECUTE FUNCTION forum.update_modified_column();
