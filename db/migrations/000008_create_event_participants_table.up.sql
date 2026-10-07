
--
-- Name: event_participants; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.event_participants (
    event_id integer NOT NULL,
    user_id integer NOT NULL,
    joined_at timestamp without time zone
);


ALTER TABLE public.event_participants OWNER TO nico;


--
-- Name: event_participants event_participants_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.event_participants
    ADD CONSTRAINT event_participants_pkey PRIMARY KEY (event_id, user_id);


--
-- Name: event_participants event_participants_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.event_participants
    ADD CONSTRAINT event_participants_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id);


--
-- Name: event_participants event_participants_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.event_participants
    ADD CONSTRAINT event_participants_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);