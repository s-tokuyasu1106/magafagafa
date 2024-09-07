FROM nginx:latest

COPY ./nginx.conf /etc/nginx/conf.d/nginx.conf
ADD _wildcard.magafagafa.local.pem /etc/certs/_wildcard.magafagafa.local.pem
ADD _wildcard.magafagafa.local-key.pem /etc/certs/_wildcard.magafagafa.local-key.pem
EXPOSE 443