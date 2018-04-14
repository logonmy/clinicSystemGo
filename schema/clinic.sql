--诊所
CREATE TABLE clinic
(
  code varchar(20) PRIMARY KEY NOT NULL,
  name varchar(40) NOT NULL,
  responsible_person varchar(40) NOT NULL,
  area varchar(40),
  status boolean NOT NULL DEFAULT true,
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP
);

