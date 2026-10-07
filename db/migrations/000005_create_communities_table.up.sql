--
-- Name: communities; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.communities (
    id integer NOT NULL,
    name character varying(150) NOT NULL,
    slug character varying(150) NOT NULL,
    description text,
    category character varying(50),
    image_url text,
    location character varying(100),
    is_active boolean DEFAULT true,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.communities OWNER TO nico;

--
-- Name: communities_id_seq; Type: SEQUENCE; Schema: public; Owner: nico
--

CREATE SEQUENCE public.communities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.communities_id_seq OWNER TO nico;

--
-- Name: communities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nico
--

ALTER SEQUENCE public.communities_id_seq OWNED BY public.communities.id;

--
-- Name: communities communities_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.communities
    ADD CONSTRAINT communities_pkey PRIMARY KEY (id);

--
-- Name: communities communities_slug_key; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.communities
    ADD CONSTRAINT communities_slug_key UNIQUE (slug);