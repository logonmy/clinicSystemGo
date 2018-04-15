--医生出诊排班
CREATE TABLE doctor_visit_schedule
(
  id integer PRIMARY KEY NOT NULL, --排班编号
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