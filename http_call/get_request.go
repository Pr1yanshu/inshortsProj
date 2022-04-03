package http_call

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"inshortsProj/constant"
	"io/ioutil"
	"net"
	"net/http"
	"runtime/debug"
	"time"
)

const (
	GetHttpVertical = "Http_get_middleware"
	HttpMiddleware  = "Http_middleware"
)

func MakeGetHttpCall(apiUrl string, headers map[string]interface{}, timeoutSeconds int) ([]byte, error) {

	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic in MakeGetHttpCall ,General Error: " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			fmt.Println(ErrorString)
		}
	}()

	req, err := http.NewRequest("GET", apiUrl, nil)
	for key, value := range headers {
		if stringValue, ok := value.(string); ok {
			req.Header.Set(key, stringValue)
		} else {
			ErrorString := "Error in adding headers to https object for  API :" + apiUrl + " and for header = " + fmt.Sprint(headers)
			fmt.Println(ErrorString, "\n ")
		}
	}

	// making an http client using pester
	dialer := &net.Dialer{
		Timeout: time.Duration(constant.DIALER_TIME_OUT) * time.Millisecond,
	}
	client := pester.New()
	client.KeepLog = true
	client.Concurrency = 1
	client.MaxRetries = 3
	client.Backoff = pester.ExponentialBackoff
	client.Transport = &http.Transport{
		MaxIdleConnsPerHost: constant.MaxIdleConnectionsPerHost,
		DialContext:         dialer.DialContext,
	}
	client.Timeout = time.Duration(time.Duration(timeoutSeconds) * time.Second)
	response, err := client.Do(req)
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	//config.PesterInfoForKibana(context, fmt.Sprint(client.LogString()))

	emptyData := make([]byte, 0)
	if err != nil {
		return emptyData, err
	} else if response.StatusCode != 200 {
		ErrorString := fmt.Sprintf("Status code is non 200 , Status code received = %d", response.StatusCode)
		fmt.Println(ErrorString)
		return emptyData, errors.New(ErrorString)
	}

	if body, err := ioutil.ReadAll(response.Body); err != nil {
		ErrorString := "Error in reading data from response.Body " + fmt.Sprint(err)
		fmt.Println(ErrorString)
		return emptyData, err
	} else {
		//fmt.Println("Response for api :"+apiUrl+" is "+string(body))
		return body, nil
	}

}
