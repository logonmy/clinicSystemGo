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
  id serial PRIMARY KEY NOT NULL, --编号
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
  register_type INTEGER NOT NULL  DEFAULT 1,--登记类型：1线下分诊，2预约
  triage_time timestamp DEFAULT LOCALTIMESTAMP,--分诊完成时间 或 报道时间
  reception_time timestamp DEFAULT LOCALTIMESTAMP,--接诊时间
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  UNIQUE (department_id, clinic_patient_id,treat_status,visit_date)--联合主键，就诊人、科室、状态、日期唯一
);

--预约记录
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
  visit_place VARCHAR(20),--就诊地址
  operation_id integer REFERENCES personnel(id),--操作人编码
  created_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP
);

--挂号记录
CREATE TABLE registration
(
  id serial PRIMARY KEY NOT NULL, --编号
  clinic_patient_id integer NOT NULL references clinic_patient(id),--患者id
  department_id integer NOT NULL REFERENCES department(id),--科室id
  clinic_triage_patient_id integer,--分诊就诊人id
  personnel_id integer NOT NULL REFERENCES personnel(id),--医生id
  visit_date DATE NOT NULL,--出诊日期
  am_pm CHAR(1) NOT NULL CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  is_today boolean NOT NULL DEFAULT false,--是否当日号
  visit_type_code integer NOT NULL REFERENCES visit_type(code),--出诊类型编码
  status VARCHAR(2) NOT NULL DEFAULT '01',--就诊状态
  visit_place VARCHAR(20),--就诊地址
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

--用户登录记录
CREATE TABLE personnel_login_record
(
  id serial PRIMARY KEY NOT NULL,--id
  personnel_id INTEGER NOT NULL references personnel(id),--人员id
  ip VARCHAR(20) NOT NULL,--角色id
  remark VARCHAR(100),
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
);

--体征
CREATE TABLE body_sign
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL UNIQUE references clinic_triage_patient(id),--分诊记录id
  weight FLOAT,--体重(kg)
  height FLOAT,--升高（m）
  bmi FLOAT,--体重（千克）/（身高（米）*身高（米））
  blood_type VARCHAR(2) CHECK(blood_type = 'A' OR blood_type = 'B' OR blood_type = 'O' OR blood_type = 'AB' OR blood_type = 'UC' ),--血型 uc: 未查
  rh_blood_type integer CHECK(rh_blood_type = 0 or rh_blood_type = 1),--RH血型 0: 阴性，1阳性
  temperature_type integer CHECK(temperature_type >0 AND temperature_type <6),--RH血型 1: 口温，2：耳温，3：额温，4：腋温，5：肛温
  temperature FLOAT,--温度
  breathe integer,--呼吸(次/分钟)
  pulse integer,--脉搏(次/分钟)
  systolic_blood_pressure integer,--血压收缩压
  diastolic_blood_pressure integer,--血压舒张压
  blood_sugar_time varchar(20),--血糖时间
  blood_sugar_type integer,--血糖时段类型 1：睡前，2，晚餐后，3晚餐前，4午餐后，5，午餐前，6，早餐后，7早餐前，8：凌晨，9空腹
  blood_sugar_concentration FLOAT,--血糖浓度(mmol/I)
  left_vision varchar(5),--左眼视力
  right_vision varchar(5),--右眼视力
  oxygen_saturation FLOAT,--氧饱和度(%)
  pain_score integer CHECK(pain_score>-1 AND pain_score<11),--疼痛评分
  remark VARCHAR(100),
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
)

--费用项目类型
CREATE TABLE charge_project_type
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) UNIQUE NOT NULL,--名称
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
);

--诊前病历
CREATE TABLE pre_medical_record
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL UNIQUE references clinic_triage_patient(id),--分诊记录id
  has_allergic_history boolean,--是否有过敏
  allergic_history text,--过敏史
  personal_medical_history text,--个人病史
  family_medical_history text,--家族病史
  vaccination text,--接种疫苗
  menarche_age integer,--月经初潮年龄
  menstrual_period_start_day varchar(10),--月经经期开始时间
  menstrual_period_end_day varchar(10),--月经经期结束时间
  menstrual_cycle_start_day varchar(10),--月经周期结束时间
  menstrual_cycle_end_day varchar(10),--月经周期结束时间
  menstrual_last_day varchar(10),--末次月经时间
  gestational_weeks integer,--孕周
  childbearing_history text,--生育史
  remark VARCHAR(100),
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
);

--诊前欲诊
CREATE TABLE pre_diagnosis
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL UNIQUE references clinic_triage_patient(id),--分诊记录id
  chief_complaint text,--主诉
  history_of_rresent_illness text,--现病史
  history_of_past_illness text,--既往史
  physical_examination text,--体格检查
  remark text,--备注
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp
);



