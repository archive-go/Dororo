package core

import (
	"time"
)

func Start()  {

	b := NewBiHandle()
	b.BigBrotherIsWatchingYou()

	time.Sleep(time.Second * 60 * 60)
}