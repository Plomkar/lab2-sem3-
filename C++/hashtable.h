#pragma once
// Задание 6.1. Хеш-таблица с открытой адресацией
#include <string>
#include <vector>
using namespace std;

enum HashMode { LINEAR_PROBING, DOUBLE_HASHING };

enum SlotState { EMPTY = 0, USED = 1, DELETED = 2 };

template <class V>
struct HTEntry {
    string key;
    V value;
    int state;
    HTEntry() : key(), value(), state(EMPTY) {}
};

template <class V>
struct HashTable {
    vector<HTEntry<V>> slots;
    int live;
    int filled;
    HashMode mode;
    bool autoResize;
};

inline bool HT_IsPrime(int n) {
    if (n < 2) return false;
    for (int i = 2; (long long)i * i <= n; i++)
        if (n % i == 0) return false;
    return true;
}

inline int HT_NextPrime(int n) {
    while (!HT_IsPrime(n)) n++;
    return n;
}

inline unsigned long long HT_Hash1(const string& s) {
    unsigned long long h = 0;
    for (unsigned char c : s) h = h * 131 + c;
    return h;
}

inline unsigned long long HT_Hash2(const string& s) {
    unsigned long long h = 1469598103934665603ULL;
    for (unsigned char c : s) { h ^= c; h *= 1099511628211ULL; }
    return h;
}

template <class V>
void HT_Init(HashTable<V>* t, HashMode mode = DOUBLE_HASHING, int capacity = 11, bool autoResize = true) {
    t->slots.assign(HT_NextPrime(capacity), HTEntry<V>());
    t->live = 0;
    t->filled = 0;
    t->mode = mode;
    t->autoResize = autoResize;
}

template <class V>
int HT_Size(const HashTable<V>* t) { return t->live; }

template <class V>
int HT_Probe(const HashTable<V>* t, const string& key, int i) {
    long long cap = (long long)t->slots.size();
    long long h1 = (long long)(HT_Hash1(key) % cap);
    long long step = 1;
    if (t->mode == DOUBLE_HASHING)
        step = 1 + (long long)(HT_Hash2(key) % (cap - 1));
    return (int)((h1 + (long long)i * step) % cap);
}

template <class V>
V* HT_Get(HashTable<V>* t, const string& key) {
    int cap = (int)t->slots.size();
    for (int i = 0; i < cap; i++) {
        HTEntry<V>& e = t->slots[HT_Probe(t, key, i)];
        if (e.state == EMPTY) return nullptr;
        if (e.state == USED && e.key == key) return &e.value;
    }
    return nullptr;
}

template <class V>
bool HT_Has(HashTable<V>* t, const string& key) { return HT_Get(t, key) != nullptr; }

template <class V>
void HT_Rehash(HashTable<V>* t, int newCap) {
    vector<HTEntry<V>> old;
    old.swap(t->slots);
    t->slots.assign(HT_NextPrime(newCap), HTEntry<V>());
    t->live = 0;
    t->filled = 0;
    HashMode m = t->mode;
    bool ar = t->autoResize;
    t->autoResize = false;
    for (auto& e : old)
        if (e.state == USED) {
            int cap = (int)t->slots.size();
            for (int i = 0; i < cap; i++) {
                HTEntry<V>& d = t->slots[HT_Probe(t, e.key, i)];
                if (d.state != USED) { d.key = e.key; d.value = e.value; d.state = USED; t->live++; t->filled++; break; }
            }
        }
    t->mode = m;
    t->autoResize = ar;
}

template <class V>
bool HT_Put(HashTable<V>* t, const string& key, const V& value) {
    if (t->autoResize && (t->filled + 1) * 10 > (int)t->slots.size() * 7)
        HT_Rehash(t, (int)t->slots.size() * 2);
    int cap = (int)t->slots.size();
    int firstFree = -1;
    for (int i = 0; i < cap; i++) {
        int idx = HT_Probe(t, key, i);
        HTEntry<V>& e = t->slots[idx];
        if (e.state == USED && e.key == key) { e.value = value; return true; }
        if (e.state == DELETED && firstFree == -1) firstFree = idx;
        if (e.state == EMPTY) { if (firstFree == -1) firstFree = idx; break; }
    }
    if (firstFree == -1) return false;
    HTEntry<V>& e = t->slots[firstFree];
    if (e.state == EMPTY) t->filled++;
    e.key = key; e.value = value; e.state = USED;
    t->live++;
    return true;
}

template <class V>
bool HT_Del(HashTable<V>* t, const string& key) {
    int cap = (int)t->slots.size();
    for (int i = 0; i < cap; i++) {
        HTEntry<V>& e = t->slots[HT_Probe(t, key, i)];
        if (e.state == EMPTY) return false;
        if (e.state == USED && e.key == key) { e.state = DELETED; t->live--; return true; }
    }
    return false;
}
