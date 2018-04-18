--就诊人
CREATE TABLE patient
(
  id serial PRIMARY KEY NOT NULL,--id
  cert_no varchar(18) UNIQUE NOT NULL,--身份证号
  name varchar(10) NOT NULL,--姓名
  birthday varchar(8) NOT NULL,--身份证号
  sex integer NOT NULL CHECK(sex = 0 OR sex = 1),--性别 0：女，1：男
  phone varchar(11) not Null,--手机号
  patient_channel_id INTEGER NOT Null references patient_channel(id),
  address varchar(40),--住址
  profession varchar(40),--职业
  remark varchar(200),--备注
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
  );