#include <iostream>
#include <string>
#include <cctype>
#include "stack.h"
using namespace std;

static int prec(char c) {
    if (c == '^') return 3;
    if (c == '*' || c == '/') return 2;
    if (c == '+' || c == '-') return 1;
    return 0;
}

static bool isOp(char c) { return c == '+' || c == '-' || c == '*' || c == '/' || c == '^'; }

bool infixToPostfix(const string& s, string& result) {
    Stack st;
    SInit(&st);
    result.clear();
    bool expectOperand = true;
    bool ok = true;

    for (size_t i = 0; i < s.size() && ok; ) {
        char c = s[i];
        if (isspace((unsigned char)c)) { i++; continue; }

        if (isalnum((unsigned char)c) || c == '.') {
            if (!expectOperand) { ok = false; break; }
            string tok;
            while (i < s.size() && (isalnum((unsigned char)s[i]) || s[i] == '.')) tok += s[i++];
            result += tok + " ";
            expectOperand = false;
        } else if (c == '(') {
            if (!expectOperand) { ok = false; break; }
            SPUSH(&st, c);
            i++;
        } else if (c == ')') {
            if (expectOperand) { ok = false; break; }
            while (!SIsEmpty(&st) && SPEAK(&st) != '(') result += string(1, (char)SPOP(&st)) + " ";
            if (SIsEmpty(&st)) { ok = false; break; }
            SPOP(&st);
            i++;
        } else if (isOp(c)) {
            if (expectOperand) { ok = false; break; }
            bool right = (c == '^');
            while (!SIsEmpty(&st) && SPEAK(&st) != '(' &&
                   (prec((char)SPEAK(&st)) > prec(c) || (prec((char)SPEAK(&st)) == prec(c) && !right)))
                result += string(1, (char)SPOP(&st)) + " ";
            SPUSH(&st, c);
            expectOperand = true;
            i++;
        } else {
            ok = false;
        }
    }

    if (ok && expectOperand) ok = false;
    while (ok && !SIsEmpty(&st)) {
        char c = (char)SPOP(&st);
        if (c == '(') { ok = false; break; }
        result += string(1, c) + " ";
    }
    while (!SIsEmpty(&st)) SPOP(&st);
    if (ok && !result.empty()) result.pop_back();
    return ok;
}

static void handle(const string& line) {
    string res;
    if (infixToPostfix(line, res)) cout << res << endl;
    else cout << "Ошибка в выражении" << endl;
}

int main(int argc, char* argv[]) {
    if (argc > 1) {
        string expr;
        for (int i = 1; i < argc; i++) expr += string(argv[i]) + " ";
        handle(expr);
    } else {
        string line;
        while (getline(cin, line)) if (!line.empty()) handle(line);
    }
    return 0;
}
