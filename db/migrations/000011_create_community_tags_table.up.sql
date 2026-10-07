
--
-- Name: community_tags; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.community_tags (
    tags_id integer NOT NULL,
    community_id integer NOT NULL
);


ALTER TABLE public.community_tags OWNER TO nico;

--
-- Name: community_tags community_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.community_tags
    ADD CONSTRAINT community_tags_pkey PRIMARY KEY (tags_id, community_id);

--
-- Name: community_tags community_tags_community_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.community_tags
    ADD CONSTRAINT community_tags_community_id_fkey FOREIGN KEY (community_id) REFERENCES public.communities(id);


--
-- Name: community_tags community_tags_tags_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.community_tags
    ADD CONSTRAINT community_tags_tags_id_fkey FOREIGN KEY (tags_id) REFERENCES public.tags(id);