# Api Cotação de Dólar em Go

Api de consulta de cotação de dólar com uma cliente que consome a API e salva os dados em um arquivo texto.

## Estrutura do Projeto

O projeto está dividido em dois módulos principais:

### Servidor

- API REST que consome dados da API externa de cotações
- Armazena as cotações em um banco SQLite
- Disponibiliza endpoint para consulta de cotações

### Cliente

- Consome a API do servidor
- Salva as cotações em arquivo texto

## Requisitos

- Go 1.24.0 ou superior
- SQLite

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/pcbrsites/golang-cotacao.git
cd golang-cotacao
```

2. Instale as dependências do servidor:

```bash
cd servidor
go mod tidy
```

3. Instale as dependências do cliente:

```bash
cd cliente
go mod tidy
```

## Uso

### Iniciando o Servidor

1. Entre no diretório do servidor:

```bash
cd servidor
```

2. Execute o servidor:

```bash
go run servidor.go
```

O servidor iniciará na porta 8080.

ou docker compose

```bash
docker-compose up --build
```

### Usando o Cliente

1. Em outro terminal, entre no diretório do cliente:

```bash
cd cliente
```

2. Execute o cliente:

```bash
go run cliente.go
```

O cliente irá:

- Consultar a cotação atual do dólar através do servidor
- Exibir a cotação no terminal
- Salvar a cotação em um arquivo texto (cotacao.txt)

## API

### Endpoint

```
GET /cotacao
```

### Resposta de Sucesso

```json
{
  "bid": "5.6435"
}
```

### Resposta de Erro

```json
{
  "status_code": 500,
  "error": "Erro ao obter cotação",
  "message": "mensagem de erro detalhada"
}
```

## Timeouts

- Servidor: 200ms para chamada externa
- Servidor: 10ms para operações no banco
- Cliente: 300ms para chamada ao servidor

## Banco de Dados

O sistema utiliza SQLite para armazenar as cotações. O arquivo do banco é criado automaticamente em:

```
servidor/cotacao.db
```

ou docker compose volume:

```
./cotacao.db
```
