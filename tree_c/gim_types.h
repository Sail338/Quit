#ifdef _gim_types.h
#define _gim_types.h
typedef struct TNode{
	TNode *Parent;
	TNode *Children;
	char *data;
	char *type;
	char *path
} TNode;

extern TNode* newTNode();

typedef struct GimMasterTree  {
	   char* MasterDir
	   TNode* MasterRoot;
	    char* MasterRecordDir;
} GimMasterTree;

extern GimMasterTree NewGimMasterTree();

typedef  struct Blob {
	    char* ShaHash;
	    int64_t ModTime;
	    char* ActualPath;
}Blob;


#endif
