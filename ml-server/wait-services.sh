#!/bin/bash

set -e

kafka_host="kafka"
kafka_port="9092"
cmd="$@"

# Wait for Kafka to be available
until ncat -z $kafka_host $kafka_port; do
   >&1 echo "Kafka is unavailable - sleeping"
   sleep 1
done

>&1 echo "Kafka is up"

# Execute the main command
exec $cmd
