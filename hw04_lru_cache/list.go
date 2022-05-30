package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	items *ListItem
}

func (list *list) Len() int {
	var i int
	firstItem := list.Front()
	for firstItem != nil {
		firstItem = firstItem.Next
		i++
	}
	return i
}

func (list *list) Front() *ListItem {
	if list.items == nil {
		return nil
	}
	for list.items.Prev != nil {
		list.items = list.items.Prev
	}
	return list.items
}

func (list *list) Back() *ListItem {
	if list.items == nil {
		return nil
	}
	for list.items.Next != nil {
		list.items = list.items.Next
	}
	return list.items
}

func (list *list) PushFront(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	if firstItem := list.Front(); firstItem == nil {
		list.items = item
	} else {
		item.Next = firstItem
		firstItem.Prev = item
		item.Prev = nil
	}
	return item
}

func (list *list) PushBack(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	if lastItem := list.Back(); lastItem == nil {
		list.items = item
	} else {
		item.Prev = lastItem
		lastItem.Next = item
		item.Next = nil
	}
	return item
}

func (list *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
		list.items = i.Prev // сдвигаем курсор с удаляемого
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
		list.items = i.Next
	}
	if i.Next == nil && i.Prev == nil {
		list.items = nil // если едиственный элемент в списке то очищаем курсор
	}
}

func (list *list) MoveToFront(i *ListItem) {
	tmp := i.Value
	list.Remove(i)
	list.PushFront(tmp)
}

func NewList() List {
	return new(list)
}
