# CloudOnboardReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the organization&#39;s display name. Ignored when personal is true, which derives the name from the caller&#39;s own username instead. | [optional] 
**Personal** | Pointer to **bool** | Personal asks for the caller&#39;s own workspace: the name is derived from their username and the slug auto-suffixes to stay unique. Meaningless — and refused — for a caller who already has an organization. | [optional] 

## Methods

### NewCloudOnboardReq

`func NewCloudOnboardReq() *CloudOnboardReq`

NewCloudOnboardReq instantiates a new CloudOnboardReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOnboardReqWithDefaults

`func NewCloudOnboardReqWithDefaults() *CloudOnboardReq`

NewCloudOnboardReqWithDefaults instantiates a new CloudOnboardReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudOnboardReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudOnboardReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudOnboardReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudOnboardReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersonal

`func (o *CloudOnboardReq) GetPersonal() bool`

GetPersonal returns the Personal field if non-nil, zero value otherwise.

### GetPersonalOk

`func (o *CloudOnboardReq) GetPersonalOk() (*bool, bool)`

GetPersonalOk returns a tuple with the Personal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonal

`func (o *CloudOnboardReq) SetPersonal(v bool)`

SetPersonal sets Personal field to given value.

### HasPersonal

`func (o *CloudOnboardReq) HasPersonal() bool`

HasPersonal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


