package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestTimePrint(t *testing.T) {
	a := time.Now().Add(time.Minute * 15).Unix()
	p := time.Now().Add(time.Hour * 2).Unix()
	at := time.Unix(a, 0)
	rt := time.Unix(p, 0)
	now := time.Now()
	unow := time.Now().Unix()
	ut := time.Unix(unow, 0)

	fmt.Println(a)
	fmt.Println(t)
	fmt.Println(at)
	fmt.Println(rt)
	fmt.Println(now)
	fmt.Println(ut)
	fmt.Println(at.Sub(now))
	fmt.Println(rt.Sub(now))
}
