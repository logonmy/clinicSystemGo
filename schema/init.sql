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
  open_flag boolean NOT NULL DEFAULT false,--是否开放
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
  clinic_patient_id INTEGER NOT NULL references clinic_patient(id),--科室就诊人id
  department_id INTEGER references department(id),--科室id
  doctor_id INTEGER references personnel(id),--接诊医生医生id
  visit_date DATE NOT NULL DEFAULT CURRENT_DATE,--日期
  am_pm varchar(1) CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  start_time varchar(10),--开始时间
  end_time varchar(10),--结束时间
  register_type INTEGER NOT NULL,--登记类型：1预约，2线下分诊
  visit_type integer NOT NULL,--出诊类型 1: 首诊， 2复诊，3：术后复诊
  status integer NOT NULL,--状态，对应 分诊就诊人操作记录表 type
  doctor_visit_schedule_id integer,--排版id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--分诊就诊人操作记录表
CREATE TABLE clinic_triage_patient_operation
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id integer NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  type INTEGER NOT NULL,--操作类型 10:登记，20：分诊(换诊)，30：接诊，40：已就诊， 100：取消
  times INTEGER NOT NULL,--操作次数
  personnel_id INTEGER references personnel(id),--操作员id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (clinic_triage_patient_id, type, times)--联合主键，就诊人、科室、状态、日期唯一
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
  id integer PRIMARY KEY NOT NULL,--id
  name varchar(20) UNIQUE NOT NULL,--名称
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊疗
CREATE TABLE diagnosis_treatment
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) UNIQUE NOT NULL,--名称
  en_name varchar(20),--英文名称
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊所诊疗
CREATE TABLE clinic_diagnosis_treatment
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  name varchar(20) UNIQUE NOT NULL,--名称
  en_name varchar(20),--英文名称
  cost integer CHECK(cost > 0), --成本价
  price integer NOT NULL CHECK(price > 0), --诊疗费金额
  status boolean NOT NULL DEFAULT true,--是否启用
  is_discount boolean DEFAULT false,--是否允许折扣
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

--门诊待缴费
CREATE TABLE mz_unpaid_orders
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  charge_project_type_id INTEGER NOT NULL references charge_project_type(id),--收费类型id
  charge_project_id INTEGER NOT NULL,--收费项目id
  order_sn varchar(50) NOT NULL,--单号
  soft_sn INTEGER NOT NULL,--序号
  name varchar(20) NOT NULL,--收费名称
  price INTEGER NOT NULL CHECK(price > 0),--单价
  amount INTEGER NOT NULL CHECK(amount > 0),--数量
  unit varchar(20),--单位
  total INTEGER NOT NULL,--折前总价格
  discount INTEGER NOT NULL DEFAULT 0,--打折总金额金额
  fee INTEGER NOT NULL,--折后金额  fee = total - discount
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
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  out_trade_no varchar(30) UNIQUE,--系统交易号
  trade_no varchar(30) UNIQUE,--第三方平台交易号(如；支付宝，微信)
  orders_ids text NOT NULL, --已缴费与未交费的记录id
  operation_id INTEGER NOT NULL references personnel(id),--操作员id
  pay_method_code varchar(2) NOT NULL,--支付方式编码，1-支付宝，2-微信, 3-银行卡, 4-现金

  status varchar(30) NOT NULL,--订单状态  WATTING_FOR_PAY 待支付 ; TRADE_SUCCESS 交易成功； PART_REFUND 部分退费； TOTAL_REFUND 全额退费
  
  derate_money INTEGER NOT NULL DEFAULT 0 CHECK(derate_money >= 0),--减免金额
  voucher_money INTEGER NOT NULL DEFAULT 0 CHECK(voucher_money >= 0) ,--抵金券金额
  discount_money INTEGER NOT NULL DEFAULT 0 CHECK(discount_money >= 0),--折扣金额
  bonus_points_money INTEGER NOT NULL DEFAULT 0 CHECK(bonus_points_money >= 0) ,--积分兑换金额
  on_credit_money INTEGER NOT NULL DEFAULT 0 CHECK(on_credit_money >= 0) ,--挂账金额
  medical_money INTEGER NOT NULL DEFAULT 0 CHECK(medical_money >= 0),--医保金额

  total_money  INTEGER NOT NULL  ,--应收金额
  balance_money INTEGER NOT NULL ,--实收金额
  refund_money INTEGER NOT NULL DEFAULT 0,--已退费金额

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
  order_sn varchar(50) NOT NULL,--单号
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
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  charge_project_type_id INTEGER NOT NULL references charge_project_type(id),--收费类型id
  charge_project_id INTEGER NOT NULL,--收费项目id
  order_sn varchar(50) NOT NULL,--单号
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
  patient_id INTEGER references patient(id),--关联的患者
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

