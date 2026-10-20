#pragma once
#include <iostream>
using namespace std;

struct QNode {
    int data;
    QNode* next;
};

struct Queue {
    QNode* head;
    QNode* tail;
};

void QInit(Queue* q);
bool QIsEmpty(Queue* q);
void QPUSH(Queue* q, int value);
int  QPOP(Queue* q);
int  QPEEK(Queue* q);
void QPRINT(Queue* q);