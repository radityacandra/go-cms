CREATE TABLE IF NOT EXISTS users (
  id VARCHAR PRIMARY KEY,
  username VARCHAR UNIQUE,
  password VARCHAR NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_by VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);

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
