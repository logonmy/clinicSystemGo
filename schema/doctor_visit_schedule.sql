--医生出诊排班
CREATE TABLE doctor_visit_schedule
(
  id integer PRIMARY KEY NOT NULL, --排班编号
  department_id VARCHAR(10) REFERENCES department_personnel(department_id),--医生编码
  personnel_id VARCHAR(10) REFERENCES department_personnel(personnel_id),--科室编码
  visit_date DATA NOT NULL,--出诊日期
  am_pm CHAR(1) NOT NULL CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  stop_flag boolean NOT NULL DEFAULT false,--停诊标识
  tatal_num INT NOT NULL DEFAULT 20,--总的接诊数
  left_num INT NOT NULL DEFAULT 20 CHECK(left_num <= tatal_num),--剩余接诊数
  visit_code integer NOT NULL REFERENCES visit_type(code),--出诊编码
  created_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP,
  UNIQUE (doctor_code,dept_code,visit_date,am_pm,visit_code)
);