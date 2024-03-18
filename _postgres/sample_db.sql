DROP TABLE IF EXISTS films CASCADE;
CREATE TABLE films (
  id          SERIAL PRIMARY KEY,
  name        VARCHAR(150) NOT NULL,
  year        INTEGER,
  description TEXT
	  CHECK (LENGTH(description) <= 1000),
  rating      NUMERIC
	  CHECK (rating >= 0 AND rating <= 10)
);

DROP TABLE IF EXISTS actors CASCADE;
CREATE TABLE actors (
  id       SERIAL PRIMARY KEY,
  name     VARCHAR(255) NOT NULL,
  gender   CHAR(1)
	  CHECK (gender IN ('M', 'F') OR gender IS NULL),
  birthday DATE
);

DROP TABLE IF EXISTS films_actors CASCADE;
CREATE TABLE films_actors (
  film_id INTEGER REFERENCES films(id),
  actor_id INTEGER REFERENCES actors(id),
  CONSTRAINT films_actors_pk PRIMARY KEY(film_id, actor_id)
);

DROP TABLE IF EXISTS accounts CASCADE;
CREATE TABLE accounts (
  username VARCHAR(255) NOT NULL PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  role INTEGER CHECK (role IN (0, 1)) NOT NULL
);

INSERT INTO accounts (username, password, role) VALUES
('admin', 'admin_password', 0),
('user', 'user_password', 1),
('user2', 'user2_password', 1);

INSERT INTO films (name, year, description, rating) VALUES
  ('Пираты Карибского моря: Проклятие Черной жемчужины', 2003, 'Фэнтези', 8.4),
  ('Жанна Дюбарри', 2023, 'Драма', 7.1),
  ('Чарли и шоколадная фабрика', 2005, 'Мюзикл', 6.9),
  ('Убийство в Восточном экспрессе', 2017, 'Детектив', 7.1),
  ('Алиса в Стране чудес', 2010, 'Фэнтези', 7.7);

INSERT INTO films (name, year, description, rating) VALUES
  ('Без обид', 2023, 'Комедия', 6.4),
  ('Голодные игры', 2012, 'Фантастика', 7.1),
  ('Голодные игры: И вспыхнет пламя', 2013, 'Фантастика', 7.2),
  ('Джой', 2015, 'Драма', 6.8),
  ('Пассажиры', 2016, 'Фантастика', 7.5);

INSERT INTO films (name, year, description, rating) VALUES
  ('Джон Уик 4', 2023, 'Боевик', 7.6),
  ('На берегу реки', 1986, 'Драма', 6.6),
  ('Джон Уик', 2014, 'Боевик', 7.0),
  ('Матрица: Воскрешение', 2021, 'Фантастика', 8.2),
  ('Джон Уик 3', 2019, 'Боевик', 7.0);

INSERT INTO actors (name, gender, birthday) VALUES
	('Джонни Депп', 'M', '1963-06-09'),
	('Дженнифер Лоуренс', 'F', '1990-08-15'),
	('Киану Ривз', 'M', '1964-09-02');

INSERT INTO films_actors (film_id, actor_id) VALUES
  (1, 1),
  (2, 1),
  (3, 1),
  (4, 1),
  (5, 1),
  (6, 2),
  (7, 2),
  (8, 2),
  (9, 2),
  (10, 2),
  (11, 3),
  (12, 3),
  (13, 3),
  (14, 3),
  (15, 3);
