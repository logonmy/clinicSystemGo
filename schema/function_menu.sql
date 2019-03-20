--
-- PostgreSQL database dump
--
delete from public.function_menu where id > 0;

--
-- Data for Name: function_menu; Type: TABLE DATA; Schema: public; Owner: clinicdb
--

INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (1, 0, NULL, '/treatment', NULL, '就诊流程', true, 0, '2018-07-11 22:24:34.848997+08', '2018-07-11 22:24:34.848997+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (2, 0, NULL, '/clinic', NULL, '诊所管理', true, 1, '2018-07-11 22:25:23.738263+08', '2018-07-11 22:25:23.738263+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (3, 0, NULL, '/finance', NULL, '财务管理', true, 2, '2018-07-11 22:25:44.538366+08', '2018-07-11 22:25:44.538366+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (4, 0, NULL, '/setting', NULL, '设置管理', true, 3, '2018-07-11 22:26:12.034463+08', '2018-07-11 22:26:12.034463+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (5, 0, NULL, '/platform', NULL, '平台管理', true, 4, '2018-07-11 22:26:46.117069+08', '2018-07-11 22:26:46.117069+08', NULL, '02');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (6, 1, 1, '/treatment/registration', '/static/icons/patient.svg', '就诊人登记', true, 0, '2018-07-11 22:28:19.965356+08', '2018-07-11 22:28:19.965356+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (7, 1, 1, '/treatment/triage/triage', '/static/icons/triage.svg', '预约分诊', true, 1, '2018-07-11 22:28:58.696536+08', '2018-07-11 22:28:58.696536+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (8, 1, 1, '/treatment/admission', '/static/icons/admission.svg', '医生接诊', true, 2, '2018-07-11 22:29:24.660641+08', '2018-07-11 22:29:24.660641+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (9, 1, 1, '/treatment/charge', '/static/icons/charge.svg', '收费管理', true, 3, '2018-07-11 22:30:31.942817+08', '2018-07-11 22:30:31.942817+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (10, 1, 1, '/treatment/drugdelivery', '/static/icons/drugdelivery.svg', '门诊发药', true, 4, '2018-07-11 22:30:48.912694+08', '2018-07-11 22:30:48.912694+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (11, 1, 1, '/treatment/exam', '/static/icons/exam.svg', '检查', true, 5, '2018-07-11 22:31:04.951655+08', '2018-07-11 22:31:04.951655+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (13, 1, 1, '/treatment/treat', '/static/icons/treat.svg', '治疗', true, 7, '2018-07-11 22:33:19.216376+08', '2018-07-11 22:33:19.216376+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (14, 1, 1, '/treatment/drugretail', '/static/icons/drugretail.svg', '药品零售', true, 8, '2018-07-11 22:33:32.514037+08', '2018-07-11 22:33:32.514037+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (15, 1, 2, '/clinic', '/static/icons/department.svg', '科室管理', true, 0, '2018-07-11 22:34:23.549898+08', '2018-07-11 22:34:23.549898+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (16, 1, 2, '/clinic/doctor', '/static/icons/doctor.svg', '医生管理', true, 1, '2018-07-11 22:34:40.029441+08', '2018-07-11 22:34:40.029441+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (17, 1, 2, '/clinic/schedule', '/static/icons/schedule.svg', '排班管理', true, 2, '2018-07-11 22:34:52.273081+08', '2018-07-11 22:34:52.273081+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (12, 1, 1, '/treatment/inspect', '/static/icons/inspect.svg', '检验', true, 6, '2018-07-11 22:31:18.770879+08', '2018-07-11 22:31:18.770879+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (18, 1, 2, '/clinic/pharmacy', '/static/icons/pharmacy.svg', '药房管理', true, 3, '2018-07-11 22:35:11.100337+08', '2018-07-11 22:35:11.100337+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (19, 1, 2, '/clinic/consumable', '/static/icons/consumable.svg', '耗材管理', true, 4, '2018-07-11 22:36:56.297531+08', '2018-07-11 22:36:56.297531+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (20, 1, 2, '/clinic/patient', '/static/icons/patient2.svg', '患者管理', true, 5, '2018-07-11 22:37:13.129473+08', '2018-07-11 22:37:13.129473+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (21, 1, 3, '/finance', '/static/icons/expenseReport.svg', '费用报表', true, 0, '2018-07-11 22:37:36.959645+08', '2018-07-11 22:37:36.959645+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (23, 1, 3, '/finance/analysis', '/static/icons/medicalReport.svg', '分析类报表', true, 2, '2018-07-11 22:38:13.249847+08', '2018-07-11 22:38:13.249847+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (22, 1, 3, '/finance/template', '/static/icons/medicalReport.svg', '医用报表', true, 1, '2018-07-11 22:37:52.725834+08', '2018-07-11 22:37:52.725834+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (24, 1, 3, '/finance/stock', '/static/icons/medicalReport.svg', '进销存统计', true, 3, '2018-07-11 22:39:49.142892+08', '2018-07-11 22:39:49.142892+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (25, 1, 4, '/setting/chargeItemSetting', '/static/icons/charge2.svg', '收费项目设置', true, 0, '2018-07-11 22:40:45.169624+08', '2018-07-11 22:40:45.169624+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (26, 1, 4, '/setting/template', '/static/icons/template.svg', '模板设置', true, 1, '2018-07-11 22:41:10.167527+08', '2018-07-11 22:41:10.167527+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (28, 1, 4, '/setting/userRights', '/static/icons/userRights.svg', '用户权限', true, 3, '2018-07-11 22:41:55.765246+08', '2018-07-11 22:41:55.765246+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (27, 1, 4, '/setting/permissionGroup', '/static/icons/permissionGroup.svg', '权限分组', true, 2, '2018-07-11 22:41:32.191884+08', '2018-07-11 22:41:32.191884+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (31, 2, 25, '/setting/chargeItemSetting/wMedicinePrescription', NULL, '西/成药处方', true, 0, '2018-07-11 22:44:09.256733+08', '2018-07-11 22:44:09.256733+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (32, 2, 25, '/setting/chargeItemSetting/cMedicinePrescription', NULL, '中药处方', true, 1, '2018-07-11 22:44:21.383082+08', '2018-07-11 22:44:21.383082+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (33, 2, 25, '/setting/chargeItemSetting/inspectionPhysician', NULL, '检验医嘱', true, 2, '2018-07-11 22:44:35.37431+08', '2018-07-11 22:44:35.37431+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (34, 2, 25, '/setting/chargeItemSetting/testItems', NULL, '检验项目', true, 3, '2018-07-11 22:44:47.813224+08', '2018-07-11 22:44:47.813224+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (35, 2, 25, '/setting/chargeItemSetting/checkAdvice', NULL, '检查医嘱', true, 4, '2018-07-11 22:45:04.244786+08', '2018-07-11 22:45:04.244786+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (36, 2, 25, '/setting/chargeItemSetting/treatDoctor', NULL, '治疗医嘱', true, 5, '2018-07-11 22:45:17.11183+08', '2018-07-11 22:45:17.11183+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (37, 2, 25, '/setting/chargeItemSetting/meterialCosts', NULL, '材料费用', true, 6, '2018-07-11 22:45:30.968364+08', '2018-07-11 22:45:30.968364+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (38, 2, 25, '/setting/chargeItemSetting/otherFee', NULL, '其他费用', true, 7, '2018-07-11 22:45:42.804974+08', '2018-07-11 22:45:42.804974+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (39, 2, 25, '/setting/chargeItemSetting/medicalTreatmentItems', NULL, '诊疗项目', true, 8, '2018-07-11 22:45:56.022988+08', '2018-07-11 22:45:56.022988+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (40, 2, 26, '/setting/template/medicalRecordTemplate', NULL, '病历模板', true, 0, '2018-07-11 22:46:24.197573+08', '2018-07-11 22:46:24.197573+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (41, 2, 26, '/setting/template/inspectionTemplate', NULL, '检验模板', true, 1, '2018-07-11 22:46:36.967684+08', '2018-07-11 22:46:36.967684+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (42, 2, 26, '/setting/template/checkTemplate', NULL, '检查模板', true, 2, '2018-07-11 22:46:48.046385+08', '2018-07-11 22:46:48.046385+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (43, 2, 26, '/setting/template/treatTemplate', NULL, '治疗模板', true, 3, '2018-07-11 22:47:00.178672+08', '2018-07-11 22:47:00.178672+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (44, 2, 26, '/setting/template/wMedicinePrescriptionTemplate', NULL, '西/成药处方模板', true, 4, '2018-07-11 22:47:13.923306+08', '2018-07-11 22:47:13.923306+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (45, 2, 26, '/setting/template/cMedicinePrescriptionTemplate', NULL, '中药处方模板', true, 5, '2018-07-11 22:47:27.144127+08', '2018-07-11 22:47:27.144127+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (46, 2, 26, '/setting/template/checkReportTemplate', NULL, '检查报告模板', true, 6, '2018-07-30 22:25:51.730681+08', '2018-07-30 22:25:51.730681+08', NULL, '01');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (30, 1, 5, '/platform/business', '/static/icons/business.svg', '业务管理', true, 1, '2018-07-11 22:42:50.245297+08', '2018-07-11 22:42:50.245297+08', NULL, '02');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (48, 1, 5, '/platform/finance/dailyReport', '/static/icons/business.svg', '财务分析', true, 2, '2018-08-04 15:53:01.111387+08', '2018-08-04 15:53:01.111387+08', NULL, '02');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (29, 1, 5, '/platform/clinique/add', '/static/icons/clinic.svg', '诊所管理', true, 0, '2018-07-11 22:42:28.390266+08', '2018-07-11 22:42:28.390266+08', NULL, '02');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (49, 1, 5, '/platform/operation/totalAmount', '/static/icons/business.svg', '运营分析', true, 3, '2018-08-04 15:53:25.718066+08', '2018-08-04 15:53:25.718066+08', NULL, '02');
INSERT INTO public.function_menu (id, level, parent_function_menu_id, url, icon, name, status, weight, created_time, updated_time, deleted_time, ascription) VALUES (50, 1, 5, '/platform/account/add', '/static/icons/business.svg', '账号设置', true, 4, '2018-08-04 15:53:42.589096+08', '2018-08-04 15:53:42.589096+08', NULL, '02');



--
-- PostgreSQL database dump complete
--

