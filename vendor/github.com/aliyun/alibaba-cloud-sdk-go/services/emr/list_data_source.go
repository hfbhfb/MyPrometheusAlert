package emr

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

// ListDataSource invokes the emr.ListDataSource API synchronously
// api document: https://help.aliyun.com/api/emr/listdatasource.html
func (client *Client) ListDataSource(request *ListDataSourceRequest) (response *ListDataSourceResponse, err error) {
	response = CreateListDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataSourceWithChan invokes the emr.ListDataSource API asynchronously
// api document: https://help.aliyun.com/api/emr/listdatasource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDataSourceWithChan(request *ListDataSourceRequest) (<-chan *ListDataSourceResponse, <-chan error) {
	responseChan := make(chan *ListDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataSource(request)
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

// ListDataSourceWithCallback invokes the emr.ListDataSource API asynchronously
// api document: https://help.aliyun.com/api/emr/listdatasource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDataSourceWithCallback(request *ListDataSourceRequest, callback func(response *ListDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataSourceResponse
		var err error
		defer close(result)
		response, err = client.ListDataSource(request)
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

// ListDataSourceRequest is the request struct for api ListDataSource
type ListDataSourceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CreateFrom      string           `position:"Query" name:"CreateFrom"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Name            string           `position:"Query" name:"Name"`
	SourceType      string           `position:"Query" name:"SourceType"`
	Id              string           `position:"Query" name:"Id"`
	ProjectId       string           `position:"Query" name:"ProjectId"`
}

// ListDataSourceResponse is the response struct for api ListDataSource
type ListDataSourceResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Total          int            `json:"Total" xml:"Total"`
	PageSize       int            `json:"PageSize" xml:"PageSize"`
	PageNumber     int            `json:"PageNumber" xml:"PageNumber"`
	DataSourceList DataSourceList `json:"DataSourceList" xml:"DataSourceList"`
}

// CreateListDataSourceRequest creates a request to invoke ListDataSource API
func CreateListDataSourceRequest() (request *ListDataSourceRequest) {
	request = &ListDataSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListDataSource", "emr", "openAPI")
	return
}

// CreateListDataSourceResponse creates a response to parse from ListDataSource response
func CreateListDataSourceResponse() (response *ListDataSourceResponse) {
	response = &ListDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}