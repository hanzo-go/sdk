# PatrolPatrolMoveIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Act** | Pointer to **string** | Act is out, in or audit. | [optional] 
**Holder** | Pointer to **string** | Holder is who the keys went to. Who SIGNED for the move is the caller, and is stamped from the request rather than read from here. | [optional] 
**Job** | Pointer to **string** | Job is the work the keys were drawn for. | [optional] 
**Name** | Pointer to **string** | Name is the key set&#39;s reference, from the path. | [optional] 

## Methods

### NewPatrolPatrolMoveIn

`func NewPatrolPatrolMoveIn() *PatrolPatrolMoveIn`

NewPatrolPatrolMoveIn instantiates a new PatrolPatrolMoveIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolMoveInWithDefaults

`func NewPatrolPatrolMoveInWithDefaults() *PatrolPatrolMoveIn`

NewPatrolPatrolMoveInWithDefaults instantiates a new PatrolPatrolMoveIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAct

`func (o *PatrolPatrolMoveIn) GetAct() string`

GetAct returns the Act field if non-nil, zero value otherwise.

### GetActOk

`func (o *PatrolPatrolMoveIn) GetActOk() (*string, bool)`

GetActOk returns a tuple with the Act field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAct

`func (o *PatrolPatrolMoveIn) SetAct(v string)`

SetAct sets Act field to given value.

### HasAct

`func (o *PatrolPatrolMoveIn) HasAct() bool`

HasAct returns a boolean if a field has been set.

### GetHolder

`func (o *PatrolPatrolMoveIn) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *PatrolPatrolMoveIn) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *PatrolPatrolMoveIn) SetHolder(v string)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *PatrolPatrolMoveIn) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetJob

`func (o *PatrolPatrolMoveIn) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *PatrolPatrolMoveIn) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *PatrolPatrolMoveIn) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *PatrolPatrolMoveIn) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolMoveIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolMoveIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolMoveIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolMoveIn) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


