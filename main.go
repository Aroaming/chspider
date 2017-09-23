package main

import (
	"chspider/base"
	"chspider/data"
	"chspider/spider"
	"fmt"
	"io/ioutil"
	"strings"

	_ "os"

	"github.com/PuerkitoBio/goquery"
)

var stopr []data.Stoper

func osage(programName string) {
	fmt.Printf(
		`
osage:
	chspider [configure file]

eg: 
	chspider conf/conf.xml
 `)
}

var (
	url = "http://jandan.net/ooxx"

	pagUrl = "http://jandan.net/ooxx/page-%d"
)

func main() {

	config := base.Confige("./conf/config.xml")

	fmt.Println(config.Log.File)

	logger := base.Mylog()
	//CInit()
	defer Release()

	client, err := spider.NewSpider(nil)
	if err != nil {
		fmt.Println(err)
	}
	client.SetUrl(url)     //设置URL
 	data, e := client.Go() //获取数据

	if e != nil {
		logger.Debug("获取数据失败")
	}
	doc, _ := base.QueryBytes(data)
	temp := doc.Find(".current-comment-page").Text()
	//查询页数
	pagenum := strings.Replace(strings.Split(temp, "]")[0], "[", "", -1)
	logger.Info(pagenum)

	//	fmt.Println(pagenum)
	num := base.Atoi(pagenum)
	spider.MkdirVendor(num) //循环创建文件夹

	//for 循环抓取
	for i := num; i >= 0; i-- {
		index := fmt.Sprintf(pagUrl, i) //循环拼接需要查询的url
		client.SetUrl(index)
		data, e = client.Go() //获取数据
		if e != nil {
			logger.Errorf("开始抓取数据，错误页[%d]：%d", i)
			continue
		}
		logger.Infof("抓取数据地址[ %s ] 完成!", index)

		doc, _ = base.QueryBytes(data)
		doc.Find(".view_img_link").Each(func(num int, node *goquery.Selection) {
			imgurl, ok := node.Attr("href")
			if !ok {
				return
			}

			//去重设置
			temp := strings.Split(imgurl, ".")
			tempnum := len(temp)
			if tempnum <= 1 {
				return
			}
			//后缀
			houzhui := temp[tempnum-1]
			//filename
			filename := base.Md5(imgurl) + "." + houzhui

			//大本营路径
			filedir := "./pic" + "/" + filename
			//页数分图
			hashfiledir := "./pic" + "/" + base.Itoa(i) + "/" + filename

			//每次判断是否大本营是否存在
			exist := base.FileExist(filedir)
			//hash 是否存在
			exist2 := base.FileExist(hashfiledir)

			if !exist2 && exist {
				if !exist2 {
					//文件中读取数据并返回文件的内容
					temp, e := ioutil.ReadFile(filedir)
					if e != nil {
						return
					}

					//写,
					ioutil.WriteFile(hashfiledir, temp, 0777)
					return
				}
				return
			}
			if exist2 {
				return
			}

			if strings.HasPrefix(imgurl, "//") {
				imgurl = "http:" + imgurl
			}

			//开始抓取
			client.SetUrl(imgurl).SetRefer(index)

			data, e = client.Go()
			if e != nil {
				logger.Error("抓取图片 ", imgurl, " 错误")
			}
			logger.Error("抓取图片 ", imgurl, " 成功")

			ioutil.WriteFile(hashfiledir, data, 0777)
		})

	}

}

//func CInit() {
//	//创建my_log单实例
//	mylog := base.Mylog()
//	//	fmt.Println(mylog.Log.File)
//	stopr = append(stopr, mylog)

//	//	mylog.Debugf("tesst", "aaaa")
//}

func Release() {
	for _, v := range stopr {
		v.Stop()
	}
}
