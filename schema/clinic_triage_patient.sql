--分诊就诊人
CREATE TABLE clinic_triage_patient
(
  id serial PRIMARY KEY NOT NULL,--id
  department_id INTEGER NOT NULL references department(id),--科室id
  clinic_patient_id INTEGER NOT NULL references clinic_patient(id),--科室就诊人id
  personnel_id INTEGER NOT NULL references personnel(id),--录入人员id
  treat_status boolean NOT NULL DEFAULT false,--是否分诊
  visit_date DATE NOT NULL DEFAULT CURRENT_DATE,--日期
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  UNIQUE (department_id, clinic_patient_id,treat_status,visit_date)--联合主键，就诊人、科室、状态、日期唯一
);