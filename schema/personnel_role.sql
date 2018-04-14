--用户角色表
CREATE TABLE role
(
  personnel_code INTEGER NOT NULL,--人员编码
  clinic_code varchar(40) NOT NULL references clinic(code),--诊所编码
  role_id INTEGER NOT NULL references role(id),--角色id
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
);

