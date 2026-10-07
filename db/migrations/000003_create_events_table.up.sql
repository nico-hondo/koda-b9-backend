
--
-- Name: events; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.events (
    id integer NOT NULL,
    title character varying(200) NOT NULL,
    description text,
    category character varying(50),
    location character varying(200),
    image_url text,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL,
    max_participants integer,
    status character varying(20) DEFAULT 'upcoming'::character varying,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    community_id integer NOT NULL,
    organizer_id integer NOT NULL
);


ALTER TABLE public.events OWNER TO nico;

--
-- Name: events_id_seq; Type: SEQUENCE; Schema: public; Owner: nico
--

CREATE SEQUENCE public.events_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.events_id_seq OWNER TO nico;

--
-- Name: events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nico
--

ALTER SEQUENCE public.events_id_seq OWNED BY public.events.id;

--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: events events_community_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_community_id_fkey FOREIGN KEY (community_id) REFERENCES public.communities(id);


--
-- Name: events events_organizer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_organizer_id_fkey FOREIGN KEY (organizer_id) REFERENCES public.users(id);