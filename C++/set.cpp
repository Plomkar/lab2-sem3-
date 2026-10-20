#include "set.h"

static int bucketOf(const Set* s, int value) {
    long long cap = (long long)s->buckets.size();
    long long r = value % cap;
    if (r < 0) r += cap;
    return (int)r;
}

void SetInit(Set* s, int capacity) {
    if (capacity < 4) capacity = 4;
    s->buckets.assign(capacity, FList());
    for (auto& b : s->buckets) FInit(&b);
    s->size = 0;
}

void SetClear(Set* s) {
    for (auto& b : s->buckets)
        while (b.head != nullptr) FDEL(&b, 0);
    s->size = 0;
}

static void rehash(Set* s) {
    vector<FList> old;
    old.swap(s->buckets);
    int newCap = (int)old.size() * 2;
    s->buckets.assign(newCap, FList());
    for (auto& b : s->buckets) FInit(&b);
    for (auto& b : old) {
        FNode* cur = b.head;
        while (cur != nullptr) {
            FNode* next = cur->next;
            int idx = bucketOf(s, cur->data);
            cur->next = s->buckets[idx].head;
            s->buckets[idx].head = cur;
            cur = next;
        }
    }
}

bool SetAt(Set* s, int value) {
    return FSEARCH(&s->buckets[bucketOf(s, value)], value) != -1;
}

bool SetAdd(Set* s, int value) {
    if (SetAt(s, value)) return false;
    if (s->size + 1 > (int)s->buckets.size()) rehash(s);
    FPUSH(&s->buckets[bucketOf(s, value)], value);
    s->size++;
    return true;
}

bool SetDel(Set* s, int value) {
    FList* b = &s->buckets[bucketOf(s, value)];
    if (FSEARCH(b, value) == -1) return false;
    FDEL_VALUE(b, value);
    s->size--;
    return true;
}

int SetSize(Set* s) { return s->size; }

void SetToArray(Set* s, Array* out) {
    MInit(out);
    for (auto& b : s->buckets)
        for (FNode* cur = b.head; cur != nullptr; cur = cur->next)
            MPUSH(out, cur->data);
}
