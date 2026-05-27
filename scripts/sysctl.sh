#!/bin/bash

sysctl -w fs.file-max=10000000

sysctl -w net.core.somaxconn=65535

sysctl -w net.ipv4.tcp_max_syn_backlog=262144

sysctl -w net.ipv4.ip_local_port_range="1024 65535"

sysctl -w net.ipv4.tcp_tw_reuse=1

sysctl -w net.ipv4.tcp_fin_timeout=15

ulimit -n 10000000
