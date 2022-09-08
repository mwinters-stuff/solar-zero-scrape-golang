FROM ubuntu:latest
WORKDIR /app
RUN mkdir /config
COPY solar-zero-scrape /app/
ENTRYPOINT [ "/app/solar-zero-scrape", "/config/solar-zero-scrape.json" ]