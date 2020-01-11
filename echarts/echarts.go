package echarts

import (
	"github.com/go-echarts/go-echarts/charts"
	"log"
	"math/rand"
	"os"
)
var  nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}


func BarDemo() {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	_ = bar.Render(f)
}

func randInt() []int {
	cnt := len(nameItems)
	r := make([]int, 0)
	for i := 0; i < cnt; i++ {
		r = append(r, int(rand.Int63()) % 50)
	}
	return r
}
