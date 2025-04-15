package utils

type Heap struct {
	heap []int
}

func NewHeap() Heap {
	return Heap{}
}

func (h Heap) Push(x int) {
	h.heap = append(h.heap, x)
	// 挿入された頂点番号
	i := len(h.heap) - 1
	for i > 0 {
		// 親の頂点番号
		p := (i - 1) / 2
		// 逆転がなければ終了
		if h.heap[p] >= x {
			break
		}
		// 自分の値を親の値にする
		h.heap[i] = h.heap[p]
		i = p
	}
	// x は最終的にこの位置に持ってくる
	h.heap[i] = x
}

func (h Heap) Top() int {
	if len(h.heap) == 0 {
		return -1
	}
	return h.heap[0]
}

func (h Heap) Pop() {
	l := len(h.heap)
	if l == 0 {
		return
	}
	x := h.heap[l-1] // 頂点に持ってくる値
	h.heap = h.heap[:l-1]
	i := 0
	for i*2+1 < l {

	}
}
