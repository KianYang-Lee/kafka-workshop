# Kafka Workshop

This repo contains material for Kafka workshop.

## Setup

Kafka broker is setup and ran in a container using `docker compose`. Client can
be ran using `bitnami/kafka` container.

## Start-up

Just run the following to start up the broker. It will be exposed via localhost
on port `9092`.

```sh
$ docker compose up
```

The same container image can be used to spin up a CLI instance with the
following:

```sh
$ docker run -it --rm --network host bitnami/kafka:latest bash
```
