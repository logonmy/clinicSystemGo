--诊所
CREATE TABLE clinic
(
  id serial PRIMARY KEY NOT NULL,--编码
  code varchar(20) UNIQUE NOT NULL,--编码
  name varchar(40) NOT NULL,--名称
  responsible_person varchar(40) NOT NULL,--负责人
  province varchar(30),--省
  city varchar(30),--市
  district varchar(30),--区
  area varchar(40),--详细地区
  phone varchar(11),--手机号
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
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
	created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
	updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
	deleted_time timestamp with time zone,
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
  image varchar(50),--头像
  detail text,--描述
  phone varchar(11),--手机号
  title varchar(10),--职称
  username varchar(20) UNIQUE,--账号
  password varchar(40),--密码
  is_appointment boolean NOT NULL DEFAULT true,--是否开放预约/挂号
  is_clinic_admin boolean NOT NULL DEFAULT false,--是否是诊所超级管理员
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
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
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (department_id, personnel_id, type)--人员id、科室id、 类别唯一
);

--医生出诊排班
CREATE TABLE doctor_visit_schedule
(
  id serial PRIMARY KEY NOT NULL, --排班编号
  department_id integer REFERENCES department(id),--医生id
  personnel_id integer REFERENCES personnel(id),--科室id
  visit_date DATE NOT NULL,--出诊日期
  am_pm varchar(1) NOT NULL CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  stop_flag boolean NOT NULL DEFAULT false,--停诊标识
  is_today boolean NOT NULL DEFAULT false,--是否当日号
  tatal_num integer NOT NULL DEFAULT 20,--总的接诊数
  left_num integer NOT NULL DEFAULT 20 CHECK(left_num <= tatal_num),--剩余接诊数
  created_time TIMESTAMP with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP with time zone,
  UNIQUE (personnel_id,department_id,visit_date,am_pm,is_today)
);

--医生出诊周排班模板
CREATE TABLE doctor_visit_schedule_model
(
  id serial PRIMARY KEY NOT NULL, --模板id
  department_id integer REFERENCES department(id),--医生id
  personnel_id integer REFERENCES personnel(id),--科室id
  weekday INTEGER NOT NULL CHECK(weekday BETWEEN -1 AND 7),--出诊 日期（周几，0 代表 周日，1 周一...）
  am_pm varchar(1) NOT NULL CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  tatal_num INTEGER NOT NULL DEFAULT 20,--总的接诊数
  created_time TIMESTAMP with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP with time zone,
  UNIQUE (department_id,personnel_id,weekday,am_pm)
);

--就诊人来源
CREATE TABLE patient_channel
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(40) NOT NULL,--名称
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--就诊人
CREATE TABLE patient
(
  id serial PRIMARY KEY NOT NULL,--id
  image varchar(50),
  cert_no varchar(18) UNIQUE,--身份证号
  name varchar(10) NOT NULL,--姓名
  birthday varchar(8) NOT NULL,--身份证号
  sex integer NOT NULL CHECK(sex = 0 OR sex = 1),--性别 0：女，1：男
  phone varchar(11) not NULL,--手机号
  patient_channel_id INTEGER NOT Null references patient_channel(id),
  province varchar(30),--省
  city varchar(30),--市
  district varchar(30),--区
  address varchar(40),--详细住址
  profession varchar(40),--职业
  remark varchar(200),--备注
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (name, birthday, phone)--唯一
);

--诊所就诊人
CREATE TABLE clinic_patient
(
  id serial PRIMARY KEY NOT NULL, --编号
  patient_id integer NOT NULL references patient(id),--患者身份证号
  clinic_id integer NOT NULL references clinic(id),--诊所编码
  personnel_id integer NOT NULL references personnel(id),--录入人员id
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (patient_id, clinic_id)--联合主键，就诊人身份证号和诊所编码唯一
);

--分诊就诊人
CREATE TABLE clinic_triage_patient
(
  id serial PRIMARY KEY NOT NULL,--id
  department_id INTEGER references department(id),--科室id
  clinic_patient_id INTEGER NOT NULL references clinic_patient(id),--科室就诊人id
  register_personnel_id INTEGER NOT NULL references personnel(id),--录入人员id
  doctor_id INTEGER references personnel(id),--接诊医生医生id
  triage_personnel_id INTEGER references personnel(id),--分诊人员id
  treat_status boolean NOT NULL DEFAULT false,--是否分诊
  visit_date DATE NOT NULL DEFAULT CURRENT_DATE,--日期
  register_type INTEGER NOT NULL,--登记类型：1预约，2线下分诊
  visit_type integer,--出诊类型 1: 首诊， 2复诊，3：术后复诊
  triage_time timestamp with time zone,--分诊完成时间 或 报道时间
  reception_time timestamp with time zone,--接诊时间
  complete_time timestamp with time zone,--完成时间
  cancelled boolean NOT NULL DEFAULT false,--取消标识
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (department_id, clinic_patient_id,treat_status,visit_date)--联合主键，就诊人、科室、状态、日期唯一
);