--挂账记录表总表
CREATE TABLE on_credit_record
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  trade_no varchar(30) UNIQUE,--第三方平台交易号(如；支付宝，微信)
  on_credit_money INTEGER NOT NULL, --挂账总金额金额
  already_pay_money INTEGER NOT NULL DEFAULT 0, --已还款金额
  operation_id INTEGER NOT NULL references personnel(id),--未交费创建人id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--挂账还款记录
CREATE TABLE on_credit_record_detail
(
  id serial PRIMARY KEY NOT NULL,--id
  on_credit_record_id INTEGER NOT NULL references on_credit_record(id),--分诊记录id
  type INTEGER NOT NULL CHECK(type = 0 or type = 1),--类型 0-挂账 1-还账 
  pay_method_code varchar(2),--支付方式编码 01-现金，02-微信，03-支付宝，04-银行卡
  should_repay_moeny INTEGER NOT NULL, --应还金额
  repay_moeny INTEGER NOT NULL, --实还金额
  remain_repay_moeny INTEGER NOT NULL, --剩余挂账金额
  operation_id INTEGER NOT NULL references personnel(id),--未交费创建人id
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

--药品分类
CREATE TABLE drug_class
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(30) NOT NULL UNIQUE,--药品分类名称
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
  name varchar(100) NOT NULL,--厂商名称
  abbr_name varchar(100),--
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
  input_type varchar(5),--护士录入类别
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
  type INTEGER NOT NULL CHECK(type = 0 or type = 1),--类型 0-西药 1-中药
  code varchar(20),--编码
  name varchar(30) NOT NULL,--药品名称
  py_code varchar(20),--拼音码
  barcode varchar(20),--条形码
  d_code varchar(20),--简码
  print_name varchar(30),--打印名称
  specification varchar(30),--规格
  spe_comment varchar(40),--规格备注
  manu_factory_name varchar(100),--生产厂商
  drug_class_id integer references drug_class(id),--药品类型id
  dose_form_name varchar(20),--药品剂型
  license_no varchar(100),--国药准字、文号
  once_dose integer,--常用剂量
  once_dose_unit_name varchar(10),--用量单位 常用剂量单位
  dosage integer,--剂量
  dosage_unit_name varchar(10),--剂量单位
  preparation_count integer,--制剂数量/包装量
  preparation_count_unit_name varchar(10),--制剂数量单位
  packing_unit_name varchar(10),--药品包装单位
  route_administration_name varchar(50),--用药途径id/默认用法
  frequency_name varchar(20),--用药频率/默认频次
  default_remark varchar(20),--默认用量用法说明
  weight integer,--重量
  weight_unit_name varchar(10),--重量单位
  volum integer,--体积
  vol_unit_name varchar(10),--体积单位
  serial varchar(20),--包装序号
  national_standard varchar(20),--国标分类
  concentration varchar(10),--浓度
  dcode varchar(20),--自定义码
  infusion_flag integer,--大输液标志,9为并开药
  country_flag integer,--进口
  divide_flag integer,--分装标志 
  low_dosage_flag integer,--大规格小剂量标志
  self_flag integer,--自费标识
  separate_flag integer,--单列标志
  suprice_flag integer,--贵重标志
  drug_flag integer,--毒麻标志
  english_name varchar(30),--英文名称
  sy_code varchar(30),--上药编码

  zy_bill_item varchar(20),--住院帐单码
  mz_bill_item varchar(20),--门诊帐单码
  zy_charge_group varchar(20),--住院用药品分组
  mz_charge_group varchar(20),--门诊用药品分组
  extend_code varchar(20),--药品与外界衔接码
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (name, specification, manu_factory_name)
);

