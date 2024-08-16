#!/bin/bash

# Start RabbitMQ server in the background
rabbitmq-server &

# Wait for the .erlang.cookie file to be created
while [ ! -f /var/lib/rabbitmq/.erlang.cookie ]; do
  sleep 1
done

# Adjust permissions on the .erlang.cookie file
chmod 600 /var/lib/rabbitmq/.erlang.cookie

# Wait for RabbitMQ server to finish
wait
