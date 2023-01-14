#pragma once

typedef struct s_attack
{
    char *path;
    char *address;
    int port;
    int power;
    int time;
} l7_attack;

extern void *http_get_flood(void *arg);
