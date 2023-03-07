


# Define a imagem base que será utilizada
FROM golang:1.17-alpine

# Cria um diretório de trabalho para a aplicação
WORKDIR /app

# Copia o arquivo go.mod para o diretório de trabalho
COPY go.mod .

# Baixa as dependências necessárias
RUN go mod download

# Copia todo o código fonte para o diretório de trabalho
COPY . .

# Compila a aplicação
RUN go build -o app

# Define o comando que será executado quando o container iniciar
CMD ["./app"]


