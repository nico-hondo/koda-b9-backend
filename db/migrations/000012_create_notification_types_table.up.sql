
--
-- Name: notification_type; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.notification_type (
    id integer NOT NULL,
    name character varying(50),
    icon character varying(200),
    bg_color character varying(50),
    icon_color character varying(50),
    created_at timestamp without time zone
);


ALTER TABLE public.notification_type OWNER TO nico;

--
-- Name: notification_type_id_seq; Type: SEQUENCE; Schema: public; Owner: nico
--

CREATE SEQUENCE public.notification_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.notification_type_id_seq OWNER TO nico;

--
-- Name: notification_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nico
--

ALTER SEQUENCE public.notification_type_id_seq OWNED BY public.notification_type.id;

--
-- Name: notification_type notification_type_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.notification_type
    ADD CONSTRAINT notification_type_pkey PRIMARY KEY (id);