package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("bili.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	contentBytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	uidList := strings.Split(string(contentBytes), "\n")
	for _, uidLine := range uidList {
		time.Sleep(time.Second)
		uid, err := strconv.ParseInt(uidLine, 10, 64)
		if err != nil {
			//fmt.Println("error: ", err, "skip this line")
			continue
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/card?mid="+fmt.Sprint(uid), nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("authority", "api.bilibili.com")
		req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
		req.Header.Set("cache-control", "max-age=0")
		//req.Header.Set("cookie", "buvid3=FEF6BB3A-1C34-7426-99E1-990F3C35DC7686095infoc; b_nut=1668787786; i-wanna-go-back=-1; _uuid=2C187453-A83A-24A5-5783-539D727E10D7194147infoc; buvid4=39C1B666-761D-0222-4133-036FE754F58201240-022111900-eAdXW4PqOliYRZl9lmWbWQ%3D%3D; PVID=1; fingerprint=818f3f46e592a502624c7fcc4848f72d; buvid_fp_plain=undefined; b_ut=5; CURRENT_FNVAL=4048; rpdid=|(JY~uJ))k)~0J'uYY)~)m|Y|; buvid_fp=818f3f46e592a502624c7fcc4848f72d; hit-new-style-dyn=0; hit-dyn-v2=1; sid=7afavwpx; CURRENT_PID=487c5300-c7c7-11ed-bbc0-6be6a96b7cd2; b_lsid=97267DFB_188DC9D65CD; bsource=search_google")
		req.Header.Set("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"macOS"`)
		req.Header.Set("sec-fetch-dest", "document")
		req.Header.Set("sec-fetch-mode", "navigate")
		req.Header.Set("sec-fetch-site", "none")
		req.Header.Set("sec-fetch-user", "?1")
		req.Header.Set("upgrade-insecure-requests", "1")
		req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%s\n", bodyText)
		fansCount(string(bodyText))
	}
}

type BilibiliResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Card struct {
			Mid         string `json:"mid"`
			Name        string `json:"name"`
			Approve     bool   `json:"approve"`
			Sex         string `json:"sex"`
			Rank        string `json:"rank"`
			Face        string `json:"face"`
			FaceNft     int    `json:"face_nft"`
			FaceNftType int    `json:"face_nft_type"`
			DisplayRank string `json:"DisplayRank"`
			Regtime     int    `json:"regtime"`
			Spacesta    int    `json:"spacesta"`
			Birthday    string `json:"birthday"`
			Place       string `json:"place"`
			Description string `json:"description"`
			Article     int    `json:"article"`
			Attentions  []any  `json:"attentions"`
			Fans        int    `json:"fans"`
			Friend      int    `json:"friend"`
			Attention   int    `json:"attention"`
			Sign        string `json:"sign"`
			LevelInfo   struct {
				CurrentLevel int `json:"current_level"`
				CurrentMin   int `json:"current_min"`
				CurrentExp   int `json:"current_exp"`
				NextExp      int `json:"next_exp"`
			} `json:"level_info"`
			Pendant struct {
				Pid               int    `json:"pid"`
				Name              string `json:"name"`
				Image             string `json:"image"`
				Expire            int    `json:"expire"`
				ImageEnhance      string `json:"image_enhance"`
				ImageEnhanceFrame string `json:"image_enhance_frame"`
			} `json:"pendant"`
			Nameplate struct {
				Nid        int    `json:"nid"`
				Name       string `json:"name"`
				Image      string `json:"image"`
				ImageSmall string `json:"image_small"`
				Level      string `json:"level"`
				Condition  string `json:"condition"`
			} `json:"nameplate"`
			Official struct {
				Role  int    `json:"role"`
				Title string `json:"title"`
				Desc  string `json:"desc"`
				Type  int    `json:"type"`
			} `json:"Official"`
			OfficialVerify struct {
				Type int    `json:"type"`
				Desc string `json:"desc"`
			} `json:"official_verify"`
			Vip struct {
				Type       int   `json:"type"`
				Status     int   `json:"status"`
				DueDate    int64 `json:"due_date"`
				VipPayType int   `json:"vip_pay_type"`
				ThemeType  int   `json:"theme_type"`
				Label      struct {
					Path                  string `json:"path"`
					Text                  string `json:"text"`
					LabelTheme            string `json:"label_theme"`
					TextColor             string `json:"text_color"`
					BgStyle               int    `json:"bg_style"`
					BgColor               string `json:"bg_color"`
					BorderColor           string `json:"border_color"`
					UseImgLabel           bool   `json:"use_img_label"`
					ImgLabelURIHans       string `json:"img_label_uri_hans"`
					ImgLabelURIHant       string `json:"img_label_uri_hant"`
					ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
					ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				Role               int    `json:"role"`
				AvatarSubscriptURL string `json:"avatar_subscript_url"`
				TvVipStatus        int    `json:"tv_vip_status"`
				TvVipPayType       int    `json:"tv_vip_pay_type"`
				TvDueDate          int    `json:"tv_due_date"`
				VipType            int    `json:"vipType"`
				VipStatus          int    `json:"vipStatus"`
			} `json:"vip"`
			IsSeniorMember int `json:"is_senior_member"`
		} `json:"card"`
		Following    bool `json:"following"`
		ArchiveCount int  `json:"archive_count"`
		ArticleCount int  `json:"article_count"`
		Follower     int  `json:"follower"`
		LikeNum      int  `json:"like_num"`
	} `json:"data"`
}

func fansCount(contentStr string) int {
	var content BilibiliResp
	err := json.Unmarshal([]byte(contentStr), &content)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	fmt.Println(content.Data.Card.Fans)
	return 0
}
