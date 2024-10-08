CREATE TABLE users (
    "id" uuid PRIMARY KEY,
    "username" varchar UNIQUE,
    "email" varchar UNIQUE,
    "email_status" boolean DEFAULT false,
    "pas_hash" varchar,
    "created_at" timestamp DEFAULT (now())
);

CREATE TABLE admins (
    "id" integer PRIMARY KEY,
    "uid" uuid,
    "name" varchar,
    "created_at" timestamp DEFAULT (now())
);

ALTER TABLE admins ADD FOREIGN KEY ("uid") REFERENCES users ("id");