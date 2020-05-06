#if !defined __AE__ERL_COMM_H
#define __AE__ERL_COMM_H

#ifdef DEBUG
#include <stdio.h>
#define DEBUG_PRINTF printf
#else
#define DEBUG_PRINTF
#endif

typedef unsigned char byte;

#define ECRECOVER_BUFFER_MAX 1024

int read_cmd(byte *buf);
int write_cmd(byte *buf, int len);
int read_exact(byte *buf, int len);
int write_exact(byte *buf, int len);

#endif
