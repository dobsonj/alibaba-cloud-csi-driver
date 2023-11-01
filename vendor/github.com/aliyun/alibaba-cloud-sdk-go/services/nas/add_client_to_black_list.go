package nas

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

// AddClientToBlackList invokes the nas.AddClientToBlackList API synchronously
func (client *Client) AddClientToBlackList(request *AddClientToBlackListRequest) (response *AddClientToBlackListResponse, err error) {
	response = CreateAddClientToBlackListResponse()
	err = client.DoAction(request, response)
	return
}

// AddClientToBlackListWithChan invokes the nas.AddClientToBlackList API asynchronously
func (client *Client) AddClientToBlackListWithChan(request *AddClientToBlackListRequest) (<-chan *AddClientToBlackListResponse, <-chan error) {
	responseChan := make(chan *AddClientToBlackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddClientToBlackList(request)
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

// AddClientToBlackListWithCallback invokes the nas.AddClientToBlackList API asynchronously
func (client *Client) AddClientToBlackListWithCallback(request *AddClientToBlackListRequest, callback func(response *AddClientToBlackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddClientToBlackListResponse
		var err error
		defer close(result)
		response, err = client.AddClientToBlackList(request)
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

// AddClientToBlackListRequest is the request struct for api AddClientToBlackList
type AddClientToBlackListRequest struct {
	*requests.RpcRequest
	ClientToken  string `position:"Query" name:"ClientToken"`
	ClientIP     string `position:"Query" name:"ClientIP"`
	FileSystemId string `position:"Query" name:"FileSystemId"`
}

// AddClientToBlackListResponse is the response struct for api AddClientToBlackList
type AddClientToBlackListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddClientToBlackListRequest creates a request to invoke AddClientToBlackList API
func CreateAddClientToBlackListRequest() (request *AddClientToBlackListRequest) {
	request = &AddClientToBlackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "AddClientToBlackList", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddClientToBlackListResponse creates a response to parse from AddClientToBlackList response
func CreateAddClientToBlackListResponse() (response *AddClientToBlackListResponse) {
	response = &AddClientToBlackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
