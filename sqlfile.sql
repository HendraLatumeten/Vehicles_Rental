--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4
-- Dumped by pg_dump version 14.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: histories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.histories (
    histories_id integer NOT NULL,
    user_id character varying NOT NULL,
    vehicles_id integer NOT NULL,
    date_order_at date NOT NULL,
    date_order_end date NOT NULL,
    price integer NOT NULL,
    payment character varying NOT NULL,
    status character varying NOT NULL,
    create_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.histories OWNER TO postgres;

--
-- Name: histories_histories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.histories_histories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.histories_histories_id_seq OWNER TO postgres;

--
-- Name: histories_histories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.histories_histories_id_seq OWNED BY public.histories.histories_id;


--
-- Name: myseq; Type: SEQUENCE; Schema: public; Owner: 29music
--

CREATE SEQUENCE public.myseq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.myseq OWNER TO "29music";

--
-- Name: mysequence; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mysequence
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mysequence OWNER TO postgres;

--
-- Name: newsequence; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.newsequence
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.newsequence OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    user_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    username character varying NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    gender character varying(2) NOT NULL,
    address character varying NOT NULL,
    phone_number integer NOT NULL,
    date_of_birth character varying NOT NULL,
    image character varying,
    create_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    role character varying NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: vehicles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vehicles (
    vehicles_id integer NOT NULL,
    name character varying NOT NULL,
    type_vehicles character varying NOT NULL,
    city character varying NOT NULL,
    capacity integer NOT NULL,
    image character varying NOT NULL,
    orders integer NOT NULL,
    create_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.vehicles OWNER TO postgres;

--
-- Name: vehicles_order; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vehicles_order
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vehicles_order OWNER TO postgres;

--
-- Name: vehicles_vehicles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vehicles_vehicles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vehicles_vehicles_id_seq OWNER TO postgres;

--
-- Name: vehicles_vehicles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vehicles_vehicles_id_seq OWNED BY public.vehicles.vehicles_id;


--
-- Name: histories histories_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.histories ALTER COLUMN histories_id SET DEFAULT nextval('public.histories_histories_id_seq'::regclass);


--
-- Name: vehicles vehicles_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicles ALTER COLUMN vehicles_id SET DEFAULT nextval('public.vehicles_vehicles_id_seq'::regclass);


--
-- Data for Name: histories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.histories (histories_id, user_id, vehicles_id, date_order_at, date_order_end, price, payment, status, create_at, updated_at) FROM stdin;
2	e672b0cd-6789-43ee-b698-2aed6931e2e1	4	2022-09-02	2022-10-02	2000000	cash	deliver	2022-09-21 02:48:49.177796+07	2022-09-21 02:48:49.177796+07
3	e672b0cd-6789-43ee-b698-2aed6931e2e1	4	2022-09-02	2022-10-02	2000000	cash	deliver	2022-09-21 02:52:30.10545+07	2022-09-21 02:52:30.10545+07
4	e672b0cd-6789-43ee-b698-2aed6931e2e1	4	2022-09-02	2022-10-02	2000000	cash	deliver	2022-09-21 02:54:05.418404+07	2022-09-21 02:54:05.418404+07
5	e672b0cd-6789-43ee-b698-2aed6931e2e1	5	2022-09-02	2022-10-02	2000000	cash	deliver	2022-09-21 02:54:13.895364+07	2022-09-21 02:54:13.895364+07
6	e672b0cd-6789-43ee-b698-2aed6931e2e1	5	2022-09-02	2022-10-02	2000000	cash	deliver	2022-09-21 02:54:30.638754+07	2022-09-21 02:54:30.638754+07
7	e672b0cd-6789-43ee-b698-2aed6931e2e1	5	2022-09-02	2022-10-02	2000000	cash	deliver	2022-09-21 02:55:17.187827+07	2022-09-21 02:55:17.187827+07
8	e672b0cd-6789-43ee-b698-2aed6931e2e1	6	2022-09-02	2022-10-02	2000000	cash	deliver	2022-09-21 02:55:39.618912+07	2022-09-21 02:55:39.618912+07
9	e672b0cd-6789-43ee-b698-2aed6931e2e1	4	2022-09-02	2022-10-02	2000000	cash	deliver	2022-09-21 02:55:50.48311+07	2022-09-21 02:55:50.48311+07
10	e672b0cd-6789-43ee-b698-2aed6931e2e1	6	2022-09-02	2022-10-02	2000000	cash	deliver	2022-09-21 02:56:03.444726+07	2022-09-21 02:56:03.444726+07
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (user_id, username, email, password, gender, address, phone_number, date_of_birth, image, create_at, updated_at, role) FROM stdin;
97dd217c-aa31-45e0-afa3-1ea1172b27f3	test	test@gmail.com	$2a$10$e30AzwC82laiuxEwFnR.IenZI9jPOnFk9hazn1dKY2x01OnQaURLi	P	Bekasi	334423	20-10-1963	test.jpg	2022-09-24 13:34:57.336956+07	2022-09-24 13:34:57.336956+07	user
0788648f-2b11-44d7-943c-579dec632d72	Hendra	Hendra@gmail.com	$2a$10$umdvWCrKF.6sJAXqN0/9auLPsTOodxnb1Y5fPhsb5InSs3UPy2uPK	P	Bekasi	334423	20-10-1963	test.jpg	2022-09-24 16:13:29.461612+07	2022-09-24 16:13:29.461612+07	user
\.


--
-- Data for Name: vehicles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.vehicles (vehicles_id, name, type_vehicles, city, capacity, image, orders, create_at, updated_at) FROM stdin;
5	vespa	Bike	Jakarta	2	vespa.jpg	3	2022-09-21 02:53:30.945194+07	2022-09-21 02:53:30.945194+07
4	Supra	Bike	Bekasi	1	Fixie.jpg	4	2022-09-21 02:17:57.289241+07	2022-09-21 02:17:57.289241+07
6	Honda Metic	Bike	Bekasi	1	Fixie.jpg	2	2022-09-21 02:55:01.873507+07	2022-09-21 02:55:01.873507+07
\.


--
-- Name: histories_histories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.histories_histories_id_seq', 10, true);


--
-- Name: myseq; Type: SEQUENCE SET; Schema: public; Owner: 29music
--

SELECT pg_catalog.setval('public.myseq', 1, true);


--
-- Name: mysequence; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mysequence', 13, true);


--
-- Name: newsequence; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.newsequence', 2, true);


--
-- Name: vehicles_order; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vehicles_order', 1, false);


--
-- Name: vehicles_vehicles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vehicles_vehicles_id_seq', 6, true);


--
-- Name: histories histories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.histories
    ADD CONSTRAINT histories_pkey PRIMARY KEY (histories_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: vehicles vehicles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicles
    ADD CONSTRAINT vehicles_pkey PRIMARY KEY (vehicles_id);


--
-- Name: histories histories_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.histories
    ADD CONSTRAINT histories_fk FOREIGN KEY (vehicles_id) REFERENCES public.vehicles(vehicles_id);


--
-- PostgreSQL database dump complete
--

