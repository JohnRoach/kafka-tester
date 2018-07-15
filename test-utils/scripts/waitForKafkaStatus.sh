#!/usr/bin/env bash

i=0

while [ $i -lt 120 ]
do
   kubectl exec $(kubectl get pod -l app=kafka-test -o jsonpath="{.items[0].metadata.name}") --container='kafka' -- cub kafka-ready -b localhost:29092 1 60
   if [ $? != 0 ]; then
       echo "Kafka hasn't started yet... let us wait a little more..."
   else
       echo "Kafka has started!"
       exit 0
   fi
   sleep 2
   i=$[$i+1]
done

echo "Kafka couldn't start in 240 seconds..."
exit 1