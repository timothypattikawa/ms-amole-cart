CREATE TABLE
    tb_amole_cart_items (
        taci_id bigserial primary key,
        taci_cart_id int not null,
        taci_product_id int not null,
        taci_qty int not null,
        taci_price int not null
    )