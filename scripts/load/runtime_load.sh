#!/bin/bash

for i in {1..1000}
do
  curl -s localhost:18080/stats > /dev/null
done
