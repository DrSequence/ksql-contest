#  /bin/sh

echo "build event producer"

cd ../event-producer && nohup make run-local &

echo "builded"

echo "build order producer"

cd ../order-producer && nohup make run-local &

echo "builded"