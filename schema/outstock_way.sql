--
-- PostgreSQL database dump
--

delete from outstock_way where id > 0;

--
-- Data for Name: outstock_way; Type: TABLE DATA; Schema: public; Owner: clinicdb
--

INSERT INTO public.outstock_way (id, name, created_time, updated_time, deleted_time) VALUES (1, '科室领用', '2018-05-27 00:07:13.652046+08', '2018-05-27 00:07:13.652046+08', NULL);
INSERT INTO public.outstock_way (id, name, created_time, updated_time, deleted_time) VALUES (2, '退货出库', '2018-05-27 00:07:13.652046+08', '2018-05-27 00:07:13.652046+08', NULL);
INSERT INTO public.outstock_way (id, name, created_time, updated_time, deleted_time) VALUES (3, '报损出库', '2018-05-27 00:07:13.652046+08', '2018-05-27 00:07:13.652046+08', NULL);
INSERT INTO public.outstock_way (id, name, created_time, updated_time, deleted_time) VALUES (4, '门诊发药', '2018-08-12 23:15:39.526287+08', '2018-08-12 23:15:39.526287+08', NULL);
INSERT INTO public.outstock_way (id, name, created_time, updated_time, deleted_time) VALUES (5, '零售发药', '2018-08-12 23:15:39.526287+08', '2018-08-12 23:15:39.526287+08', NULL);


--
-- PostgreSQL database dump complete
--

