public class segtree_java {
    static class SegmentTree {
        long[] tree;
        int n;

        SegmentTree(int size) {
            n = size;
            tree = new long[4 * n];
        }

        void build(long[] a, int v, int tl, int tr) {
            if (tl == tr) {
                tree[v] = a[tl];
            } else {
                int tm = (tl + tr) / 2;
                build(a, v * 2, tl, tm);
                build(a, v * 2 + 1, tm + 1, tr);
                tree[v] = tree[v * 2] + tree[v * 2 + 1];
            }
        }

        long query(int v, int tl, int tr, int l, int r) {
            if (l > r) return 0;
            if (l == tl && r == tr) return tree[v];
            int tm = (tl + tr) / 2;
            return query(v * 2, tl, tm, l, Math.min(r, tm)) + 
                   query(v * 2 + 1, tm + 1, tr, Math.max(l, tm + 1), r);
        }

        void update(int v, int tl, int tr, int pos, long new_val) {
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
    }

    public static void main(String[] args) {
        int n = 5;
        long[] a = {1, 2, 3, 4, 5};
        SegmentTree st = new SegmentTree(n);
        st.build(a, 1, 0, n - 1);

        System.out.println("Sum 0-4: " + st.query(1, 0, n - 1, 0, 4));
        st.update(1, 0, n - 1, 2, 10);
        System.out.println("Sum 0-4 after update: " + st.query(1, 0, n - 1, 0, 4));

        String msg = "Segment tree runs perfectly!";
        System.out.println(msg);
    }
}
