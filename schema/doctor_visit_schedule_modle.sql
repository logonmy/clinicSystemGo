--医生出诊周排班末班
CREATE TABLE doctor_visit_schedule
(
  id integer PRIMARY KEY NOT NULL, --模板id
  department_id VARCHAR(10) REFERENCES department_personnel(department_id),--医生编码
  personnel_id VARCHAR(10) REFERENCES department_personnel(personnel_id),--科室编码
  weekday INTEGER NOT NULL CHECK(weekday BETWEEN -1 AND 7),--出诊 日期（周几，0 代表 周日，1 周一...）
  am_pm CHAR(1) NOT NULL CHECK(am_pm = 'a' OR am_pm = 'p'),--出诊上下午
  tatal_num INT NOT NULL DEFAULT 20,--总的接诊数
  visit_type_code integer NOT NULL REFERENCES visit_type(code),--出诊类别编码
  created_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP,
  UNIQUE (department_id,personnel_id,weekday,am_pm,visit_type_code)
);