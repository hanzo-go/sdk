# TrainPublishIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Force** | Pointer to **bool** | Force publishes a rejected job&#39;s output, and is recorded. | [optional] 
**Id** | Pointer to **string** | ID names the job. | [optional] 
**Name** | Pointer to **string** | Name is the published name; the output&#39;s name when absent. Publishing a name the org already published replaces that capability. | [optional] 
**Route** | Pointer to [**TrainRoute**](TrainRoute.md) | Route is which questions the capability answers: required to publish a capability, refused for any other output. | [optional] 

## Methods

### NewTrainPublishIn

`func NewTrainPublishIn() *TrainPublishIn`

NewTrainPublishIn instantiates a new TrainPublishIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainPublishInWithDefaults

`func NewTrainPublishInWithDefaults() *TrainPublishIn`

NewTrainPublishInWithDefaults instantiates a new TrainPublishIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForce

`func (o *TrainPublishIn) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *TrainPublishIn) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *TrainPublishIn) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *TrainPublishIn) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetId

`func (o *TrainPublishIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainPublishIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainPublishIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainPublishIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TrainPublishIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrainPublishIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrainPublishIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrainPublishIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoute

`func (o *TrainPublishIn) GetRoute() TrainRoute`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *TrainPublishIn) GetRouteOk() (*TrainRoute, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *TrainPublishIn) SetRoute(v TrainRoute)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *TrainPublishIn) HasRoute() bool`

HasRoute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


