
CREATE TABLE clinic(
  code 				varchar(20)			PRIMARY KEY     NOT NULL,
  name        varchar(40)			NOT NULL,
  responsible_person        varchar(40)			NOT NULL,
  area        varchar(40),
  status        boolean			NOT NULL		DEFAULT true,
  create_time   timestamp				NOT NULL		DEFAULT LOCALTIMESTAMP
);

