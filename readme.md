# 各平台情况说明

## bilibili

下面这个链接可以获取，fans 字段即为粉丝数：
https://api.bilibili.com/x/web-interface/card?mid=9618426

```shell
curl 'https://api.bilibili.com/x/web-interface/card?mid=9618426' \
  -H 'authority: api.bilibili.com' \
  -H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7' \
  -H 'accept-language: zh-CN,zh;q=0.9' \
  -H 'cache-control: max-age=0' \
  -H $'cookie: buvid3=FEF6BB3A-1C34-7426-99E1-990F3C35DC7686095infoc; b_nut=1668787786; i-wanna-go-back=-1; _uuid=2C187453-A83A-24A5-5783-539D727E10D7194147infoc; buvid4=39C1B666-761D-0222-4133-036FE754F58201240-022111900-eAdXW4PqOliYRZl9lmWbWQ%3D%3D; PVID=1; fingerprint=818f3f46e592a502624c7fcc4848f72d; buvid_fp_plain=undefined; b_ut=5; CURRENT_FNVAL=4048; rpdid=|(JY~uJ))k)~0J\'uYY)~)m|Y|; buvid_fp=818f3f46e592a502624c7fcc4848f72d; hit-new-style-dyn=0; hit-dyn-v2=1; sid=7afavwpx; CURRENT_PID=487c5300-c7c7-11ed-bbc0-6be6a96b7cd2; b_lsid=97267DFB_188DC9D65CD; bsource=search_google' \
  -H 'sec-ch-ua: "Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: document' \
  -H 'sec-fetch-mode: navigate' \
  -H 'sec-fetch-site: none' \
  -H 'sec-fetch-user: ?1' \
  -H 'upgrade-insecure-requests: 1' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36' \
  --compressed ;
curl 'https://api.bilibili.com/favicon.ico' \
  -H 'authority: api.bilibili.com' \
  -H 'accept: image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8' \
  -H 'accept-language: zh-CN,zh;q=0.9' \
  -H $'cookie: buvid3=FEF6BB3A-1C34-7426-99E1-990F3C35DC7686095infoc; b_nut=1668787786; i-wanna-go-back=-1; _uuid=2C187453-A83A-24A5-5783-539D727E10D7194147infoc; buvid4=39C1B666-761D-0222-4133-036FE754F58201240-022111900-eAdXW4PqOliYRZl9lmWbWQ%3D%3D; PVID=1; fingerprint=818f3f46e592a502624c7fcc4848f72d; buvid_fp_plain=undefined; b_ut=5; CURRENT_FNVAL=4048; rpdid=|(JY~uJ))k)~0J\'uYY)~)m|Y|; buvid_fp=818f3f46e592a502624c7fcc4848f72d; hit-new-style-dyn=0; hit-dyn-v2=1; sid=7afavwpx; CURRENT_PID=487c5300-c7c7-11ed-bbc0-6be6a96b7cd2; b_lsid=97267DFB_188DC9D65CD; bsource=search_google' \
  -H 'referer: https://api.bilibili.com/x/web-interface/card?mid=9618426' \
  -H 'sec-ch-ua: "Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: image' \
  -H 'sec-fetch-mode: no-cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36' \
  --compressed
```

## douyin
抖音获取粉丝数需要进驻 openapi 开放平台，目前还在审核中：
https://developer.open-douyin.com/apply?service=

## weibo

https://weibo.com/ajax/profile/info?uid=3868613846

## 小红书


## 快手

参考：
https://github.com/reader0421/FetchFans/tree/master

