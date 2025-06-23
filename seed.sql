-- default role
INSERT INTO public.roles(id, name, is_deleted, created_by, created_at)
VALUES ('c9bb254c-6e59-4aec-a5ff-4ac093b80ddb', 'default', false, 'system', EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000);

INSERT INTO public.role_acls(id, role_id, access, is_deleted, created_by, created_at)
VALUES (gen_random_uuid(), 'c9bb254c-6e59-4aec-a5ff-4ac093b80ddb', 'get-profile', false, 'system', EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000);

INSERT INTO public.role_acls(id, role_id, access, is_deleted, created_by, created_at)
VALUES (gen_random_uuid(), 'c9bb254c-6e59-4aec-a5ff-4ac093b80ddb', 'create-article', false, 'system', EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000);

INSERT INTO public.role_acls(id, role_id, access, is_deleted, created_by, created_at)
VALUES (gen_random_uuid(), 'c9bb254c-6e59-4aec-a5ff-4ac093b80ddb', 'create-article-published', false, 'system', EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000);

INSERT INTO public.role_acls(id, role_id, access, is_deleted, created_by, created_at)
VALUES (gen_random_uuid(), 'c9bb254c-6e59-4aec-a5ff-4ac093b80ddb', 'update-article', false, 'system', EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000);

INSERT INTO public.role_acls(id, role_id, access, is_deleted, created_by, created_at)
VALUES (gen_random_uuid(), 'c9bb254c-6e59-4aec-a5ff-4ac093b80ddb', 'update-article-published', false, 'system', EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000);

INSERT INTO public.role_acls(id, role_id, access, is_deleted, created_by, created_at)
VALUES (gen_random_uuid(), 'c9bb254c-6e59-4aec-a5ff-4ac093b80ddb', 'get-article-revision', false, 'system', EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000);

INSERT INTO public.role_acls(id, role_id, access, is_deleted, created_by, created_at)
VALUES (gen_random_uuid(), 'c9bb254c-6e59-4aec-a5ff-4ac093b80ddb', 'create-tag', false, 'system', EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000);

INSERT INTO public.role_acls(id, role_id, access, is_deleted, created_by, created_at)
VALUES (gen_random_uuid(), 'c9bb254c-6e59-4aec-a5ff-4ac093b80ddb', 'list-tag', false, 'system', EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000);

-- default user
INSERT INTO "public"."users"("id","username","password", "full_name", "is_deleted","created_by","created_at","updated_at","updated_by")
VALUES
('bfafeee1-1557-4ecd-bfc7-94271ad1246c','admin','$2a$10$Jokvhr.vEdynykSKiZHsm.3cZG7gp2/ah0Z6p4.4kDwVACS.3BBKG', 'John Doe', FALSE,'bfafeee1-1557-4ecd-bfc7-94271ad1246c',EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000,NULL,NULL);

INSERT INTO "public"."user_roles"("id","role_id","user_id","is_deleted","created_by","created_at","updated_at","updated_by")
VALUES
('bbaf231a-4f5e-4fe7-bf2f-8a000938687a','c9bb254c-6e59-4aec-a5ff-4ac093b80ddb','bfafeee1-1557-4ecd-bfc7-94271ad1246c',FALSE,'bfafeee1-1557-4ecd-bfc7-94271ad1246c',EXTRACT(EPOCH FROM NOW() AT TIME ZONE 'UTC') * 1000,NULL,NULL);