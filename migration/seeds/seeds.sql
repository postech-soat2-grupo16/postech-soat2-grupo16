-- DELETE CURRENT DATA
DELETE FROM clientes;
DELETE FROM pedido_items;
DELETE FROM items;
DELETE FROM pedidos;

TRUNCATE pagamentos, clientes, pedidos, items, pedido_items RESTART IDENTITY;

-- INSERT CLIENTES
INSERT INTO clientes (email, cpf, name)
VALUES ('cliente_teste_1@gmail.com', '12312312312', 'cliente teste 1'),
    ('cliente_teste_2@gmail.com', '22312312312', 'cliente teste 2'),
    ('cliente_teste_3@gmail.com', '22312312312', 'cliente teste 3');

-- INSERT ITEMS
INSERT INTO items (name, category, description, price, created_at, updated_at, deleted_at)
VALUES ('Lanche', 'LANCHE', 'Lanche', 1, 'NOW'::timestamptz, 'NOW'::timestamptz, null),
       ('Bebida', 'BEBIDA', 'Bebida', 1, 'NOW'::timestamptz, 'NOW'::timestamptz, null),
       ('item_3', 'BEBIDA', 'update', 7.20, 'NOW'::timestamptz, 'NOW'::timestamptz, null),
       ('item_4', 'SOBREMESA', 'get', 10.30, 'NOW'::timestamptz, 'NOW'::timestamptz, null),
       ('item_5', 'ACOMPANHAMENTO', 'get', 30.20, 'NOW'::timestamptz, 'NOW'::timestamptz, null),
       ('item_6', 'ACOMPANHAMENTO', 'used for delete', 30.20, 'NOW'::timestamptz, 'NOW'::timestamptz, null),
       ('item_7_deleted', 'ACOMPANHAMENTO', 'deleted', 28.74, 'NOW'::timestamptz, 'NOW'::timestamptz, 'NOW'::timestamptz);

-- INSERT PEDIDOS
INSERT INTO pedidos (status, notes, cliente_id)
VALUES ('CRIADO', '', 1),
    ('CRIADO', '', 2),
    ('AGUARDANDO_PAGAMENTO', '', 2),
    ('RECEBIDO', '', 3);

-- INSERT PEDIDOS ITENS
INSERT INTO pedido_items (pedido_id, item_id, quantity)
VALUES (1, 1, 1),
       (1, 2, 1),
       (3, 2, 1);


-- INSERT PAGAMENTOS
INSERT INTO pagamentos (pedido_id, amount, status)
VALUES (3, 2, 'NEGADO'),
    (3, 2, 'APROVADO');