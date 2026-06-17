#!/bin/bash

while true
do
	printf "helix\n" | nc localhost 9100
	sleep 0.01
done
