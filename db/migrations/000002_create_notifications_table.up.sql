
--
-- Name: notifications; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.notifications (
    id integer NOT NULL,
    user_id integer NOT NULL,
    type_id integer NOT NULL,
    title character varying(200) NOT NULL,
    message text,
    read_at timestamp without time zone
);


ALTER TABLE public.notifications OWNER TO nico;

--
-- Name: notifications_id_seq; Type: SEQUENCE; Schema: public; Owner: nico
--

CREATE SEQUENCE public.notifications_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.notifications_id_seq OWNER TO nico;

--
-- Name: notifications_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nico
--

ALTER SEQUENCE public.notifications_id_seq OWNED BY public.notifications.id;

--
-- Name: notifications notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);


--
-- Name: notifications notifications_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_type_id_fkey FOREIGN KEY (type_id) REFERENCES public.notification_type(id);


--
-- Name: notifications notifications_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);