delete from ranking;
delete from activity;
delete from person_challenge;
delete from challenge;
delete from person;

drop table ranking;
drop table activity;
drop type activity_type;
drop table person_challenge;
drop table challenge;
drop table person;


CREATE TABLE person (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  email VARCHAR UNIQUE NOT NULL,
  phone VARCHAR(11) UNIQUE NOT NULL,
  password VARCHAR NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE challenge (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  description VARCHAR,
  goal INT NOT NULL,
  max_per_day INT NOT NULL,
  start_date DATE NOT NULL,
  finish_date DATE NOT NULL,
  owner_id INTEGER REFERENCES person(id) NOT null,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE person_challenge (
  person_id INTEGER REFERENCES person(id),
  challenge_id INTEGER REFERENCES challenge(id),
  PRIMARY KEY (person_id, challenge_id),
  UNIQUE (person_id, challenge_id)
);

CREATE TYPE activity_type AS ENUM (
  'basketball', 'calisthenics', 'crossfit', 'cycling', 'football', 'gym', 'hike', 'handball', 'martialArts', 'run', 'swimming', 'volleyball', 'other'
);

CREATE TABLE activity (
  id SERIAL PRIMARY KEY,
  type activity_type NOT NULL,
  date TIMESTAMP NOT NULL,
  person_id INTEGER REFERENCES person(id) NOT NULL,
  challenge_id INTEGER REFERENCES challenge(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CHECK (type IN ('basketball', 'calisthenics', 'crossfit', 'cycling', 'football', 'gym', 'hike', 'handball', 'martialArts', 'run', 'swimming', 'volleyball', 'other'))
);

CREATE TABLE ranking (
  id SERIAL PRIMARY KEY,
  completed INT NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  person_id INTEGER REFERENCES person(id),
  challenge_id INTEGER REFERENCES challenge(id),
  UNIQUE (person_id, challenge_id)
);

CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER ranking_update_trigger
BEFORE UPDATE ON ranking
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();


