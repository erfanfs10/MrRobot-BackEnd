BEGIN;

ALTER TABLE products DROP CONSTRAINT products_status_check;

ALTER TABLE products ALTER COLUMN status SET DEFAULT 'available';

ALTER TABLE products ALTER COLUMN status SET NOT NULL;

ALTER TABLE products
ADD CONSTRAINT products_status_check CHECK (
    status IN (
        'available',
        'notAvailable',
        'comingSoon'
    )
);

COMMIT;