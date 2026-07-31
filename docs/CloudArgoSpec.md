# CloudArgoSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**CloudArgoDestination**](CloudArgoDestination.md) |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**CloudArgoSource**](CloudArgoSource.md) |  | [optional] 

## Methods

### NewCloudArgoSpec

`func NewCloudArgoSpec() *CloudArgoSpec`

NewCloudArgoSpec instantiates a new CloudArgoSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoSpecWithDefaults

`func NewCloudArgoSpecWithDefaults() *CloudArgoSpec`

NewCloudArgoSpecWithDefaults instantiates a new CloudArgoSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *CloudArgoSpec) GetDestination() CloudArgoDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CloudArgoSpec) GetDestinationOk() (*CloudArgoDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CloudArgoSpec) SetDestination(v CloudArgoDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *CloudArgoSpec) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetProject

`func (o *CloudArgoSpec) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudArgoSpec) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudArgoSpec) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudArgoSpec) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSource

`func (o *CloudArgoSpec) GetSource() CloudArgoSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudArgoSpec) GetSourceOk() (*CloudArgoSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudArgoSpec) SetSource(v CloudArgoSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudArgoSpec) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


