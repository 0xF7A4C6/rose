#define _GNU_SOURCE

#include <dirent.h>
#include <stdlib.h>
#include <stdio.h>
#include "util.h"

/**
 * @brief Read a file and return it's content as an array of lines
 *
 * @param path File path
 * @return char**
 */
char **read_file(char *path)
{
    FILE *fp = fopen(path, "r");
    if (fp == NULL)
        return NULL;

    char **lines = NULL;
    char *line = NULL;
    size_t len = 0;
    ssize_t read;

    while ((read = getline(&line, &len, fp)) != FAILURE)
    {
        add_string(lines, &len, line);
    }

    fclose(fp);
    if (line)
        free(line);

    return lines;
}
