#define GO_C_LIB_EXPORTS

#include "../include/go4c.h"
#include <stdio.h>

FnCallBackLib_C g_callback = 0;

GO_C_LIB_API int32_t Go4CInit_C(FnCallBackLib_C callback)
{
    printf("hello world\n");
    callback((char*)"hello", 4);
    g_callback = callback;

    return 0;
}

GO_C_LIB_API int32_t Go4CRelease_C()
{
    return 0;
}

GO_C_LIB_API int32_t Go4CInitCommand_C(char* data, int32_t len, FnCallBackCmd_C cmd_callback, int32_t cmdId)
{
    printf("hello cmd: %s\n", data);
    if (cmd_callback) {
        cmd_callback(cmdId, data, len);
    }
    if (g_callback) {
        g_callback(data, len);
    }
    return 0;
}
