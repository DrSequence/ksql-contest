kill $(ps aux | grep 'event-producer' | awk '{print $2}') 
kill $(ps aux | grep 'order-producer' | awk '{print $2}')