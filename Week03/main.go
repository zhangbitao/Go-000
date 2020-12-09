package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"Week03/server/grpc"
	"Week03/server/http"

	"golang.org/x/sync/errgroup"
)

func registerSignal(shutdown chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, []os.Signal{syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM}...)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]
				fmt.Printf("errgroup: panic recovered: %s\n%s", r, buf)
			}
		}()

		switch <-c {
		case syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
			close(shutdown)
			return
		}
	}()
}

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	s1 := http.New(8081)
	s2 := grpc.New(9000)

	group.Go(func() error {
		return s1.Start(ctx)
	})
	group.Go(func() error {
		return s2.Start(ctx)
	})
	group.Go(func() error {
		// wait for quit signal
		shutdown := make(chan struct{})
		registerSignal(shutdown)
		<-shutdown

		s1.Stop(ctx)
		s2.Stop(ctx)
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}
