-- name: GetAllProduct :many
SELECT tbap_id, tbap_name, tbap_price, tbap_description, tbap_created_at, tbap_updated_at
FROM tb_amole_product;

-- name: GetProductById :one
SELECT tbap_id, tbap_name, tbap_price, tbap_description, tbap_created_at, tbap_updated_at
FROM tb_amole_product WHERE tbap_id = $1;