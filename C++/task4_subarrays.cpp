#include <iostream>
#include <sstream>
#include <string>
#include <vector>
#include "array.h"
#include "set.h"
#include "hashtable.h"
using namespace std;

static Array          g_ids;
static vector<string> g_names;
static vector<string> g_out;

static void emit(Array* cur) {
    string s = "{";
    for (int i = 0; i < MLEN(cur); i++) {
        if (i) s += ", ";
        s += g_names[MGET(cur, i)];
    }
    s += "}";
    g_out.push_back(s);
}

static void gen(int start, int need, Array* cur) {
    if (need == 0) { emit(cur); return; }
    int n = MLEN(&g_ids);
    Set tried;
    SetInit(&tried, 8);
    for (int i = start; i + need <= n; i++) {
        int id = MGET(&g_ids, i);
        if (SetAt(&tried, id)) continue;
        SetAdd(&tried, id);
        MPUSH(cur, id);
        gen(i + 1, need - 1, cur);
        MDEL(cur, MLEN(cur) - 1);
    }
    SetClear(&tried);
}

int main() {
    string line;
    getline(cin, line);
    istringstream iss(line);
    string tok;

    HashTable<int> dict;
    HT_Init(&dict, DOUBLE_HASHING);
    MInit(&g_ids);
    while (iss >> tok) {
        int* p = HT_Get(&dict, tok);
        int id;
        if (p) id = *p;
        else { id = (int)g_names.size(); g_names.push_back(tok); HT_Put(&dict, tok, id); }
        MPUSH(&g_ids, id);
    }
    int n = MLEN(&g_ids);
    if (n == 0) { cout << "[{}]" << endl; return 0; }
    if (n > 20) { cout << "Слишком длинный массив (результат содержит до 2^n подмассивов), максимум 20 элементов" << endl; return 1; }

    Array cur;
    MInit(&cur);
    for (int len = 0; len <= n; len++) gen(0, len, &cur);

    cout << "[";
    for (size_t i = 0; i < g_out.size(); i++) cout << (i ? ", " : "") << g_out[i];
    cout << "]" << endl;
    return 0;
}
