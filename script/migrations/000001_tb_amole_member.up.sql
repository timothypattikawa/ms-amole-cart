CREATE TABLE tb_amole_member(
    id bigserial PRIMARY KEY,
    name varchar(254) not null,
    email varchar(254) not null,
    password varchar not null,
    address varchar not null
);