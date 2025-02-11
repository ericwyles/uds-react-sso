# === STAGE 1: Build the React app ===
FROM node:20-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy package.json and package-lock.json before running npm install (better caching)
COPY sampleapp/package.json sampleapp/package-lock.json ./

# Install dependencies
RUN npm install --frozen-lockfile

# Copy the rest of the application files
COPY sampleapp/. .

# Build the React app
RUN npm run build


FROM nginx:alpine

COPY index.html /usr/share/nginx/html/index.html
COPY nginx.conf /etc/nginx/nginx.conf

RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /app/build /usr/share/nginx/html

EXPOSE 8080

# Start NGINX
CMD ["nginx", "-g", "daemon off;"]