--挂号记录
CREATE TABLE registration
(
  id serial PRIMARY KEY NOT NULL, --编号
  clinic_patient_id integer NOT NULL references clinic_patient(id),--患者id
  department_id integer NOT NULL REFERENCES department(id),--科室id
  clinic_triage_patient_id integer UNIQUE NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  personnel_id integer NOT NULL REFERENCES personnel(id),--医生id
  visit_date DATE NOT NULL,--出诊日期
  am_pm varchar(1) NOT NULL CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  is_today boolean NOT NULL DEFAULT false,--是否当日号
  status VARCHAR(2) NOT NULL DEFAULT '01',--就诊状态
  visit_place VARCHAR(20),--就诊地址
  operation_id integer REFERENCES personnel(id),--操作人编码
  created_time TIMESTAMP with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP with time zone
);


--角色表
CREATE TABLE role
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) UNIQUE NOT NULL,--名称
  clinic_id integer NOT NULL references clinic(id),--诊所编码
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);


--用户角色表
CREATE TABLE personnel_role
(
  personnel_id INTEGER NOT NULL references personnel(id),--人员id
  role_id INTEGER NOT NULL references role(id),--角色id
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  PRIMARY KEY (personnel_id, role_id)--联合主键，人员编码、角色编码唯一
);

