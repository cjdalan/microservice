FROM nginx:latest

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./nginx/sites-enabled/default /etc/nginx/sites-enabled/default
RUN touch /usr/share/nginx/html/error.html

RUN service nginx restart