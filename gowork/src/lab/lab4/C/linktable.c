#include <stdio.h>
#include <stdlib.h>
#include "linktable.h"

tLinkTable *CreateLinkTable()
{
    tLinkTable *pLinkTable = (tLinkTable*)malloc(sizeof(tLinkTable));
    if(pLinkTable == NULL)
    {
        return NULL;
    }
    pLinkTable->pHead = NULL;
    pLinkTable->pTail = NULL;
    pLinkTable->SumOfNode = 0;
    return pLinkTable;
}

int DeleteLinkTable(tLinkTable *pLinkTable)
{
    if(pLinkTable == NULL)
    {
        return FAILURE;
    }
    while(pLinkTable->pHead != NULL)
    {
        tLinkTableNode *tmp = pLinkTable->pHead;
        pLinkTable->pHead = pLinkTable->pHead->pNext;
        free(tmp);
    }
    pLinkTable->pHead = NULL;
    pLinkTable->pTail = NULL;
    pLinkTable->SumOfNode = -1;
    free(pLinkTable);
    return SUCCESS;
}

int AddLinkTableNode(tLinkTable *pLinkTable,tLinkTableNode *pNode)
{
    if(pLinkTable == NULL || pNode == NULL)
    {
        return FAILURE;
    }
    if(pLinkTable->pHead == NULL)
    {
        pLinkTable->pHead = pNode;
        pLinkTable->pTail = pNode;
        pLinkTable->SumOfNode = 1;
    }
    else
    {
        pLinkTable->pTail->pNext = pNode;
        pLinkTable->pTail = pNode;
        pLinkTable->SumOfNode ++;
    }
    return SUCCESS;
}

int DelLinkTableNode(tLinkTable *pLinkTable,tLinkTableNode *pNode)
{
    if(pLinkTable == NULL ||pNode == NULL)
    {
        return FAILURE;
    }
    if(pLinkTable->pHead == pNode)
    {
        tLinkTableNode *tmp = pLinkTable->pHead;
        pLinkTable->pHead = pLinkTable->pHead->pNext;
        free(tmp);
        pLinkTable->SumOfNode--;
    }
    if(pLinkTable->SumOfNode == 0)
    {
        pLinkTable->pTail = NULL;
        return FAILURE;
    }
    tLinkTableNode *tmp = pLinkTable->pHead;
    while(tmp != NULL)
    {
        if(tmp->pNext == pNode)
        {
            tmp->pNext = pNode ->pNext;
            pLinkTable->SumOfNode --;
            if(pLinkTable->SumOfNode == 0)
            {
                pLinkTable->pTail = NULL;
            }
            return SUCCESS;
        }
        tmp = tmp->pNext;
    }
    return FAILURE;
}

tLinkTableNode * GetLinkTableHead(tLinkTable *pLinkTable)
{
    if(pLinkTable == NULL||pLinkTable->pHead == NULL)
    {
        return NULL;
    }
    return pLinkTable->pHead;
}

tLinkTableNode *GetNextLinkTableNode(tLinkTable *pLinkTable,tLinkTableNode *pNode)
{
    if(pLinkTable == NULL||pNode == NULL)
    {
    return NULL;
    }
    tLinkTableNode *tmp = pLinkTable->pHead;
    while(tmp != NULL)
    {
        if(tmp == pNode)
        {
            return tmp->pNext;
        }
        tmp = tmp->pNext;
    }
    return NULL;
}
