#include <stdio.h>
#include <stdlib.h>
#include "linklist.h"

tDataNode* FindCmd(tDataNode *head, char *cmd) {
    if (head == NULL || cmd == NULL)
        return NULL;
    tDataNode *p = head;
    while (p != NULL) {
        if (strcmp(cmd, p->cmd) == 0)
            return p;
        p = p->next;
    }
    return NULL;
}

int ShowAllCmd(tDataNode *head) {
    printf("Menu list:\n");
    tDataNode *p = head;
    while (p != NULL) {
        printf("%s - %s\n", p->cmd, p->desc);
        p = p->next;
    }
    return 0;
}

