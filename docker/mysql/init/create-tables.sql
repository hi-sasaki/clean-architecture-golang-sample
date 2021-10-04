drop table if exists users;
CREATE TABLE users (
  id VARCHAR(45) NOT NULL,
  first_name VARCHAR(45) NULL,
  last_name VARCHAR(45) NULL,
  password  VARCHAR(45) NOT NULL NULL,
  PRIMARY KEY (id));