--科室
CREATE TABLE department
(
	id serial PRIMARY KEY NOT NULL,--id
	code varchar(20) NOT NULL,--科室编码
	name varchar(20) NOT NULL,--科室名称
	clinic_code varchar(20) NOT NULL references clinic(code),--所属诊所
	status boolean NOT NULL DEFAULT true,--是否启用
	is_appointment boolean NOT NULL,--是否开放预约/挂号
	is_personnel boolean NOT NULL,--是否人事科室
	created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
	updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
	deleted_time timestamp,
	UNIQUE (code, clinic_code)
);