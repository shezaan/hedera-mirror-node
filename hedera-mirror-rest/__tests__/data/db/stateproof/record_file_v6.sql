--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.18
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
-- The following line is auto generated by pg_dump and there is no option to disable it. Clearing search_path globally
-- will cause the sql queries in our implementation fail (relation not found). So comment it out.
-- SELECT pg_catalog.set_config('search_path', '', false, 0);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: record_file; Type: TABLE DATA; Schema: public; Owner: mirror_node
--
INSERT INTO public.record_file (name, load_start, load_end, hash, prev_hash, consensus_start, consensus_end, node_account_id, count, digest_algorithm, hapi_version_major, hapi_version_minor, hapi_version_patch, version, file_hash, bytes, index) VALUES
('2022-07-26T02_36_36.804825645Z.rcd.gz', 1658803000, 1658803000, '54307db19d55f291a9b1797f10e366ac3fca31f8e7a9d376bd4b8aed47c3d14dd5b5974c8369eb141beaf56948b86243', 'e7330b0e994486f2c760e61ef21fb008006b1b375286a33b2df5557f2c05dfc63b171e24fed43aa91dc9953e67169ec0', 1658802996804825645, 1658802997993309003, 3, 8, 0, 0, 28, 1, 6, '8a10cc197b1885b43542cfd1aa25b037f1a81d6caaf0b6ed11c106c76df10c914087ac0a2ce4c0d7b4d894b9b3c4ad9d', NULL, 311114);