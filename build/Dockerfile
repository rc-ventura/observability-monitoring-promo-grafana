
FROM harbor.sflabs.com.br/processo/petshop:latest

# Define a pasta de trabalho do container
WORKDIR /usr/src/app

# Copia o arquivo package.json para a pasta de trabalho

COPY package.json /usr/src/app/

#Instala o node e faza atualização
RUN apk add --update nodejs npm

# Instala as dependências
RUN npm install

# Copia o código-fonte da aplicação para a pasta de trabalho
COPY . .

# Define a porta que a aplicação vai escutar
EXPOSE 3000

# Define o comando que será executado quando o container for iniciado
CMD ["npm", "start"]
