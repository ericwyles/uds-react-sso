# Use the official lightweight NGINX image
#FROM cgr.dev/du-uds-defenseunicorns/nginx:1.27.4
FROM nginx:alpine

# Copy the static HTML file into the container
COPY index.html /usr/share/nginx/html/index.html
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 8080

# Start NGINX
CMD ["nginx", "-g", "daemon off;"]