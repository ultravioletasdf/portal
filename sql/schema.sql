PRAGMA foreign_keys = ON;

CREATE TABLE users (
  id integer PRIMARY KEY,
  email text NOT NULL,
  username text NOT NULL,
  name text NOT NULL,
  password text NOT NULL
);

CREATE INDEX idx_users_id ON users(id);
CREATE UNIQUE INDEX idx_users_email ON users(email);
CREATE UNIQUE INDEX idx_users_username ON users(username);

CREATE TABLE apps (
  id integer PRIMARY KEY,
  owner_id integer NOT NULL,

  name text NOT NULL,
  scope integer NOT NULL,
  FOREIGN KEY(owner_id) REFERENCES users(id)
);
CREATE TABLE allowed_addresses (
  id integer PRIMARY KEY,
  address text NOT NULL,

  app_id integer NOT NULL,
  FOREIGN KEY(app_id) REFERENCES apps(id) ON DELETE CASCADE
);
