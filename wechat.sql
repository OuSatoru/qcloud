CREATE TABLE accesstoken
(
  id          INTEGER PRIMARY KEY NOT NULL,
  jikan       TIMESTAMP           NOT NULL,
  accesstoken VARCHAR(512),
  expiresin   INTEGER,
  errcode     INTEGER,
  errmsg      VARCHAR(20)
);

CREATE SEQUENCE addone
  START WITH 1
  INCREMENT BY 1
  NO MAXVALUE
  NO MINVALUE
  CACHE 1;

ALTER TABLE accesstoken ALTER COLUMN id SET DEFAULT nextval('addone');

SELECT setval('addone', 1, FALSE);

SELECT * FROM accesstoken;

DELETE FROM accesstoken WHERE id BETWEEN 2 AND (select max(id) FROM accesstoken);

SELECT accesstoken
FROM accesstoken
WHERE id = (SELECT max(id)
            FROM
              (SELECT *
               FROM accesstoken
               WHERE accesstoken IS NOT NULL) a)