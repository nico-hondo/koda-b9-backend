
--
-- Name: testimonies; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.testimonies (
    id integer NOT NULL,
    user_id integer NOT NULL,
    comment text,
    created_at timestamp without time zone
);


ALTER TABLE public.testimonies OWNER TO nico;

--
-- Name: testimonies_id_seq; Type: SEQUENCE; Schema: public; Owner: nico
--

CREATE SEQUENCE public.testimonies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.testimonies_id_seq OWNER TO nico;

--
-- Name: testimonies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nico
--

ALTER SEQUENCE public.testimonies_id_seq OWNED BY public.testimonies.id;

--
-- Name: testimonies testimonies_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.testimonies
    ADD CONSTRAINT testimonies_pkey PRIMARY KEY (id);


--
-- Name: testimonies testimonies_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.testimonies
    ADD CONSTRAINT testimonies_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);
