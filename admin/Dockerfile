FROM node:20-slim AS builder

WORKDIR /app
COPY . .

RUN npm install --legacy-peer-deps
RUN npm run build

FROM nginx:1.25 as deploy

WORKDIR /app
COPY --from=builder /app/dist ./dist
COPY nginx.config /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]
