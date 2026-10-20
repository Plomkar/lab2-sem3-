#include "stack.h"

void SInit(Stack* s) { s->head = nullptr; }

bool SIsEmpty(Stack* s) { return s->head == nullptr; }

void SPUSH(Stack* s, int value) {
    SNode* newNode = new SNode{value, s->head};
    s->head = newNode;
}

int SPOP(Stack* s) {
    if (SIsEmpty(s)) { cout << "Stack is empty" << endl; return -1; }
    SNode* del = s->head;
    int val = del->data;
    s->head = del->next;
    delete del;
    return val;
}

int SPEAK(Stack* s) {
    if (SIsEmpty(s)) return -1;
    return s->head->data;
}

void SPRINT(Stack* s) {
    SNode* cur = s->head;
    while (cur != nullptr) {
        cout << cur->data << " ";
        cur = cur->next;
    }
    cout << endl;
}