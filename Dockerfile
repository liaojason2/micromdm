#Build MicroMDM
FROM golang:buster AS build

WORKDIR /src

COPY . /src

RUN apt update

RUN apt-get install -y make

RUN make

FROM debian:buster

RUN apt update

RUN apt-get install -y make gnupg curl

RUN echo "deb http://packages.cloud.google.com/apt gcsfuse-buster main" | tee /etc/apt/sources.list.d/gcsfuse.list \
  && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -

RUN apt update

RUN apt-get install -y make gcsfuse

COPY --from=build /src/build/linux/micromdm /usr/bin

COPY --from=build /src/entrypoint.sh .

COPY --from=build /src/bricks-micromdm.json .

EXPOSE 80 443 2195 2196 2197 5223

ENTRYPOINT [ "/entrypoint.sh" ]

#VOLUME ["/var/db/micromdm"]

CMD ["micromdm", "serve", "-server-url", "https://bricks-micromdm-test-u6idxxekbq-de.a.run.app", "-api-key", "12345", "-tls=false"]
