package service

import (
	"fmt"
	"time"

	"github.com/htmluz/worklog/internal/domain"
	"github.com/htmluz/worklog/internal/storage"
	"github.com/htmluz/worklog/pkg/idgen"
)

type WorklogService struct {
	storage storage.Storage
}

func NewWorklogService(s storage.Storage) *WorklogService {
	return &WorklogService{storage: s}
}

func (ws *WorklogService) Start(windowID string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}

	wdw := domain.Window{
		ID: windowID,
	}
	wdw.Start()

	taskID := idgen.New()

	task := &domain.Task{
		ID:        taskID,
		Name:      windowID,
		CreatedAt: time.Now(),
	}
	task.AddWindow(wdw)

	store.Tasks[taskID] = task

	return ws.persist(store)
}

func (ws *WorklogService) Child(parentTaskID, windowID string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}

	task, ok := store.Tasks[parentTaskID]
	if !ok {
		return fmt.Errorf("Tarefa %s nao encontrada", parentTaskID)
	}

	wdw := domain.Window{
		ID: windowID,
	}
	wdw.Start()

	task.AddWindow(wdw)

	return ws.persist(store)
}

func (ws *WorklogService) Pause(taskID, windowID string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}

	task, ok := store.Tasks[taskID]
	if !ok {
		return fmt.Errorf("task %s nao encontrada", taskID)
	}

	// TODO
	for _, wdw := range task.Windows {
		if wdw.ID == windowID {
			wdw.Pause()
		}
	}
	// win.Pause()
	return ws.persist(store)
}

func (ws *WorklogService) Resume(taskID, windowID string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}

	task, ok := store.Tasks[taskID]
	if !ok {
		return fmt.Errorf("Window %s nao encontrada", windowID)
	}

	for _, wdw := range task.Windows {
		if wdw.ID == windowID {
			wdw.Resume()
		}
	}
	return ws.persist(store)
}

func (ws *WorklogService) Switch(fromTask, fromWindow, toTask, toWindow string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}
	if from, ok := store.Tasks[fromTask]; ok {
		for _, wdw := range from.Windows {
			if wdw.ID == fromWindow {
				wdw.Pause()
			}
		}
	}
	if to, ok := store.Tasks[toTask]; ok {
		for _, wdw := range to.Windows {
			wdw.Resume()
		}
	}
	return ws.persist(store)
}

func (ws *WorklogService) Stop(taskID, windowID string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}

	task, ok := store.Tasks[taskID]
	if !ok {
		return fmt.Errorf("task %s nao encontrada", taskID)
	}

	for _, wdw := range task.Windows {
		if wdw.ID == windowID {
			wdw.Close()
			z := task.RemoveWindow(windowID)
			if z {
				task.Close()
			}
		}
	}

	return ws.persist(store)
}

func (ws *WorklogService) List() error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}
	if len(store.Tasks) < 1 {
		fmt.Println("Sem tasks")
		return nil
	}
	for _, t := range store.Tasks {
		fmt.Printf("%v", t)
	}
	return nil
}

// helpers

func (ws *WorklogService) persist(store *domain.Store) error {
	if err := ws.storage.Save(store); err != nil {
		return err
	}
	fmt.Println("stored task")
	return nil
}
