package edas

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

// UpdateFlowControl invokes the edas.UpdateFlowControl API synchronously
// api document: https://help.aliyun.com/api/edas/updateflowcontrol.html
func (client *Client) UpdateFlowControl(request *UpdateFlowControlRequest) (response *UpdateFlowControlResponse, err error) {
	response = CreateUpdateFlowControlResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFlowControlWithChan invokes the edas.UpdateFlowControl API asynchronously
// api document: https://help.aliyun.com/api/edas/updateflowcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateFlowControlWithChan(request *UpdateFlowControlRequest) (<-chan *UpdateFlowControlResponse, <-chan error) {
	responseChan := make(chan *UpdateFlowControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFlowControl(request)
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

// UpdateFlowControlWithCallback invokes the edas.UpdateFlowControl API asynchronously
// api document: https://help.aliyun.com/api/edas/updateflowcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateFlowControlWithCallback(request *UpdateFlowControlRequest, callback func(response *UpdateFlowControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFlowControlResponse
		var err error
		defer close(result)
		response, err = client.UpdateFlowControl(request)
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

// UpdateFlowControlRequest is the request struct for api UpdateFlowControl
type UpdateFlowControlRequest struct {
	*requests.RoaRequest
	ConsumerAppId string           `position:"Query" name:"ConsumerAppId"`
	Granularity   string           `position:"Query" name:"Granularity"`
	RuleType      string           `position:"Query" name:"RuleType"`
	AppId         string           `position:"Query" name:"AppId"`
	UrlVar        string           `position:"Query" name:"UrlVar"`
	ServiceName   string           `position:"Query" name:"ServiceName"`
	Threshold     requests.Integer `position:"Query" name:"Threshold"`
	RuleId        string           `position:"Query" name:"RuleId"`
	Strategy      string           `position:"Query" name:"Strategy"`
	MethodName    string           `position:"Query" name:"MethodName"`
}

// UpdateFlowControlResponse is the response struct for api UpdateFlowControl
type UpdateFlowControlResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateFlowControlRequest creates a request to invoke UpdateFlowControl API
func CreateUpdateFlowControlRequest() (request *UpdateFlowControlRequest) {
	request = &UpdateFlowControlRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateFlowControl", "/pop/v5/flowControl", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateFlowControlResponse creates a response to parse from UpdateFlowControl response
func CreateUpdateFlowControlResponse() (response *UpdateFlowControlResponse) {
	response = &UpdateFlowControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}