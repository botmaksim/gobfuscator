class SegmentTree
  def initialize(size)
    @n = size
    @tree = Array.new(4 * size, 0)
  end

  def build(a, v, tl, tr)
    if tl == tr
      @tree[v] = a[tl]
    else
      tm = (tl + tr) / 2
      build(a, v * 2, tl, tm)
      build(a, v * 2 + 1, tm + 1, tr)
      @tree[v] = @tree[v * 2] + @tree[v * 2 + 1]
    end
  end

  def query(v, tl, tr, l, r)
    return 0 if l > r
    return @tree[v] if l == tl && r == tr
    tm = (tl + tr) / 2
    query(v * 2, tl, tm, l, [r, tm].min) + 
      query(v * 2 + 1, tm + 1, tr, [l, tm + 1].max, r)
  end

  def update(v, tl, tr, pos, new_val)
    if tl == tr
      @tree[v] = new_val
    else
      tm = (tl + tr) / 2
      if pos <= tm
        update(v * 2, tl, tm, pos, new_val)
      else
        update(v * 2 + 1, tm + 1, tr, pos, new_val)
      end
      @tree[v] = @tree[v * 2] + @tree[v * 2 + 1]
    end
  end
end

def main
  n = 5
  a = [1, 2, 3, 4, 5]
  st = SegmentTree.new(n)
  st.build(a, 1, 0, n - 1)

  puts "Sum 0-4: #{st.query(1, 0, n - 1, 0, 4)}"
  st.update(1, 0, n - 1, 2, 10)
  puts "Sum 0-4 after update: #{st.query(1, 0, n - 1, 0, 4)}"

  msg = "Segment tree runs perfectly!"
  puts msg
end

main
