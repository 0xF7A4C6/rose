#define _GNU_SOURCE

#include <stddef.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <assert.h>

/**
 * @brief Add a string to an array of strings
 *
 * @param array Array to fill
 * @param len Array length
 * @param string String to add
 */
void add_string(char **array, int *len, char *string)
{
    array = realloc(array, sizeof(char *) * (*len + 1));
    array[*len] = string;
    *len += 1;
}

/**
 * @brief Fetch logs into SDTOUT
 *
 * @param str String to print
 * @param fatal Is a fatal error (exit)
 */
void debug(char *str, int fatal)
{
    if (fatal)
    {
#ifdef DEBUG
        perror(str);
#endif
        exit(EXIT_FAILURE);
    }
    else
    {
#ifdef DEBUG
        printf("%s\n", str);
#endif
    }
}

/**
 * @brief Generate a random string
 *
 * @param str String to fill
 * @param size Random string size
 * @return char*
 */
char *rand_string(char *str, size_t size)
{
    const char charset[] = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";

    if (size)
    {
        --size;
        for (size_t n = 0; n < size; n++)
        {
            int key = rand() % (int)(sizeof charset - 1);
            str[n] = charset[key];
        }
        str[size] = '\0';
    }

    return str;
}

/**
 * @brief Get the CPU arch
 *
 * @return char*
 */
char *get_build()
{
#if defined(__x86_64__) || defined(_M_X64)
    return "x86_64";
#elif defined(i386) || defined(__i386__) || defined(__i386) || defined(_M_IX86)
    return "x86_32";
#elif defined(__ARM_ARCH_2__)
    return "ARM2";
#elif defined(__ARM_ARCH_3__) || defined(__ARM_ARCH_3M__)
    return "ARM3";
#elif defined(__ARM_ARCH_4T__) || defined(__TARGET_ARM_4T)
    return "ARM4T";
#elif defined(__ARM_ARCH_5_) || defined(__ARM_ARCH_5E_)
    return "ARM5"
#elif defined(__ARM_ARCH_6T2_) || defined(__ARM_ARCH_6T2_)
    return "ARM6T2";
#elif defined(__ARM_ARCH_6__) || defined(__ARM_ARCH_6J__) || defined(__ARM_ARCH_6K__) || defined(__ARM_ARCH_6Z__) || defined(__ARM_ARCH_6ZK__)
    return "ARM6";
#elif defined(__ARM_ARCH_7__) || defined(__ARM_ARCH_7A__) || defined(__ARM_ARCH_7R__) || defined(__ARM_ARCH_7M__) || defined(__ARM_ARCH_7S__)
    return "ARM7";
#elif defined(__ARM_ARCH_7A__) || defined(__ARM_ARCH_7R__) || defined(__ARM_ARCH_7M__) || defined(__ARM_ARCH_7S__)
    return "ARM7A";
#elif defined(__ARM_ARCH_7R__) || defined(__ARM_ARCH_7M__) || defined(__ARM_ARCH_7S__)
    return "ARM7R";
#elif defined(__ARM_ARCH_7M__)
    return "ARM7M";
#elif defined(__ARM_ARCH_7S__)
    return "ARM7S";
#elif defined(__aarch64__) || defined(_M_ARM64)
    return "ARM64";
#elif defined(mips) || defined(__mips__) || defined(__mips)
    return "MIPS";
#elif defined(__sh__)
    return "SUPERH";
#elif defined(__powerpc) || defined(__powerpc__) || defined(__powerpc64__) || defined(__POWERPC__) || defined(__ppc__) || defined(__PPC__) || defined(_ARCH_PPC)
    return "POWERPC";
#elif defined(__PPC64__) || defined(__ppc64__) || defined(_ARCH_PPC64)
    return "POWERPC64";
#elif defined(__sparc__) || defined(__sparc)
    return "SPARC";
#elif defined(__m68k__)
    return "M68K";
#else
    return "UNKNOWN";
#endif
}