--药品别名
CREATE TABLE drug_print
(
  id serial PRIMARY KEY NOT NULL,--id
  drug_id INTEGER NOT NULL references drug(id),--药品id
  name varchar(30) NOT NULL,--药品名称
  py_code varchar(20),--拼音码
  d_code varchar(20),--简码
  status boolean,--启用标志
  print_name varchar(30),--药品别名
  name_type varchar(10),--类型
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--供应商
CREATE TABLE supplier
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(30) NOT NULL,--供应商名称
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--入库方式
CREATE TABLE instock_way
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) NOT NULL,--入库方式名称
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--出库方式
CREATE TABLE outstock_way
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) NOT NULL,--出库方式名称
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--库房
CREATE TABLE storehouse
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) NOT NULL,--库房名称
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊所药品
CREATE TABLE clinic_drug
(
  id serial PRIMARY KEY NOT NULL,--id
  type INTEGER NOT NULL CHECK(type = 0 or type = 1),--类型 0-西药 1-中药
  drug_class_id INTEGER references drug_class(id),--药品分类
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  name varchar(50) NOT Null,--药名
  specification varchar(50) NOT Null,--规格
  manu_factory_name varchar(50) NOT Null,--生产厂商
  dose_form_name varchar(50) NOT Null,--剂型
  print_name varchar(50),--商品名
  license_no varchar(50) NOT Null,--国药准字
  py_code varchar(50),--拼音码
  barcode varchar(50) NOT Null,--条形码
  status boolean NOT NULL DEFAULT true,--是否启用
  dosage integer,--剂量
  dosage_unit_name varchar(10),--剂量单位
  preparation_count integer,--制剂数量/包装量
  preparation_count_unit_name varchar(10),--制剂数量单位
  packing_unit_name varchar(10),--药品包装单位
  ret_price integer,--零售价
  buy_price integer,--成本价
  mini_dose integer,--最小剂量
  is_discount boolean DEFAULT false,--是否允许折扣
  is_bulk_sales boolean DEFAULT false,--是否允许拆零销售
  bulk_sales_price integer,--拆零售价/最小剂量售价
  fetch_address integer DEFAULT 0,--取药地点 0 本诊所，1外购 2， 代购
  once_dose integer,--常用剂量
  once_dose_unit_name varchar(10),--用量单位 常用剂量单位
  route_administration_name varchar(50),--用药途径id/默认用法
  frequency_name varchar(20),--用药频率/默认频次
  illustration text,--说明
  day_warning integer,--效期预警天数
  stock_warning integer,--库存预警数
  english_name varchar(30),--英文名称
  sy_code varchar(30),--上药编码
  country_flag boolean,--进口
  self_flag boolean,--自费标识
  drug_flag boolean,--毒麻标志
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (clinic_id, name, specification, manu_factory_name)
);

--药品库存
CREATE TABLE drug_stock
(
  id serial PRIMARY KEY NOT NULL,--id
  storehouse_id integer NOT NULL references storehouse(id),--库房id
  clinic_drug_id INTEGER NOT NULL references clinic_drug(id),--诊所药品id
  supplier_name varchar(100),--供应商
  stock_amount INTEGER NOT NULL DEFAULT 0,--库存数量
  serial varchar(20),--批号
  eff_date DATE NOT NULL,--有效日期
  buy_price integer,--成本价
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (storehouse_id, clinic_drug_id, supplier_name, serial, eff_date)
);

