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

func (ws *WorklogService) Start(windowName string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}

	taskID := idgen.New()

	task := &domain.Task{
		ID:        taskID,
		Name:      windowName,
		CreatedAt: time.Now(),
	}
	task.AddWindow(windowName)

	window := &domain.Window{
		ID: windowName,
	}
	window.Start()

	store.Tasks[taskID] = task
	store.Windows[windowName] = window

	return ws.persist(store)
}

func (ws *WorklogService) Child(parentTaskID, windowName string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}

	task, ok := store.Tasks[parentTaskID]
	if !ok {
		return fmt.Errorf("Tarefa %s nao encontrada", parentTaskID)
	}

	window := &domain.Window{
		ID: windowName,
	}
	window.Start()

	task.AddWindow(windowName)
	store.Windows[windowName] = window

	return ws.persist(store)
}

func (ws *WorklogService) Pause(taskID, windowName string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}

	win, ok := store.Windows[windowName]
	if !ok {
		return fmt.Errorf("Window %s nao encontrada", windowName)
	}
	win.Pause()
	return ws.persist(store)
}

func (ws *WorklogService) Resume(windowName string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}

	win, ok := store.Windows[windowName]
	if !ok {
		return fmt.Errorf("Window %s nao encontrada", windowName)
	}
	win.Resume()
	return ws.persist(store)
}

func (ws *WorklogService) Switch(fromName, toName string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}
	if from, ok := store.Windows[fromName]; ok {
		from.Pause()
	}
	if to, ok := store.Windows[toName]; ok {
		to.Resume()
	}
	return ws.persist(store)
}

func (ws *WorklogService) Stop(taskID, windowName string) error {
	store, err := ws.storage.Load()
	if err != nil {
		return err
	}
	win, ok := store.Windows[windowName]
	if !ok {
		return fmt.Errorf("window %s nao encontrada", windowName)
	}
	win.Close()

	task, ok := store.Tasks[taskID]
	if !ok {
		return fmt.Errorf("task %s nao encontrada", taskID)
	}
	for _, wname := range task.WindowIDs {
		if wname == windowName {
			z := task.RemoveWindow(windowName)
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
