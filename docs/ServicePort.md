# ServicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NodePort** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewServicePort

`func NewServicePort() *ServicePort`

NewServicePort instantiates a new ServicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortWithDefaults

`func NewServicePortWithDefaults() *ServicePort`

NewServicePortWithDefaults instantiates a new ServicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServicePort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodePort

`func (o *ServicePort) GetNodePort() int32`

GetNodePort returns the NodePort field if non-nil, zero value otherwise.

### GetNodePortOk

`func (o *ServicePort) GetNodePortOk() (*int32, bool)`

GetNodePortOk returns a tuple with the NodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePort

`func (o *ServicePort) SetNodePort(v int32)`

SetNodePort sets NodePort field to given value.

### HasNodePort

`func (o *ServicePort) HasNodePort() bool`

HasNodePort returns a boolean if a field has been set.

### GetPort

`func (o *ServicePort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServicePort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServicePort) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ServicePort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *ServicePort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServicePort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServicePort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ServicePort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUrl

`func (o *ServicePort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServicePort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServicePort) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServicePort) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


