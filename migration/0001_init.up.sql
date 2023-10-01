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
	cliente_id int8 NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT pedidos_pkey PRIMARY KEY (id),
	CONSTRAINT fk_pedidos_clientes FOREIGN KEY (cliente_id) REFERENCES clientes(id)
);
CREATE INDEX idx_pedidos_deleted_at ON pedidos USING btree (deleted_at);

CREATE TABLE IF NOT EXISTS pedido_items (
	pedido_id int8 NULL,
	item_id int8 NULL,
	quantity int8 NULL,
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT pedidos_items_pkey PRIMARY KEY (id),
	CONSTRAINT fk_pedidos_items_items FOREIGN KEY (item_id) REFERENCES items(id),
	CONSTRAINT fk_pedidos_items FOREIGN KEY (pedido_id) REFERENCES pedidos(id)
);
CREATE INDEX idx_pedidos_items_deleted_at ON pedido_items USING btree (deleted_at);

-- Criação da tabela 'pagamentos'
CREATE TABLE IF NOT EXISTS pagamentos (
    id SERIAL PRIMARY KEY,
    pedido_id INT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    status TEXT NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL,
    CONSTRAINT fk_pagamentos_pedidos FOREIGN KEY (pedido_id) REFERENCES pedidos(id)
    );
CREATE INDEX idx_pagamentos_deleted_at ON pagamentos USING btree (deleted_at);