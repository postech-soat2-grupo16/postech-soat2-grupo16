CREATE TABLE IF NOT EXISTS clientes (
	id bigserial NOT NULL,
	email text NULL,
	cpf text NULL,
	"name" text NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT clientes_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_clientes_deleted_at ON clientes USING btree (deleted_at);

CREATE TABLE IF NOT EXISTS items (
	id bigserial NOT NULL,
	"name" varchar(255) NOT NULL,
	category varchar(100) NOT NULL,
	description varchar(255) NOT NULL,
	price numeric NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT items_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_items_deleted_at ON items USING btree (deleted_at);

CREATE TABLE IF NOT EXISTS pedidos (
	id bigserial NOT NULL,
	status text NOT NULL,
	notes text NULL,
	clientes_id int8 NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT pedidos_pkey PRIMARY KEY (id),
	CONSTRAINT fk_pedidos_clientes FOREIGN KEY (clientes_id) REFERENCES clientes(id)
);
CREATE INDEX idx_pedidos_deleted_at ON pedidos USING btree (deleted_at);

CREATE TABLE IF NOT EXISTS pedidos_items (
	pedidos_id int8 NULL,
	items_id int8 NULL,
	quantity int8 NULL,
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT pedidos_items_pkey PRIMARY KEY (id),
	CONSTRAINT fk_pedidos_items_items FOREIGN KEY (items_id) REFERENCES items(id),
	CONSTRAINT fk_pedidos_items FOREIGN KEY (pedidos_id) REFERENCES pedidos(id)
);
CREATE INDEX idx_pedidos_items_deleted_at ON pedidos_items USING btree (deleted_at);