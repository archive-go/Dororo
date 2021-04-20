package core

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"time"
)



func Start()  {

	b := GetBrowser("user/bili" , false)
	page := b.MustPage("http://bilibili.com")
	login := page.MustWaitLoad().MustHasR(".unlogin-avatar" , "登录")

	//需要登陆操作
	if login == true {
		fmt.Println("-----need login --------")
		b.MustPage("http://bilibili.com").MustElement(".bili-icon_dingdao_dengchu")
		fmt.Println("--------log finish--------")
	}
	playAdnShare(page)


	time.Sleep(time.Second * 60 * 60)
}

//返回包含本地头文件的浏览器端

func GetBrowser(path string , headless bool) *rod.Browser {

	b := rod.New().ControlURL(launcher.New().Headless(headless).UserDataDir(path).MustLaunch()).MustConnect()

	return b

}

//判断用户是否登录

func isLog()  {
	//b := GetBrowser("http://jd.com")
}

//自动播放 && 分享视频

func playAdnShare(page *rod.Page)  {
	el , err := page.ElementX(".video-card-reco")
}

//直播项目签到

func LiveLog  (page *rod.Page) {
	page.MustElementR("a" ,"直播").MustClick()
	fmt.Println("bark")
}