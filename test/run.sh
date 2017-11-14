#! /bin/bash

./server &

rm -f /tmp/proxy.sock
../proxy -l /tmp/proxy.sock &
