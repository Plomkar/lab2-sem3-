#include "array.h"

void MInit(Array* a) { a->data.clear(); }

void MPUSH(Array* a, int value) { a->data.push_back(value); }

void MPUSH(Array* a, int index, int value) {
    if (index < 0 || index > (int)a->data.size()) return;
    a->data.insert(a->data.begin() + index, value);
}

int MGET(Array* a, int index) {
    if (index < 0 || index >= (int)a->data.size()) return -1;
    return a->data[index];
}

void MDEL(Array* a, int index) {
    if (index < 0 || index >= (int)a->data.size()) return;
    a->data.erase(a->data.begin() + index);
}

void MSET(Array* a, int index, int value) {
    if (index < 0 || index >= (int)a->data.size()) return;
    a->data[index] = value;
}

int MLEN(Array* a) { return (int)a->data.size(); }