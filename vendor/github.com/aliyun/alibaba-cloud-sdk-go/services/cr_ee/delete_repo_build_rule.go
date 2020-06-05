package cr_ee

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

// DeleteRepoBuildRule invokes the cr.DeleteRepoBuildRule API synchronously
// api document: https://help.aliyun.com/api/cr/deleterepobuildrule.html
func (client *Client) DeleteRepoBuildRule(request *DeleteRepoBuildRuleRequest) (response *DeleteRepoBuildRuleResponse, err error) {
	response = CreateDeleteRepoBuildRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRepoBuildRuleWithChan invokes the cr.DeleteRepoBuildRule API asynchronously
// api document: https://help.aliyun.com/api/cr/deleterepobuildrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRepoBuildRuleWithChan(request *DeleteRepoBuildRuleRequest) (<-chan *DeleteRepoBuildRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteRepoBuildRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRepoBuildRule(request)
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

// DeleteRepoBuildRuleWithCallback invokes the cr.DeleteRepoBuildRule API asynchronously
// api document: https://help.aliyun.com/api/cr/deleterepobuildrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRepoBuildRuleWithCallback(request *DeleteRepoBuildRuleRequest, callback func(response *DeleteRepoBuildRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRepoBuildRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteRepoBuildRule(request)
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

// DeleteRepoBuildRuleRequest is the request struct for api DeleteRepoBuildRule
type DeleteRepoBuildRuleRequest struct {
	*requests.RpcRequest
	RepoId      string `position:"Query" name:"RepoId"`
	BuildRuleId string `position:"Query" name:"BuildRuleId"`
	InstanceId  string `position:"Query" name:"InstanceId"`
}

// DeleteRepoBuildRuleResponse is the response struct for api DeleteRepoBuildRule
type DeleteRepoBuildRuleResponse struct {
	*responses.BaseResponse
	DeleteRepoBuildRuleIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                         string `json:"Code" xml:"Code"`
	RequestId                    string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRepoBuildRuleRequest creates a request to invoke DeleteRepoBuildRule API
func CreateDeleteRepoBuildRuleRequest() (request *DeleteRepoBuildRuleRequest) {
	request = &DeleteRepoBuildRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "DeleteRepoBuildRule", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteRepoBuildRuleResponse creates a response to parse from DeleteRepoBuildRule response
func CreateDeleteRepoBuildRuleResponse() (response *DeleteRepoBuildRuleResponse) {
	response = &DeleteRepoBuildRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}