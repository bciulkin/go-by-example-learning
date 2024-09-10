DROP TABLE IF EXISTS animal;
CREATE TABLE animal (
  id         VARCHAR(36) NOT NULL,
  name       VARCHAR(30) NOT NULL,
  age        SMALLINT(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO animal
  (id, name, age)
VALUES
  ('1417fdab-6cde-4b8c-b6ec-e5bb68abf3f7', 'Salsa', 4),
  ('6fbf5c0f-4f70-4b7b-be12-5250d6bccaa7', 'Krewetka', 4),
  ('5e7e63fd-d86b-4145-9f3a-f251a6f988dd', 'Wegorz', 2);

