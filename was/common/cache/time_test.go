package cache

import (
	"fmt"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
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

	fmt.Println("duration :: ", enum.JWTAccessDuration)
	fmt.Println("time :: ", time.Minute*15)

	fmt.Println("duration :: ", enum.JWTRefreshDuration)
	fmt.Println("time :: ", time.Hour*24*7)

	fmt.Println("a", a)
	fmt.Println("p", p)
	fmt.Println("at", at)
	fmt.Println("rt", rt)
	fmt.Println("now", now)
	fmt.Println("ut", ut)
	fmt.Println("at.Sub(now)", at.Sub(now))
	fmt.Println("rt.Sub(now)", rt.Sub(now))
}
