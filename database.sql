CREATE TABLE IF NOT EXISTS users (
  id VARCHAR PRIMARY KEY,
  username VARCHAR NOT NULL,
  full_name VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);

CREATE UNIQUE INDEX uk_users_username_active
ON users (username)
WHERE is_deleted = FALSE;

CREATE TABLE IF NOT EXISTS roles (
  id VARCHAR PRIMARY KEY,
  name VARCHAR NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);

CREATE TABLE IF NOT EXISTS role_acls (
  id VARCHAR PRIMARY KEY,
  role_id VARCHAR NOT NULL,
  access VARCHAR NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR,

  CONSTRAINT fk_role_acls_roles FOREIGN KEY (role_id) REFERENCES roles (id)
);

CREATE TABLE IF NOT EXISTS user_roles (
  id VARCHAR PRIMARY KEY,
  role_id VARCHAR NOT NULL,
  user_id VARCHAR NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR,

  CONSTRAINT fk_user_roles_roles FOREIGN KEY (role_id) REFERENCES roles (id),
  CONSTRAINT fk_user_roles_users FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS articles (
  id VARCHAR PRIMARY KEY,
  content TEXT NOT NULL,
  title VARCHAR NOT NULL,
  author_id VARCHAR NOT NULL,
  status VARCHAR NOT NULL,
  parent_id VARCHAR NULL DEFAULT NULL,
  article_tag_relationship_score FLOAT NOT NULL DEFAULT 0,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR,

  CONSTRAINT fk_articles_users FOREIGN KEY (author_id) REFERENCES users (id),
  CONSTRAINT fk_articles_articles FOREIGN KEY (parent_id) REFERENCES articles (id)
);

CREATE TABLE IF NOT EXISTS tags (
  id VARCHAR PRIMARY KEY,
  name VARCHAR NOT NULL,
  trending_score FLOAT NOT NULL DEFAULT 0,
  usage_count integer NOT NULL DEFAULT 0,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);

CREATE UNIQUE INDEX uk_tags_name_active
ON tags (name)
WHERE is_deleted = FALSE;

CREATE TABLE IF NOT EXISTS article_tags (
  id VARCHAR PRIMARY KEY,
  article_id VARCHAR NOT NULL,
  tag_id VARCHAR NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR,

  CONSTRAINT fk_article_tags_articles FOREIGN KEY (article_id) REFERENCES articles (id),
  CONSTRAINT fk_article_tags_tags FOREIGN KEY (tag_id) REFERENCES tags (id)
);

CREATE UNIQUE INDEX uk_article_tags_active
ON article_tags (article_id, tag_id)
WHERE is_deleted = FALSE;

CREATE TABLE IF NOT EXISTS tag_associations (
  id VARCHAR PRIMARY KEY,
  tag1_id VARCHAR NOT NULL,
  tag2_id VARCHAR NOT NULL,
  score FLOAT NOT NULL DEFAULT 0,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR,

  CONSTRAINT fk_tag_associations_tags_1 FOREIGN KEY (tag1_id) REFERENCES tags (id),
  CONSTRAINT fk_tag_associations_tags_2 FOREIGN KEY (tag2_id) REFERENCES tags (id)
);

CREATE UNIQUE INDEX uk_tag_associations_active
ON tag_associations (tag1_id, tag2_id)
WHERE is_deleted = FALSE;