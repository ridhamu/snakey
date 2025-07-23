package game

import "github.com/ridhamu/snakey/entity"

type World struct {
	entities []entity.Entity
}

func NewWorld() *World {
	return &World{
		entities: []entity.Entity{},
	}
}

func (w *World) AddEntity(e entity.Entity) {
	w.entities = append(w.entities, e)
}

func (w *World) Entities() []entity.Entity {
	return w.entities
}

func (w World) GetEntities(tag string) []entity.Entity {
	var result []entity.Entity

	for _, entity := range w.entities {
		if entity.Tag() == tag {
			result = append(result, entity)
		}
	}

	return result
}

func (w World) GetFirstEntity(tag string) (entity.Entity, bool) {
	for _, entity := range w.entities {
		if entity.Tag() == tag {
			return entity, true
		}
	}

	return nil, false
}
