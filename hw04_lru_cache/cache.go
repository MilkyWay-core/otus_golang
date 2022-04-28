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
	finded_element := cache.items[key]
	if finded_element == nil {
		return nil, false
	}
	cache.queue.MoveToFront(finded_element)
	return finded_element.Value.(*cacheItem).value, true
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	//логику выталкивания элементов из-за размера очереди
	if cache.queue.Len()+1 > cache.capacity {
		removing_element := cache.queue.Back()                       //получаем последний
		delete(cache.items, removing_element.Value.(*cacheItem).key) //удаляем из кеша
		cache.queue.Remove(removing_element)                         //удаляем из очереди
	}
	//если элемент присутствует в словаре, то обновить его значение и переместить элемент в начало очереди
	finded_item := cache.items[key]
	if finded_item != nil {
		finded_item.Value.(*cacheItem).value = value
		cache.queue.MoveToFront(finded_item)
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
	first_item := cache.queue.Front()
	for first_item != nil {
		delete(cache.items, first_item.Value.(*cacheItem).key) //удаляем их кеша
		if first_item.Next == nil {                            //если последний элемент очереди удялем и выходим из ципла
			cache.queue.Remove(first_item)
			break
		}
		first_item = first_item.Next //перемещаем курсор на следующий элемент
		if first_item.Prev != nil {
			cache.queue.Remove(first_item.Prev) //удаляем предыдущий
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
