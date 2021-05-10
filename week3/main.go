package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

//reference https://pkg.go.dev/golang.org/x/sync/errgroup#WithContext
//笔记 errorgroup 的作用 多个goroutine执行时 记录各个groutine发生的错误
//每个子任务遇到错误时都应该抛出，g.Wait()等待所有子任务运行完成 最终返回结果集和错误信息

func main() {
	var g errgroup.Group

	// 启动第一个子任务,它执行成功
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		// return errors.New("failed to exec #1")
		return nil
	})
	// 启动第二个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		return errors.New("failed to exec #2")
	})

	// 启动第三个子任务，它执行成功
	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		return nil
	})
	// 等待三个任务都完成
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Println("failed:", err)
	}
}
