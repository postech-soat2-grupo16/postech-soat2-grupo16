# POSTECH CHALLENGE SOAT2 GRUPO 16
![Go v1.20](https://img.shields.io/badge/go-v1.20-blue)

<p align="center">
 <a href="#sobre">Sobre</a> •
 <a href="#como-executar">Como Executar</a> •
 <a href="#como-executar-testes">Como Executar Testes</a>
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

## Como Executar

Um makefile é disponibilizado para ajudar com algumas atividades rotineiras, para checar a lista de receitas disponíveis, basta acessar o [makefile aqui](./Makefile) ou simplesmente, na raiz do projeto, executar o comando `make help`.

Execute `make build-all` e em seguida `make run-all`, para executar todas as imagens necessárias e subir o projeto localmente. Caso deseje executar somente o banco de dados, é possível através da receita `make run-db` (também após o `make build-all`).

## Como Executar Testes

Esse projeto atualmente possui testes unitários e de integração, para executar todos, é necessário que ao menos o container do banco de dados esteja disponível (para saber como, veja a seção: [Como Executar](#como-executar)).

Após isso, basta executar o comando `make test`, que fará com que a base de dados atual seja recriada (para execução dos testes de integração) e os testes sejam executados.

## Como Executar Linter

Para executar os linters, é necessário ter instalada, localmente, a ferramenta `golangci`. Disponível [neste link](https://golangci-lint.run/usage/install/).

Para executar, basta utilizar o comando `make lint`, que executará o comando com o parâmetro de autofix. Ou então o comando `make ci`, que executará os testes e o linter.

