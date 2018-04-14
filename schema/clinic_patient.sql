--诊所就诊人
CREATE TABLE clinic_patient
(
  id integer PRIMARY KEY NOT NULL, --排班编号
  patient_cert_no varchar(18) NOT NULL references patient(cert_no),--患者身份证号
  clinic_code varchar(40) NOT NULL references clinic(code),--诊所编码
  personnel_code varchar(10) NOT NULL references personnel(code),--录入人员编码
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  UNIQUE (patient_cert_no, clinic_code)--联合主键，就诊人身份证号和诊所编码唯一
);