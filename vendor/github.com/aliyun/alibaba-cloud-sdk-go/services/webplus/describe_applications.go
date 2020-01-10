package webplus

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeApplications invokes the webplus.DescribeApplications API synchronously
// api document: https://help.aliyun.com/api/webplus/describeapplications.html
func (client *Client) DescribeApplications(request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
	response = CreateDescribeApplicationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApplicationsWithChan invokes the webplus.DescribeApplications API asynchronously
// api document: https://help.aliyun.com/api/webplus/describeapplications.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApplicationsWithChan(request *DescribeApplicationsRequest) (<-chan *DescribeApplicationsResponse, <-chan error) {
	responseChan := make(chan *DescribeApplicationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApplications(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeApplicationsWithCallback invokes the webplus.DescribeApplications API asynchronously
// api document: https://help.aliyun.com/api/webplus/describeapplications.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApplicationsWithCallback(request *DescribeApplicationsRequest, callback func(response *DescribeApplicationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApplicationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeApplications(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeApplicationsRequest is the request struct for api DescribeApplications
type DescribeApplicationsRequest struct {
	*requests.RoaRequest
	AppName        string           `position:"Query" name:"AppName"`
	AppId          string           `position:"Query" name:"AppId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	CategorySearch string           `position:"Query" name:"CategorySearch"`
	StackSearch    string           `position:"Query" name:"StackSearch"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	AppSearch      string           `position:"Query" name:"AppSearch"`
	EnvSearch      string           `position:"Query" name:"EnvSearch"`
}

// DescribeApplicationsResponse is the response struct for api DescribeApplications
type DescribeApplicationsResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	PageNumber   int          `json:"PageNumber" xml:"PageNumber"`
	PageSize     int          `json:"PageSize" xml:"PageSize"`
	TotalCount   int          `json:"TotalCount" xml:"TotalCount"`
	Applications Applications `json:"Applications" xml:"Applications"`
}

// CreateDescribeApplicationsRequest creates a request to invoke DescribeApplications API
func CreateDescribeApplicationsRequest() (request *DescribeApplicationsRequest) {
	request = &DescribeApplicationsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DescribeApplications", "/pop/v1/wam/application", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeApplicationsResponse creates a response to parse from DescribeApplications response
func CreateDescribeApplicationsResponse() (response *DescribeApplicationsResponse) {
	response = &DescribeApplicationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}