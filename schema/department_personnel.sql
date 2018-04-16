--人员科室关系表
CREATE TABLE department_personnel
(
  id serial PRIMARY KEY NOT NULL, --id
  department_id INTEGER NOT NULL references department(id),--科室id
  personnel_id INTEGER NOT NULL references personnel(id),--人员id
  type INTEGER NOT NULL CHECK(type > 0 AND type < 3),-- 关系类型 1：人事科室， 2：出诊科室
  status boolean NOT NULL DEFAULT true,--是否启用
  created_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  updated_time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
  deleted_time timestamp,
  UNIQUE (department_id, personnel_id, type)--人员id、科室id、 类别唯一
);