#include <assert.h>
#include <string.h>

#include "libmessage.c"

int
main()
{
    char* m1 = alloc_message("");
    char* m2 = alloc_message("hogehoge");

    assert(strcmp(m1, "Hello, ") == 0);
    assert(strcmp(m2, "Hello, hogehoge") == 0);

    return 0;
}
