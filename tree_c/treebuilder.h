#ifndef _treebuilder
#define  _treebuilder
extern GimMasterTree* InitGimMasterTree (char *root);
extern TNode* GenerateTree(char *path);
extern TNode* GenerateTNode(char *fpath);
extern void WalkNode(TNode *node, int level);
char *writeBlob(Blob *blolb);
char * writeTree(TNode* root);
#endif
