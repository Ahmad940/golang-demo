CREATE TABLE IF NOT EXISTS Todo
(
  id        uuid                 DEFAULT uuid_generate_v4(),
  title     VARCHAR(15) NOT NULL,
  completed boolean     NOT NULL DEFAULT false,
  date      timestamp   NOT NULL,
  PRIMARY KEY (id)
);
