--人员
CREATE TABLE personnel
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(10) NOT NULL,--编码
  name varchar(10) NOT NULL,--名称
  clinic_code varchar(20) NOT NULL references clinic(code),--所属诊所
  weight integer NOT NULL DEFAULT 1,--权重
  title varchar(10),--职称
  username varchar(20),--账号
  password varchar(40),--密码
  is_clinic_admin boolean NOT NULL DEFAULT false,--是否是诊所超级管理员
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  UNIQUE (code, clinic_code, is_clinic_admin)
);