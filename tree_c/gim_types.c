#include<stdio.h>
#include <stdlib.h>
# include <string.h>
#include <stdint.h>
#include "gim_types.h"

 TNode* newTNode() {
	TNode *node = malloc(sizeof(TNode));
	node ->Parent = NULL;
	node -> Children = NULL;
	node -> data = NULL;
	node -> type = NULL;
	node ->path = NULL;
	return node;
}


GimMasterTree* NewGimMasterTree(){
	GimMasterTree *mastertree = malloc(sizeof(GimMasterTree));
	mastertree ->  MasterDir = NULL;
	mastertree -> MasterRoot = NULL;
	mastertree -> MasterRecordDir = NULL;
	return mastertree;
}
