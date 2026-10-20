#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include "set.h"
using namespace std;

static Set g_set;

static void loadFromFile(const string& filename) {
    ifstream in(filename);
    if (!in.is_open()) return;
    string line;
    while (getline(in, line)) {
        istringstream iss(line);
        string tag;
        iss >> tag;
        if (tag != "SET") continue;
        int v;
        while (iss >> v) SetAdd(&g_set, v);
    }
}

static void saveToFile(const string& filename) {
    ofstream out(filename);
    if (!out.is_open()) return;
    Array arr;
    SetToArray(&g_set, &arr);
    out << "SET";
    for (int i = 0; i < MLEN(&arr); i++) out << " " << MGET(&arr, i);
    out << "\n";
}

static void processQuery(const string& query) {
    istringstream iss(query);
    string cmd;
    iss >> cmd;
    int x;
    if (cmd == "SETADD" && (iss >> x))
        cout << (SetAdd(&g_set, x) ? "TRUE" : "FALSE") << endl;
    else if (cmd == "SETDEL" && (iss >> x))
        cout << (SetDel(&g_set, x) ? "TRUE" : "FALSE") << endl;
    else if (cmd == "SET_AT" && (iss >> x))
        cout << (SetAt(&g_set, x) ? "TRUE" : "FALSE") << endl;
    else if (cmd == "PRINT") {
        Array arr;
        SetToArray(&g_set, &arr);
        for (int i = 0; i < MLEN(&arr); i++) cout << MGET(&arr, i) << " ";
        cout << endl;
    } else
        cout << "Unknown command: " << query << endl;
}

int main(int argc, char* argv[]) {
    string filename, query;
    for (int i = 1; i < argc; i++) {
        string arg = argv[i];
        if (arg == "--file" && i + 1 < argc) filename = argv[++i];
        else if (arg == "--query" && i + 1 < argc) query = argv[++i];
    }
    SetInit(&g_set);
    if (!filename.empty()) loadFromFile(filename);
    if (query.empty()) {
        cout << "Usage: ./task2_set --file <file.data> --query '<SETADD|SETDEL|SET_AT|PRINT> [x]'" << endl;
        return 0;
    }
    processQuery(query);
    if (!filename.empty()) saveToFile(filename);
    SetClear(&g_set);
    return 0;
}
