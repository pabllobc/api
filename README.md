# Projeto API

## Visão Geral

Bem-vindo ao Projeto API! Este repositório hospeda uma API simples, porém poderosa, projetada para lidar com várias tarefas de forma eficiente usando a linguagem Go. Este documento irá guiá-lo através do processo de configuração, uso e contribuição.

## Tabela de Conteúdos

- [Recursos](#recursos)
- [Primeiros Passos](#primeiros-passos)
- [Instalação](#instalação)
- [Uso](#uso)
- [Endpoints da API](#endpoints-da-api)
- [Configuração](#configuração)
- [Testes](#testes)
- [Contribuindo](#contribuindo)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Licença](#licença)

## Recursos

- Arquitetura RESTful
- Autenticação e autorização de usuários
- Operações CRUD para várias entidades
- Tratamento de erros e validação
- Otimização de escalabilidade e desempenho
- Banco de dados SQLite

## Primeiros Passos

### Pré-requisitos

Certifique-se de ter os seguintes itens instalados em sua máquina:

- [Go](https://golang.org/) (v1.22.3 ou superior)

### Instalação

1. Clone o repositório:

    ```bash
    git clone https://github.com/pabllobc/api.git
    cd api
    ```

2. Instale as dependências:

    ```bash
    go mod tidy
    ```

3. Configure as variáveis de ambiente:

    Crie um arquivo `.env` no diretório raiz e adicione as variáveis de ambiente necessárias conforme mostrado no arquivo `.env.example`.

4. Inicie o servidor de desenvolvimento:

    ```bash
    go run main.go
    ```

## Uso

### Executando o Servidor

Para executar o servidor em modo de desenvolvimento, use:

```bash
# go run main.go
# Para executar todo o projeto utilize:
go run .
```

O servidor iniciará em `http://localhost:3000`.

### Endpoints da API

Abaixo estão alguns dos principais endpoints disponíveis:

- **GET** `/api/v1/resource` - Buscar todos os recursos
- **GET** `/api/v1/resource/:id` - Buscar um recurso específico por ID
- **POST** `/api/v1/resource` - Criar um novo recurso
- **PUT** `/api/v1/resource/:id` - Atualizar um recurso específico por ID
- **DELETE** `/api/v1/resource/:id` - Excluir um recurso específico por ID

Para documentação detalhada, consulte a [Documentação da API](docs/API.md).

## Configuração

A aplicação usa variáveis de ambiente para configuração. As seguintes variáveis são necessárias:

- `PORT` - O número da porta na qual o servidor será executado
- `MONGODB_URI` - A string de conexão para o banco de dados MongoDB
- `JWT_SECRET` - A chave secreta para geração de JSON Web Token (JWT)

## Configuração do BD SQLite
- Abra o terminal
- go get modernc.org/sqlite
- Crie um arquivo de inicialização do BD - db.go

## Testes

Para executar os testes, use o seguinte comando:

```bash
go test ./...
```

Isso executará todos os testes localizados no diretório `tests`.

## Contribuindo

Contribuições da comunidade são bem-vindas! Para contribuir:

1. Faça um fork do repositório
2. Crie uma nova branch (`git checkout -b feature-branch`)
3. Faça suas alterações
4. Comite suas alterações (`git commit -m 'Add new feature'`)
5. Faça o push para a branch (`git push origin feature-branch`)
6. Abra um pull request

Por favor, certifique-se de que seu código segue nossos padrões de codificação e inclui os testes apropriados.

## Tecnologias Utilizadas

- Go 1.22.3
- [github.com/bytedance/sonic v1.12.0](https://pkg.go.dev/github.com/bytedance/sonic) 
- [github.com/bytedance/sonic/loader v0.2.0](https://pkg.go.dev/github.com/bytedance/sonic/loader) 
- [github.com/cloudwego/base64x v0.1.4](https://pkg.go.dev/github.com/cloudwego/base64x) 
- [github.com/cloudwego/iasm v0.2.0](https://pkg.go.dev/github.com/cloudwego/iasm) 
- [github.com/dustin/go-humanize v1.0.1](https://pkg.go.dev/github.com/dustin/go-humanize) 
- [github.com/gabriel-vasile/mimetype v1.4.5](https://pkg.go.dev/github.com/gabriel-vasile/mimetype) 
- [github.com/gin-contrib/sse v0.1.0](https://pkg.go.dev/github.com/gin-contrib/sse) 
- [github.com/gin-gonic/gin v1.10.0](https://pkg.go.dev/github.com/gin-gonic/gin) 
- [github.com/go-playground/locales v0.14.1](https://pkg.go.dev/github.com/go-playground/locales) 
- [github.com/go-playground/universal-translator v0.18.1](https://pkg.go.dev/github.com/go-playground/universal-translator) 
- [github.com/go-playground/validator/v10 v10.22.0](https://pkg.go.dev/github.com/go-playground/validator/v10) 
- [github.com/goccy/go-json v0.10.3](https://pkg.go.dev/github.com/goccy/go-json) 
- [github.com/google/uuid v1.6.0](https://pkg.go.dev/github.com/google/uuid) 
- [github.com/hashicorp/golang-lru/v2 v2.0.7](https://pkg.go.dev/github.com/hashicorp/golang-lru/v2) 
- [github.com/json-iterator/go v1.1.12](https://pkg.go.dev/github.com/json-iterator/go) 
- [github.com/klauspost/cpuid/v2 v2.2.8](https://pkg.go.dev/github.com/klauspost/cpuid/v2) 
- [github.com/leodido/go-urn v1.4.0](https://pkg.go.dev/github.com/leodido/go-urn) 
- [github.com/mattn/go-isatty v0.0.20](https://pkg.go.dev/github.com/mattn/go-isatty) 
- [github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd](https://pkg.go.dev/github.com/modern-go/concurrent) 
- [github.com/modern-go/reflect2 v1.0.2](https://pkg.go.dev/github.com/modern-go/reflect2) 
- [github.com/ncruces/go-strftime v0.1.9](https://pkg.go.dev/github.com/ncruces/go-strftime) 
- [github.com/pelletier/go-toml/v2 v2.2.2](https://pkg.go.dev/github.com/pelletier/go-toml/v2) 
- [github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec](https://pkg.go.dev/github.com/remyoudompheng/bigfft) 
- [github.com/twitchyliquid64/golang-asm v0.15.1](https://pkg.go.dev/github.com/twitchyliquid64/golang-asm) 
- [github.com/ugorji/go/codec v1.2.12](https://pkg.go.dev/github.com/ugorji/go/codec) 
- [golang.org/x/arch v0.8.0](https://pkg.go.dev/golang.org/x/arch) 
- [golang.org/x/crypto v0.25.0](https://pkg.go.dev/golang.org/x/crypto) 
- [golang.org/x/net v0.27.0](https://pkg.go.dev/golang.org/x/net) 
- [golang.org/x/sys v0.22.0](https://pkg.go.dev/golang.org/x/sys) 
- [golang.org/x/text v0.16.0](https://pkg.go.dev/golang.org/x/text) 
- [google.golang.org/protobuf v1.34.2](https://pkg.go.dev/google.golang.org/protobuf) 
- [gopkg.in/yaml.v3 v3.0.1](https://pkg.go.dev/gopkg.in/yaml.v3) 
- [modernc.org/gc/v3 v3.0.0-20240107210532-573471604cb6](https://pkg.go.dev/modernc.org/gc/v3) 
- [modernc.org/libc v1.55.3](https://pkg.go.dev/modernc.org/libc) 
- [modernc.org/mathutil v1.6.0](https://pkg.go.dev/modernc.org/mathutil) 
- [modernc.org/memory v1.8.0](https://pkg.go.dev/modernc.org/memory) 
- [modernc.org/sqlite v1.31.1](https://pkg.go.dev/modernc.org/sqlite) 
- [modernc.org/strutil v1.2.0](https://pkg.go.dev/modernc.org/str