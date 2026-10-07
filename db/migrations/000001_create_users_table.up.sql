
--
-- Name: user_role; Type: TYPE; Schema: public; Owner: nico
--

CREATE TYPE public.user_role AS ENUM (
    'attendee',
    'organizer',
    'admin'
);


ALTER TYPE public.user_role OWNER TO nico;

--
-- Name: users; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(255) NOT NULL,
    avatar_url text,
    bio text,
    location character varying(100),
    role public.user_role DEFAULT 'attendee'::public.user_role,
    job character varying(30),
    workplace character varying(50),
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO nico;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: nico
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO nico;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nico
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: communities id; Type: DEFAULT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.communities ALTER COLUMN id SET DEFAULT nextval('public.communities_id_seq'::regclass);


--
-- Name: events id; Type: DEFAULT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.events ALTER COLUMN id SET DEFAULT nextval('public.events_id_seq'::regclass);


--
-- Name: notification_type id; Type: DEFAULT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.notification_type ALTER COLUMN id SET DEFAULT nextval('public.notification_type_id_seq'::regclass);


--
-- Name: notifications id; Type: DEFAULT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.notifications ALTER COLUMN id SET DEFAULT nextval('public.notifications_id_seq'::regclass);


--
-- Name: speakers id; Type: DEFAULT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.speakers ALTER COLUMN id SET DEFAULT nextval('public.speakers_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- Name: testimonies id; Type: DEFAULT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.testimonies ALTER COLUMN id SET DEFAULT nextval('public.testimonies_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);

--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);