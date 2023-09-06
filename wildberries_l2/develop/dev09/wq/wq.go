// Package wq - Work queue implementation
package wq

import (
	"container/list"
	"context"
	"sync"
)

// Task - функция задачи
type Task func(context.Context)

// WorkPool - пул задач
type WorkPool struct {
	maxWorkers int
	tasks      *list.List
	mutex      sync.Mutex
	wg         sync.WaitGroup
}

// New - создает пул задач
func New(maxWorkers int) *WorkPool {
	wp := WorkPool{
		maxWorkers: maxWorkers,
		tasks:      list.New(),
	}

	return &wp
}

// AddTask - добавляет задачу в пул
func (wp *WorkPool) AddTask(task Task) {
	wp.mutex.Lock()
	defer wp.mutex.Unlock()
	wp.tasks.PushBack(task)
	wp.wg.Add(1)
}

// RunAndWait - запускает воркеров и ожидает завершения всех задач
func (wp *WorkPool) RunAndWait(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for i := 0; i < wp.maxWorkers; i++ {
		go wp.workerLoop(ctx)
	}
	go wp.wgGuard(ctx)
	wp.wg.Wait()
	return ctx.Err()
}

// HasTasks - проверяет есть ли задачи
func (wp *WorkPool) HasTasks() bool {
	wp.mutex.Lock()
	defer wp.mutex.Unlock()

	return wp.tasks.Len() > 0
}

func (wp *WorkPool) workerLoop(ctx context.Context) {
	for ctx.Err() == nil && wp.HasTasks() {
		wp.performNextTask(ctx)
		wp.wg.Done()
	}
}

func (wp *WorkPool) performNextTask(ctx context.Context) {
	wp.mutex.Lock()
	if wp.tasks.Len() == 0 {
		wp.mutex.Unlock()
		return
	}

	front := wp.tasks.Front()
	task := front.Value.(Task)
	wp.tasks.Remove(front)
	wp.mutex.Unlock()

	task(ctx)
}

func (wp *WorkPool) wgGuard(ctx context.Context) {
	// release wait lock on ctx done
	<-ctx.Done()
	for i := 0; i < wp.tasks.Len(); i++ {
		wp.wg.Done()
	}
}
