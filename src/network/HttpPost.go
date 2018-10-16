package network

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"encoding/json"
	"encoding/base64"
)
/*图片识别文字*/
func HttpDo(filePath string) (str string,err error) {
	//图片转换base64
	base64img, err := TurnToBase64(filePath)
	if err != nil {
		fmt.Println("img turn err = ", err)
		return str,err
	}

	//构建httpClient
	client := &http.Client{}
	//鉴权token，有效期一个月，start:2018/9/18
	accessToken := "24.39e76639c0acfd0ea01849a4f908366e.2592000.1539828022.282335-14219223"
	//百度云文字识别api
	imgUrl := "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate" + "?access_token=" + accessToken
	//请求体body参数
	var bodyInfo = url.Values{}
	bodyInfo.Add("image_type", "BASE64")
	bodyInfo.Add("image", base64img)
	bodyInfo.Add("user_id", "1234")
	bodyInfo.Add("group_id", "123")
	data := bodyInfo.Encode()
	//构造request请求
	request, err := http.NewRequest("POST", imgUrl, strings.NewReader(data))

	if err != nil {
		fmt.Println("request create err:", err)
		return str,err
	}
	//设置请求头部
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//发送请求，返回结果
	response, err := client.Do(request)
	//使用完毕关闭流
	defer response.Body.Close()
	//读取返回结果json
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return str,err
	}
	//构造结构体获取json数据拼接所有文字
	resultInfo := ResultInfo{}
	err = json.Unmarshal(body, &resultInfo)
	if err != nil {
		return str,err
	}
	for _, word := range resultInfo.WordsResult {
		str += word.Words
	}
	return str,err
}

func TurnToBase64(filePath string) (base64img string, err error) {
	ff, err := ioutil.ReadFile(filePath)
	base64img = base64.StdEncoding.EncodeToString(ff)
	return base64img, err
}

type ResultInfo struct {
	LogId          int64         `json:"log_id"`
	WordsResultNum int           `json:"words_result_num"`
	WordsResult    []WordsResult `json:"words_result"`
}

type WordsResult struct {
	Location struct {
		Width  int `json:"width"`
		Top    int `json:"top"`
		Left   int `json:"left"`
		Height int `json:"height"`
	} 			   `json:"location"`
	Words string   `json:"words"`
}
