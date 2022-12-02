CREATE TABLE "users" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "username" varchar NOT NULL,
  "age" int NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "createdAt" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "routes" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "admin_id" int NOT NULL,
  "title" varchar NOT NULL,
  "description" text NOT NULL,
  "location" varchar NOT NULL,
  "destination" float NOT NULL,
  "height" float NOT NULL,
  "level" varchar NOT NULL,
  "createdAt" timestamptz NOT NULL DEFAULT 'now()',
  "updatedAt" timestamptz
);

CREATE TABLE "comments" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "post_id" int NOT NULL,
  "user_id" int NOT NULL,
  "comment_text" text NOT NULL,
  "createdAt" timestamptz NOT NULL DEFAULT 'now()',
  "updatedAt" timestamptz
);

ALTER TABLE "routes" ADD FOREIGN KEY ("admin_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "routes" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
