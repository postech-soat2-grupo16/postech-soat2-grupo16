-- DELETE CURRENT DATA
DELETE FROM clientes;
DELETE FROM pedido_items;
DELETE FROM items;
DELETE FROM pedidos;

TRUNCATE clientes, pedidos, items, pedido_items RESTART IDENTITY;

-- INSERT CLIENTES
INSERT INTO clientes (email, cpf, name)
VALUES ('cliente_teste_1@gmail.com', '12312312312', 'cliente teste 1'),
    ('cliente_teste_2@gmail.com', '22312312312', 'cliente teste 2');

-- INSERT ITEMS
INSERT INTO items ("name", category, description, price)
VALUES ('Lanche', 'LANCHE', 'Lanche', 1),
       ('Bebida', 'BEBIDA', 'Bebida', 1);


-- INSERT PEDIDOS
INSERT INTO pedidos (status, notes, cliente_id)
VALUES ('CRIADO', '', 1),
    ('CRIADO', '', 2),
    ('AGUARDANDO_PAGAMENTO', '', 2);

-- INSERT PEDIDOS ITENS
INSERT INTO pedido_items (pedido_id, item_id, quantity)
VALUES (1, 1, 1),
       (1, 2, 1);