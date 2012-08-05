#include <stdlib.h>
#include <string.h>
#include "robosolver.h"

static hashEntry* table[1<<24];

static int match(location* robot, unsigned char* hashed) {
    for (int i=0; i<4; i++) {
        if (hashed[i]!=(unsigned char) robot[i])
            return 0;
    }

    return 1;
}

static hashEntry* mkhash(location* robot, unsigned remainingDepth) {
    hashEntry* entry = (hashEntry*) malloc(sizeof(hashEntry));
    entry->next = NULL;
    for (int i=0; i<4; i++) {
        entry->robots[i] = (unsigned char) robot[i];
    }
    entry->remainingDepth = remainingDepth;
}

int lookup(location* robot, unsigned remainingDepth) {
    unsigned hash = 0;
    for (int i=0; i<4; i++) {
        hash = 37*hash+robot[i];
    }
    
    hash &= (1<<24)-1;
    if (!table[hash]) {
        table[hash] = mkhash(robot, remainingDepth);
        return 0;
    }

    hashEntry* entry = table[hash];
    while (1) {
        if (match(robot, entry->robots)) {
            if (remainingDepth<=entry->remainingDepth) {
                return 1;
            }

            entry->remainingDepth = remainingDepth;
            return 0;
        }

        if (!entry->next) {
            entry->next = mkhash(robot, remainingDepth);
            return 0;
        }

        entry = entry->next;
    }
}