--药品入库记录
CREATE TABLE drug_instock_record
(
  id serial PRIMARY KEY NOT NULL,--id
  storehouse_id integer NOT NULL references storehouse(id),--库房id
  order_number varchar(20) NOT NULL,--入库单号
  instock_way_name varchar(20),--入库方式
  supplier_name varchar(100),--供应商
  instock_date DATE NOT NULL DEFAULT CURRENT_DATE,--入库日期
  remark text,--备注
  instock_operation_id INTEGER NOT NULL references personnel(id),--入库人员id
  verify_operation_id INTEGER references personnel(id),--审核人员id
  verify_status varchar(2) NOT NULL DEFAULT '01',--审核状态 01 未审核 02 已审核 
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--药品入库记录item
CREATE TABLE drug_instock_record_item
(
  id serial PRIMARY KEY NOT NULL,--id
  drug_instock_record_id integer NOT NULL references drug_instock_record(id),--药品入库记录id
  clinic_drug_id INTEGER NOT NULL references clinic_drug(id),--诊所药品id
  instock_amount INTEGER NOT NULL CHECK(instock_amount > 0),--入库数量
  serial varchar(20) NOT NULL,--批号
  buy_price integer,--成本价
  eff_date DATE NOT NULL,--有效日期
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--药品出库记录
CREATE TABLE drug_outstock_record
(
  id serial PRIMARY KEY NOT NULL,--id
  storehouse_id integer NOT NULL references storehouse(id),--库房id
  department_id INTEGER NOT NULL references department(id),--领用科室id
  personnel_id INTEGER NOT NULL references personnel(id),--领用人员id
  order_number varchar(20) NOT NULL,--出库单号
  outstock_date DATE NOT NULL DEFAULT CURRENT_DATE,--出库日期
  outstock_way_name varchar(20),--出库方式
  remark text,--备注
  outstock_operation_id INTEGER NOT NULL references personnel(id),--出库人员id
  verify_operation_id INTEGER references personnel(id),--审核人员id
  verify_status varchar(2) NOT NULL DEFAULT '01',--审核状态 01 未审核 02 已审核
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--药品出库记录item
CREATE TABLE drug_outstock_record_item
(
  id serial PRIMARY KEY NOT NULL,--id
  drug_outstock_record_id integer NOT NULL references drug_outstock_record(id),--药品出库记录id
  drug_stock_id INTEGER NOT NULL references drug_stock(id),--药品库存id
  outstock_amount INTEGER NOT NULL CHECK(outstock_amount > 0),--出库数量
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊断字典
CREATE TABLE diagnosis
(
  id serial PRIMARY KEY NOT NULL,--id
  py_code varchar(10) NOT NULL,--拼音码
  name varchar(20) NOT NULL,--诊断名称
  icd_code varchar(10) ,--国际疾病分类
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--病历
CREATE TABLE medical_record
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL UNIQUE references clinic_triage_patient(id),--预约编号
  morbidity_date varchar(10), --发病日期
  chief_complaint text NOT NULL, --主诉
  history_of_present_illness text,-- 现病史
  history_of_past_illness text, --既往史
  family_medical_history text, --家族史
  allergic_history text,--过敏史
  allergic_reaction text,--过敏反应
  immunizations text,--疫苗接种史
  body_examination text,--体格检查
  diagnosis text,--诊断
  cure_suggestion text, --治疗建议
  remark text,--备注
  files text,--上传的文件
  operation_id integer REFERENCES personnel(id),--操作人编码
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--病历模板
CREATE TABLE medical_record_model
(
  id serial PRIMARY KEY NOT NULL,--id
  model_name varchar(20) NOT NULL,--模板名称
  is_common boolean NOT NULL DEFAULT false,--是否通用
  chief_complaint text NOT NULL, --主诉
  history_of_present_illness text,-- 现病史
  history_of_past_illness text, --既往史
  family_medical_history text, --家族史
  allergic_history text,--过敏史
  allergic_reaction text,--过敏反应
  immunizations text,--疫苗接种史
  body_examination text,--体格检查
  diagnosis text,--诊断
  cure_suggestion text, --治疗建议
  remark text,--备注
  operation_id integer REFERENCES personnel(id),--操作人编码
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--检查部位
CREATE TABLE examination_organ
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(10) NOT NULL,--部位名称
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--检查医嘱
CREATE TABLE examination
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(100) UNIQUE NOT NULL,--检查名称
  en_name varchar(100),--英文名称
  py_code varchar(100),--拼音码
  idc_code varchar(100),--国际编码
  unit_name varchar(20),--单位名称
  organ varchar(100),--检查部位
  remark text,--备注
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊所检查医嘱
CREATE TABLE clinic_examination
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  name varchar(100) NOT NULL,--检查名称
  en_name varchar(100),--英文名称
  py_code varchar(100),--拼音码
  idc_code varchar(100),--国际编码
  unit_name varchar(20),--单位名称
  organ varchar(100),--检查部位
  remark text,--备注
  cost integer, --成本价
  price integer NOT NULL,--销售价
  status boolean NOT NULL DEFAULT true,--是否启用
  is_discount boolean NOT NULL DEFAULT false,--是否允许折扣
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE(clinic_id,name)
);

--检验医嘱标本种类
CREATE TABLE laboratory_sample
(
  id serial PRIMARY KEY NOT NULL,--id
  code varchar(20),--编码
  name varchar(20) NOT NULL,--标本名称
  py_code varchar(20),--拼音码
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--试管颜色
CREATE TABLE cuvette_color
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) UNIQUE NOT NULL,--颜色名称
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--检验医嘱
CREATE TABLE laboratory
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(100) UNIQUE NOT NULL,--检验医嘱名称
  en_name varchar(20),--英文名称
  py_code varchar(20),--拼音码
  idc_code varchar(20),--国际编码
  unit_name varchar(20),--单位名称
  time_report varchar(30),--报告所需时间
  clinical_significance text,--临床意义
  remark text,--备注
  laboratory_sample varchar(30),--检验物
  laboratory_sample_dosage varchar(30),--检验物计量
  cuvette_color_name varchar(20),--试管颜色
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊所检验医嘱
CREATE TABLE clinic_laboratory
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  name varchar(100) NOT NULL,--检验医嘱名称
  en_name varchar(20),--英文名称
  py_code varchar(20),--拼音码
  idc_code varchar(20),--国际编码
  unit_name varchar(20),--单位名称
  time_report varchar(30),--报告所需时间
  clinical_significance text,--临床意义
  laboratory_sample varchar(30),--检验物
  laboratory_sample_dosage varchar(30),--检验物计量
  cuvette_color_name varchar(20),--试管颜色
  merge_flag integer,--合并标记
  cost integer, --成本价
  price integer NOT NULL,--销售价
  status boolean NOT NULL DEFAULT true,--是否启用
  is_discount boolean NOT NULL DEFAULT false,--是否允许折扣
  is_delivery boolean NOT NULL DEFAULT false,--是否允许外送
  remark text,--备注
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE(clinic_id,name)
);

--检验项目
CREATE TABLE laboratory_item
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(100) UNIQUE NOT NULL,--检验名称
  en_name varchar(100),--英文名称
  instrument_code varchar(20),--仪器编码
  unit_name varchar(20),--单位名称
  clinical_significance text,--临床意义
  data_type integer,--数据类型 1 定性 2 定量
  is_special boolean,--参考值是否特殊
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--检验项目参考值
CREATE TABLE laboratory_item_reference
(
  id serial PRIMARY KEY NOT NULL,--id
  laboratory_item_id integer NOT NULL references laboratory_item(id),--检验项目id
  reference_max varchar(20), --定量/定性参考值最大值
  reference_min varchar(20), --定量/定性参考值最小值
  age_max integer, --参考值年龄段最大值
  age_min integer, --参考值年龄段最小值
  reference_sex varchar(5),--参考值性别 男、女、通用
  stomach_status varchar(5),--空腹、餐后 1h、餐后 2h
  is_pregnancy boolean,--是否妊娠期
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊所检验项目
CREATE TABLE clinic_laboratory_item
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  name varchar(100) UNIQUE NOT NULL,--检验名称
  en_name varchar(100),--英文名称
  instrument_code varchar(20),--仪器编码
  unit_name varchar(20),--单位名称
  clinical_significance text,--临床意义
  data_type integer,--数据类型 1 定性 2 定量
  is_special boolean,--参考值是否特殊
  status boolean NOT NULL DEFAULT true,--是否启用
  is_delivery boolean NOT NULL DEFAULT false,--是否允许外送
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE(clinic_id,name)
);

--诊所检验项目参考值
CREATE TABLE clinic_laboratory_item_reference
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_laboratory_item_id integer NOT NULL references clinic_laboratory_item(id),--诊所检验项目id
  reference_max varchar(20), --定量/定性参考值最大值
  reference_min varchar(20), --定量/定性参考值最小值
  age_max integer, --参考值年龄段最大值
  age_min integer, --参考值年龄段最小值
  reference_sex varchar(5),--参考值性别 男、女、通用
  stomach_status varchar(5),--空腹、餐后 1h、餐后 2h
  is_pregnancy boolean,--是否妊娠期
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊所检验医嘱关联的检验项
CREATE TABLE clinic_laboratory_association
(
  clinic_laboratory_id integer NOT NULL references clinic_laboratory(id),--诊所医嘱id
  clinic_laboratory_item_id INTEGER references clinic_laboratory_item(id),--诊所检验项目id
  default_result varchar(10),--默认结果
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--治疗医嘱
CREATE TABLE treatment
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) UNIQUE NOT NULL,--名称
  en_name varchar(20),--英文名称
  py_code varchar(20),--拼音码
  idc_code varchar(20),--国际编码
  unit_name varchar(20),--单位名称
  remark text,--备注
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊所治疗项目
CREATE TABLE clinic_treatment
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  name varchar(20) UNIQUE NOT NULL,--名称
  en_name varchar(20),--英文名称
  py_code varchar(20),--拼音码
  idc_code varchar(20),--国际编码
  unit_name varchar(20),--单位名称
  remark text,--备注
  cost integer CHECK(cost > 0), --成本价
  price integer NOT NULL CHECK(price > 0), --销售价
  status boolean NOT NULL DEFAULT true,--是否启用
  is_discount boolean NOT NULL DEFAULT false,--是否允许折扣
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--材料费用
CREATE TABLE material
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) NOT NULL,--名称
  en_name varchar(20),--英文名称
  py_code varchar(20),--拼音码
  idc_code varchar(20),--国际编码
  manu_factory_name varchar(100),--生产厂商
  specification varchar(30),--规格
  unit_name varchar(20),--单位
  remark text,--备注
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE(name,manu_factory_name,specification)
);

--诊所材料
CREATE TABLE clinic_material
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  name varchar(20) NOT NULL,--名称
  en_name varchar(20),--英文名称
  py_code varchar(20),--拼音码
  idc_code varchar(20),--国际编码
  manu_factory_name varchar(100),--生产厂商
  specification varchar(30),--规格
  unit_name varchar(20),--单位
  remark text,--备注
  ret_price integer,--零售价
  buy_price integer,--成本价
  is_discount boolean DEFAULT false,--是否允许折扣
  day_warning integer,--效期预警天数
  stock_warning integer,--库存预警数
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (clinic_id, name, manu_factory_name,specification)
);

--诊所材料库存
CREATE TABLE material_stock
(
  id serial PRIMARY KEY NOT NULL,--id
  storehouse_id integer NOT NULL references storehouse(id),--库房id
  clinic_material_id INTEGER NOT NULL references clinic_material(id),--诊所药品id
  supplier_name varchar(100),--供应商
  stock_amount INTEGER NOT NULL DEFAULT 0,--库存数量
  serial varchar(20),--批号
  eff_date DATE NOT NULL,--有效日期
  buy_price integer,--成本价
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (storehouse_id, clinic_material_id, supplier_name, serial, eff_date)
);

--耗材入库记录
CREATE TABLE material_instock_record
(
  id serial PRIMARY KEY NOT NULL,--id
  storehouse_id integer NOT NULL references storehouse(id),--库房id
  order_number varchar(20) NOT NULL,--入库单号
  instock_way_name varchar(20),--入库方式
  supplier_name varchar(100),--供应商
  instock_date DATE NOT NULL DEFAULT CURRENT_DATE,--入库日期
  remark text,--备注
  instock_operation_id INTEGER NOT NULL references personnel(id),--入库人员id
  verify_operation_id INTEGER references personnel(id),--审核人员id
  verify_status varchar(2) NOT NULL DEFAULT '01',--审核状态 01 未审核 02 已审核 
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--耗材入库记录item
CREATE TABLE material_instock_record_item
(
  id serial PRIMARY KEY NOT NULL,--id
  material_instock_record_id integer NOT NULL references material_instock_record(id),--材料入库记录id
  clinic_material_id INTEGER NOT NULL references clinic_material(id),--诊所材料id
  instock_amount INTEGER NOT NULL CHECK(instock_amount > 0),--入库数量
  serial varchar(20) NOT NULL,--批号
  buy_price integer,--成本价
  eff_date DATE NOT NULL,--有效日期
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--耗材出库记录
CREATE TABLE material_outstock_record
(
  id serial PRIMARY KEY NOT NULL,--id
  storehouse_id integer NOT NULL references storehouse(id),--库房id
  department_id INTEGER NOT NULL references department(id),--领用科室id
  personnel_id INTEGER NOT NULL references personnel(id),--领用人员id
  order_number varchar(20) NOT NULL,--出库单号
  outstock_date DATE NOT NULL DEFAULT CURRENT_DATE,--出库日期
  outstock_way_name varchar(20),--出库方式
  remark text,--备注
  outstock_operation_id INTEGER NOT NULL references personnel(id),--出库人员id
  verify_operation_id INTEGER references personnel(id),--审核人员id
  verify_status varchar(2) NOT NULL DEFAULT '01',--审核状态 01 未审核 02 已审核
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--耗材出库记录item
CREATE TABLE material_outstock_record_item
(
  id serial PRIMARY KEY NOT NULL,--id
  material_outstock_record_id integer NOT NULL references material_outstock_record(id),--材料出库记录id
  material_stock_id INTEGER NOT NULL references material_stock(id),--材料库存id
  outstock_amount INTEGER NOT NULL CHECK(outstock_amount > 0),--出库数量
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--其它费用
CREATE TABLE other_cost
(
  id serial PRIMARY KEY NOT NULL,--id
  name varchar(20) NOT NULL,--名称
  en_name varchar(20),--英文名称
  py_code varchar(20),--拼音码
  unit_name varchar(20),--单位名称
  remark text,--备注
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--诊所其它费用项目
CREATE TABLE clinic_other_cost
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_id integer NOT NULL references clinic(id),--所属诊所
  name varchar(50) NOT NULL,--名称
  en_name varchar(50),--英文名称
  py_code varchar(50),--拼音码
  unit_name varchar(20),--单位名称
  remark text,--备注
  cost integer CHECK(cost > 0), --成本价
  price integer NOT NULL CHECK(price > 0), --销售价
  status boolean NOT NULL DEFAULT true,--是否启用
  is_discount boolean NOT NULL DEFAULT false,--是否允许折扣
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);


--开治疗
CREATE TABLE treatment_patient
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  clinic_treatment_id INTEGER NOT NULL references clinic_treatment(id),--治疗项目id
  order_sn varchar(20) NOT NULL,--单号
  soft_sn INTEGER NOT NULL,--序号
  times INTEGER NOT NULL CHECK(times > 0),--次数
  illustration text,--说明
  operation_id INTEGER NOT NULL references personnel(id),--操作员id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (order_sn, soft_sn)
);

--开检验
CREATE TABLE laboratory_patient
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  clinic_laboratory_id INTEGER NOT NULL references clinic_laboratory(id),--检验项目id
  order_sn varchar(50) NOT NULL,--单号
  soft_sn INTEGER NOT NULL,--序号
  times INTEGER NOT NULL CHECK(times > 0),--次数
  illustration text,--说明
  operation_id INTEGER NOT NULL references personnel(id),--操作员id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (order_sn, soft_sn)
);

--开西/成药处方
CREATE TABLE prescription_western_patient
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  clinic_drug_id INTEGER NOT NULL references clinic_drug(id),--诊所药品id
  order_sn varchar(50) NOT NULL,--单号
  soft_sn INTEGER NOT NULL,--序号
  once_dose integer,--单次剂量
  once_dose_unit_name varchar(10),--用量单位 单次剂量单位
  route_administration_name varchar(50),--用法
  frequency_name varchar(20),--用药频率/默认频次
  amount INTEGER NOT NULL CHECK(amount > 0),--总量
  illustration text,--用药说明
  fetch_address integer,--取药地点 0 本诊所 1 外购
  eff_day integer,--天数
  operation_id INTEGER NOT NULL references personnel(id),--操作员id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (order_sn, soft_sn)
);

--开中药处方
CREATE TABLE prescription_chinese_patient
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  order_sn varchar(50) UNIQUE NOT NULL,--单号
  route_administration_name varchar(50),--用法
  frequency_name varchar(20),--用药频率/默认频次
  amount INTEGER NOT NULL CHECK(amount > 0),--总剂量
  medicine_illustration text,--服药说明
  fetch_address integer,--取药地点 0 本诊所 1 外购
  eff_day integer,--天数
  operation_id INTEGER NOT NULL references personnel(id),--操作员id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--中药处方item
CREATE TABLE prescription_chinese_item
(
  id serial PRIMARY KEY NOT NULL,--id
  prescription_chinese_patient_id INTEGER NOT NULL references prescription_chinese_patient(id),--中药处方id
  clinic_drug_id INTEGER NOT NULL references clinic_drug(id),--诊所药品id
  order_sn varchar(50) NOT NULL,--单号
  soft_sn INTEGER NOT NULL,--序号
  once_dose integer,--单次剂量
  once_dose_unit_name varchar(20),--用量单位 单次剂量单位id
  amount INTEGER NOT NULL CHECK(amount > 0),--总量
  special_illustration text,--特殊要求
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (prescription_chinese_patient_id, soft_sn)
);

--开检查
CREATE TABLE examination_patient
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  clinic_examination_id INTEGER NOT NULL references clinic_examination(id),--检查项目id
  order_sn varchar(20) NOT NULL,--单号
  soft_sn INTEGER NOT NULL,--序号
  times INTEGER NOT NULL CHECK(times > 0),--次数
  organ varchar(20),--检查部位
  illustration text,--说明
  operation_id INTEGER NOT NULL references personnel(id),--操作员id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (clinic_triage_patient_id, clinic_examination_id, order_sn, soft_sn)
);

--开材料费
CREATE TABLE material_patient
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  clinic_material_id INTEGER NOT NULL references clinic_material(id),--诊所药品id
  order_sn varchar(20) NOT NULL,--单号
  soft_sn INTEGER NOT NULL,--序号
  amount INTEGER NOT NULL CHECK(amount > 0),--数量
  illustration text,--说明
  operation_id INTEGER NOT NULL references personnel(id),--操作员id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (clinic_triage_patient_id, clinic_material_id, order_sn, soft_sn)
);

