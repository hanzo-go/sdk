# IamListOrganizationsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Organizations** | Pointer to [**[]IamOrganization**](IamOrganization.md) |  | [optional] 

## Methods

### NewIamListOrganizationsOutput

`func NewIamListOrganizationsOutput() *IamListOrganizationsOutput`

NewIamListOrganizationsOutput instantiates a new IamListOrganizationsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamListOrganizationsOutputWithDefaults

`func NewIamListOrganizationsOutputWithDefaults() *IamListOrganizationsOutput`

NewIamListOrganizationsOutputWithDefaults instantiates a new IamListOrganizationsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IamListOrganizationsOutput) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IamListOrganizationsOutput) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IamListOrganizationsOutput) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IamListOrganizationsOutput) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOrganizations

`func (o *IamListOrganizationsOutput) GetOrganizations() []IamOrganization`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *IamListOrganizationsOutput) GetOrganizationsOk() (*[]IamOrganization, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *IamListOrganizationsOutput) SetOrganizations(v []IamOrganization)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *IamListOrganizationsOutput) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


