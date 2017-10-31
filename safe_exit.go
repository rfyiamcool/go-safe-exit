package safe_exit

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	GlobalWaitGroup = &ControlGroup{
		IsRunning: true,
		WG:        sync.WaitGroup{},
		Q:         make(chan os.Signal),
		ExitQ:     make(chan bool, 2000),
	}
)

type ControlGroup struct {
	IsRunning bool
	WG        sync.WaitGroup
	Q         chan os.Signal
	ExitQ     chan bool
}

func NewControlGroup() *ControlGroup {
	return GlobalWaitGroup
}

func NewSetControlGroup() *ControlGroup {
	return &ControlGroup{
		IsRunning: true,
		WG:        sync.WaitGroup{},
		Q:         make(chan os.Signal),
		ExitQ:     make(chan bool, 2000),
	}
}

func (c *ControlGroup) MakeSignal() {
	signal.Notify(c.Q, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
}

func (c *ControlGroup) RecvSignal() os.Signal {
	select {
	case s := <-c.Q:
		log.Println("recv signale: ", s)
		c.IsRunning = false
		c.PushCastExit()
		return s
	}
}

func (c *ControlGroup) MakeRecvSignal() os.Signal {
	c.MakeSignal()
	return c.RecvSignal()
}

func (c *ControlGroup) WaitTimeout(timeout int) bool {
	gg := make(chan struct{})
	go func() {
		defer close(gg)
		c.WG.Wait()
	}()

	select {
	case <-gg:
		return false
	case <-time.After(time.Duration(timeout) * time.Second):
		return true
	}
}

func (c *ControlGroup) CheckRunning() bool {
	if c.IsRunning {
		return true
	}
	return false
}

func (c *ControlGroup) Add() {
	c.WG.Add(1)
}

func (c *ControlGroup) Done() {
	c.WG.Done()
}

func (c *ControlGroup) PushCastExit() {
	for i := 0; i < 1000; i++ {
		c.ExitQ <- false
	}
}
func (c *ControlGroup) PullExit() bool {
	d, ok := <-c.ExitQ
	if ok {
		return d
	}
	return true
}

