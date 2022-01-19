FROM golang

EXPOSE 10000

WORKDIR /app

ADD . /app

RUN apt-get --yes update
RUN apt-get --yes install sqlite3
RUN go get /app/server
RUN go get /app/client

CMD [ "go", "run", "/app/server/" ]
