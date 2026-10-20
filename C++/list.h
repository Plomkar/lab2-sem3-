#pragma once
#include <iostream>
using namespace std;

struct LNode {
    int data;
    LNode* prev;
    LNode* next;
};

struct LList {
    LNode* head;
    LNode* tail;
};

void LInit(LList* l);
void LPUSH(LList* l, int value);
void LPUSH_END(LList* l, int value);
void LPUSH(LList* l, int index, int value);
void LPUSH_AFTER(LList* l, int index, int value);
void LDEL(LList* l, int index);
void LDEL_VALUE(LList* l, int value);
int  LGET(LList* l, int index);
int  LSEARCH(LList* l, int value);
void LPRINT(LList* l);
void LPRINT_REVERSE(LList* l);