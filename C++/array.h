#pragma once
#include <vector>
#include <string>
using namespace std;

struct Array {
    vector<int> data;
};

void MInit(Array* a);
void MPUSH(Array* a, int value);
void MPUSH(Array* a, int index, int value);
int  MGET(Array* a, int index);
void MDEL(Array* a, int index);
void MSET(Array* a, int index, int value);
int  MLEN(Array* a);