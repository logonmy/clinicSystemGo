
insert into patient_channel (name) values ('运营推荐'), ('会员介绍'), ('网络宣传'), ('社区患者');
insert into supplier (name) VALUES ('广州白云药厂'),('云南白药药厂');
insert into instock_way (name) VALUES ('采购入库'),('公益捐赠');
insert into outstock_way (name) VALUES ('科室领用'),('退货出库'),('报损出库');
insert into cuvette_color (name) VALUES ('红'), ('黑'), ('紫'),('蓝'), ('黄'), ('绿'),('灰'), ('橙');
INSERT INTO charge_project_type (name) VALUES ('西/成药处方'),('中药处方'),('检验医嘱'),('检验项目'),('检查医嘱'),('材料费用'),('其他费用'),('诊疗项目');
insert into laboratory_sample (code,name) VALUES (001,'标本1'),(002,'标本2');

insert into dose_unit (name,code) values ('小盒','0002'),('箱','0001'),('桶','0003');
insert into storehouse (name,clinic_id) VALUES ('专业药房',1);
