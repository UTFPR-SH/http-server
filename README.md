# Portal DIREC

Sistema de integração de setores da _DIREC_ e oportunidades da região.

# Como compilar?

Será necessário, para a compilação, o seguintes requisitos:

- Compilador da linguagem de programação [Go](http://golang.org/)
- Motor de páginas estático [Jekyll](https://jekyllrb.com/)
- Programa *make* para executar o script de construção do programa

Após ter os requisitos necessários basta escrever, no terminal,  `make` na raiz
do diretório do servidor (Onde o arquivo `Makefile` se encontra) e aguardar
a compilação. Após a compilação um binário chamado *direc-server* será
construído, em seguinte basta executá-lo para iniciar o servidor.

## Parâmetros do servidor

Os parâmetros disponíveis estão listados na tabela abaixo.


| Parâmetro | Descrição                                        | Uso                        |
| --------- | -------------------------------------            | -------------------------- |
| --port    | Modifica a porta de uso pelo servidor            | `direc-server -p 9954`     |
| -d        | Define que o servidor executará em modo de debug | `direc-server -d`          |
