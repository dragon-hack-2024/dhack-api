CREATE TABLE users (
  id SERIAL PRIMARY KEY, 
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL, 
  weight SMALLINT NOT NULL,
  birth_date DATE NOT NULL, 
  created_at TIMESTAMP NOT NULL DEFAULT(now())
);

CREATE TABLE challenges (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  steps JSONB NOT NULL,
  file_name VARCHAR(255) NOT NULL,
  duration INT NOT NULL
);

CREATE TABLE stats (
  id SERIAL PRIMARY KEY,
  calories_burned INT NOT NULL,
  rpm REAL NOT NULL,
  duration INT NOT NULL,
  score REAL NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT(now()),
  challenge_id INT NOT NULL,
  user_id INT NOT NULL,
  FOREIGN KEY (challenge_id) REFERENCES challenges(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);
