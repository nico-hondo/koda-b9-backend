
--
-- Name: tags; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.tags OWNER TO nico;

--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: nico
--

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tags_id_seq OWNER TO nico;

--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nico
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;

--
-- Name: tags tags_name_key; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_name_key UNIQUE (name);

--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);