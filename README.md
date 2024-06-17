# Kafka Workshop

This repo contains material for Kafka workshop.

## Setup

Kafka broker is setup and ran in a container using `docker compose`. Client can
be ran using `bitnami/kafka` container.

## Start-up

First create the docker network `app-tier`. We will be connecting to the broker
on this network.

```sh
$ docker network create app-tier --driver bridge
```

Just run the following to start up the broker. It will be exposed using hostname
`kafka-server` on port `9092`.

```sh
$ docker compose up
```

The same container image can be used to spin up a CLI instance with the
following:

```sh
$ docker run -it --rm --network app-tier bitnami/kafka:latest bash
```
