package interfaces

type Sortable interface {
	Less(b Sortable) bool
}
