package dfs

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

// DeleteMountPoint invokes the dfs.DeleteMountPoint API synchronously
func (client *Client) DeleteMountPoint(request *DeleteMountPointRequest) (response *DeleteMountPointResponse, err error) {
	response = CreateDeleteMountPointResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMountPointWithChan invokes the dfs.DeleteMountPoint API asynchronously
func (client *Client) DeleteMountPointWithChan(request *DeleteMountPointRequest) (<-chan *DeleteMountPointResponse, <-chan error) {
	responseChan := make(chan *DeleteMountPointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMountPoint(request)
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

// DeleteMountPointWithCallback invokes the dfs.DeleteMountPoint API asynchronously
func (client *Client) DeleteMountPointWithCallback(request *DeleteMountPointRequest, callback func(response *DeleteMountPointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMountPointResponse
		var err error
		defer close(result)
		response, err = client.DeleteMountPoint(request)
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

// DeleteMountPointRequest is the request struct for api DeleteMountPoint
type DeleteMountPointRequest struct {
	*requests.RpcRequest
	InputRegionId string `position:"Query" name:"InputRegionId"`
	MountPointId  string `position:"Query" name:"MountPointId"`
	FileSystemId  string `position:"Query" name:"FileSystemId"`
}

// DeleteMountPointResponse is the response struct for api DeleteMountPoint
type DeleteMountPointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteMountPointRequest creates a request to invoke DeleteMountPoint API
func CreateDeleteMountPointRequest() (request *DeleteMountPointRequest) {
	request = &DeleteMountPointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DFS", "2018-06-20", "DeleteMountPoint", "alidfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteMountPointResponse creates a response to parse from DeleteMountPoint response
func CreateDeleteMountPointResponse() (response *DeleteMountPointResponse) {
	response = &DeleteMountPointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
