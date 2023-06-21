-- DELETE CURRENT DATA
DELETE FROM clientes;
DELETE FROM pedido_items;
DELETE FROM items;
DELETE FROM pedidos;

-- INSERT CLIENTES
INSERT INTO clientes (email, cpf, name) 
VALUES ('cliente_teste_1@gmail.com', '12312312312', 'cliente teste 1'),
    ('cliente_teste_2@gmail.com', '22312312312', 'cliente teste 2');

-- INSERT ITEMS

-- INSERT PEDIDOS
INSERT INTO pedidos (status, notes, clientes_id)
VALUES ('CRIADO', '', 1),
    ('CRIADO', '', 2),
    ('AGUARDANDO_PAGAMENTO', '', 2);

/* -- INSERT PEDIDOS_ITEMS
INSERT INTO pedido_items (pedidos_id, items_id, quantity)
VALUES */