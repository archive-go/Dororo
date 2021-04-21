package core

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

//返回包含本地头文件的浏览器端

func GetBrowser(path string , headless bool) *rod.Browser {

	chrome , _ := launcher.LookPath()

	b := rod.New().ControlURL(launcher.New().Bin(chrome).Headless(headless).UserDataDir(path).MustLaunch()).MustConnect()

	return b

}
