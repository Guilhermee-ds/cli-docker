# Use a imagem oficial do Node.js
FROM node:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos package.json e package-lock.json antes de instalar as dependências
COPY package*.json ./


# Copie o restante do código da aplicação
COPY . .

# Exponha a porta que a aplicação utilizará
EXPOSE 3000

# Comando para iniciar a aplicação
CMD ["npm", "start"]


