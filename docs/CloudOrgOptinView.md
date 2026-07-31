# CloudOrgOptinView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanManage** | Pointer to **bool** | may the caller edit the org opt-in | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**Listed** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudOrgOptinView

`func NewCloudOrgOptinView() *CloudOrgOptinView`

NewCloudOrgOptinView instantiates a new CloudOrgOptinView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOrgOptinViewWithDefaults

`func NewCloudOrgOptinViewWithDefaults() *CloudOrgOptinView`

NewCloudOrgOptinViewWithDefaults instantiates a new CloudOrgOptinView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanManage

`func (o *CloudOrgOptinView) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *CloudOrgOptinView) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *CloudOrgOptinView) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *CloudOrgOptinView) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetDisplay

`func (o *CloudOrgOptinView) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CloudOrgOptinView) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CloudOrgOptinView) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CloudOrgOptinView) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetListed

`func (o *CloudOrgOptinView) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *CloudOrgOptinView) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *CloudOrgOptinView) SetListed(v bool)`

SetListed sets Listed field to given value.

### HasListed

`func (o *CloudOrgOptinView) HasListed() bool`

HasListed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


