--挂号记录
CREATE TABLE appointment
(
  id serial PRIMARY KEY NOT NULL, --编号
  clinic_patient_id integer NOT NULL references clinic_patient(id),--患者id
  department_id VARCHAR(10) REFERENCES department_personnel(department_id),--医生编码
  personnel_id VARCHAR(10) REFERENCES department_personnel(personnel_id),--科室编码
  visit_date DATA NOT NULL,--出诊日期
  am_pm CHAR(1) NOT NULL CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  is_today boolean NOT NULL DEFAULT false,--是否当日号
  visit_type_code integer NOT NULL REFERENCES visit_type(code),--出诊类型编码
  status VARCHAR(2) NOT NULL DEFAULT '01',--就诊状态
  visit_place VARCHAR(20),--就诊未知
  sort_no SMALLINT,--就诊序号
  operationCode VARCHAR(10),--操作人编码
  created_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP
);