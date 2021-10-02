package maxHeap

// parenr(i) = i/2
// left child (i) = 2*i
// right child (i) = 2*i + 1
type MaxHeap struct {
	data  []int
	count int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data:  []int{0},
		count: 0,
	}
}

func (m *MaxHeap) Size() int {
	return m.count
}

func (m *MaxHeap) IsEmpty() bool {
	return m.count == 0
}

//插入元素
func (m *MaxHeap) Insert(item int) {
	m.data = append(m.data, item)
	m.count++
	m.shiftUp(m.count)
}

func (m *MaxHeap) shiftUp(k int) {
	for k > 1 && m.data[k/2] < m.data[k] {
		m.data[k/2], m.data[k] = m.data[k], m.data[k/2]
		k /= 2
	}
}

func (m *MaxHeap) shiftDown(k int) {
	for 2*k <= m.count {
		j := 2 * k //在本轮循环中 data[k] 和 data[j] 交换位置
		if j+1 <= m.count && m.data[j+1] > m.data[j] {
			j = j + 1
		}

		if m.data[k] >= m.data[j] {
			break
		}

		m.data[k], m.data[j] = m.data[j], m.data[k]
		k = j
	}
}

//推出堆中最大的元素
func (m *MaxHeap) ExtractMax() (int, bool) {
	if m.count < 0 {
		return 0, false
	}
	item := m.data[1]
	m.data[1], m.data[m.count] = m.data[m.count], m.data[1]
	m.count--
	m.shiftDown(1)

	return item, true
}

//将数组构建成堆
//O(n)
func (m *MaxHeap) Heapify(arr []int) {
	for i := range arr {
		m.data = append(m.data, arr[i])
	}
	m.count = len(arr)

	for k := m.count / 2; k >= 1; k-- {
		m.shiftDown(k)
	}

	//fmt.Println(m.data)
}
