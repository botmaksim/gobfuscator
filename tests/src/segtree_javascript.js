class SegmentTree {
    constructor(size) {
        this.n = size;
        this.tree = new Array(4 * size).fill(0);
    }

    build(a, v, tl, tr) {
        if (tl === tr) {
            this.tree[v] = a[tl];
        } else {
            let tm = Math.floor((tl + tr) / 2);
            this.build(a, v * 2, tl, tm);
            this.build(a, v * 2 + 1, tm + 1, tr);
            this.tree[v] = this.tree[v * 2] + this.tree[v * 2 + 1];
        }
    }

    query(v, tl, tr, l, r) {
        if (l > r) return 0;
        if (l === tl && r === tr) return this.tree[v];
        let tm = Math.floor((tl + tr) / 2);
        return this.query(v * 2, tl, tm, l, Math.min(r, tm)) + 
               this.query(v * 2 + 1, tm + 1, tr, Math.max(l, tm + 1), r);
    }

    update(v, tl, tr, pos, new_val) {
        if (tl === tr) {
            this.tree[v] = new_val;
        } else {
            let tm = Math.floor((tl + tr) / 2);
            if (pos <= tm)
                this.update(v * 2, tl, tm, pos, new_val);
            else
                this.update(v * 2 + 1, tm + 1, tr, pos, new_val);
            this.tree[v] = this.tree[v * 2] + this.tree[v * 2 + 1];
        }
    }
}

function main() {
    let n = 5;
    let a = [1, 2, 3, 4, 5];
    let st = new SegmentTree(n);
    st.build(a, 1, 0, n - 1);

    console.log("Sum 0-4: " + st.query(1, 0, n - 1, 0, 4));
    st.update(1, 0, n - 1, 2, 10);
    console.log("Sum 0-4 after update: " + st.query(1, 0, n - 1, 0, 4));

    let msg = "Segment tree runs perfectly!";
    console.log(msg);
}

main();
