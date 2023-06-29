# POSTECH CHALLENGE SOAT2 GRUPO 16
![Go v1.20](https://img.shields.io/badge/go-v1.20-blue)

<p align="center">
 <a href="#sobre">Sobre</a> •
 <a href="#entregas-fase-1">Entregas Fase 1</a> •
 <a href="#domain-driven-development-event-storm">Sobre</a> •
 <a href="#interpretação-apis">Interpretação APIs</a> •
 <a href="#como-executar">Como Executar</a> •
 <a href="#como-executar-testes">Como Executar Testes</a> •
 <a href="#como-visualizar-o-swagger">Como Visualizar o Swagger</a> •
 <a href="#como-atualizar-o-swagger">Como Executar Testes</a> •
</p>

## Sobre
Projeto do curso de pós-graduação em Software Architecture da FIAP apresentado pelo Grupo 16 da Turma 2 de 01/2023.

Detalhes/Enúnciado: [Link](https://on.fiap.com.br/mod/conteudoshtml/view.php?id=314758&c=8960&sesskey=CSIu2psPsh)

Equipe/Contribuidores:
- Pedro Vitor Jhum Haramoto
- Jorge Eugenio Souza de Melo
- Joao Vitor Campari Racchetti
- Thiago Oliveira Camargo
- Rodrigo Luiz Pedroza Bezerra

### Entregas Fase 1

1. Documentação do sistema (DDD) utilizando a linguagem ubíqua, dos seguintes fluxos:
a. Realização do pedido e pagamento
b. Preparação e entrega do pedido

2. Uma aplicação para todo sistema de backend (monolito) que deverá ser desenvolvido seguindo os padrões apresentados nas aulas:
a. Utilizando arquitetura hexagonal
b. APIs:
    - Cadastro do Cliente
    - Identificação do Cliente via CPF
    - Criar, editar e remover de produto
    - Buscar produtos por categoria
    - Fake checkout, apenas enviar os produtos escolhidos para a fila
    - Listar os pedidos
c. Aplicação deverá ser escalável para atender grandes volumes nos horários de pico
d. Banco de dados a sua escolha
    - Inicialmente deveremos trabalhar e organizar a fila dos pedidos apenas em banco de dados

3. A aplicação deve ser entregue com um Dockerfile configurado para executá-la corretamente.
Para validação da POC, temos a seguinte limitação de infraestrutura:
a. 1 instância para banco de dados
b. 1 instâncias para executar aplicação

Não será necessário o desenvolvimento de interfaces para o frontend, o foco deve ser total no backend.

#### Domain Driven Development Event Storm

Abaixo o diagrama gerado durante o Event Storm e Dicionário de Linguagem Ubíqua, realizados como parte da entrega do projeto:

![ddd event storm](./docs/srcs/ddd.png)

Também disponível [neste link](https://miro.com/app/board/uXjVMBVJX7I=/), com todas as etapas da dinâmica de grupo realizada, nesse caso, pode ser necessário solicitar permissão de acesso.

#### Interpretação APIs 

Nessa seção, gostaríamos de descrever como interpretamos e realizamos a entrega dos requisitos (APIs) solicitados nesta fase:

- Cadastro do Cliente
    - É possível realizar o cadastro do cliente através do método `POST /clientes`, sendo que este pode ser realizado utilizando NOME (obrigatório) EMAIL (opcional) CPF (opcional);
- Identificação do Cliente via CPF
    - Caso o cliente opte por se identificar por CPF, é possível recupera-lo a partir desse dado `GET /clientes` utilizando o parâmetro opcional via `query`: "CPF"
- Criar, editar e remover de produto
    - CRUD de produtos disponível a partir dos métodos `GET`, `POST`, `PUT`e `DELETE`
- Buscar produtos por categoria
    - A busca de produtos possui o parâmetro opcional "category", trazendo todos os produtos com aquela categoria
- Fake checkout, apenas enviar os produtos escolhidos para a fila
    - O checkout (fake checkout no momento) realizamos através da atualização do status do pedido, que ao passar para o status `RECEBIDO`, pode ser recuperado na listagem a partir desse filtro opcional "status" no endpoint e possivelmente ser exibido em uma interface de cozinha para lista de preparação
- Listar os pedidos
    - A listagem de pedidos está disponível com filtro opcional de "status", conforme mencionado acima

Para mais informações sobre contratos/API, é possível acessar através do swagger, como mencionado na seção [como visualizar o swagger](#como-visualizar-o-swagger).


## Como Executar

Um makefile é disponibilizado para ajudar com algumas atividades rotineiras, para checar a lista de receitas disponíveis, basta acessar o [makefile aqui](./Makefile) ou simplesmente, na raiz do projeto, executar o comando `make help`.

Execute `make build-all` e em seguida `make run-all`, para executar todas as imagens necessárias e subir o projeto localmente. Caso deseje executar somente o banco de dados, é possível através da receita `make run-db` (também após o `make build-all`).

## Como Executar Testes

Esse projeto atualmente possui testes unitários e de integração, para executar todos, é necessário que ao menos o container do banco de dados esteja disponível (para saber como, veja a seção: [Como Executar](#como-executar)).

Após isso, basta executar o comando `make test`, que fará com que a base de dados atual seja recriada (para execução dos testes de integração) e os testes sejam executados.

## Como Executar Linter

Para executar os linters, é necessário ter instalada, localmente, a ferramenta `golangci`. Disponível [neste link](https://golangci-lint.run/usage/install/).

Para executar, basta utilizar o comando `make lint`, que executará o comando com o parâmetro de autofix. Ou então o comando `make ci`, que executará os testes e o linter.

## Como Visualizar o Swagger

Este projeto conta com Swagger para especificação e documentação da API. Para visualizar, basta executar localmente o projeto, como indicado na seção [Como Executar](#como-executar). E então acessar o link abaixo:

`http://localhost:8000/swagger/index.html`

## Como Atualizar o Swagger

Para atualizar o Swagger após a criação de um novo endpoint ou alteração de um endpoint existente, basta executar as anotações conforma a [documentação](https://github.com/swaggo/http-swagger#a-practical-example) do `swag`. E em seguida, executar a receita do makefile conforme o exemplo abaixo:

`make update-docs` 