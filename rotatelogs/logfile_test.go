package rotatelogs

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"
)

func TestEnqueue(t *testing.T) {
	f := LogFile{
		Node:      "xxx",
		Timestamp: time.Now(),
		Filepath:  "/sdf/sdf",
	}
	m := New(ByTime)
	m.Enqueue(f)
	q := m.GetQueue("xxx")
	fmt.Println(q.Peek())
}

func TestTimeParser(t *testing.T) {
	fmt.Println(hex.EncodeToString(NewLine))

	s := "2023-04-19T11-16-37.209"
	ts, err := time.Parse("2006-01-02T15-04-05.999999999", s)
	if err != nil {
		t.Error(err)
	}
	t.Log(ts)
}
