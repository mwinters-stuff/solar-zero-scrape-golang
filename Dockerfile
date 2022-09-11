FROM ubuntu:latest
RUN apt-get update && apt-get upgrade && apt-get -y install ca-certificates
WORKDIR /app
RUN mkdir /config
COPY solar-zero-scrape /app/
ENTRYPOINT [ "/app/solar-zero-scrape", "/config/solar-zero-scrape.json" ]