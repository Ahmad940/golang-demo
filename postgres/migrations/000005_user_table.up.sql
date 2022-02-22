CREATE TABLE "users"
(
  id         uuid                 DEFAULT uuid_generate_v4(),
  username   VARCHAR(15) NOT NULL,
  email      VARCHAR     NOT NULL,
  password   VARCHAR     NOT NULL,
  created_at timestamp   NOT NULL DEFAULT NULL,
  updated_at timestamp            DEFAULT null,
  PRIMARY KEY (id)
);
