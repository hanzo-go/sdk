# PubsubListObjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]PubsubObjectMeta**](PubsubObjectMeta.md) |  | [optional] 

## Methods

### NewPubsubListObjects200Response

`func NewPubsubListObjects200Response() *PubsubListObjects200Response`

NewPubsubListObjects200Response instantiates a new PubsubListObjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubListObjects200ResponseWithDefaults

`func NewPubsubListObjects200ResponseWithDefaults() *PubsubListObjects200Response`

NewPubsubListObjects200ResponseWithDefaults instantiates a new PubsubListObjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *PubsubListObjects200Response) GetObjects() []PubsubObjectMeta`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PubsubListObjects200Response) GetObjectsOk() (*[]PubsubObjectMeta, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PubsubListObjects200Response) SetObjects(v []PubsubObjectMeta)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *PubsubListObjects200Response) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


