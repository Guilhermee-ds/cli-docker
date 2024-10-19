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




FROM python:3.11-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the Python requirements file
# COPY requirements.txt ./

# Install dependencies
# RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application code
COPY . .

# Expose the application port (se aplicável)
EXPOSE 8000

# Run the Python application
CMD ["python", "app.py"]

