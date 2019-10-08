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

// MetastoreDropDatabase invokes the emr.MetastoreDropDatabase API synchronously
// api document: https://help.aliyun.com/api/emr/metastoredropdatabase.html
func (client *Client) MetastoreDropDatabase(request *MetastoreDropDatabaseRequest) (response *MetastoreDropDatabaseResponse, err error) {
	response = CreateMetastoreDropDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// MetastoreDropDatabaseWithChan invokes the emr.MetastoreDropDatabase API asynchronously
// api document: https://help.aliyun.com/api/emr/metastoredropdatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreDropDatabaseWithChan(request *MetastoreDropDatabaseRequest) (<-chan *MetastoreDropDatabaseResponse, <-chan error) {
	responseChan := make(chan *MetastoreDropDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MetastoreDropDatabase(request)
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

// MetastoreDropDatabaseWithCallback invokes the emr.MetastoreDropDatabase API asynchronously
// api document: https://help.aliyun.com/api/emr/metastoredropdatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreDropDatabaseWithCallback(request *MetastoreDropDatabaseRequest, callback func(response *MetastoreDropDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MetastoreDropDatabaseResponse
		var err error
		defer close(result)
		response, err = client.MetastoreDropDatabase(request)
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

// MetastoreDropDatabaseRequest is the request struct for api MetastoreDropDatabase
type MetastoreDropDatabaseRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DbName          string           `position:"Query" name:"DbName"`
	DatabaseId      string           `position:"Query" name:"DatabaseId"`
}

// MetastoreDropDatabaseResponse is the response struct for api MetastoreDropDatabase
type MetastoreDropDatabaseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMetastoreDropDatabaseRequest creates a request to invoke MetastoreDropDatabase API
func CreateMetastoreDropDatabaseRequest() (request *MetastoreDropDatabaseRequest) {
	request = &MetastoreDropDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "MetastoreDropDatabase", "emr", "openAPI")
	return
}

// CreateMetastoreDropDatabaseResponse creates a response to parse from MetastoreDropDatabase response
func CreateMetastoreDropDatabaseResponse() (response *MetastoreDropDatabaseResponse) {
	response = &MetastoreDropDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}