#define _GNU_SOURCE

// include needed library
#include <sys/sysinfo.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string.h>
#include "network.h"
#include "../utils/util.h"

/**
 * @brief Make http-get requests
 *
 * @param path resource path
 * @param address address of webserver
 * @param port port of webserver
 * @return char* response-text
 */
HttpResponse *http_get(char *path, char *address, int port)
{
    // create socket
    int sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock == -1)
    {
        return NULL;
    }

    // create socket address
    struct sockaddr_in *server = malloc(sizeof(struct sockaddr_in));
    if (server == NULL)
    {
        return NULL;
    }

    // setup socket address
    server->sin_family = AF_INET;
    server->sin_port = htons(port);
    server->sin_addr.s_addr = inet_addr(address);

    // connect to server
    if (connect(sock, (struct sockaddr *)server, sizeof(struct sockaddr_in)) == -1)
    {
        return NULL;
    }

    // create http request
    /**
     * Todo:
     *  - fix error
     *  - need to remove 1x \r\n
     */
    char *request = NULL;
    if (asprintf(&request, "GET %s HTTP/1.1\r\nHost: %s:%d\r\n\r\n\r\n", path, address, port) == -1)
    {
        return NULL;
    }

    // HTTP/1.1\r\nHost: 127.0.0.1:1337\r\nUser-Agent: curl/7.83.1\r\nAccept: */*\r\n\r\n

    printf("Request: %s", request);

    // send http request
    if (send(sock, request, strlen(request), 0) == -1)
    {
        return NULL;
    }

    // create buffer to store response
    char *response = malloc(1);
    if (response == NULL)
    {
        return NULL;
    }

    // create buffer to store response data
    char *buffer = malloc(1);
    if (buffer == NULL)
    {
        return NULL;
    }

    // read response data
    int bytes_read = 0;
    while ((bytes_read = recv(sock, buffer, 1, 0)) > 0)
    {
        // append response data to response
        response = realloc(response, strlen(response) + bytes_read + 1);
        if (response == NULL)
        {
            return NULL;
        }
        strncat(response, buffer, bytes_read);
    }

    printf("========================\n%s\n========================\n", response);

    // create http response
    HttpResponse *http_response = malloc(sizeof(HttpResponse));

    /*
        response format:
            HTTP/1.1 200 OK
            Date: Mon, 07 Nov 2022 11:21:35 GMT
            Content-Length: 58
            Content-Type: text/plain; charset=utf-8

            0.0.2|http://85.31.44.75:3333/download?arch=test|rose_testHTTP/1.1 400 Bad Request
            Content-Type: text/plain; charset=utf-8
            Connection: close

            400 Bad Request
    */

    // get status code
    char *status_code = strtok(response, " ");
    status_code = strtok(NULL, " ");
    http_response->StatusCode = atoi(status_code);

    // get content
    char *content = NULL;

    // get response body
    while ((content = strtok(NULL, "\r\n")) != NULL)
    {
        printf("content: %s\n", content);
    }

    printf("Status Code: %d\n", http_response->StatusCode);
    printf("Content: %s\n", http_response->Content);

    // free memory
    free(response);
    free(request);
    free(buffer);
    free(server);

    // return response
    return http_response;
}

/**
 * @brief Download content from url
 *
 * @param filename filename to save content to
 * @param url url to download content from
 * @return int (BOOL)
 */
int download_bin(char *filename, char *url)
{
    char *cmd = NULL;

    if (asprintf(&cmd, "wget -O %s %s || curl -o %s %s", filename, url, filename, url) == -1)
    {
        return FALSE;
    }

    if (system(cmd) == -1)
    {
        return FALSE;
    }

    free(cmd);
    return TRUE;
}
