package productquery

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

	SEL_ALL_PRODUCTS_ALL_FIELDS = `
	SELECT
			*
	FROM
		(
		SELECT
			1 AS kind,
			sp.*
		FROM
			single_products sp
		WHERE
			deleted_at IS NULL
	UNION
		SELECT
			2 AS kind,
			vp.*
		FROM
			variant_products vp
		WHERE
			deleted_at IS NULL ) tablea
	ORDER BY
		created_at LIMIT ?,?`

	SEL_PRODUCT_ITEM_BY_PRODUCT_ID = `
	SELECT 
	*
	FROM (
	SELECT
		1 AS kind,
		sp.created_at,
		spi.id AS id,
		spi.SP_id as product_id,
		spi.sku,
		spi.weight 
	FROM
		single_products as sp
	INNER JOIN SP_items as spi on
		sp.id = spi.SP_id
	INNER JOIN SP_item_prices sip on 
		spi.id = sip.SP_item_id
	WHERE
		sp.id IN (?) AND sp.deleted_at IS NULL 
	UNION 
	SELECT
		2 AS kind,
		vp.created_at,
		vpi.id AS id,
		vpi.VP_id as product_id,
		vpi.sku,
		vpi.weight 
	FROM
		variant_products as vp
	INNER JOIN VP_items as vpi on
		vp.id = vpi.vp_id
		INNER JOIN VP_item_prices vip  on 
		vpi.id = vip.VP_item_id
	WHERE
		vp.id IN (?) AND vp.deleted_at IS NULL 
		
	) tablea ORDER BY created_at`
	SEL_PRODUCT_PRICES_BY_ID = `
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

	SEL_PRODUCT_STOCKS_BY_ID = `
	SELECT
	*
	FROM
	(
	SELECT
		1 AS kind,
		sp.name,
		sp.created_at,
		spi.id AS item_id,
		spi.sku,
		sii.id AS item_inventory_id,
		sii.stock
	FROM
		single_products as sp
	INNER JOIN SP_items as spi on
		sp.id = spi.SP_id
	INNER JOIN SP_item_inventories sii ON
		spi.id = sii.SP_item_id
	WHERE
		sp.id IN (?)
		AND sp.deleted_at IS NULL
	UNION
	SELECT
		2 AS kind,
		vp.name,
		vp.created_at,
		vpi.id AS item_id,
		vpi.sku,
		vii.id AS item_inventory_id,
		vii.stock
	FROM
		variant_products as vp
	INNER JOIN VP_items as vpi on
		vp.id = vpi.vp_id
	INNER JOIN VP_item_inventories vii ON
		vpi.id = vii.VP_item_id
	WHERE
		vp.id IN (?)
		AND vp.deleted_at IS NULL ) tablea
	ORDER BY
	created_at;
	`

	SEL_SP_BY_ID = `
	SELECT 1 as product_kind_id, single_products.* 
	FROM
		single_products WHERE
	single_products.id IN (?)
	`

	SEL_VP_BY_ID = `Select 2 as product_kind_id, variant_products.* from variant_products where variant_products.id IN (?)`

	SEL_SP_BY_ITEM_INVENTORY_ID = `SELECT sii.* FROM SP_item_inventories sii 
	where sii.id in (?)`

	SEL_VP_BY_ITEM_INVENTORY_ID = `SELECT vii.* FROM VP_item_inventories vii 
	where vii.id in (?)`

	SEL_SP_BY_ITEM_ID = `SELECT si.* FROM SP_items si 
	where si.id in (?)`

	SEL_VP_BY_ITEM_ID = `SELECT vi.* FROM VP_items vi 
	where vi.id in (?)`
)
