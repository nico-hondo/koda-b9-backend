
--
-- Name: event_tags; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.event_tags (
    tags_id integer NOT NULL,
    event_id integer NOT NULL
);


ALTER TABLE public.event_tags OWNER TO nico;

--
-- Name: event_tags event_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.event_tags
    ADD CONSTRAINT event_tags_pkey PRIMARY KEY (tags_id, event_id);


--
-- Name: event_tags event_tags_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.event_tags
    ADD CONSTRAINT event_tags_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id);


--
-- Name: event_tags event_tags_tags_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.event_tags
    ADD CONSTRAINT event_tags_tags_id_fkey FOREIGN KEY (tags_id) REFERENCES public.tags(id);