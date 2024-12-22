-- name: InsertCartItemsByCartId :one
INSERT INTO public.tb_amole_cart_items
(taci_cart_id, taci_product_id, taci_qty, taci_price)
VALUES($1, $2, $3, $4) RETURNING *;

-- name: GetCarItemsByCartIdAmdProductid :one
SELECT taci_id, taci_cart_id, taci_product_id, taci_qty, taci_price
FROM public.tb_amole_cart_items WHERE taci_cart_id=$1 AND taci_product_id = $2;

-- name: UpdateCartItemByCartId :exec 
UPDATE public.tb_amole_cart_items
SET taci_product_id=$2, taci_qty=$3, taci_price=$4
WHERE taci_id=$1;
