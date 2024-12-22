CREATE TABLE
    tb_amole_cart (
        tac_id bigserial primary key,
        tac_member_id int not null,
        tac_total_price int not null,
        tac_status varchar not null
    )