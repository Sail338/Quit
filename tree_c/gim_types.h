#ifndef _gim_types
#define _gim_types

typedef struct TNode{
	struct	TNode* Parent;
	struct TNode **Children;
	char *data;
	char *type;
	char *path;
} TNode;

extern TNode* newTNode();

typedef struct GimMasterTree  {
	   char* MasterDir;
	   TNode* MasterRoot;
	    char* MasterRecordDir;
} GimMasterTree;

extern GimMasterTree* NewGimMasterTree();

typedef  struct Blob {
	    char* ShaHash;
	    int64_t ModTime;
	    char* ActualPath;
}Blob;


#endif
