#!/sbin/runscript

depend() {
  need localmount
}

start() {
    ebegin "Starting headached."
      start-stop-daemon --start \
          --pidfile /var/run/headached.pid --background --make-pidfile \
	  --exec /bin/bash -- -c "exec /usr/sbin/headached > /var/log/headached.log 2>&1"
    eend $?
}

stop() {
  ebegin "Stopping headached"
    start-stop-daemon --stop --exec /usr/sbin/headached \
        --pidfile /var/run/headached.pid
  eend $?
  
}