DROP TRIGGER IF EXISTS update_product_rating ON reviews;
DROP FUNCTION IF EXISTS automatic_average_rating();
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS products;