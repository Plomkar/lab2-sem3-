#include <iostream>
#include <algorithm>
#include <string>
#include <sstream>
#include "set.h"
using namespace std;

static Array g_vals;
static Array g_group;
static Set   g_used;
static int   g_n, g_S, g_k;

static bool dfs(int start, int cur, int idx) {
    if (idx == g_k) return true;
    if (cur == g_S) return dfs(0, 0, idx + 1);
    for (int i = start; i < g_n; i++) {
        if (SetAt(&g_used, i)) continue;
        int v = MGET(&g_vals, i);
        if (cur + v > g_S) continue;
        SetAdd(&g_used, i);
        MSET(&g_group, i, idx);
        if (dfs(i + 1, cur + v, idx)) return true;
        SetDel(&g_used, i);
        if (cur == 0) break;
    }
    return false;
}

static bool tryPartition(long long total, int S, Array* vals, bool print) {
    int maxV = MGET(vals, 0);
    if (S < maxV || total % S != 0) return false;
    g_k = (int)(total / S);
    g_S = S;
    if (g_k < 2) return false;
    SetInit(&g_used);
    MInit(&g_group);
    for (int i = 0; i < g_n; i++) MPUSH(&g_group, -1);
    bool ok = dfs(0, 0, 0);
    if (ok && print) {
        cout << "Сумма каждого подмножества = " << S << ", подмножеств: " << g_k << endl;
        for (int j = 0; j < g_k; j++) {
            cout << "{";
            bool first = true;
            for (int i = 0; i < g_n; i++)
                if (MGET(&g_group, i) == j) { cout << (first ? "" : ", ") << MGET(&g_vals, i); first = false; }
            cout << "}" << endl;
        }
    }
    SetClear(&g_used);
    return ok;
}

int main() {
    int n;
    if (!(cin >> n) || n <= 0) { cout << "Ошибка: некорректный ввод" << endl; return 1; }

    Set s;
    SetInit(&s);
    for (int i = 0; i < n; i++) {
        long long x;
        if (!(cin >> x) || x <= 0) { cout << "Ошибка: нужны натуральные числа" << endl; return 1; }
        SetAdd(&s, (int)x);
    }
    int S = -1;
    string rest;
    getline(cin, rest);
    istringstream iss(rest);
    int tmp;
    if (iss >> tmp) {
        if (tmp <= 0) { cout << "Ошибка: сумма S должна быть натуральным числом" << endl; return 1; }
        S = tmp;
    }

    SetToArray(&s, &g_vals);
    sort(g_vals.data.begin(), g_vals.data.end(), greater<int>());
    g_n = MLEN(&g_vals);
    long long total = 0;
    for (int i = 0; i < g_n; i++) total += MGET(&g_vals, i);

    bool any = false;
    if (S > 0) {
        any = tryPartition(total, S, &g_vals, true);
    } else {
        int maxV = MGET(&g_vals, 0);
        for (long long cand = maxV; cand <= total / 2; cand++)
            if (total % cand == 0)
                any |= tryPartition(total, (int)cand, &g_vals, true);
    }
    if (!any) cout << "Ошибка: множество нельзя разбить на подмножества с равной суммой" << endl;
    SetClear(&s);
    return 0;
}
