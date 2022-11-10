package test

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse() // 读取命令行参数
	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
}
