# O11yO11ySessionOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthNSupport** | Pointer to [**O11yO11yAuthNSupport**](O11yO11yAuthNSupport.md) | AuthNSupport lists the org&#39;s open sign-in routes. | [optional] 
**Id** | Pointer to **string** | ID is the org id. | [optional] 
**Name** | Pointer to **string** | Name is the org&#39;s display name. | [optional] 
**Warning** | Pointer to [**O11yO11yErrorDetail**](O11yO11yErrorDetail.md) | Warning reports an org whose SSO is configured but not currently usable, in the platform&#39;s error shape. | [optional] 

## Methods

### NewO11yO11ySessionOrg

`func NewO11yO11ySessionOrg() *O11yO11ySessionOrg`

NewO11yO11ySessionOrg instantiates a new O11yO11ySessionOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySessionOrgWithDefaults

`func NewO11yO11ySessionOrgWithDefaults() *O11yO11ySessionOrg`

NewO11yO11ySessionOrgWithDefaults instantiates a new O11yO11ySessionOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthNSupport

`func (o *O11yO11ySessionOrg) GetAuthNSupport() O11yO11yAuthNSupport`

GetAuthNSupport returns the AuthNSupport field if non-nil, zero value otherwise.

### GetAuthNSupportOk

`func (o *O11yO11ySessionOrg) GetAuthNSupportOk() (*O11yO11yAuthNSupport, bool)`

GetAuthNSupportOk returns a tuple with the AuthNSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNSupport

`func (o *O11yO11ySessionOrg) SetAuthNSupport(v O11yO11yAuthNSupport)`

SetAuthNSupport sets AuthNSupport field to given value.

### HasAuthNSupport

`func (o *O11yO11ySessionOrg) HasAuthNSupport() bool`

HasAuthNSupport returns a boolean if a field has been set.

### GetId

`func (o *O11yO11ySessionOrg) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11ySessionOrg) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11ySessionOrg) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11ySessionOrg) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yO11ySessionOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11ySessionOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11ySessionOrg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11ySessionOrg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWarning

`func (o *O11yO11ySessionOrg) GetWarning() O11yO11yErrorDetail`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *O11yO11ySessionOrg) GetWarningOk() (*O11yO11yErrorDetail, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *O11yO11ySessionOrg) SetWarning(v O11yO11yErrorDetail)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *O11yO11ySessionOrg) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


