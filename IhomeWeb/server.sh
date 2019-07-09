redis-server ./conf/redis.conf

fdfs_trackerd  /home/lcc/go/src/renting/IhomeWeb/conf/tracker.conf restart

fdfs_storaged  /home/lcc/go/src/renting/IhomeWeb/conf/storage.conf restart

sudo nginx
