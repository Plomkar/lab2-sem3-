#include "list.h"

void LInit(LList* l) { l->head = nullptr; l->tail = nullptr; }

void LPUSH(LList* l, int value) {
    LNode* newNode = new LNode{value, nullptr, l->head};
    if (l->head != nullptr) l->head->prev = newNode;
    else l->tail = newNode;
    l->head = newNode;
}

void LPUSH_END(LList* l, int value) {
    LNode* newNode = new LNode{value, l->tail, nullptr};
    if (l->tail != nullptr) l->tail->next = newNode;
    else l->head = newNode;
    l->tail = newNode;
}

void LPUSH(LList* l, int index, int value) {
    if (index <= 0) { LPUSH(l, value); return; }
    LNode* cur = l->head;
    for (int i = 0; i < index - 1 && cur != nullptr; i++) cur = cur->next;
    if (cur == nullptr) { LPUSH_END(l, value); return; }
    LNode* newNode = new LNode{value, cur, cur->next};
    if (cur->next != nullptr) cur->next->prev = newNode;
    else l->tail = newNode;
    cur->next = newNode;
}

void LPUSH_AFTER(LList* l, int index, int value) {
    LNode* cur = l->head;
    for (int i = 0; i < index && cur != nullptr; i++) cur = cur->next;
    if (cur == nullptr) return;
    LNode* newNode = new LNode{value, cur, cur->next};
    if (cur->next != nullptr) cur->next->prev = newNode;
    else l->tail = newNode;
    cur->next = newNode;
}

void LDEL(LList* l, int index) {
    LNode* cur = l->head;
    for (int i = 0; i < index && cur != nullptr; i++) cur = cur->next;
    if (cur == nullptr) return;
    if (cur->prev != nullptr) cur->prev->next = cur->next;
    else l->head = cur->next;
    if (cur->next != nullptr) cur->next->prev = cur->prev;
    else l->tail = cur->prev;
    delete cur;
}

void LDEL_VALUE(LList* l, int value) {
    LNode* cur = l->head;
    while (cur != nullptr && cur->data != value) cur = cur->next;
    if (cur == nullptr) return;
    if (cur->prev != nullptr) cur->prev->next = cur->next;
    else l->head = cur->next;
    if (cur->next != nullptr) cur->next->prev = cur->prev;
    else l->tail = cur->prev;
    delete cur;
}

int LGET(LList* l, int index) {
    LNode* cur = l->head;
    for (int i = 0; i < index && cur != nullptr; i++) cur = cur->next;
    return cur ? cur->data : -1;
}

int LSEARCH(LList* l, int value) {
    LNode* cur = l->head;
    int idx = 0;
    while (cur != nullptr) {
        if (cur->data == value) return idx;
        cur = cur->next;
        idx++;
    }
    return -1;
}

void LPRINT(LList* l) {
    LNode* cur = l->head;
    while (cur != nullptr) {
        cout << cur->data << " ";
        cur = cur->next;
    }
    cout << endl;
}

void LPRINT_REVERSE(LList* l) {
    LNode* cur = l->tail;
    while (cur != nullptr) {
        cout << cur->data << " ";
        cur = cur->prev;
    }
    cout << endl;
}