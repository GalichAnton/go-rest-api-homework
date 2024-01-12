package storage

import (
	"sync"

	"github.com/GalichAnton/go-rest-api-homework/internal/model"
)

type SyncMap struct {
	Elems map[string]model.Task
	M     sync.RWMutex
}

var Tasks = map[string]model.Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

var TaskStorage = SyncMap{
	Elems: Tasks,
	M:     sync.RWMutex{},
}
