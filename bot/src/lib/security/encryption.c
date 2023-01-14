#define _GNU_SOURCE

#include <string.h>
#include "../utils/util.h"

/**
 * @brief Encrypt a string using XOR algorithm
 *
 * @param str Data to encrypt
 * @return char*
 */
char *encrypt_str(char *str)
{
    int str_len = strlen(str);
    int key_len = strlen(ENCRYPTKEY);
    char *result = malloc(sizeof(char) * str_len);

    for (int i = 0; i < str_len; i++)
        result[i] = str[i] ^ ENCRYPTKEY[i % key_len];

    return result;
}
