CREATE TABLE admin(
	username        varchar(20)			NOT NULL,
 	phone        varchar(11)			NOT NULL,
	password        varchar(40)			NOT NULL,
	clinic_code        varchar(40)		NOT NULL  references clinic(code),
	status        boolean			NOT NULL		DEFAULT true,
	create_time   timestamp				NOT NULL		DEFAULT LOCALTIMESTAMP,
 	update_time   timestamp				NOT NULL		DEFAULT LOCALTIMESTAMP,
 	PRIMARY KEY (username, clinic_code)
);