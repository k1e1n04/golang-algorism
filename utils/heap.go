package utils

//ヒープの条件
//  ・頂点 v の親頂点を p としたとき，key[p]key[v] が成立する．
//  ・木の高さを h としたとき，木の深さ h－1 以下の部分については，完全二分木を形成している.
//  ・木の高さを h としたとき，木の深さ h の部分については，頂点が左詰めされている．

type Heap struct {
	heap []int
}

func NewHeap() Heap {
	return Heap{}
}

// Push はヒープに要素を追加
func (h *Heap) Push(x int) {
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

// Top はヒープの最大値を返す
func (h *Heap) Top() int {
	if len(h.heap) == 0 {
		return -1
	}
	return h.heap[0]
}

// Pop はヒープの最大値を削除
func (h *Heap) Pop() {
	l := len(h.heap)
	if l == 0 {
		return
	}

	x := h.heap[l-1]
	h.heap = h.heap[:l-1]
	l--

	if l == 0 {
		return
	}

	i := 0
	for i*2+1 < l {
		// 左の子
		a := i*2 + 1
		// 右の子
		b := i*2 + 2
		// 大きい方の子を選ぶ
		if b < l && h.heap[b] > h.heap[a] {
			a = b
		}
		if h.heap[a] <= x {
			break
		}
		h.heap[i] = h.heap[a]
		i = a
	}
	h.heap[i] = x
}
