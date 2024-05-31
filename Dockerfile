# Imagem base
FROM golang

# Adiciona informações da pessoa que mantem a imagem

# diretoria um diretorio de trabalho
WORKDIR /app

# aponta a variavel gopath do go para o diretorio app
ENV GOPATH=/app

# copia os arquivos do projeto para o workdir do container
COPY . /app

# execulta o main.go e baixa as dependencias do projeto
RUN go build main.go

# Comando para rodar o executavel
ENTRYPOINT ["./main"]

# expõe a pota 8080
EXPOSE 8080