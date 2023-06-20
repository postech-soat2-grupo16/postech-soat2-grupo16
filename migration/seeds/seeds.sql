-- DELETE CURRENT DATA
DELETE FROM clientes;
DELETE FROM pedidos_items;
DELETE FROM items;
DELETE FROM pedidos;

-- INSERT CLIENTES
INSERT INTO clientes (email, cpf, name) 
VALUES ('cliente_teste_1@gmail.com', '12312312312', 'cliente teste 1'),
    ('cliente_teste_2@gmail.com', '22312312312', 'cliente teste 2');

-- INSERT ITEMS
-- INSERT PEDIDOS
-- INSERT PEDIDOS_ITEMS