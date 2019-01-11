--
-- PostgreSQL database dump
--

delete from supplier where id > 0;

--
-- Data for Name: supplier; Type: TABLE DATA; Schema: public; Owner: clinicdb
--

INSERT INTO public.supplier (id, name, created_time, updated_time, deleted_time) VALUES (1, '广州白云药厂', '2018-05-27 00:07:13.652046+08', '2018-05-27 00:07:13.652046+08', NULL);
INSERT INTO public.supplier (id, name, created_time, updated_time, deleted_time) VALUES (2, '云南白药药厂', '2018-05-27 00:07:13.652046+08', '2018-05-27 00:07:13.652046+08', NULL);


--
-- PostgreSQL database dump complete
--

