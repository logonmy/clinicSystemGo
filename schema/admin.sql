--
-- PostgreSQL database dump
--
delete from admin where id > 0;

--
-- Data for Name: admin; Type: TABLE DATA; Schema: public; Owner: clinicdb
--

INSERT INTO public.admin (id, name, phone, title, username, password, is_clinic_admin, status, created_time, updated_time, deleted_time) VALUES (1, '平台管理员', '13211112222', '平台经理', 'pt_admin', 'e10adc3949ba59abbe56e057f20f883e', true, true, '2018-08-04 16:26:13.950835+08', '2018-08-20 17:59:51.471799+08', NULL);

--
-- PostgreSQL database dump complete
--