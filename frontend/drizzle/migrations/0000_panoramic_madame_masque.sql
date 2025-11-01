CREATE TABLE "urls" (
	"id" serial PRIMARY KEY NOT NULL,
	"short_code" varchar(8) NOT NULL,
	"original_url" text NOT NULL,
	"clicks" integer DEFAULT 0,
	"created_at" timestamp DEFAULT now(),
	"user_id" integer,
	CONSTRAINT "urls_short_code_unique" UNIQUE("short_code")
);
