# CloudOrgOptinReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | Pointer to **string** | Display is the name shown for the org on that board: 1-40 characters of letters, digits, space, dot, underscore, apostrophe or hyphen. Left empty on a listing opt-in it defaults to the org id. | [optional] 
**Listed** | Pointer to **bool** | Listed publishes the org on the cross-org global board when true, and withdraws it when false. | [optional] 

## Methods

### NewCloudOrgOptinReq

`func NewCloudOrgOptinReq() *CloudOrgOptinReq`

NewCloudOrgOptinReq instantiates a new CloudOrgOptinReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOrgOptinReqWithDefaults

`func NewCloudOrgOptinReqWithDefaults() *CloudOrgOptinReq`

NewCloudOrgOptinReqWithDefaults instantiates a new CloudOrgOptinReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplay

`func (o *CloudOrgOptinReq) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CloudOrgOptinReq) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CloudOrgOptinReq) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CloudOrgOptinReq) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetListed

`func (o *CloudOrgOptinReq) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *CloudOrgOptinReq) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *CloudOrgOptinReq) SetListed(v bool)`

SetListed sets Listed field to given value.

### HasListed

`func (o *CloudOrgOptinReq) HasListed() bool`

HasListed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