--用户登录记录
CREATE TABLE personnel_login_record
(
  id serial PRIMARY KEY NOT NULL,--id
  personnel_id INTEGER NOT NULL references personnel(id),--人员id
  ip VARCHAR(20) NOT NULL,--角色id
  remark VARCHAR(100),
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
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
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--费用项目类型
CREATE TABLE charge_project_type
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) UNIQUE NOT NULL,--名称
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--收费项目-诊疗
CREATE TABLE charge_project_treatment
(
  id serial PRIMARY KEY NOT NULL,--id
  project_type_id INTEGER NOT NULL references charge_project_type(id),--收费类型id
  name varchar(20) UNIQUE NOT NULL,--名称
  name_en varchar(20),--英文名称
  cost integer CHECK(cost > 0), --成本价
  fee integer NOT NULL CHECK(fee > 0), --销售价
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
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
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
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
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--门诊待缴费缴费
CREATE TABLE mz_unpaid_orders
(
  id serial PRIMARY KEY NOT NULL,--id
  registration_id INTEGER NOT NULL references registration(id),--分诊记录id
  charge_project_type_id INTEGER NOT NULL references charge_project_type(id),--收费类型id
  charge_project_id INTEGER NOT NULL,--收费项目id
  order_sn varchar(20) NOT NULL,--单号
  soft_sn INTEGER NOT NULL,--序号
  name varchar(20) NOT NULL,--收费名称
  price INTEGER NOT NULL CHECK(price > 0),--单价
  amount INTEGER NOT NULL CHECK(amount > 0),--数量
  unit varchar(20),--单位
  total INTEGER NOT NULL,--总价格
  discount INTEGER NOT NULL DEFAULT 0,--打折金额
  fee INTEGER NOT NULL,--缴费金额
  operation_id INTEGER NOT NULL references personnel(id),--操作员id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (order_sn, soft_sn)
);

--门诊缴费记录
CREATE TABLE mz_paid_record
(
  id serial PRIMARY KEY NOT NULL,--id
  registration_id INTEGER NOT NULL references registration(id),--分诊记录id
  out_trade_no varchar(20) UNIQUE,--第三方交易号
  soft_sns varchar(30) NOT NULL,--序号
  order_sn varchar(20) NOT NULL,--单号
  confrim_id INTEGER NOT NULL references personnel(id),--操作员id
  pay_type_code varchar(2) NOT NULL,--支付类型编码 01-医保支付，02-挂账金额，03-抵金券，04-积分
  pay_method_code varchar(2) NOT NULL,--支付方式编码 01-现金，02-微信，03-支付宝，04-银行卡

  status varchar(30) NOT NULL,--订单状态  --SUCCESS 缴费成功 --HALF-REFUND 半退费  --REFUND-SUCCESS 全退
  
  derate_money INTEGER NOT NULL DEFAULT 0 CHECK(derate_money >= 0),--减免金额
  voucher_money INTEGER NOT NULL DEFAULT 0 CHECK(voucher_money >= 0) ,--抵金券金额
  discount_money INTEGER NOT NULL DEFAULT 0 CHECK(discount_money >= 0),--折扣金额
  bonus_points_money INTEGER NOT NULL DEFAULT 0 CHECK(bonus_points_money >= 0) ,--积分兑换金额
  on_credit_money INTEGER NOT NULL DEFAULT 0 CHECK(on_credit_money >= 0) ,--挂账金额
  medical_money INTEGER NOT NULL DEFAULT 0 CHECK(medical_money >= 0),--医保金额
  total_money  INTEGER NOT NULL  ,--应收金额
  balance_money INTEGER NOT NULL ,--实收金额

  derate_money_refund INTEGER NOT NULL DEFAULT 0 CHECK(derate_money_refund <= 0),--减免金额（退）
  voucher_money_refund INTEGER NOT NULL DEFAULT 0 CHECK(voucher_money_refund <= 0),--抵金券金额（退）
  discount_money_refund INTEGER NOT NULL DEFAULT 0 CHECK(discount_money_refund <= 0),--折扣金额（退）
  bonus_points_money_refund INTEGER NOT NULL DEFAULT 0 CHECK(bonus_points_money_refund <= 0),--积分兑换金额（退）
  on_credit_money_refund INTEGER NOT NULL DEFAULT 0 CHECK(on_credit_money_refund <= 0),--挂账金额（退）
  medical_money_refund INTEGER NOT NULL DEFAULT 0 CHECK(medical_money_refund <= 0),--医保金额（退）
  total_money_refund  INTEGER NOT NULL DEFAULT 0  CHECK(total_money_refund <= 0),--应退金额（退）
  balance_money_refund INTEGER NOT NULL DEFAULT 0 CHECK(balance_money_refund <= 0),--实退金额（退）

  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--门诊缴费记录子表
CREATE TABLE mz_paid_record_detail
(
  id serial PRIMARY KEY NOT NULL,--id
  mz_paid_record_id INTEGER NOT NULL references mz_paid_record(id),
  out_trade_no varchar(20) UNIQUE,--第三方交易号
  refund_status boolean NOT NULL DEFAULT false,--退费标识
  soft_sns varchar(30) NOT NULL,--序号
  order_sn varchar(20) NOT NULL,--单号
  confrim_id INTEGER NOT NULL references personnel(id),--确认操作员id

  derate_money INTEGER NOT NULL DEFAULT 0 ,--减免金额
  voucher_money INTEGER NOT NULL DEFAULT 0 ,--抵金券金额
  discount_money INTEGER NOT NULL DEFAULT 0,--折扣金额
  bonus_points_money INTEGER NOT NULL DEFAULT 0 ,--积分兑换金额
  on_credit_money INTEGER NOT NULL DEFAULT 0 ,--挂账金额
  medical_money INTEGER NOT NULL DEFAULT 0 ,--医保金额
  total_money  INTEGER NOT NULL  ,--应收金额
  balance_money INTEGER NOT NULL ,--实收金额

  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--门诊已缴费缴费
CREATE TABLE mz_paid_orders
(
  id serial PRIMARY KEY NOT NULL,--id
  mz_paid_record_id INTEGER NOT NULL references mz_paid_record(id),
  registration_id INTEGER NOT NULL references registration(id),--分诊记录id
  charge_project_type_id INTEGER NOT NULL references charge_project_type(id),--收费类型id
  charge_project_id INTEGER NOT NULL,--收费项目id
  order_sn varchar(20) NOT NULL,--单号
  soft_sn INTEGER NOT NULL,--序号
  name varchar(20) NOT NULL,--收费名称
  price INTEGER NOT NULL CHECK(price > 0),--单价
  amount INTEGER NOT NULL CHECK(amount > 0),--数量
  unit varchar(20),--单位
  total INTEGER NOT NULL,--总价格
  discount INTEGER NOT NULL DEFAULT 0,--打折金额
  fee INTEGER NOT NULL,--缴费金额
  operation_id INTEGER NOT NULL references personnel(id),--未交费创建人id
  confrim_id INTEGER NOT NULL references personnel(id),--确认操作员id
  refund_id INTEGER references personnel(id), --退费操作员
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (order_sn, soft_sn)
);

--记录盈利详情，产生的支付和退费都得插入这个表
CREATE TABLE charge_detail
(
  id serial PRIMARY KEY NOT NULL,--id
  pay_record_id INTEGER NOT NULL,--缴费记录
  out_trade_no varchar(20),--第三方交易号
  out_refund_no varchar(20) UNIQUE,--第三方退费交易号
  in_out varchar(3) NOT NULL CHECK(in_out = 'in' OR in_out = 'out'),--收入或支出
  patient_id INTEGER references personnel(id),--关联的患者
  department_id INTEGER references department(id),--关联的科室
  doctor_id INTEGER references personnel(id),--关联的医生信息
  pay_type_code varchar(2) NOT NULL,--支付类型编码 01-门诊缴费，02-挂号费，03-挂账还款，04-住院缴费
  pay_type_code_name varchar(10) NOT NULL,--支付类型名称
  pay_method_code varchar(2) NOT NULL,--支付方式编码 01-现金，02-微信，03-支付宝，04-银行卡
  pay_method_code_name varchar(10) NOT NULL,--支付方式名称
  discount_money INTEGER NOT NULL DEFAULT 0 ,--折扣金额
  derate_money INTEGER NOT NULL DEFAULT 0 ,--减免金额
  medical_money INTEGER NOT NULL DEFAULT 0 ,--医保金额
  voucher_money INTEGER NOT NULL DEFAULT 0 ,--抵金券金额
  bonus_points_money INTEGER NOT NULL DEFAULT 0,--积分兑换金额
  on_credit_money INTEGER NOT NULL,--挂账金额
  total_money  INTEGER NOT NULL ,--应收金额
  balance_money INTEGER NOT NULL,--实收金额
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
  
);

--平台管理人员
CREATE TABLE admin
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(10) NOT NULL,--名称
  phone varchar(11),--手机号
  title varchar(10),--职称
  username varchar(20) UNIQUE,--账号
  password varchar(40),--密码
  is_clinic_admin boolean NOT NULL DEFAULT false,--是否是诊所超级管理员
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--一级菜单功能项
CREATE TABLE parent_functionMenu
(
  id serial PRIMARY KEY NOT NULL,--id
  url varchar(20) NOT NULL,--功能路由
  ascription VARCHAR(2) NOT NULL DEFAULT '01',--菜单所属类型 01 诊所 02 平台
  name varchar(20),--菜单名
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--二级菜单功能项
CREATE TABLE children_functionMenu
(
  id serial PRIMARY KEY NOT NULL,--id
  parent_functionMenu_id INTEGER references parent_functionMenu(id),--上级菜单id
  url varchar(20) NOT NULL,--功能路由
  name varchar(20),--菜单名
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊所菜单项
CREATE TABLE clinic_children_functionMenu
(
  id serial PRIMARY KEY NOT NULL,--id
  children_functionMenu_id INTEGER NOT NULL references children_functionMenu(id),--菜单id
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (children_functionMenu_id, clinic_id)
);

--诊所角色菜单项
CREATE TABLE role_clinic_functionMenu
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_children_functionMenu_id INTEGER NOT NULL references clinic_children_functionMenu(id),--诊所菜单id
  role_id integer NOT NULL references role(id),--所属角色
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (clinic_children_functionMenu_id, role_id)
);

--平台管理员菜单项
CREATE TABLE admin_functionMenu
(
  id serial PRIMARY KEY NOT NULL,--id
  children_functionMenu_id INTEGER NOT NULL references children_functionMenu(id),--平台菜单id
  admin_id integer NOT NULL references admin(id),--所属平台管理员
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (children_functionMenu_id, admin_id)
);

--用药频率
CREATE TABLE frequency
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(20),--频率码
  name varchar(20) NOT NULL,--名称
  py_code varchar(10),--拼音码
  define_code varchar(10),--自定义码
  print_code varchar(10),--打印名称
  week_day_flag integer,--周日标志
  update_flag integer,--允许修改标志
  delete_flag integer,--删除标志
  weight integer,--排序码/权重
  in_out_flag integer,--门诊住院标记
  medical_bill_code integer,--医保账单码
  doctor_flag integer,--医生标记
  input_frequency varchar(20),--护士录入频率
  times integer,--次数
  days integer,--天数
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--药品包装单位表
CREATE TABLE dose_unit
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(20),--编码
  name varchar(20) NOT NULL,--名称
  py_code varchar(10),--拼音码
  d_code varchar(10),--简码
  deleted_flag integer,--删除标志
  change_flag integer,--修改标志
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--药物类型
CREATE TABLE drug_class
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(20),--编码
  name varchar(30) NOT NULL,--药品名称
  py_code varchar(20),--拼音码
  d_code varchar(20),--简码
  drug_kind integer,--什么？
  extend_code varchar(10),--什么？
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--药品类别
CREATE TABLE drug_type
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(20),--编码
  name varchar(30) NOT NULL,--药品名称
  py_code varchar(20),--拼音码
  d_code varchar(20),--简码
  deleted_flag integer,--删除标志
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--药品剂型
CREATE TABLE dose_form
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(20),--编码
  name varchar(30) NOT NULL,--药品名称
  py_code varchar(20),--拼音码
  d_code varchar(20),--简码
  deleted_flag integer,--删除标志
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--药品生产厂商
CREATE TABLE manu_factory
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(20),--编码
  name varchar(20) NOT NULL,--厂商名称
  abbr_name varchar(30),--
  zip_code varchar(10),--
  address varchar(40),--地址
  tel varchar(20),--电话
  py_code varchar(20),--拼音码
  d_code varchar(20),--简码
  product_range varchar(20),--
  comment varchar(40),--
  deleted_flag integer,--删除标志
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--用药途径
CREATE TABLE route_administration
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(20) UNIQUE NOT NULL,--编码
  name varchar(20) NOT NULL,--名称
  print_name varchar(20),--打印名称
  py_code varchar(20),--拼音码
  d_code varchar(20),--五笔
  type_code varchar(20),--分类编码
  is_print integer,--是否打印
  input_type integer,--护士录入类别
  deleted_flag integer,--删除标志
  weight integer,--排序码/权重
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);


--药品
CREATE TABLE drug
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(20),--编码
  name varchar(30) NOT NULL,--药品名称
  english_name varchar(30),--英文名称
  sy_code varchar(30),--上药编码
  py_code varchar(20),--拼音码
  barcode varchar(20),--条形码
  d_code varchar(20),--简码
  print_name varchar(30),--打印名称
  serial varchar(20),--包装序号
  concentration varchar(10),--浓度
  specification varchar(30),--规格
  spe_comment varchar(40),--规格备注
  fix_price integer,--批发价
  ret_price integer,--零售价
  buy_price integer,--购入价
  manu_factory varchar(20),--生产厂商
  self_flag integer,--自费标识
  separate_flag integer,--单列标志
  suprice_flag integer,--贵重标志
  drug_flag integer,--毒麻标志
  dcode varchar(20),--自定义码
  infusion_flag integer,--大输液标志,9为并开药
  country_flag integer,--进口
  divide_flag integer,--分装标志
  deleted_flag integer,--删除标志
  zy_bill_item varchar(20),--住院帐单码
  mz_bill_item varchar(20),--门诊帐单码
  zy_charge_group varchar(20),--住院用药品分组
  mz_charge_group varchar(20),--门诊用药品分组
  extend_code varchar(20),--药品与外界衔接码
  license_no varchar(20),--国药准字、文号
  low_dosage_flag integer,--大规格小剂量标志
  eff_month varchar(2),--有效月份
  eff_day integer,--有效天数
  stock_warning integer,--库存预警数
  weight integer,--重量
  volum integer,--体积
  once_dose integer,--常用剂量
  stock_amount integer,--库存量
  virtual_stock_amount integer,--虚拟库存量
  is_discount boolean DEFAULT false,--是否允许折扣
  is_bulk_sales boolean DEFAULT false,--是否允许拆零销售
  bulk_sales_price integer,--拆零售价
  status boolean NOT NULL DEFAULT true,--是否启用
  drug_class_id integer references drug_class(id),--药物类型id
  drug_type_id integer references drug_type(id),--药品分类id
  dose_form_id integer references dose_form(id),--药品剂型id
  dose_unit_id integer references dose_unit(id),--药品包装id
  frequency_id integer references frequency(id),--用药频率id
  route_administration_id integer references route_administration(id),--用药途径id
  manu_factory_id integer references manu_factory(id),--生产厂商id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--药品别名
CREATE TABLE drug_print
(
  id serial PRIMARY KEY NOT NULL,--id
  drug_id INTEGER NOT NULL references drug(id),--药品id
  name varchar(30) NOT NULL,--药品名称
  py_code varchar(20),--拼音码
  d_code varchar(20),--简码
  flag integer,--什么标志
  print_name varchar(30),--药品别名
  name_type varchar(10),--类型
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);