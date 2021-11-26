FROM alpine

RUN mkdir /var/www

ADD licensephobia /var/www
WORKDIR /var/www

RUN ls -l

RUN ./licensephobia