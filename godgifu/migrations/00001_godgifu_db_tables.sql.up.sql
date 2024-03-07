CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS accounts_employee
(
  id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  email varchar(255) NOT NULL UNIQUE,
  password varchar(512) NOT NULL,
  phone_number varchar(50) UNIQUE,
  employment_title text,
  office_address text,
  employment_date_start timestamp with time zone,
  employment_date_end timestamp with time zone,
  verified boolean,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS accounts_identity
(
  id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  first_name text NOT NULL,
  middle_name text NOT NULL,
  last_name text NOT NULL,
  age bigint,
  sex varchar(20),
  gender text,
  height text,
  home_address text,
  brithdate timestamp,
  brithplace text,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS telephone_numbers
(
  id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  account_id uuid NOT NULL DEFAULT uuid_generate_v4(),
  phone_number varchar(50)
);

CREATE TABLE IF NOT EXISTS documents
(
  id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  document_title text NOT NULL,
  author_name text[],
  author_id uuid NOT NULL DEFAULT uuid_generate_v4(),
  description text,
  languages text[],
  security_access_level varchar(50),
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS files
(
  id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  document_id uuid NOT NULL DEFAULT uuid_generate_v4(),
  title text,
  author_name text[],
  authord_id uuid NOT NULL DEFAULT uuid_generate_v4(),
  security_access_level text,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL
);

-- CREATE TABLE IF NOT EXISTS users_device_data
-- (
--   account_id uuid NOT NULL DEFAULT uuid_generate_v4(),
--   ip_address text,
--   user_agent text,
--   operating_system text,
--   screen_resolution text,
--   created_at timestamp with time zone NOT NULL,
--   updated_at timestamp with time zone NOT NULL
-- )

-- CREATE TABLE IF NOT EXISTS tablename
-- (
  
-- )