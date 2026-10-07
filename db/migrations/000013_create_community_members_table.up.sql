
--
-- Name: community_members; Type: TABLE; Schema: public; Owner: nico
--

CREATE TABLE public.community_members (
    community_id integer NOT NULL,
    user_id integer NOT NULL,
    community_role character varying(20) DEFAULT 'member'::character varying,
    joined_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.community_members OWNER TO nico;

--
-- Name: community_members community_members_pkey; Type: CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.community_members
    ADD CONSTRAINT community_members_pkey PRIMARY KEY (community_id, user_id);

--
-- Name: community_members community_members_community_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.community_members
    ADD CONSTRAINT community_members_community_id_fkey FOREIGN KEY (community_id) REFERENCES public.communities(id) ON DELETE CASCADE;


--
-- Name: community_members community_members_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nico
--

ALTER TABLE ONLY public.community_members
    ADD CONSTRAINT community_members_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;