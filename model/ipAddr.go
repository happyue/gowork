package model

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

//ServerIPInfo 全局ServerIPInfo
var ServerIPInfo IPInfo

//IPInfo 获取ip信息，外网ip，内网ip，所在城市，运营商。通过太平洋网站接口。
type IPInfo struct {
	IntranetIP   string        `json:"intranet_ip"`
	ExternalIP   string        `json:"external_ip"`
	PconlineInfo *PconlineInfo `json:"pconlineInfo"`
}

//PconlineInfo 太平洋网站信息接口结构体
type PconlineInfo struct {
	IP          string `json:"ip"`
	Pro         string `json:"pro"`
	ProCode     string `json:"proCode"`
	City        string `json:"city"`
	CityCode    string `json:"cityCode"`
	Region      string `json:"region"`
	RegionCode  string `json:"regionCode"`
	Addr        string `json:"addr"`
	RegionNames string `json:"regionNames"`
	Err         string `json:"err"`
}

//GetIPInfo 获取ip信息，外网ip，内网ip，所在城市，运营商。通过太平洋网站接口。
func GetIPInfo() error {
	// ipInfo := &IPInfo{}
	pconlineInfo := &PconlineInfo{}
	var err error

	if ServerIPInfo.ExternalIP, err = GetExternal(); err != nil {
		return err
	}

	url := "http://whois.pconline.com.cn/ipJson.jsp?ip=" + ServerIPInfo.ExternalIP + "&json=true"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder()))
	// out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// log.Debug("i`m error!!!!!!1")
		return err
	}

	// out, err = simplifiedchinese.GBK.NewEncoder().Bytes(out)
	// if err != nil {
	// 	log.Debug("i`m error!!!!!!1")
	// 	return nil, err
	// }

	// out = []byte(mahonia.NewEncoder("gbk").ConvertString(string(out)))
	if err := json.Unmarshal(out, &pconlineInfo); err != nil {
		return err
	}
	// fmt.Printf("%v\n", pconlineInfo)
	ServerIPInfo.PconlineInfo = pconlineInfo
	ServerIPInfo.IntranetIP = GetIntranetIP()

	return nil
}

//GetExternal 获取外网ip 通过http://myexternalip.com/raw接口
func GetExternal() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)

	return strings.Replace(string(content), "\n", "", -1), nil
}

//GetIntranetIP 获取内网ip
func GetIntranetIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

// func GetIntranetIp() {
// 	addrs, err := net.InterfaceAddrs()

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	for _, address := range addrs {

// 		// 检查ip地址判断是否回环地址
// 		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
// 			if ipnet.IP.To4() != nil {
// 				fmt.Println("ip:", ipnet.IP.String())
// 			}

// 		}
// 	}
// }

// func TabaoAPI(ip string) *IPInfo {
// 	url := "http://ip.taobao.com/service/getIpInfo.php?ip="
// 	url += ip

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil
// 	}
// 	defer resp.Body.Close()

// 	out, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil
// 	}
// 	var result IPInfo
// 	if err := json.Unmarshal(out, &result); err != nil {
// 		return nil
// 	}

// 	return &result
// }

// func inet_ntoa(ipnr int64) net.IP {
// 	var bytes [4]byte
// 	bytes[0] = byte(ipnr & 0xFF)
// 	bytes[1] = byte((ipnr >> 8) & 0xFF)
// 	bytes[2] = byte((ipnr >> 16) & 0xFF)
// 	bytes[3] = byte((ipnr >> 24) & 0xFF)

// 	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
// }

// func inet_aton(ipnr net.IP) int64 {
// 	bits := strings.Split(ipnr.String(), ".")

// 	b0, _ := strconv.Atoi(bits[0])
// 	b1, _ := strconv.Atoi(bits[1])
// 	b2, _ := strconv.Atoi(bits[2])
// 	b3, _ := strconv.Atoi(bits[3])

// 	var sum int64

// 	sum += int64(b0) << 24
// 	sum += int64(b1) << 16
// 	sum += int64(b2) << 8
// 	sum += int64(b3)

// 	return sum
// }

// func IpBetween(from net.IP, to net.IP, test net.IP) bool {
// 	if from == nil || to == nil || test == nil {
// 		fmt.Println("An ip input is nil") // or return an error!?
// 		return false
// 	}

// 	from16 := from.To16()
// 	to16 := to.To16()
// 	test16 := test.To16()
// 	if from16 == nil || to16 == nil || test16 == nil {
// 		fmt.Println("An ip did not convert to a 16 byte") // or return an error!?
// 		return false
// 	}

// 	if bytes.Compare(test16, from16) >= 0 && bytes.Compare(test16, to16) <= 0 {
// 		return true
// 	}
// 	return false
// }

// func IsPublicIP(IP net.IP) bool {
// 	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
// 		return false
// 	}
// 	if ip4 := IP.To4(); ip4 != nil {
// 		switch true {
// 		case ip4[0] == 10:
// 			return false
// 		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
// 			return false
// 		case ip4[0] == 192 && ip4[1] == 168:
// 			return false
// 		default:
// 			return true
// 		}
// 	}
// 	return false
// }
