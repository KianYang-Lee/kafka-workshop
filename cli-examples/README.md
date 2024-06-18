# CLI Examples

Let's go through a list of common commands for Kafka administration.

## Topics

- List topic

```sh
$ kafka-topics.sh --bootstrap-server localhost:9092 --list
```

- Create topic

```sh
$ kafka-topics.sh --bootstrap-server localhost:9092 --create --topic "my-topic"
```

- Describe topic details

```sh
$ kafka-topics.sh --boostrap-server localhost:9092 --describe --topic my-topic
```

- Adding partitions

```sh
$ kafka-topics.sh --bootstrap-server localhost:9092 --alter --topic my-topic \
--partitions 16
```

## Consumer Groups

- List consumers

```sh
$ kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list
```

- Describe consumer

```sh
$ kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe \
--group my-consumer
```

- Delete Group

```sh
$ kafka-consumer-groups.sh --bootstrap-server localhost:9092 --delete --group \
my-consumer
```

## Dynamic Configuration Changes

- Changing retention topic to 1 hour

```sh
$ kafka-configs.sh --bootstrap-server localhost:9092 --alter --entity-type \
topics --entity-name my-topic --add-config retention.ms=3600000
```

- Describe configuration overrides

```sh
$ kafka-configs.sh --bootstrap-server localhost:9092 --describe --entity-type \
topics --entity-name my-topic
```

- Remove configuration overrides

```sh
$ kafka-configs.sh --bootstrap-server localhost:9092 --alter --entity-type \
topics --entity-name my-topic --delete-config retention.ms
```

## Console Producer

Convenience utility to perform producer actions using CLI instead of writing an
application. Used mainly for testing purpose.

- Producing messages

```sh
$ kafka-console-producer.sh --bootstrap-server localhost:9092 --topic my-topic

> Msg 1
> Msg 2
> Msg 3
>^D
```

We need to send EOF to close the client (`Control-D` in most cases).

To modify properties of producer client, we can either use a producer
configuration file by specifying `--producer.config <config-file>` or pass
additional arguments in the form `--producer-property <key>=<value>`, where key
is configuration option name.

## Console Consumer

Convenience utility to perform consumer actions using CLI.

- Consume messages

```sh
$ kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic \
my-topic --from-beginning
```

To exit, press `Control-C`.

Apart from choosing a specific topic, we can consume from a list of topics using
the `--whitelist` flag.

To modify properties of consumer client, we can either use a consumer
configuration file by specifying `--consumer.config <config-file>` or pass
additional arguments in the form `--consumer-property <key>=<value>`, where key
is configuration option name.

## Partition Management

We can manage partitions to do works like choosing a new leader and assigning
partitions to brokers.

- Initiate a leader reelection for all topics in a cluster

```sh
$ kafka-leader-election.sh --bootstrap-server localhost:9092 --election-type \
PREFERRED --all-topic-partitions
```

## Offset Management

We can manage offset using `kafka-consumer-groups.sh`.

```sh
$ kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group \
"wiki-test-group" --topic "wiki-test" --reset-offsets --to-offset 0 --execute
```

We can also choose certain partitions or topics to perform elections. We can
either use `--partition` flag or use a JSON file.
