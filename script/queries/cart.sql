-- name: CreateCart :one
INSERT INTO public.tb_amole_cart(tac_member_id, tac_total_price, tac_status)
VALUES($1, $2, $3) RETURNING *;

-- name: UpdateCart :exec
UPDATE public.tb_amole_cart
SET tac_member_id=$2, tac_total_price=$3,  tac_status=$4
WHERE tac_id = $1;

-- name: GetCartAndCartItemsByMemberIdAndActiveStatus :one
SELECT tac_id, tac_member_id, tac_total_price tac_status, taci.*
FROM public.tb_amole_cart as tac 
JOIN public.tb_amole_cart_items as taci ON taci.taci_cart_id=tac.tac_id
WHERE tac_member_id = $1 AND tac_status = 'ACTIVATE';

-- name: GetCartByMemberId :one
SELECT tac_id, tac_member_id, tac_total_price, tac_status
FROM public.tb_amole_cart
WHERE tac_member_id = $1 AND tac_status = 'ACTIVATE' LIMIT 1;

-- name: GetCountCartActiveProduct :one
SELECT COUNT(taci.taci_id)
FROM public.tb_amole_cart as tac 
JOIN public.tb_amole_cart_items as taci ON taci.taci_cart_id=tac.tac_id
WHERE tac.tac_member_id = $1 AND tac.tac_status = 'ACTIVATE';


