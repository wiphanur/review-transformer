#!/bin/bash

set -e

mongo_host="mongodb"
mongo_port="27017"
kafka_host="kafka"
kafka_port="9092"
cmd="$@"

# Wait for MongoDB to be available
until ncat -z $mongo_host $mongo_port; do
   >&2 echo "MongoDB is unavailable - sleeping"
   sleep 1
done

>&2 echo "MongoDB is up"

# Wait for Kafka to be available
until ncat -z $kafka_host $kafka_port; do
   >&2 echo "Kafka is unavailable - sleeping"
   sleep 1
done

>&2 echo "Kafka is up"

# Execute the main command
exec $cmd
