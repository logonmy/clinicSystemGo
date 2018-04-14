--诊所管理员
CREATE TABLE admin
(
  username varchar(20) NOT NULL,
  --用户名
  phone varchar(11) NOT NULL,
  --手机号
  email VARCHAR(30),
  --邮箱
  password varchar(40) NOT NULL,
  --密码
  clinic_code varchar(40) NOT NULL references clinic(code),
  --诊所编码
  status boolean NOT NULL DEFAULT true,
  --是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  PRIMARY KEY (username, clinic_code)
  --联合主键，用户名和诊所编码唯一
);