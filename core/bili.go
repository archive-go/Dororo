package core

import (
	"Dororo/utils"
	"encoding/json"
	"fmt"
	"github.com/go-rod/rod"
	"math/rand"
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

	childPage.MustElement(".bilibili-player-upinfo-span.restart")

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

type Dynamic struct {
	Time string `json:"time"`
	Content string `json:"content"`
}
// 老大哥在看你

func (b *BiliHandle)BigBrotherIsWatchingYou(id string)  {

	var  dynamicArr []Dynamic
	var dynamicArrItme Dynamic

	scpage := b.B.MustPage(fmt.Sprintf("https://space.bilibili.com/%s/dynamic" ,id))
	for {

		scpage.Reload()
		dynamicArr = nil
		scpage.MustElement(".main-content")

		arr  := scpage.MustElements(".main-content")


		for _ , v := range  arr {

			dyTime := v.MustElement(".detail-link.tc-slate").MustText()
			content := v.MustElement(".card-content").MustText()

			dynamicItem := Dynamic{
				Time:    dyTime,
				Content: content,
			}

			dynamicArr = append(dynamicArr , dynamicItem)

		}
		//程序第一次运行
		if(dynamicArrItme.Time == "" ){

			dynamicArrItme = dynamicArr[0]

			utils.WxSendMsg(string(`老大哥已经开始看着他啦 ， 最后一条动态更新为` + dynamicArrItme.Time))
			fmt.Println("fir")

		}


		if dynamicArrItme.Time != dynamicArr[0].Time  || dynamicArrItme.Content != dynamicArr[0].Content {
			fmt.Println("更新")
			utils.WxSendMsg(fmt.Sprintf("老大哥发现Ta的动态已经更新啦 ！！！ 时间为 ： %s , 内容为 : %s "  , dynamicArr[0].Time ,dynamicArr[0].Content))
			dynamicArrItme = dynamicArr[0]
		}

		rand.Seed(time.Now().UnixNano())


		time.Sleep(time.Second * 3 )

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

//我的直播关注列表

func (b *BiliHandle) GetLiveFocus (){

	type LiveItme struct {
		Name string `json:"name"`
		Status string `json:"status"`
	}

	var rsString string

	cpage := b.B.MustPage("https://link.bilibili.com/p/center/index#/user-center/follow/1")

	for{
		var liveArr []LiveItme
		cpage.Reload()
		cpage.MustElement(".favourite-card")
		arr := cpage.MustElements(".favourite-card")

		for _ , v := range arr {
			liveArr = append(liveArr , LiveItme{
				Name:   v.MustElement(".anchor-name").MustText(),
				Status: v.MustElement(".anchor-status").MustText(),
			})
		}
		rsByte , err  := json.Marshal(liveArr)
		if err != nil {
			fmt.Println(err)
		}
		if rsString == "" {

			rsString = string(rsByte)
			utils.WxSendMsg(string(`直播列表正在监听` + rsString))
			fmt.Println("fir")

		}

		if rsString != string(rsString) {
			rsString = string(rsByte)
			utils.WxSendMsg(string(`直播列表已经更新` + rsString))
		}

		rand.Seed(time.Now().UnixNano())

		r := rand.Intn(5) + 5
		fmt.Println(r)
		time.Sleep(time.Minute * time.Duration(r))
	}

}


//直播项目签到

func LiveSignIn  (page *rod.Page) {
	page.MustElementR("a" ,"直播").MustClick()
	fmt.Println("bark")
}
