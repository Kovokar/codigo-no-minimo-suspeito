
## Requisitos

- [Go](https://golang.org/dl/) 1.18 ou superior
- [Git](https://git-scm.com/)

## Instalação

### 1. Clone o repositório

```bash
git clone https://github.com/Kovokar/codigo-no-minimo-suspeito.git
cd codigo-no-minimo-suspeito.git
```

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto:

```bash
cp .env.example .env
```

Ou crie manualmente o arquivo `.env` com as seguintes variáveis:

```env
# Configurações do servidor
PORT=8080
HOST=localhost

# Configurações do banco de dados
DB_HOST=localhost
DB_PORT=5432
DB_NAME=meu_banco
DB_USER=usuario
DB_PASSWORD=senha


CREATE TABLE acessos (
	id serial PRIMARY KEY,
	nome varchar(100) NOT NULL,
	momento timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)
```

### 3. Instale as dependências (IMPORTANTE)

```bash
go mod tidy
```

### 4. Execute o projeto

```bash
go run cmd/app/main.go
```

## Estrutura do Projeto

```
.
├── cmd/
│   └── app/
│       ├── deleta_banco_linux      # Script para deletar banco (Linux)
│       ├── deleta_banco_mac        # Script para deletar banco (macOS)
│       ├── deleta_banco_windows.exe # Script para deletar banco (Windows)
│       └── main.go                 # Arquivo principal da aplicação
├── internal/
│   ├── db/
│   │   └── conn.go                 # Conexão com banco de dados
│   ├── models/
│   │   └── acessos.go              # Modelo de dados para acessos
│   └── repository/
│       ├── delete.go               # Operações de delete no banco
│       └── get.go                  # Operações de consulta no banco
├── go.mod                          # Dependências do módulo Go
├── go.sum                          # Checksums das dependências
├── .env                            # Variáveis de ambiente (não versionado)
├── .env.example                    # Exemplo de configuração
└── README.md                       # Este arquivo
```

## Comandos Úteis

```bash
# Executar testes
go test ./...

# Compilar o projeto
go build -o app cmd/app/main.go

# Executar o binário compilado
./app

# Verificar dependências não utilizadas
go mod tidy

# Executar scripts de limpeza do banco (conforme seu sistema)
./cmd/app/deleta_banco_linux    # Linux
./cmd/app/deleta_banco_mac      # macOS
./cmd/app/deleta_banco_windows.exe # Windows
```

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).