FROM node:20-slim AS builder

WORKDIR /app
COPY package*.json ./
RUN npm install --legacy-peer-deps

COPY . .
RUN npm run build

FROM nginx:1.25 as deploy

WORKDIR /app
COPY --from=builder /app/dist ./dist

CMD ["nginx", "-g", "daemon off;"]
