# Financial Statistics Module

Este módulo em Go é parte integrante de uma plataforma de controle financeiro moderna baseada em microserviços. Ele é responsável por autenticar usuários via JWT, extrair o `user_id` do token e calcular estatísticas financeiras com base nas transações cadastradas no banco de dados.

---

## Índice

- [Objetivo](#objetivo)
- [Principais Funcionalidades](#principais-funcionalidades)
- [Arquitetura e Tecnologias](#arquitetura-e-tecnologias)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Pré-requisitos](#pré-requisitos)
- [Instalação e Configuração](#instalação-e-configuração)
- [Como Executar](#como-executar)
- [Endpoints Disponíveis](#endpoints-disponíveis)
- [Fluxo de Dados e Integração](#fluxo-de-dados-e-integração)
- [Contribuição](#contribuição)
- [Licença](#licença)

---

## Objetivo

Desenvolver um microserviço para:
- **Validação de JWT:** Assegurar que somente usuários autenticados acessem os dados.
- **Cálculo de Estatísticas Financeiras:** Agregar e calcular receitas, despesas e saldo com base em um intervalo de datas especificado.
- **Integração com Banco de Dados:** Utilizar o GORM para conexão com bancos de dados PostgreSQL ou MySQL e realizar migrações de modelos.

---

## Principais Funcionalidades

1. **Autenticação via JWT**  
   - Verifica a presença e validade do token de autenticação.
   - Extrai o `user_id` dos claims do token para identificar o usuário.

2. **Cálculo de Estatísticas**  
   - Consulta as transações financeiras do usuário no banco de dados.
   - Agrupa os valores por tipo (receita/income e despesa/expense).
   - Retorna as estatísticas: Total de receitas, total de despesas e saldo.

3. **Tratamento de Datas**  
   - Permite filtrar transações através de parâmetros `start_date` e `end_date` (formato: `YYYY-MM-DD`).

4. **Configuração do Pool de Conexões**  
   - Otimiza o acesso ao banco de dados através da configuração adequada do pool de conexões.

---

## Arquitetura e Tecnologias

| **Módulo**                        | **Tecnologia**                              | **Descrição**                                                    |
| --------------------------------- | ------------------------------------------- | ---------------------------------------------------------------- |
| **HTTP Server**                   | [Gin](https://github.com/gin-gonic/gin)     | Roteamento e criação de APIs RESTful.                            |
| **Autenticação**                  | [JWT](https://github.com/golang-jwt/jwt)      | Validação e gerenciamento de tokens de autenticação.             |
| **ORM e Banco de Dados**          | [GORM](https://gorm.io/)                     | Conexão e migração com bancos de dados PostgreSQL/MySQL.         |
| **Ambiente e Configuração**       | [godotenv](https://github.com/joho/godotenv)  | Carregamento de variáveis de ambiente a partir de um arquivo `.env`. |

---

## Estrutura do Projeto

```
financial-statistics/
├── controllers/
│   ├── auth_middleware.go        # Middleware para validação de JWT
│   └── statistics_controller.go  # Controlador para calcular estatísticas financeiras
├── database/
│   └── database.go               # Conexão e configuração do banco de dados
├── models/
│   └── transactions.go           # Modelo da transação financeira
├── services/
│   └── statistics.go             # Lógica para cálculo das estatísticas financeiras
├── main.go                       # Configuração e inicialização do servidor
└── .env                          # Arquivo de configuração de variáveis de ambiente (não incluso no repositório)
```

---

## Pré-requisitos

- **Go 1.16+** ou versão superior instalada.
- Banco de dados PostgreSQL ou MySQL.
- Variáveis de ambiente definidas:
  - `DATABASE_DSN`: DSN para conexão com o banco de dados.
  - `JWT_SECRET`: Chave secreta para assinatura dos tokens JWT.
  - `PORT` (opcional): Porta na qual o servidor irá rodar (padrão: 8080).

---

## Instalação e Configuração

1. **Clone o Repositório**

   ```bash
   git clone https://seurepositorio.com/financial-statistics.git
   cd financial-statistics
   ```

2. **Instale as Dependências**

   Utilize o gerenciador de pacotes do Go:

   ```bash
   go mod download
   ```

3. **Configuração do Ambiente**

   Crie um arquivo `.env` na raiz do projeto e defina as variáveis necessárias:

   ```dotenv
   DATABASE_DSN=postgres://usuario:senha@localhost:5432/nomedobanco?sslmode=disable
   JWT_SECRET=sua_chave_secreta
   PORT=8080
   ```

---

## Como Executar

1. **Iniciar o Servidor**

   Execute o comando:

   ```bash
   go run main.go
   ```

2. **Verifique os Logs**

   O servidor deverá iniciar e conectar ao banco de dados. Se `JWT_SECRET` não estiver definido, um valor padrão será utilizado.

---

## Endpoints Disponíveis

### GET `/statistics`

- **Descrição:** Retorna as estatísticas financeiras do usuário autenticado.
- **Parâmetros de Query:**
  - `start_date` (opcional): Data de início (formato `YYYY-MM-DD`). Padrão: `1970-01-01`.
  - `end_date` (opcional): Data de fim (formato `YYYY-MM-DD`). Padrão: data atual.
- **Headers:**
  - `Authorization: Bearer <token>`
- **Resposta de Sucesso:**

  ```json
  {
    "total_receitas": 10000.00,
    "total_despesas": 7500.00,
    "saldo": 2500.00
  }
  ```

- **Exemplo de Requisição com cURL:**

  ```bash
  curl -H "Authorization: Bearer SEU_TOKEN_AQUI" "http://localhost:8080/statistics?start_date=2023-01-01&end_date=2023-12-31"
  ```

---

## Fluxo de Dados e Integração

O módulo GO integra-se à arquitetura geral da plataforma de controle financeiro conforme o seguinte fluxo:

1. **Autenticação:**  
   - O usuário realiza login na interface (ex.: React) e obtém um token JWT.
2. **Envio do Token:**  
   - O cliente armazena o token e o utiliza para realizar requisições autenticadas.
3. **Processamento no Módulo Go:**  
   - O middleware valida o token, extrai o `user_id` e permite acesso ao endpoint `/statistics`.
   - O controlador consulta o banco de dados e utiliza o serviço para calcular as estatísticas.
4. **Retorno ao Cliente:**  
   - As estatísticas financeiras são retornadas para a interface do usuário para visualização e análise.

> **Diagrama do Sistema**  
> ![Diagrama do Sistema](./Diagram.png)

---

## Contribuição

Contribuições são bem-vindas! Se desejar melhorar o projeto, siga as diretrizes abaixo:

1. Faça um fork do repositório.
2. Crie uma branch com a sua feature: `git checkout -b minha-feature`.
3. Realize suas alterações e faça commits com mensagens claras.
4. Envie suas alterações com `git push origin minha-feature`.
5. Abra um Pull Request para revisão.

---

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---

Este README fornece uma visão abrangente do módulo GO para controle financeiro, detalhando sua configuração, execução e integração com a plataforma. Caso tenha dúvidas ou sugestões, sinta-se à vontade para entrar em contato ou abrir uma issue no repositório.