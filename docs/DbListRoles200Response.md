# DbListRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]DbRole**](DbRole.md) |  | [optional] 

## Methods

### NewDbListRoles200Response

`func NewDbListRoles200Response() *DbListRoles200Response`

NewDbListRoles200Response instantiates a new DbListRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbListRoles200ResponseWithDefaults

`func NewDbListRoles200ResponseWithDefaults() *DbListRoles200Response`

NewDbListRoles200ResponseWithDefaults instantiates a new DbListRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *DbListRoles200Response) GetRoles() []DbRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DbListRoles200Response) GetRolesOk() (*[]DbRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DbListRoles200Response) SetRoles(v []DbRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *DbListRoles200Response) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


