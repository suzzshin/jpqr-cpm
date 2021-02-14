package jpqrcpm

import (
	"testing"
	
	"time"
	"strconv"
	"strings"
)

func Test_Generate(t *testing.T) {
	q := NewJPQR1("12345678", "87654321", "{\"key\":\"value\"}")
	c, err := q.Encode()

	if err == nil {
		t.Logf("ok: %s", c)
	} else {
		t.Errorf("failed encode: %s", err)
	}
}

func Test_EncodedHeader(t *testing.T) {
	q := NewJPQR1(strconv.FormatInt(time.Now().Unix(), 10), strconv.FormatInt(time.Now().Unix(), 10), "{\"key\":\"value\"}")
	c, err := q.Encode()

	if err == nil {
		if strings.HasPrefix(c, "hQVKUFFS") {
			t.Logf("ok: %s", c)
		} else {
			t.Errorf("bad header: %s", c)
		}
	} else {
		t.Errorf("failed encode: %s", err)
	}
}

func Benchmark_Generate(b *testing.B) {
	b.ResetTimer()

	q := NewJPQR1(strconv.FormatInt(time.Now().Unix(), 10), strconv.FormatInt(time.Now().Unix(), 10), "{\"key\":\"value\"}")
	for i := 0; i < b.N; i++ {
		q.Encode()
	}

	b.StopTimer()
}