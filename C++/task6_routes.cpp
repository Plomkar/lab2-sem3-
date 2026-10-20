#include <iostream>
#include <sstream>
#include <string>
#include <vector>
#include "hashtable.h"
#include "forward_list.h"
#include "queue.h"
#include "stack.h"
#include "array.h"
using namespace std;

static const int INF = 1000000000;

struct Network {
    HashTable<int> ids;
    HashTable<int> delay;
    vector<string> names;
    vector<FList>  adj;
};

static Network net;

static int getId(const string& name, bool create) {
    int* p = HT_Get(&net.ids, name);
    if (p) return *p;
    if (!create) return -1;
    int id = (int)net.names.size();
    net.names.push_back(name);
    FList f;
    FInit(&f);
    net.adj.push_back(f);
    HT_Put(&net.ids, name, id);
    return id;
}

static void addRoute(const string& a, const string& b, int d) {
    int u = getId(a, true), v = getId(b, true);
    bool exists = HT_Has(&net.delay, a + "|" + b);
    HT_Put(&net.delay, a + "|" + b, d);
    HT_Put(&net.delay, b + "|" + a, d);
    if (!exists && u != v) {
        FPUSH(&net.adj[u], v);
        FPUSH(&net.adj[v], u);
    }
}

static void findPath(const string& from, const string& to) {
    int s = getId(from, false), t = getId(to, false);
    if (s == -1 || t == -1) { cout << "Сервер не найден" << endl; return; }

    int n = (int)net.names.size();
    Array dist, parent, inQ;
    MInit(&dist); MInit(&parent); MInit(&inQ);
    for (int i = 0; i < n; i++) { MPUSH(&dist, INF); MPUSH(&parent, -1); MPUSH(&inQ, 0); }

    Queue q;
    QInit(&q);
    MSET(&dist, s, 0);
    QPUSH(&q, s);
    MSET(&inQ, s, 1);
    while (!QIsEmpty(&q)) {
        int u = QPOP(&q);
        MSET(&inQ, u, 0);
        for (FNode* e = net.adj[u].head; e != nullptr; e = e->next) {
            int v = e->data;
            int w = *HT_Get(&net.delay, net.names[u] + "|" + net.names[v]);
            if (MGET(&dist, u) + w < MGET(&dist, v)) {
                MSET(&dist, v, MGET(&dist, u) + w);
                MSET(&parent, v, u);
                if (!MGET(&inQ, v)) { QPUSH(&q, v); MSET(&inQ, v, 1); }
            }
        }
    }

    if (MGET(&dist, t) == INF) { cout << "Маршрута между серверами нет" << endl; return; }

    Stack path;
    SInit(&path);
    for (int v = t; v != -1; v = MGET(&parent, v)) SPUSH(&path, v);
    cout << "минимальная задержка = " << MGET(&dist, t) << " (маршрут ";
    bool first = true;
    while (!SIsEmpty(&path)) {
        cout << (first ? "" : " -> ") << "\"" << net.names[SPOP(&path)] << "\"";
        first = false;
    }
    cout << ")" << endl;
}

int main() {
    HT_Init(&net.ids, DOUBLE_HASHING);
    HT_Init(&net.delay, DOUBLE_HASHING);

    string line;
    while (getline(cin, line)) {
        istringstream iss(line);
        string cmd;
        if (!(iss >> cmd)) continue;
        if (cmd == "ADD" || cmd == "UPDATE") {
            string a, b; int d;
            if (iss >> a >> b >> d && d >= 0) { addRoute(a, b, d); cout << "OK" << endl; }
            else cout << "Ошибка: ожидалось  " << cmd << " A B d  (d >= 0)" << endl;
        } else if (cmd == "PATH") {
            string a, b;
            if (iss >> a >> b) findPath(a, b);
            else cout << "Ошибка: ожидалось  PATH A B" << endl;
        } else if (cmd == "EXIT") {
            break;
        } else {
            cout << "Unknown command: " << cmd << endl;
        }
    }
    return 0;
}
