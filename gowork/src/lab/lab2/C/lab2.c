#include<stdio.h>
#include<stdlib.h>

int cmdHelp()
{
    printf( "***********************Help**********************\n"
            "*\t\t\t\t\t\t*\n"
            "*\thelp\tShow help list\t\t\t*\n"
            "*\tls\tList files\t\t\t*\n"
            "*\tpwd\tPrint working directory\t\t*\n"
            "*\tmkdir\tCreate directory\t\t*\n"
            "*\tcp\tCopy file\t\t\t*\n"
            "*\tmv\tMove or rename\t\t\t*\n"
            "*\trm\tDelete file\t\t\t*\n"
            "*\tprint\tPrint Code\t\t\t*\n"
            "*\tquit\tQuit menu program\t\t*\n"
            "*\t\t\t\t\t\t*\n"
            "*************************************************\n\n"
    );
    return 0;
}

int printCode()
{
    char c;
    FILE *fp = NULL;
    fp = fopen("menu.c", "r");
    if (NULL == fp)
    {
        printf("Error: menu.c does not exist!\n");
        return -1;
    }
    while (fscanf(fp, "%c", &c) != EOF)
        printf("%c", c);
    fclose(fp);
    fp = NULL;
    return 0;
}

int main()
{
    char cmd[128];
    cmdHelp();
    while (1)
    {
        printf("Menu->");
        scanf("%s", cmd);
        if (strcmp(cmd, "help") == 0)
        {
            cmdHelp();
        }
        else if (strcmp(cmd, "ls") == 0)
        {
            system(cmd);
        }
        else if (strcmp(cmd, "pwd") == 0)
        {
            system(cmd);
        }
        else if (strcmp(cmd, "mkdir") == 0)
        {
            system(cmd);
        }
        else if (strcmp(cmd, "cp") == 0)
        {
            system(cmd);
        }
        else if (strcmp(cmd, "mv") == 0)
        {
            system(cmd);
        }
        else if (strcmp(cmd, "rm") == 0)
        {
            system(cmd);
        }
        else if (strcmp(cmd, "print") == 0)
        {
            printCode();
        }
        else if (strcmp(cmd, "quit") == 0)
        {
            exit(0);
        }
        else
        {
            printf("Error: unsupported command, you can use 'help' to list the available commands\n");
        }
    }
}
