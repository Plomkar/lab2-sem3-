#include <iostream>
#include "tree.h"
#include "stack.h"
#include "array.h"
using namespace std;

static void inorder(CBTree* t, Array* order) {
    MInit(order);
    int n = (int)t->data.size();
    Stack st;
    SInit(&st);
    int cur = 0;
    while (cur < n || !SIsEmpty(&st)) {
        while (cur < n) { SPUSH(&st, cur); cur = 2 * cur + 1; }
        cur = SPOP(&st);
        MPUSH(order, cur);
        cur = 2 * cur + 2;
    }
}

static bool isSorted(CBTree* t, Array* order) {
    for (int i = 1; i < MLEN(order); i++)
        if (TGET(t, MGET(order, i - 1)) >= TGET(t, MGET(order, i))) return false;
    return true;
}

int main() {
    int n;
    if (!(cin >> n) || n <= 0) { cout << "Ошибка: некорректный ввод" << endl; return 1; }
    CBTree t;
    TInit(&t);
    for (int i = 0; i < n; i++) {
        int v;
        if (!(cin >> v)) { cout << "Ошибка: ожидалось " << n << " чисел" << endl; return 1; }
        TINSERT(&t, v);
    }

    Array order;
    inorder(&t, &order);

    int first = -1, second = -1;
    for (int i = 1; i < MLEN(&order); i++) {
        int a = MGET(&order, i - 1), b = MGET(&order, i);
        if (TGET(&t, a) > TGET(&t, b)) {
            if (first == -1) first = a;
            second = b;
        }
    }

    if (first == -1) {
        cout << "Дерево уже является корректным BST:" << endl;
        TPRINT(&t);
        return 0;
    }

    swap(t.data[first], t.data[second]);

    if (!isSorted(&t, &order)) {
        cout << "Ошибка: дерево нельзя исправить обменом двух узлов" << endl;
        return 1;
    }
    cout << "Поменяны местами узлы с индексами " << first << " и " << second << endl;
    cout << "Исправленное дерево (по уровням): ";
    TPRINT(&t);
    return 0;
}
