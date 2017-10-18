#! /bin/bash

./server &

rm -f /tmp/proxy-*.sock
for I in $(seq 100); do
	./pidproxy &
done
