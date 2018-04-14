--诊所
CREATE TABLE clinic
(
  code varchar(20) PRIMARY KEY NOT NULL,
  --编码
  name varchar(40) NOT NULL,
  --名称
  responsible_person varchar(40) NOT NULL,
  --负责人
  area varchar(40),
  --地区
  status boolean NOT NULL DEFAULT true,
  --是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP
);

