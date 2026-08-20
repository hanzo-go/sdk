# IamSearchOrganizationsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Organizations** | Pointer to [**[]IamOrganization**](IamOrganization.md) |  | [optional] 

## Methods

### NewIamSearchOrganizationsOutput

`func NewIamSearchOrganizationsOutput() *IamSearchOrganizationsOutput`

NewIamSearchOrganizationsOutput instantiates a new IamSearchOrganizationsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSearchOrganizationsOutputWithDefaults

`func NewIamSearchOrganizationsOutputWithDefaults() *IamSearchOrganizationsOutput`

NewIamSearchOrganizationsOutputWithDefaults instantiates a new IamSearchOrganizationsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *IamSearchOrganizationsOutput) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *IamSearchOrganizationsOutput) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *IamSearchOrganizationsOutput) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *IamSearchOrganizationsOutput) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetOrganizations

`func (o *IamSearchOrganizationsOutput) GetOrganizations() []IamOrganization`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *IamSearchOrganizationsOutput) GetOrganizationsOk() (*[]IamOrganization, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *IamSearchOrganizationsOutput) SetOrganizations(v []IamOrganization)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *IamSearchOrganizationsOutput) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


