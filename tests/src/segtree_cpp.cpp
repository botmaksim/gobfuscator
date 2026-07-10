#include <iostream>
#include <vector>
#include <string>

using namespace std;

class SegmentTree {
    vector<long long> tree;
    int n;

public:
    SegmentTree(int size) {
        n = size;
        tree.assign(4 * n, 0);
    }

    void build(const vector<long long>& a, int v, int tl, int tr) {
        if (tl == tr) {
            tree[v] = a[tl];
        } else {
            int tm = (tl + tr) / 2;
            build(a, v * 2, tl, tm);
            build(a, v * 2 + 1, tm + 1, tr);
            tree[v] = tree[v * 2] + tree[v * 2 + 1];
        }
    }

    long long query(int v, int tl, int tr, int l, int r) {
        if (l > r) return 0;
        if (l == tl && r == tr) return tree[v];
        int tm = (tl + tr) / 2;
        return query(v * 2, tl, tm, l, min(r, tm)) + 
               query(v * 2 + 1, tm + 1, tr, max(l, tm + 1), r);
    }

    void update(int v, int tl, int tr, int pos, long long new_val) {
        if (tl == tr) {
            tree[v] = new_val;
        } else {
            int tm = (tl + tr) / 2;
            if (pos <= tm)
                update(v * 2, tl, tm, pos, new_val);
            else
                update(v * 2 + 1, tm + 1, tr, pos, new_val);
            tree[v] = tree[v * 2] + tree[v * 2 + 1];
        }
    }
};

int main() {
    int n = 5;
    vector<long long> a = {1, 2, 3, 4, 5};
    SegmentTree st(n);
    st.build(a, 1, 0, n - 1);

    cout << "Sum 0-4: " << st.query(1, 0, n - 1, 0, 4) << endl;
    st.update(1, 0, n - 1, 2, 10);
    cout << "Sum 0-4 after update: " << st.query(1, 0, n - 1, 0, 4) << endl;

    string msg = "Segment tree runs perfectly!";
    cout << msg << endl;

    return 0;
}
