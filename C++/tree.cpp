#include "tree.h"
#include <queue>

void TInit(CBTree* t) { t->data.clear(); }

void TINSERT(CBTree* t, int value) {
    t->data.push_back(value);
}

int TGET(CBTree* t, int index) {
    if (index < 0 || index >= (int)t->data.size()) return -1;
    return t->data[index];
}

int TSEARCH(CBTree* t, int value) {
    for (int i = 0; i < (int)t->data.size(); i++)
        if (t->data[i] == value) return i;
    return -1;
}

bool TCHECK_COMPLETE(CBTree* t) {
    int n = t->data.size();
    if (n == 0) return true;
    bool sawNull = false;
    for (int i = 0; i < n; i++) {
        int left = 2 * i + 1;
        int right = 2 * i + 2;
        if (left < n && sawNull) return false;
        if (left >= n) sawNull = true;
        if (right < n && sawNull) return false;
        if (right >= n) sawNull = true;
    }
    return true;
}

void TPRINT(CBTree* t) {
    for (int v : t->data) cout << v << " ";
    cout << endl;
}