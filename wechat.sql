CREATE TABLE accesstoken
(
  id          INTEGER PRIMARY KEY NOT NULL,
  jikan       TIMESTAMP            NOT NULL,
  accesstoken VARCHAR(512),
  expiresin   INTEGER,
  errcode     INTEGER,
  errmsg      VARCHAR(20)
);