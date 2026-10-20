#pragma once
#include <iostream>
using namespace std;

struct SNode {
    int data;
    SNode* next;
};

struct Stack {
    SNode* head;
};

void SInit(Stack* s);
bool SIsEmpty(Stack* s);
void SPUSH(Stack* s, int value);
int  SPOP(Stack* s);
int  SPEAK(Stack* s);
void SPRINT(Stack* s);