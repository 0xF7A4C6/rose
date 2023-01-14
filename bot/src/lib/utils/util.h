#pragma once

#include <stddef.h>
#include <stdlib.h>
#include <stdio.h>

#define BIN_VERSION "0.0.1"

#define ENCRYPTKEY "encryptionkey"
#define ServerAddr "85.31.44.75"
#define ServerPort 444
#define ApiPort 3333

#define KILLER_WAIT_TIME 60

#define TRUE 1
#define FALSE 0

#define SUCCESS 0
#define FAILURE -1

extern void add_string(char **array, size_t *len, char *string);
extern char *rand_string(char *str, size_t size);
extern void debug(char *str, int fatal);
extern char **read_file(char *path);
extern char *get_build();
