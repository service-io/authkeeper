// Package recorderx
// @author tabuyos
// @since 2023/8/22
// @description recorderx
package recorderx

import (
	"context"
	"deepsea/config/constant"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"
)

func TestWithContext(t *testing.T) {
	now := time.Now()
	recorder := WithContext(context.Background(), nil, slog.String("name", "tabuyos"))
	for i := 0; i < 10000; i++ {
		recorder.Trace("hello...")
		recorder.Debug("hello...")
		recorder.Info("hello...")
		recorder.Warn("hello...")
		recorder.Error("hello...")
		recorder.Fatal("hello...")

		fmt.Println("----------------------------------")
		recorder.With(slog.String("gender", "male")).With(slog.String("birthday", "2023/11/11")).Info("ok")
		fmt.Println("----------------------------------")

		// recorder = WithContext(context.Background(), slog.String("name", "tabuyos"), slog.Int64("age", 321))
		recorder.Trace("world...")
		recorder.Debug("world...")
		recorder.Info("world...")
		recorder.Warn("world...")
		recorder.Error("world...")
		recorder.Fatal("world...")
	}
	fmt.Println(time.Since(now))
}

func TestWithDefault(t *testing.T) {
	recorder := WithDefault(slog.String("name", "tabuyos"))
	recorder.Info("hello...")
	recorder.Infof("hello %+v...", "tabuyos")
}

func TestWithRecorder(t *testing.T) {
	now := time.Now()
	handler := &consoleHandler{
		opts:   slog.HandlerOptions{AddSource: true},
		writer: os.Stdout,
	}
	// handler := slog.NewTextHandler(os.Stdout, nil)
	var recorder Recorder = &commonRecorder{
		ctx:    context.Background(),
		logger: slog.New(handler),
	}
	for i := 0; i < 10000; i++ {
		// recorder.Info("hello...")
		// recorder.Error("hello...")
		// recorder.Infof("hello %+v...", "tabuyos")
		// recorder.With("name", "aaron liew").Info("hi...")
		// recorder.With(slog.Group("props", slog.String("age", "321"))).Info("hsi...")
		// recorder.With(slog.Group("props", slog.String("age", "123"))).Info("hei...")

		recorder.Trace("hello...")
		recorder.Debug("hello...")
		recorder.Info("hello...")
		recorder.Warn("hello...")
		recorder.Error("hello...")
		recorder.Fatal("hello...")

		fmt.Println("----------------------------------")
		recorder.With(slog.String("gender", "male")).With(slog.String("birthday", "2023/11/11")).Info("ok")
		fmt.Println("----------------------------------")

		// recorder = WithContext(context.Background(), slog.String("name", "tabuyos"), slog.Int64("age", 321))
		recorder.Trace("world...")
		recorder.Debug("world...")
		recorder.Info("world...")
		recorder.Warn("world...")
		recorder.Error("world...")
		recorder.Fatal("world...")
	}

	fmt.Println(time.Since(now))
}

func TestWithRecorderWith(t *testing.T) {
	now := time.Now()

	handler := NewConsoleHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}, nil)

	// handler = slog.NewTextHandler(os.Stdout, nil)

	handler = handler.WithGroup("op")

	var recorder Recorder = &commonRecorder{
		ctx:    context.Background(),
		logger: slog.New(handler),
	}

	// recorder.Info("hello...")
	// recorder.Error("hello...")
	// recorder.Infof("hello %+v...", "tabuyos")
	recorder.With("name", "tabuyos").With("name", "fdsa").Info("hi...")
	recorder.With(slog.Group("props", slog.String("age", "321"))).Info("hsi...")
	// recorder.With(slog.Group("props", slog.String("age", "123"))).Info("hei...")
	recorder.WithGroup("ppp").With(slog.String("age", "123321")).With(slog.String("name", "张三")).Info("hei...")

	fmt.Println(time.Since(now))
}

func TestBuilder(t *testing.T) {
	var buf strings.Builder
	buf.WriteString("hello")
	buf.WriteString(" ")
	buf.WriteString("world")
	println(buf.String())
}

