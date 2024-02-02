-- Add down migration script here
DROP TYPE forum.gender;
DROP TABLE forum.user CASCADE;
