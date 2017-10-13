// hesan project main.go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
	//	"strings"
	"strconv"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Config struct {
	xLineNum int
	fileName string
	left     float64
	right    float64
}
type Point struct {
	index int
	x     float64
	y     float64
	group int
}
type Side struct {
	A       Point
	B       Point
	SideLen float64
}

var result = make([]string, 0) //统计结果
var config Config

func main() {
	var pointList = make([]Point, 0)     //所有点的集合
	var sideListAll = make([]Side, 0)    //存放所有的边
	var sideListTarget = make([]Side, 0) //存放全部过滤好的边
	var treeLen float64                  //树的总长
	var treeAverage float64              //树的平均长度
	var allLen float64                   //算数总长
	var allAverage float64               //算数平均长度
	t1 := time.Now()
	config = readConfig()
	fmt.Printf("config is %v \n", config)
	readData(&pointList)
	fmt.Printf("晶粒的总数为 = %d \n", len(pointList))
	result = append(result, fmt.Sprintf("晶粒的总数为 = %d \n", len(pointList)))
	//	fmt.Print(pointList)
	rendSide(&sideListAll, &pointList, &config)
	//对边进行排序
	for i := 0; i < len(sideListAll); i++ {
		for j := 0; j < len(sideListAll)-i-1; j++ {
			if sideListAll[j].SideLen > sideListAll[j+1].SideLen {
				temp := sideListAll[j]
				//				fmt.Println(temp)
				sideListAll[j] = sideListAll[j+1]
				sideListAll[j+1] = temp
			}
		}
	}
	var flag = make([]int, len(pointList))
	for i, p := range pointList {
		flag[i] = i
		p.index = i
	}
	for _, side := range sideListAll {
		foot1 := flag[side.A.index]
		foot2 := flag[side.B.index]
		if foot1 != foot2 {
			sideListTarget = append(sideListTarget, side)
			for j, value := range flag {
				if value == foot2 {
					flag[j] = foot1
				}
			}
		}
		allLen += side.SideLen
	}
	allAverage = allLen / float64(len(sideListAll))
	fmt.Printf("所有边长总间距 = %f \n", allLen)
	result = append(result, fmt.Sprintf("所有边长总间距 = %f \n", allLen))
	fmt.Printf("所有边长的平均间距 = %f \n", allAverage)
	result = append(result, fmt.Sprintf("所有边长的平均间距 = %f \n", allAverage))
	fmt.Printf("总共采集的边数量 = %d \n", len(sideListAll))
	result = append(result, fmt.Sprintf("总共采集的边数量 = %d \n", len(sideListAll)))
	fmt.Printf("最小生产树采集的边数量 = %d \n", len(sideListTarget))
	result = append(result, fmt.Sprintf("最小生产树采集的边数量 = %d \n", len(sideListTarget)))
	for _, t := range sideListTarget {
		treeLen += t.SideLen
	}
	treeAverage = treeLen / (float64)(len(sideListTarget))
	fmt.Printf("树的总间距 = %f \n", treeLen)
	result = append(result, fmt.Sprintf("树的总间距 = %f \n", treeLen))
	fmt.Printf("树的平均间距 = %f \n", treeAverage)
	result = append(result, fmt.Sprintf("树的平均间距 = %f \n", treeAverage))

	resultOutPut()
	sideOutPut(&sideListAll)
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}

