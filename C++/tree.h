#pragma once
#include <vector>
#include <iostream>
using namespace std;

struct CBTree {
    vector<int> data;
};

void TInit(CBTree* t);
void TINSERT(CBTree* t, int value);
bool TCHECK_COMPLETE(CBTree* t);
int  TGET(CBTree* t, int index);
int  TSEARCH(CBTree* t, int value);
void TPRINT(CBTree* t);