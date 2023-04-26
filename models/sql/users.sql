CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  first_name TEXT,
  last_name TEXT,
  password_hash TEXT,
  dob timestamp(0) with time zone NOT NULL,
  activated bool NOT NULL DEFAULT 'f',
  banned bool NOT NULL DEFAULT 'f',
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  version integer NOT NULL DEFAULT 1
);

INSERT INTO
  users (
    email,
    first_name,
    last_name,
    password_hash,
    dob,
    activated,
    banned
  )
VALUES
  (
    'awdawd@@@@@@oidjawo',
    'Lara',
    'Mundo',
    '2a898a8wd88b8s8nb8sf8342n',
    '2002-04-26 08:32:19+00',
    '0',
    'FALSE'
  );