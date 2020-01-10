package trademark

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

// QueryQrCodeUploadStatus invokes the trademark.QueryQrCodeUploadStatus API synchronously
// api document: https://help.aliyun.com/api/trademark/queryqrcodeuploadstatus.html
func (client *Client) QueryQrCodeUploadStatus(request *QueryQrCodeUploadStatusRequest) (response *QueryQrCodeUploadStatusResponse, err error) {
	response = CreateQueryQrCodeUploadStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryQrCodeUploadStatusWithChan invokes the trademark.QueryQrCodeUploadStatus API asynchronously
// api document: https://help.aliyun.com/api/trademark/queryqrcodeuploadstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryQrCodeUploadStatusWithChan(request *QueryQrCodeUploadStatusRequest) (<-chan *QueryQrCodeUploadStatusResponse, <-chan error) {
	responseChan := make(chan *QueryQrCodeUploadStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryQrCodeUploadStatus(request)
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

// QueryQrCodeUploadStatusWithCallback invokes the trademark.QueryQrCodeUploadStatus API asynchronously
// api document: https://help.aliyun.com/api/trademark/queryqrcodeuploadstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryQrCodeUploadStatusWithCallback(request *QueryQrCodeUploadStatusRequest, callback func(response *QueryQrCodeUploadStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryQrCodeUploadStatusResponse
		var err error
		defer close(result)
		response, err = client.QueryQrCodeUploadStatus(request)
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

// QueryQrCodeUploadStatusRequest is the request struct for api QueryQrCodeUploadStatus
type QueryQrCodeUploadStatusRequest struct {
	*requests.RpcRequest
	OssKey   string `position:"Query" name:"OssKey"`
	FieldKey string `position:"Query" name:"FieldKey"`
	Uuid     string `position:"Query" name:"Uuid"`
}

// QueryQrCodeUploadStatusResponse is the response struct for api QueryQrCodeUploadStatus
type QueryQrCodeUploadStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OssKey    string `json:"OssKey" xml:"OssKey"`
	OssUrl    string `json:"OssUrl" xml:"OssUrl"`
	Status    int    `json:"Status" xml:"Status"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateQueryQrCodeUploadStatusRequest creates a request to invoke QueryQrCodeUploadStatus API
func CreateQueryQrCodeUploadStatusRequest() (request *QueryQrCodeUploadStatusRequest) {
	request = &QueryQrCodeUploadStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "QueryQrCodeUploadStatus", "trademark", "openAPI")
	return
}

// CreateQueryQrCodeUploadStatusResponse creates a response to parse from QueryQrCodeUploadStatus response
func CreateQueryQrCodeUploadStatusResponse() (response *QueryQrCodeUploadStatusResponse) {
	response = &QueryQrCodeUploadStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}