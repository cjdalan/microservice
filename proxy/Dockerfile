FROM nginx:latest

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./nginx/sites-enabled/default /etc/nginx/sites-enabled/default

RUN service nginx restart