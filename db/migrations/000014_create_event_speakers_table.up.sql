
--
-- Name: event_speakers; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.event_speakers (
    speaker_id integer NOT NULL,
    event_id integer NOT NULL
);


ALTER TABLE public.event_speakers OWNER TO nico;

--
-- Name: event_speakers event_speakers_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.event_speakers
    ADD CONSTRAINT event_speakers_pkey PRIMARY KEY (speaker_id, event_id);


--
-- Name: event_speakers event_speakers_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.event_speakers
    ADD CONSTRAINT event_speakers_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id);


--
-- Name: event_speakers event_speakers_speaker_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.event_speakers
    ADD CONSTRAINT event_speakers_speaker_id_fkey FOREIGN KEY (speaker_id) REFERENCES public.speakers(id);