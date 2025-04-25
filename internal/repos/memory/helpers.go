package memory

import (
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/pkg/interfaces"
	"slices"
)

func getPage[T interfaces.Sortable](page models.PageMetadata, data []T) []T {
	if page.OrderDir == "" {
		page.OrderDir = "desc"
	}

	slices.SortFunc(data, func(a, b T) int {
		if page.OrderDir == "desc" {
			if a.Less(b) {
				return 1
			}
			return -1
		}

		if a.Less(b) {
			return -1
		}
		return 1
	})

	if page.Size == 0 {
		return data
	}

	idFrom := page.Page * page.Size
	idTo := (page.Page + 1) * page.Size

	if len(data) < idFrom {
		return nil
	}
	if len(data) < idTo {
		return data[idFrom:len(data)]
	}

	return data[idFrom:idTo]
}
