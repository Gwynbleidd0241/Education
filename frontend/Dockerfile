FROM nginx:alpine
COPY ./pages /usr/share/nginx/html/pages
COPY ./image /usr/share/nginx/html/image
COPY ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]