CREATE TABLE patient(
 	cert_no        varchar(18)			PRIMARY KEY			NOT NULL,
	name        varchar(10)			NOT NULL,
 	birthday        varchar(8)			NOT NULL,
	sex        integer			NOT NULL			CHECK (sex > -1 && sex < 2),
	phone		varchar(11)			not Null,
	address		varchar(40),
	profession		varchar(40),
	remark			varchar(200),
 	status        boolean			NOT NULL		DEFAULT true,
 	create_time   timestamp				NOT NULL		DEFAULT LOCALTIMESTAMP,
	update_time   timestamp				NOT NULL		DEFAULT LOCALTIMESTAMP
);