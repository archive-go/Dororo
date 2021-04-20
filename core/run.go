package core

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"time"
)


var b *rod.Browser
func Start()  {

	b = GetBrowser("user/bili" , false)
	page := b.MustPage("http://bilibili.com")
	login := page.MustWaitLoad().MustHasR(".unlogin-avatar" , "登录")

	//需要登陆操作
	if login == true {
		fmt.Println("-----need login --------")
		b.MustPage("http://bilibili.com").MustElement(".bili-icon_dingdao_dengchu")
		fmt.Println("--------log finish--------")
	}

	//playAdnShare(page)

	spacePage := b.MustPage("http://space.bilibili.com")
	GetFocus(spacePage)

	time.Sleep(time.Second * 60 * 60)
}

//返回包含本地头文件的浏览器端

func GetBrowser(path string , headless bool) *rod.Browser {

	chrome , _ := launcher.LookPath()

	b := rod.New().ControlURL(launcher.New().Bin(chrome).Headless(headless).UserDataDir(path).MustLaunch()).MustConnect()

	return b

}

//判断用户是否登录

func isLog()  {
	//b := GetBrowser("http://jd.com")
}

//自动播放 && 分享视频

func playAdnShare(page *rod.Page)  {
	v := page.MustElement(".video-card-reco a").MustAttribute("href")

	childPage := b.MustPage(fmt.Sprintf("http:%s" , *v))
	childPage.MustElement("#bilibiliPlayer").MustClick()
	videoTime := childPage.MustElement(".bilibili-player-video-time-total").MustText()
	fmt.Println(videoTime)  // 视频总场地 "09:22" 后期转为秒 在处理  更好的思路 等 开始时间和结束时间相等时候  关闭这个标签页
	time.Sleep(time.Second * 10)
	childPage.Close()
	//b.MustPage()
}

//获取我关注的列表

func GetFocus(spacePage *rod.Page)  {

	spacePage.MustElement(".n-statistics a").MustClick()
	spacePage.MustElement(".list-item")
	arr := spacePage.MustElements(".list-item")
	for _ , v := range  arr {
		name , _ := v.MustElement(".fans-name").Text()
		fmt.Println(name)
	}
	if len(arr) == 20 {
		spacePage.MustElement(".be-pager-next").MustClick()
		time.Sleep(time.Second)
		GetFocus(spacePage)
	}
}

//直播项目签到

func LiveLog  (page *rod.Page) {
	page.MustElementR("a" ,"直播").MustClick()
	fmt.Println("bark")
}