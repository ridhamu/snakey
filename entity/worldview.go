package entity


type WorldView interface {
	GetEntities(tag string) []Entity
}
