# DeclaredResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]Declared**](Declared.md) |  | [optional] 
**CdUnavailable** | Pointer to [**Unreadable**](Unreadable.md) |  | [optional] 
**Org** | Pointer to **string** | Org is the directory read — the caller&#39;s own, or another when a SuperAdmin asked to act as it. | [optional] 

## Methods

### NewDeclaredResp

`func NewDeclaredResp() *DeclaredResp`

NewDeclaredResp instantiates a new DeclaredResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclaredRespWithDefaults

`func NewDeclaredRespWithDefaults() *DeclaredResp`

NewDeclaredRespWithDefaults instantiates a new DeclaredResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *DeclaredResp) GetApps() []Declared`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *DeclaredResp) GetAppsOk() (*[]Declared, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *DeclaredResp) SetApps(v []Declared)`

SetApps sets Apps field to given value.

### HasApps

`func (o *DeclaredResp) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetCdUnavailable

`func (o *DeclaredResp) GetCdUnavailable() Unreadable`

GetCdUnavailable returns the CdUnavailable field if non-nil, zero value otherwise.

### GetCdUnavailableOk

`func (o *DeclaredResp) GetCdUnavailableOk() (*Unreadable, bool)`

GetCdUnavailableOk returns a tuple with the CdUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdUnavailable

`func (o *DeclaredResp) SetCdUnavailable(v Unreadable)`

SetCdUnavailable sets CdUnavailable field to given value.

### HasCdUnavailable

`func (o *DeclaredResp) HasCdUnavailable() bool`

HasCdUnavailable returns a boolean if a field has been set.

### GetOrg

`func (o *DeclaredResp) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *DeclaredResp) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *DeclaredResp) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *DeclaredResp) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


