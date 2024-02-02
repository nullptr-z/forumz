-- Add down migration script here

DROP EXTENSION btree_gist;
DROP TABLE forum.user CASCADE;
