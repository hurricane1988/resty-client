package utils

import (
	"github.com/go-resty/resty/v2"
)

// HttpResponse resty返回函数
func HttpResponse(err error, resp *resty.Response) (result map[string]interface{}) {
	// Explore response object
	// fmt.Println("Response Info:")
	// fmt.Println("  Error      :", err)
	// fmt.Println("  Status Code:", resp.StatusCode())
	// fmt.Println("  Status     :", resp.Status())
	// fmt.Println("  Proto      :", resp.Proto())
	// fmt.Println("  Time       :", resp.Time())
	// fmt.Println("  Received At:", resp.ReceivedAt())
	// fmt.Println("  Body       :\n", resp)
	// fmt.Println()
	var respBody = make(map[string]interface{})

	respBody["Error"] = err
	respBody["StatusCode"] = resp.StatusCode()
	respBody["Status"] = resp.Status()
	respBody["Proto"] = resp.Proto
	respBody["Time"] = resp.Time()
	respBody["ReceivedAt"] = resp.ReceivedAt()
	respBody["respBody"] = resp
	return respBody
}

// HttpTrace resty 打印trace信息
func HttpTrace(resp *resty.Response) (result map[string]interface{}) {
	ti := resp.Request.TraceInfo()
	// 组合结构体
	var traceBody = make(map[string]interface{})
	traceBody["DNSLookup"] = ti.DNSLookup
	traceBody["ConnTime"] = ti.ConnTime
	traceBody["TCPConnTime"] = ti.TCPConnTime
	traceBody["TLSHandshake"] = ti.TLSHandshake
	traceBody["ServerTime"] = ti.ServerTime
	traceBody["ResponseTime"] = ti.ResponseTime
	traceBody["TotalTime"] = ti.TotalTime
	traceBody["IsConnReused"] = ti.IsConnReused
	traceBody["ConnIdleTime"] = ti.ConnIdleTime
	traceBody["RequestAttempt"] = ti.RequestAttempt
	traceBody["RemoteAddr"] = ti.RemoteAddr.String()
	return traceBody
}
