--诊所管理员
CREATE TABLE admin
(
  username varchar(20) NOT NULL,
  phone varchar(11) NOT NULL,
  password varchar(40) NOT NULL,
  clinic_code varchar(40) NOT NULL references clinic(code),
  status boolean NOT NULL DEFAULT true,
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  PRIMARY KEY (username, clinic_code)
);