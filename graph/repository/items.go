package repository

import (
	"fmt"
	"sync"
)

type Item struct {
	ID   int
	Name string
}

type ListItemsRepository struct {
	*sync.Mutex
	items  []Item
	lastID int
}

func NewListItemsRepository() *ListItemsRepository {
	return &ListItemsRepository{Mutex: &sync.Mutex{}}
}

func (r *ListItemsRepository) InitDefaultList() {
	r.lastID = 0

	for i := 1; i <= 5; i++ {
		r.Add(fmt.Sprintf("Item %d", i))
	}
}

func (r *ListItemsRepository) Add(name string) Item {
	r.Lock()
	defer r.Unlock()

	r.lastID++
	item := Item{ID: r.lastID, Name: name}

	r.items = append(r.items, item)

	return item
}

func (r *ListItemsRepository) GetByIDs(ids []int) (result []Item) {
	r.Lock()
	defer r.Unlock()

	if len(ids) == 0 {
		return result
	}

	for _, item := range r.items {
		for _, id := range ids {
			if id == item.ID {
				result = append(result, item)

				break
			}
		}
	}

	return result
}

func (r *ListItemsRepository) GetAll() (result []Item) {
	return r.items
}
