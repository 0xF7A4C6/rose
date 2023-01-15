## Technical description of new persistence capabilities

Apart from its known extensive P2P and DDoS abilities, we have recently observed several new and unique capabilities of the Mozi botnet.

Targeting Netgear, Huawei, and ZTE gateways, the malware now takes specific actions to increase its chances of survival upon reboot or any other attempt by other malware or responders to interfere with its operation. Here are some examples:

### Achieving privileged persistence

A specific check is conducted for the existence of the **/overlay** folder, and whether the malware does not have write permissions to the folder **/etc**. In this case, it will try to exploit **[CVE-2015-1328](https://nvd.nist.gov/vuln/detail/CVE-2015-1328)**.

Successful exploitation of the vulnerability will grant the malware access to the following folders:

*   /etc/rc.d
*   /etc/init.d

Then the following actions are taken:

*   It places the script file named **S95Baby.sh** in these folders.
*   The script runs the files **/usr/networks** or **/user/networktmp**. These are copies of the executable.
*   It adds the script to **/etc/rcS.d** and **/etc/rc.loca**l in case it lacks privileges.

### ZTE devices

A specific check is conducted for the existence of the **/usr/local/ct** folder; this serves as an indicator of the device being a ZTE modem/router device.

The following actions are taken:

*   It copies its other instance **(/usr/networks)** to **/usr/local/ct/ctadmin0**; this provides persistency for the malware.
*   It deletes the file **/home/httpd/web_shell_cmd.gch**. This file can be used to gain access through exploitation of the vulnerability [**CVE-2014-2321**](https://nvd.nist.gov/vuln/detail/CVE-2014-2321); deleting it prevents future attacks.
*   It executes the following commands. These disable **Tr-069** and its ability to connect to auto-configuration server (ACS). **Tr-069** is a protocol for remote configuration of network devices; it’s usually utilized by service providers to configure customers’ equipment.

<pre style="padding-left: 80px;">sendcmd 1 DB set MgtServer 0 Tr069Enable 1 
sendcmd 1 DB set PdtMiddleWare 0 Tr069Enable 0 
sendcmd 1 DB set MgtServer 0 URL http://127.0.0.1 
sendcmd 1 DB set MgtServer 0 UserName notitms 
sendcmd 1 DB set MgtServer 0 ConnectionRequestUsername notitms 
sendcmd 1 DB set MgtServer 0 PeriodicInformEnable 0 
sendcmd 1 DB save</pre>

### Huawei devices

Execution of the following commands changes the password and disables the management server for Huawei modem/router devices. It also prevents others from gaining access to the device through the management server.

<pre style="padding-left: 80px;">cfgtool set /mnt/jffs2/hw_ctree.xml 
InternetGatewayDevice.ManagementServer URL http://127.0.0.1
cfgtool set /mnt/jffs2/hw_ctree.xml 
InternetGatewayDevice.ManagementServer ConnectionRequestPassword acsMozi</pre>

To provide an additional level of persistence it also creates the following files if needed and appends an instruction to run its copy from **/usr/networks**.

<pre style="padding-left: 80px;">/mnt/jffs2/Equip.sh
/mnt/jffs2/wifi.sh
/mnt/jffs2/WifiPerformance.sh</pre>

### Preventing remote access

The malware blocks the following TCP ports:

*   23—Telnet
*   2323—Telnet alternate port
*   7547—Tr-069 port
*   35000—Tr-069 port on Netgear devices
*   50023—Management port on Huawei devices
*   58000—Unknown usage

These ports are used to gain remote access to the device. Shutting them increases the malware’s chances of survival.

### Script infector

It scans for **.sh** files in the filesystem, excluding the following paths:

<pre style="padding-left: 80px;">/tmp /dev /var /lib /haha /proc /sys</pre>

It also appends a line to each file. The line instructs the script to run a copy of the malware from **/usr/networks**. This increases its chances of survival on various devices.

### Traffic injection and DNS spoofing capabilities

The malware receives commands from its distributed hash table (DHT) network. The latter is a P2P protocol for decentralized communications. The commands are received and stored in a file, of which parts are encrypted. This module works only on devices capable of IPv4 forwarding. It checks whether **/proc/sys/net/ipv4/ip_forward** is set to 1; such positive validation is characteristic of routers and gateways. This module works on ports UDP 53 (DNS) and TCP 80 (HTTP).

### Configuration commands

Apart from the previously documented commands in Table 1—for more information, read [A New Botnet Attack Just Mozied Into Town](https://malware.news/t/a-new-botnet-attack-just-mozied-into-town/43210)—we also discovered these commands:

<pre style="padding-left: 80px;">[hi] – Presence of the command indicates it needs to use the MiTM module.
[set] – Contains encrypted portion which describes how to use the MiTM module.</pre>

<figure class="wp-block-table">

<table>

<tbody>

<tr>

<td>**Command**</td>

<td>**Description**</td>

</tr>

<tr>

<td>**[ss]**</td>

<td>Bot role</td>

</tr>

<tr>

<td>**[ssx]**</td>

<td>enable/disable tag [ss]</td>

</tr>

<tr>

<td>**[cpu]**</td>

<td>CPU architecture</td>

</tr>

<tr>

<td>**[cpux]**</td>

<td>enable/disable tag [cpu]</td>

</tr>

<tr>

<td>**[nd]**</td>

<td>new DHT node</td>

</tr>

<tr>

<td>**[hp]**</td>

<td>DHT node hash prefix</td>

</tr>

<tr>

<td>**[atk]**</td>

<td>DDoS attack type</td>

</tr>

<tr>

<td>**[ver]**</td>

<td>Value in V section in DHT protocol</td>

</tr>

<tr>

<td>**[sv]**</td>

<td>Update config</td>

</tr>

<tr>

<td>**[ud]**</td>

<td>Update bot</td>

</tr>

<tr>

<td>**[dr]**</td>

<td>Download and execute payload from the specified URL</td>

</tr>

<tr>

<td>**[rn]**</td>

<td>Execute specified command</td>

</tr>

<tr>

<td>**[dip]**</td>

<td>ip:port to download Mozi bot</td>

</tr>

<tr>

<td>**[idp]**</td>

<td>report bot</td>

</tr>

<tr>

<td>**[count]**</td>

<td>URL that used to report bot</td>

</tr>

</tbody>

</table>

</figure>

_Table 1\. Previously documented Mozi commands._

