FROM nginx:alpine
COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY cmd/general /usr/share/nginx/html
WORKDIR /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]