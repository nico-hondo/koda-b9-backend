INSERT INTO public.notifications (id, user_id, type_id, title, message, read_at) VALUES
-- NOTIF EVENT JOIN - Upcoming yang besok-besok
(1, 2, 1, 'Berhasil Join Event', 'Kamu berhasil join React Workshop Tomorrow yang dimulai besok!', false, NOW() - INTERVAL '2 hours'),
(2, 2, 1, 'Berhasil Join Event', 'Kamu berhasil join Ngopi Bareng Frontend Dev lusa di Jakarta', false, NOW() - INTERVAL '3 hours');

SELECT setval('public.notifications_id_seq', (SELECT MAX(id) FROM public.notifications), true);