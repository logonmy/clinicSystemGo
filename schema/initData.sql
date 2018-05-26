
insert into patient_channel (name) values ('运营推荐'), ('会员介绍'), ('网络宣传'), ('社区患者');
insert into supplier (name) VALUES ('广州白云药厂'),('云南白药药厂');
insert into instock_way (name) VALUES ('采购入库'),('公益捐赠');
insert into outstock_way (name) VALUES ('科室领用'),('退货出库'),('报损出库');
insert into cuvette_color (name) VALUES ('红'), ('黑'), ('紫'),('蓝'), ('黄'), ('绿'),('灰'), ('橙');
INSERT INTO charge_project_type (id,name) VALUES (1,'西/成药处方'),(2,'中药处方'),(3,'检验医嘱'),(4,'检查医嘱'),(5,'材料费用'),(6,'其他费用'),(7,'治疗医嘱');
insert into examination_organ (name) values 
('肝'),
('胆'),
('胰'),
('脾'),
('肾'),
('肾上腺'),
('腹腔'),
('盆腔'),
('肠管'),
('阑尾'),
('左下腹'),
('右下腹'),
('腹腔淋巴结'),
('腹膜后'),
('左上腹'),
('胃'),
('胆囊'),
('残余尿'),
('前列腺（经腹）'),
('肾'),
('输尿管'),
('膀胱'),
('残余尿'),
('胃窗（口服造影剂）');

insert into storehouse (name,clinic_id) VALUES ('专业药房',1);