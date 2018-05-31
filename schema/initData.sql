
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
('前列腺（经腹）'),
('肾'),
('输尿管'),
('膀胱'),
('残余尿'),
('胃窗（口服造影剂）');

insert into storehouse (name,clinic_id) VALUES ('专业药房',1);

insert into parent_function_menu (url,ascription,name) VALUES 
('/treatment','01','就诊流程'),
('/clinic','01','诊所管理'),
('/finance','01','财务管理'),
('/setting','01','设置管理'),
('/platform','01','平台管理')

insert into children_function_menu (parent_function_menu_id,url,name) VALUES
(1,'/treatment/registration','就诊人登记'), 
(1,'/treatment/triage/triage','预约分诊'), 
(1,'/treatment/admission','医生接诊'),
(1,'/treatment/charge','收费管理'),
(1,'/treatment/drugdelivery','门诊发药'),
(1,'/treatment/exam','检查'),
(1,'/treatment/inspect','检验'),
(1,'/treatment/treat','治疗'),
(1,'/treatment/drugretail','药品零售'),
(2,'/clinic','科室管理'),
(2,'/clinic/doctor','医生管理'),
(2,'/clinic/schedule','排班管理'),
(2,'/clinic/pharmacy','药房管理'),
(2,'/clinic/consumable','耗材管理'),
(2,'/clinic/patient','患者管理'),
(3,'/finance','费用报表'),
(3,'/finance/template','医用报表'),
(4,'/setting/chargeItemSetting/wMedicinePrescription','西/成药处方'),
(4,'/setting/chargeItemSetting/cMedicinePrescription','中药处方'),
(4,'/setting/chargeItemSetting/inspectionPhysician','检验医嘱'),
(4,'/setting/chargeItemSetting/testItems','检验项目'),
(4,'/setting/chargeItemSetting/checkAdvice','检查医嘱'),
(5,'/platform','诊所管理'),
(5,'/platform/business','业务管理');

insert into chief_complaint (name) values 
('发热'),('头痛'),('头晕'),('鼻塞'),('流涕'),('声嘶'),('咽痛'),('咽充血'),('咳嗽'),('咳喘'),('咳痰'),('呼吸困难'),
('恶心'),('乏力'),('反酸'),('腹痛'),('腹胀'),('便秘'),('胸闷'),('腹泻'),('胸痛'),('心悸'),('腰痛'),('腰背痛'),('关节痛'),('尿频'),('尿急'),('尿痛'),('多尿'),
('排尿困难'),('水肿'),('皮疹'),('疱疹'),('红疹');