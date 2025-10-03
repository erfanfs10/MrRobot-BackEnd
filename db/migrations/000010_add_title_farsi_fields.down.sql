BEGIN;

ALTER TABLE brands DROP COLUMN IF EXISTS title_farsi;

ALTER TABLE categories DROP COLUMN IF EXISTS title_farsi;

ALTER TABLE product_types DROP COLUMN IF EXISTS title_farsi;

COMMIT;