func TestCustomizeHandler(t *testing.T) {

	handler := NewConsoleHandler(os.Stdout, &slog.HandlerOptions{
		Level: LevelTrace,
	}, slog.NewJSONHandler(NewChanWriter(nil, 0), nil))

	recorder := WithHandler(handler, slog.String(constant.TraceIdKey, "123456789"))

	recorder.Trace("Trace...")
	fmt.Println("---------------------------")
	recorder.Debug("Debug...")
	fmt.Println("---------------------------")
	recorder.Info("Info...")
	fmt.Println("---------------------------")
	recorder.Warn("Warn...")
	fmt.Println("---------------------------")
	recorder.Error("Error...")
	fmt.Println("---------------------------")
	recorder.Fatal("Fatal...")
}

func TestCustomizeHandlerLoop(t *testing.T) {

	now := time.Now()
	handler := NewConsoleHandler(os.Stdout, &slog.HandlerOptions{
		Level: LevelTrace,
	}, slog.NewJSONHandler(NewChanWriter(nil, 0), nil))

	recorder := WithHandler(handler, slog.String(constant.TraceIdKey, "123456789"))
	for i := 0; i < 10000; i++ {
		recorder.Trace("Trace...")
		fmt.Println("---------------------------")
		recorder.Debug("Debug...")
		fmt.Println("---------------------------")
		recorder.Info("Info...")
		fmt.Println("---------------------------")
		recorder.Warn("Warn...")
		fmt.Println("---------------------------")
		recorder.Error("Error...")
		fmt.Println("---------------------------")
		recorder.Fatal("Fatal...")
	}
	fmt.Println(time.Since(now))
}

func TestCustomizeHandlerWithChan(t *testing.T) {

	InitRecorder()

	now := time.Now()
	handler := NewConsoleHandler(os.Stdout, &slog.HandlerOptions{
		Level:     LevelTrace,
		AddSource: true,
	}, jsonChanHandler(nil, 1))

	go func() {
		queueChan := QueueChan()
		for e := range queueChan {
			fmt.Printf("e:> %+v", string(e.content))
		}
	}()

	recorder := WithHandler(handler, slog.String(constant.TraceIdKey, "123456789"))
	recorder.Trace("Trace...")
	// fmt.Println("---------------------------")
	recorder.Debug("Debug...")
	// fmt.Println("---------------------------")
	recorder.Info("Info...")
	// fmt.Println("---------------------------")
	recorder.Warn("Warn...")
	// fmt.Println("---------------------------")
	recorder.Error("Error...")
	// fmt.Println("---------------------------")
	recorder.Fatal("Fatal...")
	fmt.Println(time.Since(now))

	time.Sleep(1 * time.Second)
}

func TestTimeFmt(t *testing.T) {
	parse, _ := time.Parse(TimeFmt, "2023-08-24T12:33:49.300")
	fmt.Println(parse.Format(time.RFC3339Nano))
	fmt.Println(parse.Format(TimeFmt))
}

func BenchmarkFillSpace(b *testing.B) {
	b.ReportAllocs()
	s := "deepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/main"
	c := 100
	now := time.Now()
	for i := 0; i < b.N; i++ {
		if len(s) < c {
			r := fmt.Sprintf("%0*s", c, s)
			fmt.Println(r)
		} else {
			r := s[len(s)-c:]
			fmt.Println(r)
		}
	}
	fmt.Println(time.Since(now))
	// 000000000000000000000000000000000000000000000000000000000000000000000000000000deepsea-plat/temp/main
}

func BenchmarkSliceSpace(b *testing.B) {
	b.ReportAllocs()
	s := "deepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/main"
	c := 100
	ss := func() string {
		ss := make([]rune, c)
		for i := 0; i < c; i++ {
			ss[i] = '0'
		}
		return string(ss)
	}()
	now := time.Now()
	for i := 0; i < b.N; i++ {
		if len(s) < c {
			_ = ss[:c-len(s)] + s
			// fmt.Println(r)
		} else {
			_ = s[len(s)-c:]
			// fmt.Println(r)
		}
	}
	fmt.Println(time.Since(now))
	// 000000000000000000000000000000000000000000000000000000000000000000000000000000deepsea-plat/temp/main
	// at/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/maindeepsea-plat/temp/main
}
