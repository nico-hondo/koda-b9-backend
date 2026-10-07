--
-- PostgreSQL database dump
--

\restrict pboymrODEAdcGTmCogyJVTD71vRtb7255SrzyIqawyc6u7Bp1GrmT8tQFyxVmkw

-- Dumped from database version 16.15 (Debian 16.15-1.pgdg13+2)
-- Dumped by pg_dump version 18.6 (Ubuntu 18.6-0ubuntu0.26.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: communities; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.communities VALUES
	(1, 'Bandung Go Community', 'bandung-go-community', 'The premier Go programming community in Bandung — weekly meetups, workshops, and mentoring for Gophers at all levels.', 'Technology', 'https://images.unsplash.com/photo-1518770660439-4636190af475', 'Bandung', true, '2026-10-01 15:12:20.692123+00'),
	(2, 'Jakarta AI & ML Club', 'jakarta-ai-ml-club', 'Researchers, practitioners, and enthusiasts exploring machine learning, LLMs, and the future of AI in Indonesia.', 'AI', 'https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe', 'Jakarta', true, '2026-10-01 15:12:20.692123+00'),
	(3, 'Indonesia Frontend Devs', 'indonesia-frontend-devs', 'The largest frontend community in Indonesia. React, Vue, Svelte, performance, accessibility — all things frontend.', 'Technology', 'https://images.unsplash.com/photo-1555066931-4365d14bab8c', 'Indonesia', true, '2026-10-01 15:12:20.692123+00'),
	(4, 'Product Minds Indonesia', 'product-minds-indonesia', 'Product managers, designers, and builders sharing frameworks, tools, and career journeys.', 'Business', 'https://images.unsplash.com/photo-1531403009284-440f080d1e12', 'Indonesia', true, '2026-10-01 15:12:20.692123+00'),
	(5, 'Creative Coders Collective', 'creative-coders-collective', 'Design and code intersect here — generative art, design systems, creative development, and beautiful UI.', 'Design', 'https://images.unsplash.com/photo-1507238691740-187a5b1d37b8', 'Indonesia', true, '2026-10-01 15:12:20.692123+00'),
	(6, 'Startup Founders Circle', 'startup-founders-circle', 'Early-stage founders in Indonesia supporting each other through fundraising, product, and growth challenges.', 'Business', 'https://images.unsplash.com/photo-1559136555-9303baea8ebd', 'Indonesia', true, '2026-10-01 15:12:20.692123+00'),
	(7, 'Data Science ID', 'data-science-id', 'Data scientists and analysts across Indonesia — sharing datasets, papers, and practical ML projects.', 'AI', 'https://images.unsplash.com/photo-1551288049-bebda4e38f71', 'Indonesia', true, '2026-10-01 15:12:20.692123+00'),
	(8, 'Music Tech Community', 'music-tech-community', 'Producers, sound engineers, and musicians who love blending technology with music creation.', 'Music', 'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4', 'Indonesia', true, '2026-10-01 15:12:20.692123+00');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.users VALUES
	(4, 'Nico Hondo', 'nicohondo01@gmail.com', '$argon2id$v=19$m=65536,t=2,p=2$Nk0P8aWDq3XzTZ07IBvJ7g$KHL2q2uN5sXn9B5kO5D5eNnwHiBi4PdK+/JnbE8M94s', 'nicohondo-pict-url.com', 'Hidup Jokowi', 'Cibubur, Jawa Barat', 'attendee', 'Fullstack Web Developer', 'Nashta', '2026-09-27 12:50:27.551802', '2026-09-28 04:28:09.593205'),
	(7, 'Given Talenta Grib Jaya', 'given@gmail.com', '$argon2id$v=19$m=65536,t=2,p=2$Cdo6lb1514zi84F0ywmY+w$TxwM8IwlRABpdeTXKbMVYz0NdmOOwozgLQttGLETJUQ', 'given-pict-url.com', 'Hidup Hercules', 'Cibubur, Jawa Barat', 'attendee', 'Fullstack Web Developer', 'Koda', '2026-09-28 04:31:39.081536', '2026-09-28 04:33:51.114635'),
	(8, 'Given Talenta Grib Jaya', 'giventok01@gmail.com', '$argon2id$v=19$m=65536,t=2,p=2$TudFVd5LOnPcpZ3YsDOZUQ$LxBfvURlKH/fJc8julH5LHDnV/OxhRMOD7a36xpYsm0', 'given-pict-url.com', 'Hidup Hercules', 'Cibubur, Jawa Barat', 'attendee', 'Fullstack Web Developer', 'Grib Jaya Comunity', '2026-09-28 04:48:37.637289', '2026-09-28 04:52:26.514668'),
	(1, 'Habib Rizki', 'habib.organizer@gmail.com', 'pass123', NULL, NULL, NULL, 'organizer', 'Software Engineer', 'Koda', '2026-09-26 11:39:33.772656', NULL),
	(2, 'Indah Fitriyani', 'indah@gmail.com', 'pass123', NULL, NULL, NULL, 'attendee', 'Tech Lead', 'Popeye', '2026-09-26 11:39:46.625179', NULL),
	(6, 'Fajar Works', 'fajar@gmail.com', '$argon2id$v=19$m=65536,t=2,p=2$/hmvf9Z7m0aFTu2ZHeN3aQ$XcoH1KY4W3LejOLVCw4cfDJltajRuvtOBlnrKK9eizI', NULL, NULL, NULL, 'attendee', 'Frontend Developer', 'Fore', '2026-09-27 23:29:22.630731', NULL),
	(9, 'Victor Dom', 'drdom@gmail.com', '$argon2id$v=19$m=65536,t=2,p=2$DSzGkh6BEB6QJBCPotid7A$LVVOwCb2F9dyneIciFM5c7T9r6FCNMUmN/H71Lwhhgg', NULL, NULL, NULL, 'attendee', 'UI/UX Designer', 'Koda', '2026-09-30 15:23:05.750943', NULL);


--
-- Data for Name: community_members; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.community_members VALUES
	(1, 1, 'admin', '2026-01-10 02:00:00+00'),
	(1, 2, 'member', '2026-02-15 03:30:00+00'),
	(1, 4, 'member', '2026-03-01 07:00:00+00'),
	(2, 6, 'admin', '2026-01-12 04:00:00+00'),
	(2, 7, 'member', '2026-02-20 01:45:00+00'),
	(2, 9, 'member', '2026-03-10 09:20:00+00'),
	(3, 4, 'admin', '2026-01-05 06:00:00+00'),
	(3, 1, 'member', '2026-01-20 08:10:00+00'),
	(3, 2, 'member', '2026-02-01 04:00:00+00'),
	(3, 8, 'member', '2026-03-15 02:30:00+00'),
	(4, 2, 'admin', '2026-01-15 03:00:00+00'),
	(4, 6, 'member', '2026-02-10 07:15:00+00'),
	(4, 9, 'member', '2026-03-05 11:00:00+00'),
	(5, 4, 'member', '2026-02-18 05:00:00+00'),
	(5, 7, 'member', '2026-03-12 03:00:00+00'),
	(6, 9, 'admin', '2026-01-08 09:00:00+00'),
	(6, 6, 'member', '2026-02-22 06:45:00+00'),
	(7, 6, 'member', '2026-02-05 02:15:00+00'),
	(7, 8, 'member', '2026-03-08 04:30:00+00'),
	(8, 7, 'admin', '2026-01-25 10:00:00+00'),
	(8, 1, 'member', '2026-02-28 13:00:00+00');


--
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.tags VALUES
	(1, 'Product'),
	(2, 'Frontend'),
	(3, 'Technology'),
	(4, 'Cloud'),
	(5, 'Workshop'),
	(6, 'Design'),
	(7, 'Data'),
	(8, 'Programming'),
	(9, 'Backend'),
	(10, 'Architecture'),
	(11, 'Security'),
	(12, 'React'),
	(13, 'Online'),
	(14, 'AI'),
	(15, 'Machine Learning'),
	(16, 'Mobile'),
	(17, 'Flutter'),
	(18, 'UI/UX'),
	(19, 'Conference'),
	(20, 'DevOps');


--
-- Data for Name: community_tags; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.community_tags VALUES
	(3, 1),
	(8, 1),
	(14, 2),
	(3, 2),
	(3, 3),
	(8, 3),
	(1, 4),
	(6, 5),
	(8, 5),
	(1, 6),
	(14, 7),
	(8, 7),
	(3, 8);


--
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.events VALUES
	(1, 'Go Concurrency Workshop', 'Deep dive into goroutines, channels, and concurrent patterns in Go.', 'Technology', 'Bandung', 'https://images.unsplash.com/photo-1518770660439-4636190af475', '2026-08-22 07:00:00', '2026-08-22 12:00:00', 100, 'upcoming', '2026-10-01 15:46:13.843946+00', '2026-10-01 15:46:13.843946+00', 1, 1),
	(2, 'AI Product Design Summit', 'Exploring machine learning, LLMs, and the future of AI in product design.', 'AI', 'Jakarta', 'https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe', '2026-09-05 08:30:00', '2026-09-05 16:00:00', 300, 'upcoming', '2026-10-01 15:46:13.843946+00', '2026-10-01 15:46:13.843946+00', 2, 6),
	(3, 'Frontend Craft Conference', 'Annual conference covering React, Vue, Svelte, web performance, and accessibility.', 'Technology', 'Bandung', 'https://images.unsplash.com/photo-1555066931-4365d14bab8c', '2026-10-12 09:00:00', '2026-10-12 17:00:00', 200, 'upcoming', '2026-10-01 15:46:13.843946+00', '2026-10-01 15:46:13.843946+00', 3, 4),
	(4, 'Product Management Masterclass', 'Learn product frameworks, strategy, user research, and career growth in tech.', 'Business', 'Online Event', 'https://images.unsplash.com/photo-1531403009284-440f080d1e12', '2026-09-18 10:00:00', '2026-09-18 13:00:00', 150, 'upcoming', '2026-10-01 15:46:13.843946+00', '2026-10-01 15:46:13.843946+00', 4, 2),
	(5, 'Music Production Bootcamp', 'Hands-on workshop for producers, sound engineers, and musicians blending tech & music.', 'Music', 'Surabaya', 'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4', '2026-11-03 13:00:00', '2026-11-03 18:00:00', 60, 'upcoming', '2026-10-01 15:46:13.843946+00', '2026-10-01 15:46:13.843946+00', 8, 7),
	(6, 'Startup Pitch Night', 'Early-stage founders pitch to investors and get constructive feedback.', 'Business', 'Jakarta', 'https://images.unsplash.com/photo-1559136555-9303baea8ebd', '2026-09-28 18:30:00', '2026-09-28 21:00:00', 120, 'upcoming', '2026-10-01 15:46:13.843946+00', '2026-10-01 15:46:13.843946+00', 6, 9);


--
-- Data for Name: event_participants; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.event_participants VALUES
	(1, 1, '2026-08-01 09:00:00'),
	(1, 2, '2026-08-02 10:15:00'),
	(1, 4, '2026-08-05 14:30:00'),
	(1, 6, '2026-08-10 11:20:00'),
	(2, 6, '2026-08-15 08:30:00'),
	(2, 7, '2026-08-18 13:45:00'),
	(2, 2, '2026-08-22 09:10:00'),
	(3, 4, '2026-09-01 10:00:00'),
	(3, 1, '2026-09-03 11:30:00'),
	(3, 8, '2026-09-05 15:20:00'),
	(4, 2, '2026-08-25 14:00:00'),
	(4, 6, '2026-08-28 09:45:00'),
	(4, 9, '2026-09-02 17:10:00'),
	(5, 7, '2026-09-10 12:00:00'),
	(5, 1, '2026-09-12 18:30:00'),
	(6, 9, '2026-09-01 08:00:00'),
	(6, 4, '2026-09-04 13:15:00'),
	(6, 6, '2026-09-08 19:00:00'),
	(6, 8, '2026-09-15 10:40:00'),
	(2, 9, '2026-10-04 11:40:12.30134'),
	(2, 8, '2026-10-04 21:25:29.631804');


--
-- Data for Name: speakers; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.speakers VALUES
	(1, 'Budi Raharjo', 'Senior Frontend Engineer', 'Tokopedia'),
	(2, 'Sarah Drasner', 'VP of Developer Experience', 'Netlify'),
	(3, 'Kenny Santana', 'Backend Architect', 'Gojek'),
	(4, 'Dr. Andi Wijaya', 'AI Research Lead', 'Telkom Indonesia'),
	(5, 'Grace Hopper Jr', 'Cloud & DevOps Engineer', 'AWS Indonesia'),
	(6, 'Eko Kurniawan', 'Mobile Lead', 'Traveloka'),
	(7, 'Fadhlan Hamami', 'Senior Product Manager', 'Bukalapak'),
	(8, 'Pravina Putri', 'Security Engineer', 'Bhinneka'),
	(9, 'Rio Permana', 'Fullstack Developer', 'Dicoding'),
	(10, 'Anindya Putri', 'Lead Data Scientist', 'GoTo Group'),
	(11, 'Veronica Tan', 'DevOps & Cloud Engineer', 'GCP Indonesia'),
	(12, 'Dimas Prasetyo', 'Google Developer Expert', 'Google');


--
-- Data for Name: event_speakers; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.event_speakers VALUES
	(3, 1),
	(9, 1),
	(4, 2),
	(7, 2),
	(1, 3),
	(2, 3),
	(7, 4),
	(9, 5),
	(7, 6),
	(10, 6);


--
-- Data for Name: event_tags; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.event_tags VALUES
	(9, 1),
	(8, 1),
	(14, 2),
	(6, 2),
	(2, 3),
	(12, 3),
	(1, 4),
	(3, 5),
	(1, 6);


--
-- Data for Name: notification_type; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.notification_type VALUES
	(1, 'Registrasi', 'BsBell', 'bg-light-green', 'green', '2026-09-27 22:47:57.976719'),
	(2, 'Testimoni', 'IoChatboxOutline', 'bg-violet-400', 'violet', '2026-10-01 20:31:21.61994');


--
-- Data for Name: notifications; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.notifications VALUES
	(1, 6, 1, 'Registration Confirmed', 'Congratulations your account has been successfully created!', NULL),
	(2, 7, 1, 'Registration Confirmed', 'Congratulations your account has been successfully created!', NULL),
	(3, 8, 1, 'Registration Confirmed', 'Congratulations your account has been successfully created!', NULL),
	(4, 9, 1, 'Registration Confirmed', 'Congratulations your account has been successfully created!', NULL),
	(5, 9, 2, 'Testimonial Confirmed', 'Congratulations your Testimonial has been successfully created!', NULL),
	(6, 9, 1, 'Join Event Confirmed', 'Hooray, You''re registered for AI Product Design Summit on 2026-09-05 08:30:00 +0000 UTC', NULL),
	(7, 9, 1, 'UnJoin Event has been Successful', 'It''s a shame you left the AI Product Design Summit event on 2026-09-05 08:30:00 +0000 UTC.', NULL),
	(8, 9, 1, 'Join Event Confirmed', 'Hooray, You''re registered for AI Product Design Summit on 2026-09-05 08:30:00 +0000 UTC', NULL),
	(9, 8, 1, 'Join Event Confirmed', 'Hooray, You''re registered for AI Product Design Summit on 2026-09-05 08:30:00 +0000 UTC', NULL),
	(10, 8, 2, 'Testimonial Confirmed', 'Congratulations your Testimonial has been successfully created!', NULL);


--
-- Data for Name: testimonies; Type: TABLE DATA; Schema: public; Owner: nico
--

INSERT INTO public.testimonies VALUES
	(4, 1, 'EventHub benar-benar membantu saya menemukan komunitas teknologi lokal yang aktif dan insightful!', '2026-09-26 17:28:27.603363'),
	(5, 2, 'Sangat mudah mendaftar event dan bergabung dengan klub AI lewat platform ini. Top banget!', '2026-09-28 17:28:27.603363'),
	(6, 6, 'Fitur pengingat upcoming events-nya sangat berguna, jadi tidak pernah ketinggalan jadwal meetup lagi.', '2026-09-30 17:28:27.603363'),
	(7, 9, 'Event Golang sangat berkualitas, dengan pembicara yang sudah berpengalaman sehingga memberikan saya ilmu baru.', '2026-10-01 20:50:48.522958'),
	(8, 8, 'Sungguh Eventhub sangat bagus untuk menemukan event - event keren', '2026-10-04 21:30:42.521871');


--
-- Name: communities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: nico
--

SELECT pg_catalog.setval('public.communities_id_seq', 8, true);


--
-- Name: events_id_seq; Type: SEQUENCE SET; Schema: public; Owner: nico
--

SELECT pg_catalog.setval('public.events_id_seq', 6, true);


--
-- Name: notification_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: nico
--

SELECT pg_catalog.setval('public.notification_type_id_seq', 2, true);


--
-- Name: notifications_id_seq; Type: SEQUENCE SET; Schema: public; Owner: nico
--

SELECT pg_catalog.setval('public.notifications_id_seq', 10, true);


--
-- Name: speakers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: nico
--

SELECT pg_catalog.setval('public.speakers_id_seq', 12, true);


--
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: nico
--

SELECT pg_catalog.setval('public.tags_id_seq', 20, true);


--
-- Name: testimonies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: nico
--

SELECT pg_catalog.setval('public.testimonies_id_seq', 8, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: nico
--

SELECT pg_catalog.setval('public.users_id_seq', 9, true);


--
-- PostgreSQL database dump complete
--

\unrestrict pboymrODEAdcGTmCogyJVTD71vRtb7255SrzyIqawyc6u7Bp1GrmT8tQFyxVmkw

