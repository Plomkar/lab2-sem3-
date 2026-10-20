#include <iostream>
#include <string>
#include "list.h"
#include "hashtable.h"
using namespace std;

struct CacheItem {
    int value;
    LNode* node;
    CacheItem() : value(0), node(nullptr) {}
};

static void unlink(LList* l, LNode* n) {
    if (n->prev) n->prev->next = n->next; else l->head = n->next;
    if (n->next) n->next->prev = n->prev; else l->tail = n->prev;
    n->prev = n->next = nullptr;
}

static void pushFront(LList* l, LNode* n) {
    n->prev = nullptr;
    n->next = l->head;
    if (l->head) l->head->prev = n; else l->tail = n;
    l->head = n;
}

class LRUCache {
    int cap, size;
    LList order;
    HashTable<CacheItem> table;
public:
    explicit LRUCache(int capacity) : cap(capacity), size(0) {
        LInit(&order);
        HT_Init(&table, DOUBLE_HASHING);
    }
    ~LRUCache() {
        while (order.head) LDEL(&order, 0);
    }

    int get(int key) {
        CacheItem* it = HT_Get(&table, to_string(key));
        if (!it) return -1;
        unlink(&order, it->node);
        pushFront(&order, it->node);
        return it->value;
    }

    void set(int key, int value) {
        string k = to_string(key);
        CacheItem* it = HT_Get(&table, k);
        if (it) {
            it->value = value;
            unlink(&order, it->node);
            pushFront(&order, it->node);
            return;
        }
        if (cap <= 0) return;
        if (size == cap) {
            LNode* old = order.tail;
            HT_Del(&table, to_string(old->data));
            unlink(&order, old);
            delete old;
            size--;
        }
        LPUSH(&order, key);
        CacheItem item;
        item.value = value;
        item.node = order.head;
        HT_Put(&table, k, item);
        size++;
    }
};

int main() {
    int cap, q;
    if (!(cin >> cap >> q)) { cout << "Ошибка ввода: ожидалось  cap Q" << endl; return 1; }
    LRUCache cache(cap);
    bool first = true;
    for (int i = 0; i < q; i++) {
        string cmd;
        if (!(cin >> cmd)) break;
        if (cmd == "SET") {
            int x, y;
            cin >> x >> y;
            cache.set(x, y);
        } else if (cmd == "GET") {
            int x;
            cin >> x;
            cout << (first ? "" : " ") << cache.get(x);
            first = false;
        } else {
            cout << endl << "Неизвестный запрос: " << cmd << endl;
            return 1;
        }
    }
    cout << endl;
    return 0;
}