//获取最终符合数据的边
func rendSide(listAll *[]Side, pointList *[]Point, config *Config) {
	var lineMax = config.xLineNum
	//对点进行排序，x
	var totalCount = len(*pointList)
	var sample = make([]Side, 0)  //初始化样本集合
	var sample2 = make([]Side, 0) //第二次样本集合
	var lastSideLen float64       //记录样本前一个的长度
	fmt.Printf("每一行晶粒数目大约为 %d \n", lineMax)
	result = append(result, fmt.Sprintf("每一行晶粒数目大约为 %d \n", lineMax))
	for i, p := range *pointList {
		//计算前后两个点的距离
		if i < totalCount-1 {
			next := (*pointList)[i+1]
			x2 := (p.x - next.x) * (p.x - next.x)
			y2 := (p.y - next.y) * (p.y - next.y)
			sideLen := math.Sqrt(x2 + y2)
			if i == 0 {
				lastSideLen = sideLen
			}
			if (math.Abs(lastSideLen-sideLen) / lastSideLen) > 1 {
				continue
			}
			side := Side{}
			side.A = p
			side.B = next
			side.SideLen = sideLen
			sample = append(sample, side)
			lastSideLen = sideLen
		}
	}
	//根据规则过滤并返回新的数据
	var totalLen float64
	var average1, average2 float64 //算术平均值
	var sampleMin1, sampleMin2 float64 = 1000, 1000
	var sampleMax1, sampleMax2 float64
	var realMin float64 = 1000
	var realMax float64
	var variance1, variance2 float64 //方差
	var varianceTotal float64
	for _, s := range sample {
		if sampleMin1 > s.SideLen {
			sampleMin1 = s.SideLen
		}
		if sampleMax1 < s.SideLen {
			sampleMax1 = s.SideLen
		}
		totalLen += s.SideLen
	}
	average1 = totalLen / (float64)(len(sample))
	for _, s := range sample {
		varianceTotal += (s.SideLen - average1) * (s.SideLen - average1)
	}
	variance1 = math.Sqrt(varianceTotal / (float64)(len(sample)))
	fmt.Printf("第一次样本标准差  = %f\n", variance1)
	fmt.Printf("第一次采集样本平均间距 average1 = %f\n", average1)
	result = append(result, fmt.Sprintf("第一次采集样本平均间距 average1 = %f\n", average1))
	fmt.Printf("第一次采集样本最小间距 sampleMin1 = %f\n", sampleMin1)
	result = append(result, fmt.Sprintf("第一次采集样本最小间距 sampleMin1 = %f\n", sampleMin1))
	fmt.Printf("第一次采集样本最大间距 sampleMax1 = %f\n", sampleMax1)
	result = append(result, fmt.Sprintf("第一次采集样本最大间距 sampleMax1 = %f\n", sampleMax1))
	//第二次采集样本数据
	var thresholdleft = average1 - variance1*1.5
	var thresholdright = average1 + variance1*3
	fmt.Printf("第二次样本过滤阈值 thresholdleft = %f\n", thresholdleft)
	result = append(result, fmt.Sprintf("第二次样本过滤阈值 thresholdleft = %f\n", thresholdleft))
	fmt.Printf("第二次样本过滤阈值 thresholdright = %f\n", thresholdright)
	result = append(result, fmt.Sprintf("第二次样本过滤阈值 thresholdright = %f\n", thresholdright))
	for _, s := range sample {
		if s.SideLen <= thresholdright && s.SideLen >= thresholdleft {
			sample2 = append(sample2, s)
		}
	}
	varianceTotal = 0
	totalLen = 0
	for _, s := range sample2 {
		if sampleMin2 > s.SideLen {
			sampleMin2 = s.SideLen
		}
		if sampleMax2 < s.SideLen {
			sampleMax2 = s.SideLen
		}
		totalLen += s.SideLen
	}
	average2 = totalLen / (float64)(len(sample2))
	for _, s := range sample2 {
		varianceTotal += (s.SideLen - average2) * (s.SideLen - average2)
	}
	variance2 = math.Sqrt(varianceTotal / (float64)(len(sample2)))
	fmt.Printf("第二次采集样本平均间距 average2 = %f\n", average2)
	result = append(result, fmt.Sprintf("第二次采集样本平均间距 average1 = %f\n", average2))
	fmt.Printf("第二次采集样本最小间距 sampleMin2 = %f\n", sampleMin2)
	result = append(result, fmt.Sprintf("第二次采集样本最小间距 sampleMin2 = %f\n", sampleMin2))
	fmt.Printf("第二次采集样本最大间距 sampleMax2 = %f\n", sampleMax2)
	result = append(result, fmt.Sprintf("第二次采集样本最大间距 sampleMax2 = %f\n", sampleMax2))
	fmt.Printf("第二次样本标准差  = %f\n", variance2)
	result = append(result, fmt.Sprintf("第二次样本标准差  = %f\n", variance2))

	result = append(result, fmt.Sprintf("左边界值 left = %f\n", average2-variance2*config.left))
	result = append(result, fmt.Sprintf("右边界值 right = %f\n", average2+variance2*config.right))
	fmt.Println(fmt.Sprintf("左边界值 left = %f\n", average2-variance2*config.left))
	fmt.Println(fmt.Sprintf("右边界值 right = %f\n", average2+variance2*config.right))
	//计算各种各样的边
	var ran = make([]int, 0)
	ran = append(ran, 1)
	for i := 0; i < lineMax; i++ {
		ran = append(ran, lineMax-i)
		ran = append(ran, lineMax+i)
	}
	for i, p := range *pointList {
		//计算前后两个点的距离
		//		var ran = [...]int{1, lineMax + 1, lineMax - 1, lineMax + 2, lineMax - 2, lineMax + 3, lineMax - 3, lineMax - 4, lineMax + 4, lineMax - 5, lineMax + 5,
		//			lineMax - 6, lineMax + 6, lineMax - 7, lineMax + 7, lineMax - 8, lineMax + 8}
		for _, step := range ran {
			realMin, realMax = getRangeData(i, totalCount, pointList, listAll, p, average2, variance2, step, realMin, realMax)
		}
	}
	fmt.Printf("最终最小间距 realMin = %f\n", realMin)
	result = append(result, fmt.Sprintf("最终最小间距 realMin = %f\n", realMin))
	fmt.Printf("最终最大间距 realMax = %f\n", realMax)
	result = append(result, fmt.Sprintf("最终最大间距 realMax = %f\n", realMax))
	fmt.Printf("最大间距 : 最小间距 = %f : 1 \n", realMax/realMin)
	result = append(result, fmt.Sprintf("最大间距 : 最小间距 = %f : 1 \n", realMax/realMin))

}
func getRangeData(i int, totalCount int, pointList *[]Point, listAll *[]Side, p Point, average float64, variance float64, step int, realMin float64, realMax float64) (float64, float64) {
	if i < totalCount-1-step {
		next := (*pointList)[i+step]
		x2 := (p.x - next.x) * (p.x - next.x)
		y2 := (p.y - next.y) * (p.y - next.y)
		sideLen := math.Sqrt(x2 + y2)
		side := Side{}
		side.A = p
		side.B = next
		side.SideLen = sideLen
		if sideLen <= (average+variance*config.right) && sideLen >= (average-variance*config.left) {
			if realMin > sideLen {
				realMin = sideLen
			}
			if realMax < sideLen {
				realMax = sideLen
			}
			*listAll = append(*listAll, side)
		}
	}
	return realMin, realMax
}

