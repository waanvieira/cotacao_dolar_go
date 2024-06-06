# Sistema de cotação de dolar
<p>
<a href="https://github.com/waanvieira/cotacao_dolar_go?tab=MIT-1-ov-file#readme"><img src="https://img.shields.io/packagist/l/laravel/framework" alt="License"></a>
</p>

# Sobre o projeto
Projeto de comunicação entre 2 serviços em GO para realizar consulta do dolar, salvar no banco de dados e salvar a cotação em um arquivo "cotacao.txt".

# Tecnologias utilizadas
- go1.22.3 linux/amd64 
- SQLITE3

# Como executar o projeto

## Pré-requisitos
GO - https://go.dev/

```bash
# clonar repositório
git clone https://github.com/waanvieira/simplified_payment_platform.git

# entrar na pasta do projeto back end
cd cotacao-dolar-go

# Executar o server
cd server

go run main.go

O projeto será executado http://localhost:8080/

# Executar o client
cd client

go run main.go

O projeto será executado http://localhost:8081/

```

# Uso do sistema

* Executar consulta do dolar no server

curl  -X GET 'http://localhost:8080/cotacao' \
  --header 'Accept: application/json' \  


* Executar consulta do dolar no client


curl  -X GET 'http://localhost:8081/price' \
  --header 'Accept: application/json' \  

# Autor

Wanderson Alves Vieira

https://www.linkedin.com/in/wanderson-alves-vieira-59b832148
