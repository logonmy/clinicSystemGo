--
-- PostgreSQL database dump
--

delete from instock_way where id > 0;

--
-- Data for Name: instock_way; Type: TABLE DATA; Schema: public; Owner: clinicdb
--

INSERT INTO public.instock_way (id, name, created_time, updated_time, deleted_time) VALUES (1, '采购入库', '2018-05-27 00:07:13.652046+08', '2018-05-27 00:07:13.652046+08', NULL);
INSERT INTO public.instock_way (id, name, created_time, updated_time, deleted_time) VALUES (2, '公益捐赠', '2018-05-27 00:07:13.652046+08', '2018-05-27 00:07:13.652046+08', NULL);
INSERT INTO public.instock_way (id, name, created_time, updated_time, deleted_time) VALUES (3, '门诊退药', '2018-08-12 22:54:14.528675+08', '2018-08-12 22:54:14.528675+08', NULL);
INSERT INTO public.instock_way (id, name, created_time, updated_time, deleted_time) VALUES (4, '零售退药', '2018-08-12 22:54:14.528675+08', '2018-08-12 22:54:14.528675+08', NULL);

--
-- PostgreSQL database dump complete
--

