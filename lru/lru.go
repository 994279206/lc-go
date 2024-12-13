package lru

type LRUCache struct {
	size       int
	cap        int
	cache      map[int]*Node
	head, tail *Node
}

type Node struct {
	k, v       int
	prev, next *Node
}

func Constructor(capacity int) *LRUCache {
	l := &LRUCache{
		cap:   capacity,
		head:  initNode(0, 0),
		tail:  initNode(0, 0),
		cache: map[int]*Node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func initNode(k, v int) *Node {
	return &Node{
		k: k,
		v: v,
	}
}
func (l *LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		l.moveToHead(node)
		return node.v
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; !ok {
		node := initNode(key, value)
		l.addToHead(node)
		l.cache[key] = node
		l.size++
		if l.size > l.cap {
			removeNode := l.removeTail()
			delete(l.cache, removeNode.k)
			l.size--
		}
		return
	}
	node := l.cache[key]
	node.v = value
	l.moveToHead(node)
}

func (l *LRUCache) addToHead(node *Node) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) moveToHead(node *Node) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeTail() *Node {
	node := l.tail.prev
	l.removeNode(node)
	return node
}
