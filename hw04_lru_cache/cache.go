package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	if cache.capacity == 0 {
		return nil, false
	}
	findedElement := cache.items[key]
	if findedElement == nil {
		return nil, false
	}
	cache.queue.MoveToFront(findedElement)
	return findedElement.Value.(*cacheItem).value, true
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	// логику выталкивания элементов из-за размера очереди
	if cache.queue.Len()+1 > cache.capacity {
		removingElement := cache.queue.Back()                       // получаем последний
		delete(cache.items, removingElement.Value.(*cacheItem).key) // удаляем из кеша
		cache.queue.Remove(removingElement)                         // удаляем из очереди
	}
	// если элемент присутствует в словаре, то обновить его значение и переместить элемент в начало очереди

	if findedItem := cache.items[key]; findedItem != nil {
		findedItem.Value.(*cacheItem).value = value
		cache.queue.MoveToFront(findedItem)
		return true
	}
	cache.items[key] = cache.queue.PushFront(&cacheItem{
		key:   key,
		value: value,
	},
	)
	return false
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (cache *lruCache) Clear() {
	firstItem := cache.queue.Front()
	for firstItem != nil {
		delete(cache.items, firstItem.Value.(*cacheItem).key) // удаляем их кеша
		if firstItem.Next == nil {                            // если последний элемент очереди удялем и выходим из ципла
			cache.queue.Remove(firstItem)
			break
		}
		firstItem = firstItem.Next // перемещаем курсор на следующий элемент
		if firstItem.Prev != nil {
			cache.queue.Remove(firstItem.Prev) // удаляем предыдущий
		}
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
