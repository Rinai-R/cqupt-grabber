package grab

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/LgoLgo/cqupt-grabber/model"
)

type SecGrabber struct {
	wg sync.WaitGroup
}

func (g *SecGrabber) SingleRob(cookie string, load model.SecCourseData) string {
	client := &http.Client{}
	data := url.Values{}
	data.Add("isOver", strconv.Itoa(load.IsOver))
	data.Add("isSksy", strconv.Itoa(load.IsSksy))
	data.Add("isTyfx", strconv.Itoa(load.IsTyfx))
	data.Add("jxb", load.Jxb)
	data.Add("kcbh", load.Kcbh)
	data.Add("kchType", load.KchType)
	data.Add("kclb", load.Kclb)
	data.Add("kcmc", load.Kcmc)
	data.Add("nj", load.Nj)
	data.Add("sdPrint", load.SdPrint)
	data.Add("syxs", strconv.Itoa(load.Syxs))
	data.Add("teacher", load.Teacher)
	data.Add("xf", load.Xf)
	data.Add("xkzt", strconv.Itoa(load.Xkzt))
	data.Add("zym", load.Zym)
	if load.SubJxb != "" {
		data.Add("subJxb", load.SubJxb)
	}

	// 注意：数组要用多次 Add(key, value)
	for _, sd := range load.Sd {
		data.Add("sd[]", sd)
	}
	for _, zc := range load.Zc {
		data.Add("zc[]", zc)
	}

	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST", "http://xk1.cqupt.edu.cn/xkPost.php", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Origin", "http://xk1.cqupt.edu.cn")
	req.Header.Set("Referer", "http://xk1.cqupt.edu.cn/xuanke.php")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(resp.Body)
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var Response model.Response
	err = json.Unmarshal(bodyText, &Response)
	if err != nil {
		log.Fatal(err)
	}
	return Response.Info
}

// 循环抢课
func (g *SecGrabber) LoopRob(cookie string, loads []model.SecCourseData) {
	for i := 1; ; i++ {
		log.Printf("第%d次抢课开始", i)
		for j, load := range loads {
			j += 1

			info := g.SingleRob(cookie, load)
			if info == "ok" {
				log.Printf("课程%d：%s\n", j, info)
				goto ok
			} else {
				log.Printf("课程%d：%s\n", j, info)
			}
			time.Sleep(250 * time.Millisecond)
		}
		log.Printf("第%d次抢课失败\n\n", i)
		time.Sleep(250 * time.Millisecond)
	}
ok:
	log.Println("抢课成功")
}

// LoopRobWithCustomTime 循环抢课，支持多个课程同时抢，支持自定义时间。不建议使用
// 传入一个 cookie 和一个 load 切片以及自定义时间
func (g *SecGrabber) LoopRobWithCustomTime(cookie string, loads []model.SecCourseData, duration float64) {
	for i := 1; ; i++ {
		log.Printf("第%d次抢课开始", i)
		for j, load := range loads {
			j += 1
			// 调用SingleRob进行循环抢课
			info := g.SingleRob(cookie, load)
			if info == "ok" {
				log.Printf("课程%d：%s\n", j, info)
				goto ok
			} else {
				log.Printf("课程%d：%s\n", j, info)
			}
			time.Sleep(time.Duration(duration*1000) * time.Millisecond)
		}
		log.Printf("第%d次抢课失败\n\n", i)
		time.Sleep(time.Duration(duration*1000) * time.Millisecond)
	}
ok:
	log.Println("抢课成功")
}

func (g *SecGrabber) highConcurrencySingleRob(cookie string, load model.SecCourseData, j int) {
	j += 1
	log.Printf("协程%d开启\n", j)
	for i := 1; ; i++ {
		log.Printf("第%d次抢课开始", i)
		// 调用SingleRob进行循环抢课
		info := g.SingleRob(cookie, load)
		if info == "ok" {
			log.Printf(info)
			g.wg.Done()
			return
		} else {
			log.Printf("课程%d：%s\n", j, info)
		}
		log.Printf("第%d次抢课失败\n\n", i)
		time.Sleep(250 * time.Millisecond)
	}
}

// LoopRobWithHighConcurrency 高并发抢课
func (g *SecGrabber) LoopRobWithHighConcurrency(cookie string, loads []model.SecCourseData) {
	for i, load := range loads {
		g.wg.Add(1)
		// 调用SingleRob进行循环抢课
		go g.highConcurrencySingleRob(cookie, load, i)
	}
	g.wg.Wait()
	log.Println("抢课成功")
}
