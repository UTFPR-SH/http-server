# HTTP Server

O servidor está intrinsecamente conectado ao projeto do Portal Direc que visa construir um website
para a distribuição de informações aos discentes da universidade. Contudo, o objetivo é fazê-lo
independente de projeto algum, e servir páginas construídas pelo construtor de páginas estático
Jekyll.

Para atingir o objetivo de independência e alcançar uma generalidade para outros projetos, uma
lista de objetivos foi criada estabelecendo os requisitos necessários. Entende-se, ainda, que não
vai existir projeto totalmente independente do servidor e, portanto, uma extensão ao servidor deverá
ser feita nesses casos

No futuro uma extensão para conexão com o banco de dados será criada para prover uma maneira de
recuperar, gravar, editar e excluir registros de um determinado BD.

# Roadmap to independence!

- [ ] Especificação do arquivo de configuração (serialização json/yaml)
- [ ] Leitura de arquivo de configuração que indica ao servidor todas as rotas do website

# Como compilar?

Será necessário, para a compilação, o seguintes requisitos:

- Compilador da linguagem de programação [Go](http://golang.org/)
- Programa *make* para executar o script de construção do programa

Após ter os requisitos necessários basta escrever, no terminal,  `make` na raiz
do diretório do servidor (Onde o arquivo `Makefile` se encontra) e aguardar
a compilação. Após a compilação um binário chamado *server* será
construído, em seguinte basta executá-lo para iniciar o servidor.

## Parâmetros do servidor

Os parâmetros disponíveis estão listados na tabela abaixo.

| Parâmetro | Descrição                                                 | Uso                    |
| --------- | -------------------------------------                     | ---------------------  |
| --port    | Modifica a porta de uso pelo servidor                     | `server -p 9954`       |
| --html    | Indica ao servidor onde se encontra os arquivos do Jekyll | `server --html ./html` |
| -d        | Define que o servidor executará em modo de debug          | `server -d`            |
