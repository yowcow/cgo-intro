#include <stdlib.h>
#include <string.h>

char*
alloc_message(char* name)
{
    int len = strlen(name) + 8;
    char* message = malloc(sizeof(char) * len);
    strcpy(message, "Hello, ");
    strcat(message, name);
    return message;
}

void
free_message(char* message)
{
    free(message);
}
