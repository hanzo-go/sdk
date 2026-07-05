# AutoCreateFlowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**FolderId** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoCreateFlowRequest

`func NewAutoCreateFlowRequest(displayName string, ) *AutoCreateFlowRequest`

NewAutoCreateFlowRequest instantiates a new AutoCreateFlowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCreateFlowRequestWithDefaults

`func NewAutoCreateFlowRequestWithDefaults() *AutoCreateFlowRequest`

NewAutoCreateFlowRequestWithDefaults instantiates a new AutoCreateFlowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AutoCreateFlowRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutoCreateFlowRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutoCreateFlowRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFolderId

`func (o *AutoCreateFlowRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AutoCreateFlowRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AutoCreateFlowRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AutoCreateFlowRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


