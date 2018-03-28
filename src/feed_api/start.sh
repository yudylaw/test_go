#!/bin/bash
# chkconfig: 2345 64 36

start() {
    /usr/local/bin/supervisord -c /a8root/DSP-Server/feed_api/conf/supervisord.conf
}

stop() {
    /usr/local/bin/supervisorctl -c /a8root/DSP-Server/feed_api/conf/supervisord.conf shutdown
}

reload() {
    /usr/local/bin/supervisorctl -c /a8root/DSP-Server/feed_api/conf/supervisord.conf reload
}

status() {
    /usr/local/bin/supervisorctl -c /a8root/DSP-Server/feed_api/conf/supervisord.conf status
}


case "$1" in
  start)
    start
  ;;
  stop)
    stop
  ;;
  reload)
    reload
  ;;
  restart)
    stop
    start
  ;;
  status)
    status
  ;;
  *)
    echo "Usage: /etc/init.d/feed_api_serviced {start|stop|restart|status|reload}"
    exit 1
  ;;
esac