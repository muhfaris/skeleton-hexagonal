CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

BEGIN;

-- CREATE TABLE "accounts" -------------------------------------
CREATE TABLE "public"."accounts" ( 
	"id" UUid DEFAULT uuid_generate_v4() NOT NULL,
	"full_name" Character Varying( 2044 ) NOT NULL,
	"username" Character Varying( 2044 ) NOT NULL,
	"email" Character Varying( 2044 ) NOT NULL,
	"password" Character Varying( 2044 ) NOT NULL,
	"role" Character Varying( 2044 ) NOT NULL,
	"created_at" Timestamp With Time Zone DEFAULT now() NOT NULL,
	"updated_at" Timestamp With Time Zone );
 ;
-- -------------------------------------------------------------

-- CREATE INDEX "index_id" -------------------------------------
CREATE INDEX "index_id" ON "public"."accounts" USING btree( "id" );
-- -------------------------------------------------------------

COMMIT;
