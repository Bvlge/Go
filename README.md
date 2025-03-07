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

Desenvolver uma plataforma para controle financeiro que permita aos usuários:
- **Monitorar receitas e despesas:** Cadastro detalhado com classificação por categorias (ex.: Salário, Alimentação, Moradia, Transporte, Lazer, Investimentos).
- **Visualizar fluxo de caixa:** Painel com resumo financeiro, comparativos e evolução do saldo ao longo do tempo.
- **Gerar relatórios analíticos:** Relatórios e dashboards com gráficos interativos e indicadores de performance.
- **Integrar com outras tecnologias:** Uso de Open Banking, API de conversão de moedas, modo offline com sincronização posterior.

No módulo em Go, o foco é na autenticação via JWT e no cálculo das estatísticas financeiras a partir das transações cadastradas.

---

## Principais Funcionalidades

1. **Autenticação via JWT**  
   - Validação do token presente no header `Authorization`.
   - Extração do `user_id` dos claims do token para identificar o usuário.
   - Respostas de erro apropriadas em casos de token ausente ou inválido.

2. **Cálculo de Estatísticas Financeiras**  
   - Consulta e agregação dos dados financeiros com base em um intervalo de datas.
   - Cálculo do total de receitas (Income) e despesas (Loss).
   - Determinação do saldo, categoria de despesa mais frequente, número total de transações e média das transações.

3. **Cadastro e Consulta de Transações**  
   - Registro de transações com detalhamento de data, status, recorrência, tipo de pagamento e anexos.
   - Agrupamento dos dados por categoria e mês para análise dos gastos.

4. **Configuração Otimizada do Banco de Dados**  
   - Uso do GORM para gerenciar a conexão e migrações em bancos PostgreSQL ou MySQL.
   - Configuração do pool de conexões para melhorar a performance.

---

## Arquitetura e Tecnologias

| **Módulo**                        | **Tecnologia**                              | **Descrição**                                                    |
| --------------------------------- | ------------------------------------------- | ---------------------------------------------------------------- |
| **HTTP Server**                   | [Gin](https://github.com/gin-gonic/gin)     | Roteamento e criação de APIs RESTful.                            |
| **Autenticação**                  | [JWT](https://github.com/golang-jwt/jwt)      | Validação e gerenciamento dos tokens de autenticação.            |
| **ORM e Banco de Dados**          | [GORM](https://gorm.io/)                     | Conexão e migração com bancos de dados PostgreSQL/MySQL.         |
| **Gerenciamento de Variáveis**    | [godotenv](https://github.com/joho/godotenv)  | Carregamento de variáveis de ambiente a partir do arquivo `.env`.  |
| **Processamento Estatístico**     | Go (serviços customizados)                  | Cálculo e agregação de dados financeiros para dashboards e relatórios. |

---

## Estrutura do Projeto

```plaintext
financial-statistics/
├── controllers/
│   ├── auth_middleware.go        # Middleware para validação de JWT
│   ├── monthlyexpenses.go        # Controlador para consulta de despesas mensais por categoria
│   └── statistics_controller.go  # Controlador para cálculo das estatísticas financeiras
├── database/
│   └── database.go               # Conexão e configuração do banco de dados
├── models/
│   └── transactions.go           # Modelo da transação financeira
├── services/
│   ├── statistics.go             # Lógica para cálculo das estatísticas financeiras
│   └── monthly_expenses.go       # Lógica para cálculo de despesas mensais por categoria
├── main.go                       # Inicialização e configuração do servidor
└── .env                          # Arquivo de configuração de variáveis de ambiente (não incluso no repositório)
```

---

## Pré-requisitos

- **Go 1.16+** ou versão superior instalada.
- Banco de dados PostgreSQL ou MySQL.
- Variáveis de ambiente definidas:
  - `DATABASE_DSN`: DSN para conexão com o banco de dados.
  - `JWT_SECRET`: Chave secreta para assinatura dos tokens JWT.
  - `PORT` (opcional): Porta para execução do servidor (padrão: 8080).

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

   O servidor deverá iniciar, conectar ao banco de dados e, se necessário, utilizar um valor padrão para `JWT_SECRET` se não estiver definido.

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
    "saldo": 2500.00,
    "categoria_mais_frequente": "Alimentação",
    "total_transacoes": 50,
    "media_transacao": 500.00
  }
  ```

- **Exemplo de Requisição com cURL:**

  ```bash
  curl -H "Authorization: Bearer SEU_TOKEN_AQUI" "http://localhost:8080/statistics?start_date=2023-01-01&end_date=2023-12-31"
  ```

### GET `/statistics/category-expenses`

- **Descrição:** Retorna a média e total das despesas agrupadas por categoria e mês para o usuário autenticado.
- **Parâmetros de Query:**
  - `start_date` (opcional): Data de início (formato `YYYY-MM-DD`). Padrão: `2023-01-01`.
  - `end_date` (opcional): Data de fim (formato `YYYY-MM-DD`). Padrão: data atual.
- **Headers:**
  - `Authorization: Bearer <token>`
- **Resposta de Sucesso:**

  ```json
  [
    {
      "category": "Alimentação",
      "year_month": "2023-03",
      "avg_expense": 250.00,
      "total_expense": 1250.00,
      "count": 5
    }
  ]
  ```

---

## Fluxo de Dados e Integração

1. **Autenticação:**  
   - O usuário realiza login na interface (ex.: React) e recebe um token JWT.
2. **Envio do Token:**  
   - O cliente utiliza o token para realizar requisições aos endpoints protegidos.
3. **Processamento no Módulo Go:**  
   - O middleware valida o token, extrai o `user_id` e permite o acesso aos endpoints.
   - Os controladores (ex.: `/statistics` e `/statistics/category-expenses`) processam as requisições, consultam o banco de dados e acionam os serviços para realizar os cálculos necessários.
4. **Retorno ao Cliente:**  
   - Os resultados das estatísticas financeiras são enviados de volta para a interface do usuário para visualização e análise.

---

## Contribuição

Contribuições são bem-vindas! Para contribuir:

1. Faça um fork do repositório.
2. Crie uma branch com sua feature:  
   ```bash
   git checkout -b minha-feature
   ```
3. Realize as alterações e faça commits com mensagens claras.
4. Envie suas alterações:  
   ```bash
   git push origin minha-feature
   ```
5. Abra um Pull Request para revisão.

---

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---
