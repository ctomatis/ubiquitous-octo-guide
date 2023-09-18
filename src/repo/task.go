package repo

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"go.vemo/src/models"
)

const fname = "task.json"

var ErrItemNotFound error = errors.New("item does not exist")

type Tasks struct {
	Data models.Tasks
}

func (t *Tasks) Init() {
	fp, err := os.Open(fname)
	defer fp.Close()
	if err != nil {
		log.Printf("Error reading '%s' file: %v", fname, err)
	}

	var tasks models.Tasks
	jbytes, _ := io.ReadAll(fp)
	json.Unmarshal(jbytes, &tasks)

	// Set task ids
	for i := range tasks {
		tasks[i].Id = i + 1
	}
	t.Data = tasks
}

func (ts *Tasks) All() models.Tasks {
	return ts.Data
}

func (ts *Tasks) Find(id int) (t models.Task, err error) {
	for _, t = range ts.Data {
		if id == t.Id {
			return
		}
	}
	return t, ErrItemNotFound
}

func (ts *Tasks) Create(t *models.Task) {
	id := 1
	n := len(ts.Data)
	if n > 0 {
		id = ts.Data[n-1].Id + 1
	}
	t.Id = id
	ts.Data = append(ts.Data, *t)
}

func (ts *Tasks) Update(id int, t *models.Task) {
	for i := range ts.Data {
		if id == ts.Data[i].Id {
			ts.Data[i].Name = t.Name
			ts.Data[i].Detail = t.Detail
			ts.Data[i].Done = t.Done
			ts.Data[i].Due = t.Due
			*t = ts.Data[i]
			break
		}
	}
}

func (ts *Tasks) Delete(id int) {
	for i := range ts.Data {
		if id == ts.Data[i].Id {
			ts.Data = append(ts.Data[:i], ts.Data[i+1:]...)
			break
		}
	}
}
