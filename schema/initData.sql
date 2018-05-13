
insert into patient_channel (name) values ('运营推荐'), ('会员介绍'), ('网络宣传'), ('社区患者');
insert into supplier (name) VALUES ('广州白云药厂'),('云南白药药厂');
insert into instock_way (name) VALUES ('采购入库'),('公益捐赠');
insert into outstock_way (name) VALUES ('科室领用'),('退货出库'),('报损出库');
insert into cuvette_color (name) VALUES ('红'), ('黑'), ('紫'),('蓝'), ('黄'), ('绿'),('灰'), ('橙');
INSERT INTO charge_project_type (id,name) VALUES (1,'西/成药处方'),(2,'中药处方'),(3,'检验医嘱'),(4,'检查医嘱'),(5,'材料费用'),(6,'其他费用'),(7,'治疗医嘱');
insert into laboratory_sample (code,name) VALUES (001,'标本1'),(002,'标本2');

insert into dose_unit (name,code) values ('小盒','0002'),('箱','0001'),('桶','0003');
insert into drug_class (name,code) values ('类型1','0001'),('类型2','0002'),('类型3','0003');
insert into dose_form (name,code) values ('剂型1','0001'),('剂型2','0002'),('剂型3','0003');
insert into storehouse (name,clinic_id) VALUES ('专业药房',1);
INSERT into route_administration (code,name) VALUES ('0001','口服<饭前>'),('0002','口服<饭后>'),('0003','含服'),('0004','肛塞');
insert into frequency (code,name) VALUES ('0001','餐后'),('0002','2次/日(8-3)'),('0003','四次/日'),('0004','1次/每晚 (8pm)');