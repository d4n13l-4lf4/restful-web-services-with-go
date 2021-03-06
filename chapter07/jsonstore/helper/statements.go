package helper

const CREATE_TABLE_WEB_URL = "CREATE TABLE IF NOT EXISTS web_url(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);"
const INSERT_WEB_URL = "INSERT INTO web_url (url) VALUES($1) RETURNING id"
const SELECT_WEB_URL_BY_ID = "SELECT url FROM web_url WHERE id = $1"
const SELECT_BY_WEIGHT_JSONB = "SELECT * FROM \"Package\" WHERE data::jsonb->>'weight'=?"