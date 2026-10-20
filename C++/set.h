#pragma once
// Задание 2. АТД «Множество».
#include <vector>
#include "forward_list.h"
#include "array.h"
using namespace std;

struct Set {
    vector<FList> buckets;
    int size;
};

void SetInit(Set* s, int capacity = 16);
void SetClear(Set* s);
bool SetAdd(Set* s, int value);
bool SetDel(Set* s, int value);
bool SetAt(Set* s, int value);
int  SetSize(Set* s);
void SetToArray(Set* s, Array* out);