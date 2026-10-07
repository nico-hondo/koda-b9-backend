
--
-- Name: speakers; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.speakers (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    role character varying(100),
    work_at character varying(100)
);


ALTER TABLE public.speakers OWNER TO nico;

--
-- Name: speakers_id_seq; Type: SEQUENCE; Schema: public; Owner: nico
--

CREATE SEQUENCE public.speakers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.speakers_id_seq OWNER TO nico;

--
-- Name: speakers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nico
--

ALTER SEQUENCE public.speakers_id_seq OWNED BY public.speakers.id;

--
-- Name: speakers speakers_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.speakers
    ADD CONSTRAINT speakers_pkey PRIMARY KEY (id);