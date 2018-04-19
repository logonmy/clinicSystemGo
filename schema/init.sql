--诊所
CREATE TABLE clinic
(
  id serial PRIMARY KEY NOT NULL,--编码
  code varchar(20) UNIQUE NOT NULL,--编码
  name varchar(40) NOT NULL,--名称
  responsible_person varchar(40) NOT NULL,--负责人
  area varchar(40),--地区
  phone varchar(11),--手机号
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
);

--科室
CREATE TABLE department
(
	id serial PRIMARY KEY NOT NULL,--id
	code varchar(20) NOT NULL,--科室编码
	name varchar(20) NOT NULL,--科室名称
	clinic_id integer NOT NULL references clinic(id),--所属诊所
	status boolean NOT NULL DEFAULT true,--是否启用
	weight integer NOT NULL DEFAULT 1,--权重
	is_appointment boolean NOT NULL DEFAULT true,--是否开放预约/挂号
	created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
	updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
	deleted_time timestamp,
	UNIQUE (code, clinic_id)
);

--人员
CREATE TABLE personnel
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(10) NOT NULL,--编码
  name varchar(10) NOT NULL,--名称
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  weight integer NOT NULL DEFAULT 1,--权重
  phone varchar(11),--手机号
  title varchar(10),--职称
  username varchar(20) UNIQUE,--账号
  password varchar(40),--密码
  is_appointment boolean NOT NULL DEFAULT true,--是否开放预约/挂号
  is_clinic_admin boolean NOT NULL DEFAULT false,--是否是诊所超级管理员
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  UNIQUE (code, clinic_id)
);

--人员科室关系表
CREATE TABLE department_personnel
(
  id serial PRIMARY KEY NOT NULL, --id
  department_id INTEGER NOT NULL references department(id),--科室id
  personnel_id INTEGER NOT NULL references personnel(id),--人员id
  type INTEGER NOT NULL CHECK(type > 0 AND type < 3),-- 关系类型 1：人事科室， 2：出诊科室
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  UNIQUE (department_id, personnel_id, type)--人员id、科室id、 类别唯一
);

--出诊类型
CREATE TABLE visit_type
(
  code integer PRIMARY KEY NOT NULL,--编码
  name VARCHAR(10) NOT NULL,--出诊类型名称
  open_flag boolean NOT NULL DEFAULT true,--启用状态
  fee integer NOT NULL DEFAULT 0 CHECK(fee > 0),
  created_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP
);

--医生出诊排班
CREATE TABLE doctor_visit_schedule
(
  id serial PRIMARY KEY NOT NULL, --排班编号
  department_id integer REFERENCES department(id),--医生id
  personnel_id integer REFERENCES personnel(id),--科室id
  visit_date DATE NOT NULL,--出诊日期
  am_pm CHAR(1) NOT NULL CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  stop_flag boolean NOT NULL DEFAULT false,--停诊标识
  is_today boolean NOT NULL DEFAULT false,--是否当日号
  tatal_num integer NOT NULL DEFAULT 20,--总的接诊数
  left_num integer NOT NULL DEFAULT 20 CHECK(left_num <= tatal_num),--剩余接诊数
  visit_type_code integer NOT NULL REFERENCES visit_type(code),--出诊类型编码
  created_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP,
  UNIQUE (personnel_id,department_id,visit_date,am_pm,visit_type_code,is_today)
);

--医生出诊周排班模板
CREATE TABLE doctor_visit_schedule_model
(
  id serial PRIMARY KEY NOT NULL, --模板id
  department_id integer REFERENCES department(id),--医生id
  personnel_id integer REFERENCES personnel(id),--科室id
  weekday INTEGER NOT NULL CHECK(weekday BETWEEN -1 AND 7),--出诊 日期（周几，0 代表 周日，1 周一...）
  am_pm CHAR(1) NOT NULL CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  tatal_num INTEGER NOT NULL DEFAULT 20,--总的接诊数
  visit_type_code integer NOT NULL REFERENCES visit_type(code),--出诊类别编码
  created_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP,
  UNIQUE (department_id,personnel_id,weekday,am_pm,visit_type_code)
);

--就诊人来源
CREATE TABLE patient_channel
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(40) NOT NULL,--名称
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
);

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

--诊所就诊人
CREATE TABLE clinic_patient
(
  id serial PRIMARY KEY NOT NULL, --排班编号
  patient_id integer NOT NULL references patient(id),--患者身份证号
  clinic_id integer NOT NULL references clinic(id),--诊所编码
  personnel_id integer NOT NULL references personnel(id),--录入人员id
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  UNIQUE (patient_id, clinic_id)--联合主键，就诊人身份证号和诊所编码唯一
);

--分诊就诊人
CREATE TABLE clinic_triage_patient
(
  id serial PRIMARY KEY NOT NULL,--id
  department_id INTEGER NOT NULL references department(id),--科室id
  clinic_patient_id INTEGER NOT NULL references clinic_patient(id),--科室就诊人id
  register_personnel_id INTEGER NOT NULL references personnel(id),--录入人员id
  doctor_id INTEGER references personnel(id),--接诊医生医生id
  triage_personnel_id INTEGER references personnel(id),--分诊人员id
  treat_status boolean NOT NULL DEFAULT false,--是否分诊
  visit_date DATE NOT NULL DEFAULT CURRENT_DATE,--日期
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  UNIQUE (department_id, clinic_patient_id,treat_status,visit_date)--联合主键，就诊人、科室、状态、日期唯一
);

--挂号记录
CREATE TABLE appointment
(
  id serial PRIMARY KEY NOT NULL, --编号
  clinic_patient_id integer NOT NULL references clinic_patient(id),--患者id
  department_id integer NOT NULL REFERENCES department(id),--医生编码
  personnel_id integer NOT NULL REFERENCES personnel(id),--科室编码
  visit_date DATE NOT NULL,--出诊日期
  am_pm CHAR(1) NOT NULL CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  is_today boolean NOT NULL DEFAULT false,--是否当日号
  visit_type_code integer NOT NULL REFERENCES visit_type(code),--出诊类型编码
  status VARCHAR(2) NOT NULL DEFAULT '01',--就诊状态
  visit_place VARCHAR(20),--就诊未知
  sort_no SMALLINT,--就诊序号
  operation_id integer REFERENCES personnel(id),--操作人编码
  created_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP
);


--角色表
CREATE TABLE role
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) UNIQUE NOT NULL,--名称
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
);


--用户角色表
CREATE TABLE personnel_role
(
  personnel_id INTEGER NOT NULL references personnel(id),--人员id
  role_id INTEGER NOT NULL references role(id),--角色id
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  PRIMARY KEY (personnel_id, role_id)--联合主键，人员编码、角色编码唯一
);