//从文件中读取数据
func readData(x *[]Point) {
	origin, err := os.Open(config.fileName + ".txt")
	var index int = 0
	var linNum int = 1
	var floatx, floaty float64
	check(err)
	defer origin.Close()
	trans := transform.NewReader(origin, simplifiedchinese.GBK.NewDecoder())
	buff := bufio.NewReader(trans)
	for {
		line, _, err := buff.ReadLine()
		if err != nil {
			break
		}
		myline := string(line)
		if linNum%2 == 1 { //奇数行代表x 偶数行代表y
			x0, _ := strconv.ParseFloat(myline, 32)
			floatx = x0
		} else {
			y0, _ := strconv.ParseFloat(myline, 32)
			floaty = y0
		}
		if floatx != 0 && floaty != 0 {
			tempPoint := Point{}
			tempPoint.x = floatx
			tempPoint.y = floaty
			tempPoint.index = index
			tempPoint.group = index
			*x = append(*x, tempPoint)
			floatx = 0
			floaty = 0
			index++
		}
		linNum++
	}
}

func check(e error) {
	if e != nil {
		fmt.Print("a")
		panic(e)
	}
}

//读取配置文件  一行大概的数目
func readConfig() Config {
	var config Config
	origin, err := os.Open("./config.txt")
	check(err)
	defer origin.Close()
	trans := transform.NewReader(origin, simplifiedchinese.GBK.NewDecoder())
	buff := bufio.NewReader(trans)
	for {
		line, _, err := buff.ReadLine()
		if err != nil {
			break
		}
		info := string(line)
		info = strings.TrimSpace(info)
		infos := strings.Split(info, ":")
		if len(infos) >= 2 && strings.HasPrefix(info, "xLineNum") {
			temp, err := strconv.Atoi(infos[1])
			if err == nil {
				config.xLineNum = temp
			}
		}
		if len(infos) >= 2 && strings.HasPrefix(info, "fileName") {
			config.fileName = infos[1]
		}
		if len(infos) >= 2 && strings.HasPrefix(info, "left") {
			temp, err := strconv.ParseFloat(infos[1], 32)
			if err == nil {
				config.left = temp
			}
		}
		if len(infos) >= 2 && strings.HasPrefix(info, "right") {
			temp, err := strconv.ParseFloat(infos[1], 32)
			if err == nil {
				config.right = temp
			}
		}
	}
	return config
}

// 结果导出数据
func resultOutPut() {
	var f *os.File
	var err1 error
	//	if checkFileIsExist(config.fileName + "_result.txt") { //如果文件存在
	//		//		os.Remove(config.fileName + "_result.txt")
	//		//		f, err1 = os.OpenFile(config.fileName+"_result.txt", os.O_WRONLY|os.O_RDONLY, 0666) //打开文件
	//		//		fmt.Println("文件存在")
	//	}
	f, err1 = os.Create(config.fileName + "_result.txt") //创建文件
	fmt.Println(err1)
	check(err1)
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, str := range result {
		w.WriteString(str)
		w.WriteString("\r\n")
		w.Flush()
	}

}

//边长输出
func sideOutPut(sides *[]Side) {
	var f *os.File
	var err1 error
	//	if checkFileIsExist(config.fileName + "_side.txt") { //如果文件存在
	//		os.Remove(config.fileName + "_result.txt")
	//		//		f, err1 = os.OpenFile(config.fileName+"_side.txt", os.O_WRONLY|os.O_RDONLY, 0666) //打开文件
	//		//		fmt.Println("文件存在")
	//	}
	f, err1 = os.Create(config.fileName + "_side.txt") //创建文件
	defer f.Close()
	check(err1)
	w := bufio.NewWriter(f)
	for _, side := range *sides {
		w.WriteString(fmt.Sprintf("%.4f", side.SideLen))
		w.WriteString("\r\n")
		w.Flush()
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
