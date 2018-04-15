--出诊类型
CREATE TABLE visit_type
(
  code integer PRIMARY KEY NOT NULL,--编码
  name VARCHAR(10) NOT NULL,--出诊类型名称
  open_flag boolean NOT NULL DEFAULT true,--启用状态
  fee integer NOT NULL DEFAULT 0 CHECK(fee > 0),
  created_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time TIMESTAMP NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time TIMESTAMP
);