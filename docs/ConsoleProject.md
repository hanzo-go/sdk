# ConsoleProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**ConsoleProjectOrganization**](ConsoleProjectOrganization.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**RetentionDays** | Pointer to **int32** |  | [optional] 

## Methods

### NewConsoleProject

`func NewConsoleProject() *ConsoleProject`

NewConsoleProject instantiates a new ConsoleProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleProjectWithDefaults

`func NewConsoleProjectWithDefaults() *ConsoleProject`

NewConsoleProjectWithDefaults instantiates a new ConsoleProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConsoleProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsoleProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *ConsoleProject) GetOrganization() ConsoleProjectOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ConsoleProject) GetOrganizationOk() (*ConsoleProjectOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ConsoleProject) SetOrganization(v ConsoleProjectOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ConsoleProject) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetMetadata

`func (o *ConsoleProject) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleProject) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleProject) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleProject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRetentionDays

`func (o *ConsoleProject) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *ConsoleProject) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *ConsoleProject) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *ConsoleProject) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


