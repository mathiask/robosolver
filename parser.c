#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "robosolver.h"

extern unsigned N;

// Example:
// XXXXXXXXXXX\
// X .X. . . X\
// X     X   X\
// X . 2 . . X\
// X         X\
// X . . . . X\
// X         X\
// X . . . . X\
// XXXXXXXXXXX
field parse(const char *p) {
    field f=(field) malloc(N*N+1);
    char s[3];
    s[2] = '\0';
    for (location xy=0; xy<N*N; xy++) {
        while (*p==' ')
            p++;
        s[0] = *p++;
        s[1] = *p++;
        f[xy] = strtol(s, NULL, 16);
    }

    return f;
}

char *prettyPrint(field f) {
    static char* t=NULL;
    if (t==NULL)
        t=malloc((3*N+1)*N+1);

    char* s = t;
    for (location xy = 0; xy<N*N; xy++) {
        sprintf(s, "%02x", f[xy]);
        s += 2;
        if (X(xy)<N-1) {
            *s++ = ' ';
        } else {
            *s++ = '\\';
            *s++ = '\n';
        }
    }

    *s++ = '\0';
    return t;
}
