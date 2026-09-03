# O11yAvailabilityPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **string** | T is the bucket start, RFC3339 in UTC. | [optional] 
**Total** | Pointer to **int64** | Total is how many services reported at all inside the bucket. It can be lower than the current total: a target added last week reported nothing the week before, and saying so is the point. | [optional] 
**Up** | Pointer to **int64** | Up is how many services were up at the end of the bucket. | [optional] 

## Methods

### NewO11yAvailabilityPoint

`func NewO11yAvailabilityPoint() *O11yAvailabilityPoint`

NewO11yAvailabilityPoint instantiates a new O11yAvailabilityPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAvailabilityPointWithDefaults

`func NewO11yAvailabilityPointWithDefaults() *O11yAvailabilityPoint`

NewO11yAvailabilityPointWithDefaults instantiates a new O11yAvailabilityPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *O11yAvailabilityPoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *O11yAvailabilityPoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *O11yAvailabilityPoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *O11yAvailabilityPoint) HasT() bool`

HasT returns a boolean if a field has been set.

### GetTotal

`func (o *O11yAvailabilityPoint) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yAvailabilityPoint) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yAvailabilityPoint) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yAvailabilityPoint) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUp

`func (o *O11yAvailabilityPoint) GetUp() int64`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *O11yAvailabilityPoint) GetUpOk() (*int64, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *O11yAvailabilityPoint) SetUp(v int64)`

SetUp sets Up field to given value.

### HasUp

`func (o *O11yAvailabilityPoint) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


