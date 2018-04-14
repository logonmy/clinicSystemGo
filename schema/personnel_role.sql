--用户角色表
CREATE TABLE role
(
  personnel_code varchar(10) NOT NULL references personnel(code),--人员编码
  clinic_code varchar(40) NOT NULL references clinic(code),--诊所编码
  role_id INTEGER NOT NULL references role(id),--角色id
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  PRIMARY KEY (personnel_code, clinic_code, role_id)--联合主键，人员编码、诊所编码、 角色编码唯一
);

