
insert into patient_channel (name) values ('运营推荐'), ('会员介绍'), ('网络宣传'), ('社区患者');
insert into supplier (name) VALUES ('广州白云药厂'),('云南白药药厂');
insert into instock_way (name) VALUES ('采购入库'),('公益捐赠');
insert into outstock_way (name) VALUES ('科室领用'),('退货出库'),('报损出库');
insert into cuvette_color (name) VALUES ('红'), ('黑'), ('紫'),('蓝'), ('黄'), ('绿'),('灰'), ('橙');
INSERT INTO charge_project_type (id,name) VALUES (1,'西/成药处方'),(2,'中药处方'),(3,'检验医嘱'),(4,'检查医嘱'),(5,'材料费用'),(6,'其他费用'),(7,'治疗医嘱');

insert into storehouse (name,clinic_id) VALUES ('专业药房',1);