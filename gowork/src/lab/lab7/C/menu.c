#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include "linktable.h"
int Help(int argc, char* argv[]);
extern char* optarg;
extern int optind;
#define CMD_MAX_LEN 128
#define CMD_MAX_ARGV 128
#define DESC_LEN    1024
#define CMD_NUM     10
/* data struct and its operations */
tLinkTable* head = NULL;
typedef struct DataNode
{
    tLinkTableNode * pNext;
    char*   cmd;
    char*   desc;
    int     (*handler)(int argc, char* argv[]);
} tDataNode;
int SearchCondition(tLinkTableNode * pLinkTableNode, void * args)
{
    tDataNode * pNode = (tDataNode *)pLinkTableNode;
    char * tempchar = args;
    if(strcmp(pNode->cmd, args) == 0)
    {
        return  SUCCESS;
    }
    return FAILURE;
}
/* find a cmd in the linklist and return the datanode pointer */
tDataNode* FindCmd(tLinkTable * head, char * cmd)
{
    return  (tDataNode*)SearchLinkTableNode(head,SearchCondition, cmd);
}
/* show all cmd in listlist */
int ShowAllCmd(tLinkTable * head)
{
    tDataNode * pNode = (tDataNode*)GetLinkTableHead(head);
    printf("---------------------------------------------\n");
    while(pNode != NULL)
    {
        printf("%s - %s\n", pNode->cmd, pNode->desc);
        pNode = (tDataNode*)GetNextLinkTableNode(head,(tLinkTableNode *)pNode);
    }
    printf("---------------------------------------------\n");
    return 0;
}
int ReceiveCMD(char* cmd, int* argc, char* argv[])
{
    char* pcmd = NULL;
    pcmd = fgets(cmd, CMD_MAX_LEN, stdin);
    int len = strlen(cmd);
    *(cmd + len - 1) = '\0';
    pcmd = strtok(pcmd, " ");
    while(pcmd == NULL)
    {
        return -1;
    }
    while(pcmd != NULL && (*argc) < CMD_MAX_ARGV)
    {
        argv[*argc] = pcmd;
        (*argc)++;
        pcmd = strtok(NULL, " ");
    }
    return 0;
}
int MenuConfig(char* cmd, char* desc, int (*handler)(int argc, char* argv[]))
{
    tDataNode* pNode = NULL;
    if(head == NULL)
    {
        head = CreateLinkTable();
        pNode = (tDataNode*)malloc(sizeof(tDataNode));
        pNode->cmd = "help";
        pNode->desc = "list all cmds";
        pNode->handler = Help;
        AddLinkTableNode(head, (tLinkTableNode*)pNode);
    }
    pNode = (tDataNode*)malloc(sizeof(tDataNode));
    pNode->cmd = cmd;
    pNode->desc = desc;
    pNode->handler = handler;
    AddLinkTableNode(head, (tLinkTableNode*)pNode);
    return 0;
}
/* menu program */
int ExecuteMenu()
{
    /* cmd line begins */
    while(1)
    {
        int argc = 0;
        char* argv[CMD_MAX_ARGV];
        char cmd[CMD_MAX_LEN];
        printf("Input a cmd number > ");
        if(ReceiveCMD(cmd, &argc, argv) == -1)
        {
            continue;
        }
        //scanf("%s", cmd);
        tDataNode *p = FindCmd(head, argv[0]);
        if( p == NULL)
        {
            printf("This is a wrong cmd!\n ");
            continue;
        }
        printf("%s - %s\n", p->cmd, p->desc);
        if(p->handler != NULL)
        {
            p->handler(argc, argv);
        }
        
    }
}
int Help(int argc, char* argv[])
{
    int ch;
    char* ch_prom;
    if(argc == 1)
    {
        ShowAllCmd(head);
        return 0;
    }
    while((ch = getopt(argc, argv, "shl:")) != -1)
    {
        switch(ch)
        {
            case 's':
                printf("This is \"-s\" mode help\n");
                ShowAllCmd(head);
                break;
            case 'h':
                printf("This is \"-h\" mode help\n");
                ShowAllCmd(head);
                break;
            case 'l':
                ch_prom = optarg;
                printf("This is \"-l\" mode help with %s\n", ch_prom);
                ShowAllCmd(head);
                break;
            case '?':
                printf("Wrong argument!\n");
        }
    }
    optind = 1;
    return 0; 
}

