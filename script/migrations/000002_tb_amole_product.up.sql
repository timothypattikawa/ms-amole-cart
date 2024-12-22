CREATE TABLE tb_amole_product(
    tbap_id bigserial primary key,
    tbap_name varchar not null,
    tbap_price int not null,
    tbap_description text not null,
    tbap_created_at date,
    tbap_updated_at date
);