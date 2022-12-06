package point

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func StringConcatCompare() {

	a := []string{"a", "b", "c"}
	//方式1：+
	now := time.Now()
	ret1 := a[0] + a[1] + a[2]
	fmt.Printf("string: %s, time: %d\n", ret1, time.Since(now).Nanoseconds())
	//方式2：fmt.Sprintf
	now = time.Now()
	ret2 := fmt.Sprintf("%s%s%s", a[0], a[1], a[2])
	fmt.Printf("string: %s, time: %d\n", ret2, time.Since(now).Nanoseconds())

	//方式3：strings.Builder
	now = time.Now()
	var sb strings.Builder
	sb.WriteString(a[0])
	sb.WriteString(a[1])
	sb.WriteString(a[2])
	ret3 := sb.String()
	fmt.Printf("string: %s, time: %d\n", ret3, time.Since(now).Nanoseconds())

	//方式4：bytes.Buffer
	now = time.Now()
	buf := new(bytes.Buffer)
	buf.Write([]byte(a[0]))
	buf.Write([]byte(a[1]))
	buf.Write([]byte(a[2]))
	ret4 := buf.String()
	fmt.Printf("string: %s, time: %d\n", ret4, time.Since(now).Nanoseconds())

	//方式5：strings.Join
	now = time.Now()
	ret5 := strings.Join(a, "")
	fmt.Printf("string: %s, time: %d\n", ret5, time.Since(now).Nanoseconds())

}
