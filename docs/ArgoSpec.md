# ArgoSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**ArgoDestination**](ArgoDestination.md) |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**ArgoSource**](ArgoSource.md) |  | [optional] 

## Methods

### NewArgoSpec

`func NewArgoSpec() *ArgoSpec`

NewArgoSpec instantiates a new ArgoSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoSpecWithDefaults

`func NewArgoSpecWithDefaults() *ArgoSpec`

NewArgoSpecWithDefaults instantiates a new ArgoSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *ArgoSpec) GetDestination() ArgoDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ArgoSpec) GetDestinationOk() (*ArgoDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ArgoSpec) SetDestination(v ArgoDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ArgoSpec) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetProject

`func (o *ArgoSpec) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ArgoSpec) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ArgoSpec) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ArgoSpec) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSource

`func (o *ArgoSpec) GetSource() ArgoSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ArgoSpec) GetSourceOk() (*ArgoSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ArgoSpec) SetSource(v ArgoSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ArgoSpec) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


