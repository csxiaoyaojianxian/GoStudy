#ifndef _LINK_TABLE_H_
#define _LINK_TABLE_H_
#define SUCCESS 0
#define FAILURE (-1)

typedef struct LinkTableNode
{
     struct LinkTableNode * pNext;
} tLinkTableNode;
typedef struct LinkTable
{
    tLinkTableNode *pHead;
    tLinkTableNode *pTail;
    int SumOfNode;
} tLinkTable;
// Create a LinkTable
tLinkTable * CreateLinkTable();
// Delete a LinkTable
int DeleteLinkTable(tLinkTable *pLinkTable);
// Add a LinktableNode to Linktable
int AddLinkTableNode(tLinkTable *pLinkTable,tLinkTableNode *pNode);
// Delete a LinktableNode from Linktable
int DelLinkTableNode(tLinkTable *pLinkTable,tLinkTableNode *pNode);
// Get LinktableHead
tLinkTableNode *GetLinkTableHead(tLinkTable *pLinkTable);
// Get next LinktableNode
tLinkTableNode *GetNextLinkTableNode(tLinkTable *pLinkTable,tLinkTableNode *pNode);

#endif // _LINK_TABLE_H_
