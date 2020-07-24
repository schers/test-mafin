BEGIN;

-- CREATE TABLE "test_files" ---------------------------------------
CREATE TABLE file (
    "id" SERIAL NOT NULL,
    "name" Character Varying( 128 ) NOT NULL,
    "size" INTEGER NOT NULL,
    "path" Character Varying( 255 ) NOT NULL,
    PRIMARY KEY ( "id" ) );
;

COMMIT;