#pragma once
#include <string>
using namespace std;

struct FNode {
    int data;
    FNode* next;
};

struct FList {
    FNode* head;
};

void FInit(FList* f);
void FPUSH(FList* f, int value);
void FPUSH(FList* f, int index, int value);
void FPUSH_AFTER(FList* f, int index, int value);
void FPUSH_END(FList* f, int value);
void FDEL(FList* f, int index);
void FDEL_VALUE(FList* f, int value);
int  FGET(FList* f, int index);
int  FSEARCH(FList* f, int value);
void FPRINT(FList* f);