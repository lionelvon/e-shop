CREATE TABLE product
(
    id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    product_name VARCHAR(50),
     image_url VARCHAR(200)
    description TEXT,
    price INT(11) COMMENT '0.01',
);
CREATE UNIQUE INDEX products_id_uindex ON product (id);