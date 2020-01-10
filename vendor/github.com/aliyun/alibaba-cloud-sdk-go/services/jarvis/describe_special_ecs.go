package jarvis

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

// DescribeSpecialEcs invokes the jarvis.DescribeSpecialEcs API synchronously
// api document: https://help.aliyun.com/api/jarvis/describespecialecs.html
func (client *Client) DescribeSpecialEcs(request *DescribeSpecialEcsRequest) (response *DescribeSpecialEcsResponse, err error) {
	response = CreateDescribeSpecialEcsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSpecialEcsWithChan invokes the jarvis.DescribeSpecialEcs API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describespecialecs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSpecialEcsWithChan(request *DescribeSpecialEcsRequest) (<-chan *DescribeSpecialEcsResponse, <-chan error) {
	responseChan := make(chan *DescribeSpecialEcsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSpecialEcs(request)
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

// DescribeSpecialEcsWithCallback invokes the jarvis.DescribeSpecialEcs API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describespecialecs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSpecialEcsWithCallback(request *DescribeSpecialEcsRequest, callback func(response *DescribeSpecialEcsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSpecialEcsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSpecialEcs(request)
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

// DescribeSpecialEcsRequest is the request struct for api DescribeSpecialEcs
type DescribeSpecialEcsRequest struct {
	*requests.RpcRequest
	TargetIp   string `position:"Query" name:"TargetIp"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
	SourceCode string `position:"Query" name:"SourceCode"`
}

// DescribeSpecialEcsResponse is the response struct for api DescribeSpecialEcs
type DescribeSpecialEcsResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	EcsFound  bool    `json:"EcsFound" xml:"EcsFound"`
	Module    string  `json:"module" xml:"module"`
	EcsInfo   EcsInfo `json:"EcsInfo" xml:"EcsInfo"`
}

// CreateDescribeSpecialEcsRequest creates a request to invoke DescribeSpecialEcs API
func CreateDescribeSpecialEcsRequest() (request *DescribeSpecialEcsRequest) {
	request = &DescribeSpecialEcsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "DescribeSpecialEcs", "jarvis", "openAPI")
	return
}

// CreateDescribeSpecialEcsResponse creates a response to parse from DescribeSpecialEcs response
func CreateDescribeSpecialEcsResponse() (response *DescribeSpecialEcsResponse) {
	response = &DescribeSpecialEcsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}