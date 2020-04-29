FROM nginx:1.17.10-alpine

COPY default.conf /etc/nginx/conf.d/
COPY build /usr/share/nginx/html/

CMD ["nginx", "-g", "daemon off;"]
