package core

import (
	"fmt"
	"github.com/go-rod/rod"
	"time"
)

type BiliHandle struct {

	B *rod.Browser  //浏览器句柄
	Page *rod.Page //页面句柄

	IsLogin bool //是否登录
	FocusNumers int

}

func NewBiHandle() *BiliHandle {

	b := GetBrowser("user/bili" , false)
	page := b.MustPage("http://bilibili.com")
	login := !page.MustWaitLoad().MustHasR(".unlogin-avatar" , "登录")

	return &BiliHandle{
		B:       b,
		Page:    page,
		IsLogin: login,
	}

}
//等待用户登录
func (b *BiliHandle) WaitLogIn(){
	if b.IsLogin {
		b.B.MustPage("http://bilibili.com").MustElement(".bili-icon_dingdao_dengchu")
	}
}


//播放视频

func (b *BiliHandle) Play()  {
	v := b.Page.MustElement(".video-card-reco a").MustAttribute("href")

	childPage := b.B.MustPage(fmt.Sprintf("http:%s" , *v))
	childPage.MustElement("#bilibiliPlayer").MustClick()

	for {
		videoStart := childPage.MustElement("bilibili-player-video-time-now").MustText()
		videoTime := childPage.MustElement(".bilibili-player-video-time-total").MustText()

		fmt.Print(videoStart)
		fmt.Print(videoTime)

		if(videoStart == videoTime) {
			break
		}
		time.Sleep(time.Second )
	}

	fmt.Print("vide play ending")
	childPage.Close()
	//b.MustPage()
}

//分享视频

func (b *BiliHandle) Share ()  {

}

//点赞视频

func (b *BiliHandle) Like() {
	
}

// 老大哥在看你

func (b *BiliHandle)BigBrotherIsWatchingYou()  {

	scpage := b.B.MustPage("https://space.bilibili.com/20076571/dynamic")
	scpage.MustElement(".card")

	arr  := scpage.MustElements(".card")


	for k , v := range  arr {
		fmt.Println(k)
		fmt.Println(v.MustText())
		fmt.Println("-----------------------------------")
	}
}


//转发动态

func (b *BiliHandle)Forward()  {
	
}


//获取我关注的列表

func (b *BiliHandle) GetFocus()  {
	var rs []string
	spacePage := b.B.MustPage("http://space.bilibili.com")
	spacePage.MustElement(".n-statistics a").MustClick()
	spacePage.MustElement(".list-item")

	for{
		arr := spacePage.MustElements(".list-item")

		for _ , v := range  arr {
			name , _ := v.MustElement(".fans-name").Text()
			rs = append(rs ,name)
		}

		if len(arr) == 20 {
			spacePage.MustElement(".be-pager-next").MustClick()
			time.Sleep(time.Second)
		}else {
			fmt.Println("ending...")
			spacePage.Close()
			break
		}
	}
	fmt.Println(rs)
	fmt.Printf( "一共关注人数为 %d" , len(rs))
}

//直播项目签到

func LiveSignIn  (page *rod.Page) {
	page.MustElementR("a" ,"直播").MustClick()
	fmt.Println("bark")
}