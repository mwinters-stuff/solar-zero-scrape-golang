FROM ubuntu:latest
RUN apt-get update && apt-get -y upgrade && apt-get -y install apt-utils ca-certificates
RUN DEBIAN_FRONTEND=noninteractive apt-get -y install tzdata
WORKDIR /app
RUN mkdir /config
COPY solar-zero-scrape /app/
ENTRYPOINT [ "/app/solar-zero-scrape", "--config", "/config/solar-zero-scrape.json" ]