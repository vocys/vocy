PRAGMA foreign_keys=OFF;
BEGIN;

CREATE TABLE app_config_variables
(
    key           VARCHAR(100) NOT NULL PRIMARY KEY,
    value         TEXT NOT NULL
);


CREATE TABLE kv
(
    "key"  TEXT NOT NULL PRIMARY KEY,
    "value" TEXT NOT NULL
);


COMMIT;
PRAGMA foreign_keys=ON;