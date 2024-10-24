CREATE TABLE "users_secret"(
    "id" UUID NOT NULL PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "email_status" BOOLEAN NOT NULL,
    "pass_hash" VARCHAR(255) NOT NULL
);

CREATE INDEX "user_public_username_index" ON "users_secret"("username");
CREATE INDEX "user_public_email_index" ON "users_secret"("email");
