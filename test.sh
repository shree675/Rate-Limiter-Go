#!/bin/bash

START_TIME=`date +%s%N`
for i in `seq 1 10`
do
    curl -XPOST 'http://localhost:8080/post' -d '{"id":'"$i"',"origin":"A"}' &
done

for i in `seq 1 10`
do
    curl -XPOST 'http://localhost:8080/post' -d '{"id":'"$i"',"origin":"B"}' &
done

for i in `seq 1 10`
do
    curl -XPOST 'http://localhost:8080/post' -d '{"id":'"$i"',"origin":"C"}' &
done

wait
END_TIME=`date +%s%N`

echo "Finished in: $((($END_TIME-$START_TIME)/1000000)) ms"