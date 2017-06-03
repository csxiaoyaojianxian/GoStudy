#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <time.h>
#include "linktable.h"
#include "menu.h"

int ShowTime(int argc, char *argv[])
{
    time_t timep;
    time(&timep);
    printf("%s", ctime(&timep));
    return 0;
}

int Quit(int argc, char *argv[])
{
    exit(0);
}

int main(int argc, char *argv[])
{
    MenuConfig("version", "menu program v3.0", NULL);
    MenuConfig("time", "show time now", ShowTime);
    MenuConfig("quit", "leave the menu program", Quit);
    
    ExecuteMenu();
    return 0;
}