--开其它费用
CREATE TABLE other_cost_patient
(
  id serial PRIMARY KEY NOT NULL,--id
  clinic_triage_patient_id INTEGER NOT NULL references clinic_triage_patient(id),--分诊就诊人id
  clinic_other_cost_id INTEGER NOT NULL references clinic_other_cost(id),--其它费用项目id
  order_sn varchar(20) NOT NULL,--单号
  soft_sn INTEGER NOT NULL,--序号
  amount INTEGER NOT NULL CHECK(amount > 0),--数量
  illustration text,--说明
  operation_id INTEGER NOT NULL references personnel(id),--操作员id
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone,
  UNIQUE (order_sn, soft_sn)
);

--西药处方模板
CREATE TABLE prescription_western_patient_model
(
  id serial PRIMARY KEY NOT NULL,--id
  model_name varchar(20) NOT NULL,--模板名称
  is_common boolean NOT NULL DEFAULT false,--是否通用
  operation_id integer REFERENCES personnel(id),--操作人编码
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--西药处方模板项目
CREATE TABLE prescription_western_patient_model_item
(
  id serial PRIMARY KEY NOT NULL,--id
  prescription_western_patient_model_id INTEGER NOT NULL references prescription_western_patient_model(id),--西药处方模板id
  clinic_drug_id INTEGER NOT NULL references clinic_drug(id),--诊所药品id
  amount INTEGER NOT NULL CHECK(amount > 0),--总量
  once_dose integer,--单次剂量
  once_dose_unit_name varchar(20),--用量单位 单次剂量单位
  route_administration_name varchar(50),--用法
  frequency_name varchar(20),--用药频率id/默认频次
  fetch_address integer,--取药地点 0 本诊所 1 外购
  eff_day integer,--有效天数
  illustration text,--用药说明
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--中药处方模板
CREATE TABLE prescription_chinese_patient_model
(
  id serial PRIMARY KEY NOT NULL,--id
  model_name varchar(20) NOT NULL,--模板名称
  is_common boolean NOT NULL DEFAULT false,--是否通用
  route_administration_name varchar(50),--用法
  frequency_name varchar(20),--用药频率id/默认频次
  amount INTEGER NOT NULL CHECK(amount > 0),--总剂量
  eff_day integer,--天数
  fetch_address integer,--取药地点 0 本诊所 1 外购
  medicine_illustration text,--服药说明
  operation_id integer REFERENCES personnel(id),--操作人编码
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--中药药处方模板项目
CREATE TABLE prescription_chinese_patient_model_item
(
  id serial PRIMARY KEY NOT NULL,--id
  prescription_chinese_patient_model_id INTEGER NOT NULL references prescription_chinese_patient_model(id),--中药处方模板id
  clinic_drug_id INTEGER NOT NULL references clinic_drug(id),--诊所药品id
  once_dose integer,--单次剂量
  once_dose_unit_name varchar(20),--用量单位 单次剂量单位
  amount INTEGER NOT NULL CHECK(amount > 0),--总量
  special_illustration text,--说明/特殊要求
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);


--检验模板
CREATE TABLE laboratory_patient_model
(
  id serial PRIMARY KEY NOT NULL,--id
  model_name varchar(20) NOT NULL,--模板名称
  is_common boolean NOT NULL DEFAULT false,--是否通用
  operation_id integer REFERENCES personnel(id),--操作人编码
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--检验模板项目
CREATE TABLE laboratory_patient_model_item
(
  laboratory_patient_model_id INTEGER NOT NULL references laboratory_patient_model(id),--检验模板id
  clinic_laboratory_id INTEGER NOT NULL references clinic_laboratory(id),--检验医嘱id
  times INTEGER NOT NULL CHECK(times > 0),--次数
  illustration text,--说明
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--检查模板
CREATE TABLE examination_patient_model
(
  id serial PRIMARY KEY NOT NULL,--id
  model_name varchar(20) NOT NULL,--模板名称
  is_common boolean NOT NULL DEFAULT false,--是否通用
  operation_id integer REFERENCES personnel(id),--操作人编码
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--检查模板项目
CREATE TABLE examination_patient_model_item
(
  examination_patient_model_id INTEGER NOT NULL references examination_patient_model(id),--检验模板id
  clinic_examination_id INTEGER NOT NULL references clinic_examination(id),--检查项目id
  times INTEGER NOT NULL CHECK(times > 0),--次数
  illustration text,--说明
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--治疗模板
CREATE TABLE treatment_patient_model
(
  id serial PRIMARY KEY NOT NULL,--id
  model_name varchar(20) NOT NULL,--模板名称
  is_common boolean NOT NULL DEFAULT false,--是否通用
  operation_id integer REFERENCES personnel(id),--操作人编码
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);

--治疗模板项目
CREATE TABLE treatment_patient_model_item
(
  treatment_patient_model_id INTEGER NOT NULL references treatment_patient_model(id),--治疗模板id
  clinic_treatment_id INTEGER NOT NULL references clinic_treatment(id),--治疗项目id
  times INTEGER NOT NULL CHECK(times > 0),--次数
  illustration text,--说明
  created_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp with time zone NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp with time zone
);