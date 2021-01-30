package productlibrary

const (
	SEL_ALL_PRODUCTS = `
	SELECT *
	FROM 
	(
	SELECT 1 AS kind, id, name, created_at,deleted_at FROM single_products sp WHERE deleted_at IS NULL
	UNION
	SELECT 2 AS kind, id, name, created_at,deleted_at FROM variant_products vp WHERE deleted_at IS NULL
	) tablea
	ORDER BY created_at`

	SEL_SP_PRICES_BY_ID = `
	SELECT 
	*
	FROM (
	SELECT
		1 AS kind,
		sp.name,
		sp.created_at,
		spi.id AS item_id,
		spi.sku
	FROM
		single_products as sp
	INNER JOIN SP_items as spi on
		sp.id = spi.SP_id

	WHERE
		sp.id IN (?) AND sp.deleted_at IS NULL 
	UNION 
	SELECT
		2 AS kind,
		vp.name,
		vp.created_at,
		vpi.id AS item_id,
		vpi.sku
	FROM
		variant_products as vp
	INNER JOIN VP_items as vpi on
		vp.id = vpi.vp_id
	WHERE
		vp.id IN (?) AND vp.deleted_at IS NULL 
		
	) tablea ORDER BY created_at;
	`
)
