# Etapa de build
FROM node:20-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependência
COPY package*.json ./
RUN npm install

# Copiar código fonte
COPY . .

# Compilar aplicação NestJS
RUN npm run build

# Etapa final
FROM node:20-alpine
WORKDIR /app

# Copiar arquivos necessários da etapa de build
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package*.json ./

EXPOSE 3000
CMD ["node", "dist/main.js"]
