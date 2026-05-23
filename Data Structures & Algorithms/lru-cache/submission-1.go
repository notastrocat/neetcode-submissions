type dll struct {
	key int
	val int
	next *dll
	prev *dll
}

type LRUCache struct {
    table map[int]*dll
	capacity int
	head *dll
	tail *dll
}

func Constructor(capacity int) LRUCache {
	tmpHead := &dll{}
	tmpTail := &dll{}

	tmpHead.next = tmpTail
	tmpTail.prev = tmpHead

    return LRUCache{
		capacity: capacity,
		head: tmpHead,
		tail: tmpTail,
		table: make(map[int]*dll),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.table[key]; ok {
    	// handle the logic to make it LRU element
		this.removeNode(node)
		this.moveToFront(node)

		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.table[key]; ok {

		node.val = value
		this.removeNode(node)
		this.moveToFront(node)

	} else {

		newNode := &dll{key: key, val: value}
		this.table[key] = newNode
		this.moveToFront(newNode)

		if len(this.table) > this.capacity {
			lru := this.tail.prev
			
			this.removeNode(lru)
			
			delete(this.table, lru.key)
		}
	}
}

func (this *LRUCache) removeNode(node *dll) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToFront(node *dll) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}
