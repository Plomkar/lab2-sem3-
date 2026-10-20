#include "forward_list.h"
#include <iostream>

void FInit(FList* f) { f->head = nullptr; }

void FPUSH(FList* f, int value) {
    FNode* newNode = new FNode{value, f->head};
    f->head = newNode;
}

void FPUSH_END(FList* f, int value) {
    FNode* newNode = new FNode{value, nullptr};
    if (f->head == nullptr) { f->head = newNode; return; }
    FNode* cur = f->head;
    while (cur->next != nullptr) cur = cur->next;
    cur->next = newNode;
}

void FPUSH(FList* f, int index, int value) {
    if (index <= 0) { FPUSH(f, value); return; }
    FNode* cur = f->head;
    for (int i = 0; i < index - 1 && cur != nullptr; i++) cur = cur->next;
    if (cur == nullptr) return;
    FNode* newNode = new FNode{value, cur->next};
    cur->next = newNode;
}

void FPUSH_AFTER(FList* f, int index, int value) {
    FNode* cur = f->head;
    for (int i = 0; i < index && cur != nullptr; i++) cur = cur->next;
    if (cur == nullptr) return;
    FNode* newNode = new FNode{value, cur->next};
    cur->next = newNode;
}

void FDEL(FList* f, int index) {
    if (f->head == nullptr) return;
    if (index == 0) {
        FNode* del = f->head;
        f->head = f->head->next;
        delete del;
        return;
    }
    FNode* cur = f->head;
    for (int i = 0; i < index - 1 && cur->next != nullptr; i++) cur = cur->next;
    if (cur->next == nullptr) return;
    FNode* del = cur->next;
    cur->next = del->next;
    delete del;
}

void FDEL_VALUE(FList* f, int value) {
    if (f->head == nullptr) return;
    if (f->head->data == value) {
        FNode* del = f->head;
        f->head = f->head->next;
        delete del;
        return;
    }
    FNode* cur = f->head;
    while (cur->next != nullptr && cur->next->data != value) cur = cur->next;
    if (cur->next == nullptr) return;
    FNode* del = cur->next;
    cur->next = del->next;
    delete del;
}

int FGET(FList* f, int index) {
    FNode* cur = f->head;
    for (int i = 0; i < index && cur != nullptr; i++) cur = cur->next;
    return cur ? cur->data : -1;
}

int FSEARCH(FList* f, int value) {
    FNode* cur = f->head;
    int idx = 0;
    while (cur != nullptr) {
        if (cur->data == value) return idx;
        cur = cur->next;
        idx++;
    }
    return -1;
}

void FPRINT(FList* f) {
    FNode* cur = f->head;
    while (cur != nullptr) {
        cout << cur->data << " ";
        cur = cur->next;
    }
    cout << endl;
}