package core

import (
	"time"
)

func Start()  {

	b := NewBiHandle()
	b.GetLiveFocus()

	//b.BigBrotherIsWatchingYou("20076571")

	time.Sleep(time.Second * 60 * 60)
}