# CloudInstallReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tool** | Pointer to **string** | Tool is the registry name of the capability to activate (or deactivate) for the caller&#39;s own org and project. Required. | [optional] 

## Methods

### NewCloudInstallReq

`func NewCloudInstallReq() *CloudInstallReq`

NewCloudInstallReq instantiates a new CloudInstallReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInstallReqWithDefaults

`func NewCloudInstallReqWithDefaults() *CloudInstallReq`

NewCloudInstallReqWithDefaults instantiates a new CloudInstallReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTool

`func (o *CloudInstallReq) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *CloudInstallReq) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *CloudInstallReq) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *CloudInstallReq) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


