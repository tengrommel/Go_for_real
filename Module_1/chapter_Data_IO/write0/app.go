package main

import "fmt"

// 声明管道并不声明管道长度
type channelWriter chan byte

// 注入方法
func (c channelWriter)Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	count := 0
	// 创建等待通道
	wait := make(chan struct{})
	go func() {
		// 发送到 channelWriter 通道
		for _, b := range p {
			c <- b
			count ++
		}
		close(c)
		close(wait)
	}()
	// 等待进程
	<-wait
	return count,nil
}

func main() {
	data := []byte("Stream me!")
	cw := channelWriter(make(chan byte, len(data)))
	cw.Write(data)
	for c:= range cw{
		fmt.Println(c)
	}
	}
