#include <stdio.h>
#include <stdlib.h>
#include "linklist.h"

#define CMD_MAX_LEN 128
#define DESC_LEN 1024
#define CMD_NUM 10

int Help();
int Quit();

/*menu program*/
static tDataNode head[] = {
    { "help","This is a help cmd.",Help,&head[1] },
    { "version","menu program v2.0.",NULL,&head[2] },
    { "quit","Quit form menu.",Quit,NULL }
};

int main() {
    /*cmd line begin*/
    while (1) {
        char cmd[CMD_MAX_LEN];
        printf("Please input a cmd:>");
        scanf("%s", cmd);
        tDataNode* p = FindCmd(head, cmd);
        if (p == NULL) {
            printf("This is a wrong cmd!");
            continue;
        }
        printf("%s - %s\n", p->cmd, p->desc);
        if (p->handler != NULL)
            p->handler();
    }
    return 0;
}

int Help() {
    ShowAllCmd(head);
    return 0;
}

int Quit() {
    exit(0);
}

