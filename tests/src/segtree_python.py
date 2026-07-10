class SegmentTree:
    def __init__(self, size):
        self.n = size
        self.tree = [0] * (4 * size)

    def build(self, a, v, tl, tr):
        if tl == tr:
            self.tree[v] = a[tl]
        else:
            tm = (tl + tr) // 2
            self.build(a, v * 2, tl, tm)
            self.build(a, v * 2 + 1, tm + 1, tr)
            self.tree[v] = self.tree[v * 2] + self.tree[v * 2 + 1]

    def query(self, v, tl, tr, l, r):
        if l > r:
            return 0
        if l == tl and r == tr:
            return self.tree[v]
        tm = (tl + tr) // 2
        return self.query(v * 2, tl, tm, l, min(r, tm)) + \
               self.query(v * 2 + 1, tm + 1, tr, max(l, tm + 1), r)

    def update(self, v, tl, tr, pos, new_val):
        if tl == tr:
            self.tree[v] = new_val
        else:
            tm = (tl + tr) // 2
            if pos <= tm:
                self.update(v * 2, tl, tm, pos, new_val)
            else:
                self.update(v * 2 + 1, tm + 1, tr, pos, new_val)
            self.tree[v] = self.tree[v * 2] + self.tree[v * 2 + 1]

def main():
    n = 5
    a = [1, 2, 3, 4, 5]
    st = SegmentTree(n)
    st.build(a, 1, 0, n - 1)

    print(f"Sum 0-4: {st.query(1, 0, n - 1, 0, 4)}")
    st.update(1, 0, n - 1, 2, 10)
    print(f"Sum 0-4 after update: {st.query(1, 0, n - 1, 0, 4)}")

    msg = "Segment tree runs perfectly!"
    print(msg)

if __name__ == "__main__":
    main()
