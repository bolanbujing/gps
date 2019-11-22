package gps

import (
	"math"
)

const (
	pi           = 3.1415926535898
	earth_radius = 6378137
)

type GpsInfo struct {
	jd float64
	wd float64
}

// 求弧度
func radian(d float64) float64 {
	return d * pi / 180.0 //角度1˚ = π / 180
}

//计算距离

func GetDistance(a, b GpsInfo) float64 {
	radLat1 := radian(a.wd)
	radLat2 := radian(b.wd)
	i := radLat1 - radLat2
	j := radian(a.jd) - radian(b.jd)

	dst := 2 * math.Asin((math.Sqrt(math.Pow(math.Sin(i/2), 2) + math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(j/2), 2))))

	dst = dst * earth_radius
	dst = math.Round(dst*10000) / 10000
	return dst
}

/*
func main() {
	var a = GpsInfo{100.309052, 28.263772}
	var b = GpsInfo{121.293990, 28.215665}

	dst := GetDistance(a, b)
	fmt.Println("dst = ", dst) //dst = 9.281km
}
*/
