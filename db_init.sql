CREATE TABLE person (
      id INT PRIMARY KEY NOT NULL
    , email VARCHAR NOT NULL
    , is_administrator BOOLEAN DEFAULT FALSE
);
CREATE TABLE person_credentials (
      id INT PRIMARY KEY NOT NULL
    , person_id INT NOT NULL REFERENCES person(id) ON DELETE CASCADE
    , email VARCHAR NOT NULL
    , hashed_password VARCHAR NOT NULL
    , salt VARCHAR NOT NULL
);
CREATE TABLE camp (
      id INT PRIMARY KEY NOT NULL
    , title VARCHAR NOT NULL
    , specification VARCHAR
    , start_lat FLOAT NOT NULL
    , start_lng FLOAT NOT NULL
    , start_date DATE NOT NULL
    , end_date DATE
);
CREATE TABLE user_camp (
      id INT PRIMARY KEY NOT NULL
    , person_id INT NOT NULL REFERENCES person(id) ON DELETE CASCADE
    , camp_id INT NOT NULL REFERENCES camp(id) ON DELETE CASCADE
    , is_paid BOOLEAN DEFAULT FALSE
    , is_person_guide BOOLEAN DEFAULT FALSE
);

CREATE FUNCTION authorise_(varchar, varchar) RETURNS boolean AS $$
  BEGIN
    IF 
      (SELECT COUNT(*) FROM person_credentials WHERE email = $1 AND hashed_password = crypt($2, salt)) > 0
      THEN RETURN true;
      ELSE RETURN false;
    END IF;
  END;
  $$ LANGUAGE plpgsql;

CREATE FUNCTION register(varchar, varchar) RETURNS boolean AS $$
  DECLARE
    salt_v VARCHAR;
  BEGIN
    IF 
      (SELECT COUNT(*) FROM person_credentials WHERE email = $1) > 0
      THEN RETURN false;
    END IF;
    salt_v = gen_salt(xdes);
    INSERT INTO person_credentials (email, hashed_password, salt) VALUES ($1, crypt($2, salt_v), salt_v);
    INSERT INTO person (email) VALUES ($1);
    RETURN true;
  END;
  $$ LANGUAGE plpgsql;