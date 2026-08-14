# InvokeReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **string** | Input is what the function is given on stdin. It is opaque to this surface. | [optional] 

## Methods

### NewInvokeReq

`func NewInvokeReq() *InvokeReq`

NewInvokeReq instantiates a new InvokeReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvokeReqWithDefaults

`func NewInvokeReqWithDefaults() *InvokeReq`

NewInvokeReqWithDefaults instantiates a new InvokeReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *InvokeReq) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *InvokeReq) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *InvokeReq) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *InvokeReq) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


