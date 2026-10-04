# DataroomTrustAsked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the request&#39;s id, so the asker can be told about it later. | [optional] 
**State** | Pointer to **string** | State is always \&quot;open\&quot;: recording an ask decides nothing. | [optional] 

## Methods

### NewDataroomTrustAsked

`func NewDataroomTrustAsked() *DataroomTrustAsked`

NewDataroomTrustAsked instantiates a new DataroomTrustAsked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomTrustAskedWithDefaults

`func NewDataroomTrustAskedWithDefaults() *DataroomTrustAsked`

NewDataroomTrustAskedWithDefaults instantiates a new DataroomTrustAsked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataroomTrustAsked) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomTrustAsked) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomTrustAsked) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomTrustAsked) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *DataroomTrustAsked) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DataroomTrustAsked) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DataroomTrustAsked) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DataroomTrustAsked) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


