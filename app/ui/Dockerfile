# === STAGE 1: Build the React app ===
FROM node:20-alpine AS builder
WORKDIR /app
COPY sampleapp/package.json sampleapp/package-lock.json ./
RUN npm install --frozen-lockfile
COPY sampleapp/. .
RUN npm run build


FROM nginx:alpine
COPY index.html /usr/share/nginx/html/index.html
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
