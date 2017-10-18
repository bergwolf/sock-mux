#! /bin/bash

./server &

for I in $(seq 100); do
	./pidproxy &
done
