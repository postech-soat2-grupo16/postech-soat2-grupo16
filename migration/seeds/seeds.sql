-- DELETE CURRENT DATA
DELETE FROM pedido_items;
DELETE FROM pagamentos;
DELETE FROM pedidos;
DELETE FROM clientes;
DELETE FROM items;

TRUNCATE pagamentos, clientes, pedidos, items, pedido_items RESTART IDENTITY;

-- INSERT CLIENTES
INSERT INTO clientes (email, cpf, name)
VALUES ('cliente_teste_1@gmail.com', '12312312312', 'cliente teste 1'),
       ('cliente_teste_2@gmail.com', '22312312312', 'cliente teste 2'),
       ('cliente_teste_3@gmail.com', '22312312312', 'cliente teste 3'),
       ('cliente_teste_4@gmail.com', '27841622609', 'cliente teste 4 order by');

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
INSERT INTO pedidos (status, notes, cliente_id, created_at, updated_at)
VALUES ('CRIADO', '', 1, timestamp '2023-07-20 12:00', null),
       ('CRIADO', '', 2, timestamp '2023-07-21 13:00', null),
       ('AGUARDANDO_PAGAMENTO', '', 2, timestamp '2023-07-22 14:50', timestamp '2023-07-22 14:54'),
       ('RECEBIDO', '', 3, timestamp '2023-07-21 11:54:13', timestamp '2023-07-21 11:55:56'),
       ('PRONTO', '', 4, timestamp '2023-07-20 02:00', timestamp '2023-07-20 02:31'),
       ('PRONTO', '', 4, timestamp '2023-07-20 01:00', timestamp '2023-07-20 01:29'),
       ('EM_PREPARACAO', '', 4, timestamp '2023-07-23 14:26', timestamp '2023-07-23 14:36'),
       ('FINALIZADO', '', 4, timestamp '2023-07-22 12:26', timestamp '2023-07-22 12:56');

-- INSERT PEDIDOS ITENS
INSERT INTO pedido_items (pedido_id, item_id, quantity)
VALUES (1, 1, 1),
       (1, 2, 1),
       (3, 2, 1),
       (5, 1, 2),
       (5, 1, 2),
       (6, 2, 1);


-- INSERT PAGAMENTOS
INSERT INTO pagamentos (pedido_id, amount, status)
VALUES (3, 2, 'NEGADO'),
       (3, 2, 'APROVADO'),
       (4, 1, 'APROVADO');