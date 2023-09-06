package main

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

/*
   === Утилита telnet ===

   Реализовать примитивный telnet клиент:
   Примеры вызовов:
     - go-telnet --timeout=10s host port
     - go-telnet mysite.ru 8080
     - go-telnet --timeout=3s 1.1.1.1 123

   Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
   После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
   Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

   При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
   При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

const (
	timeoutPrefix  = "--timeout="
	defaultTimeout = 10 * time.Second
)

var (
	errTooFewArguments    = errors.New("too few arguments")
	errInvalidKeyArgument = errors.New("invalid key argument")
	errInvalidTimeout     = errors.New("invalid timeout")

	errClosedByServer = errors.New("closed by server")
	errClosedByUs     = errors.New("closed by us")
)

type args struct {
	timeout time.Duration
	host    string
	port    string
}

func main() {
	args, err := parseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), args.timeout)
	defer cancel()

	err = runTelnet(ctx, os.Stdout, os.Stdin, args.host, args.port)
	if err != nil && err != errClosedByUs {
		log.Fatal(err)
	}
}

func parseArgs(rawArgs []string) (args, error) {
	timeout := defaultTimeout

	if len(rawArgs) == 3 {
		if !strings.HasPrefix(rawArgs[0], timeoutPrefix) {
			return args{}, errInvalidKeyArgument
		}

		dStr := rawArgs[0][len(timeoutPrefix):]
		var err error
		timeout, err = time.ParseDuration(dStr)
		if err != nil {
			return args{}, errInvalidTimeout
		}
		rawArgs = rawArgs[1:]
	}

	if len(rawArgs) != 2 {
		return args{}, errTooFewArguments
	}

	host := rawArgs[0]
	port := rawArgs[1]
	return args{timeout, host, port}, nil
}

func runTelnet(ctx context.Context, out io.Writer, in io.Reader, host, port string) error {
	con, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}

	readerChan := bufferedReaderToChan(ctx, in)
	serverChan := bufferedReaderToChan(ctx, con)

	for err == nil {

		select {
		case <-ctx.Done():
			err = ctx.Err()
		case data, ok := <-readerChan:
			if ok {
				_, err = con.Write(data)
			} else {
				err = errClosedByUs
			}
		case data, ok := <-serverChan:
			if ok {
				_, err = out.Write(data)
			} else {
				err = errClosedByServer
			}
		}
	}
	return err
}

func bufferedReaderToChan(ctx context.Context, r io.Reader) <-chan []byte {
	ch := make(chan []byte)

	go func() {
		sc := bufio.NewScanner(r)
		for sc.Scan() && sc.Err() == nil && ctx.Err() == nil {
			data := sc.Text()
			if len(data) == 0 {
				break
			}
			ch <- []byte(data + "\n")
		}
		close(ch)
	}()

	return ch
}
