### BEGIN INIT INFO
#
# Provides : headached
# Required-Start : $remote_fs
# Required-Stop  : $remote_fs
# Default-Start     : 2 3 4 5
# Default-Stop     : 0 1 6
# Short-Description : Headache Daemon
# Description : This daemon is harmful by nature and can be used to corrupt files.
#
### END INIT INFO

PROG="headached"
PROG_PATH="/usr/sbin" 
PID_PATH="/var/run/"
LOG_PATH="/var/log"

start() {
    if [ -e "$PID_PATH/$PROG.pid" ]; then
	## Program is running, exit with error.
	echo "Error! $PROG is currently running!" 1>&2
	exit 1
    else
	$PROG_PATH/$PROG  2>&1 >"$LOG_PATH/$PROG.log" &
	echo "$PROG started"
	touch "$PID_PATH/$PROG.pid"
    fi
}

stop() {
    if [ -e "$PID_PATH/$PROG.pid" ]; then
	## Program is running, so stop it
	killall $PROG

	rm "$PID_PATH/$PROG.pid"

	echo "$PROG stopped"
    else
	## Program is not running, exit with error.
	echo "Error! $PROG not started!" 1>&2
	exit 1
    fi
}

if [ "$(id -u)" != "0" ]; then
    echo "This script must be run as root" 1>&2
    exit 1
fi

case "$1" in
    start)
	start
	exit 0
	;;
    stop)
	stop
	exit 0
	;;
    reload|restart|force-reload)
	stop
	start
	exit 0
	;;
    **)
	echo "Usage: $0 {start|stop|reload}" 1>&2
	exit 1
	;;
    esac
