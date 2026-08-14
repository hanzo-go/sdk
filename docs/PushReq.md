# PushReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch to advance; empty means \&quot;main\&quot;. A fresh branch that is the repo&#39;s first also becomes HEAD. | [optional] 
**Files** | Pointer to [**[]PushFile**](PushFile.md) | Files are added to or overwritten on the branch tip — files already there and not listed SURVIVE. At least one, at most 5000, 32 MiB each. | [optional] 
**Message** | Pointer to **string** | Message is the commit message; empty gets a generated one. | [optional] 
**Name** | Pointer to **string** | Name is the repo to push into, from the :name path segment. It is CREATED on first push if it does not exist. | [optional] 

## Methods

### NewPushReq

`func NewPushReq() *PushReq`

NewPushReq instantiates a new PushReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushReqWithDefaults

`func NewPushReqWithDefaults() *PushReq`

NewPushReqWithDefaults instantiates a new PushReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *PushReq) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PushReq) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PushReq) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PushReq) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetFiles

`func (o *PushReq) GetFiles() []PushFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *PushReq) GetFilesOk() (*[]PushFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *PushReq) SetFiles(v []PushFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *PushReq) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetMessage

`func (o *PushReq) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PushReq) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PushReq) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PushReq) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *PushReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PushReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PushReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PushReq) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


