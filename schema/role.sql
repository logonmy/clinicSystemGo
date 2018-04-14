--用户角色表
CREATE TABLE role
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) UNIQUE NOT NULL,--名称
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
);

