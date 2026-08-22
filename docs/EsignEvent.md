# EsignEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when it happened, in unix milliseconds. | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
**Email** | Pointer to **string** | Email is the email of whoever caused it, null for an event with no actor — the sender&#39;s own calls are recorded without one. | [optional] 
**Id** | Pointer to **string** | ID is the entry id. | [optional] 
**Name** | Pointer to **string** | Name is the name of whoever caused it, null when it was not recorded. | [optional] 
**Type** | Pointer to **string** | Type is what happened: DOCUMENT_CREATED, RECIPIENT_CREATED, FIELD_CREATED, DOCUMENT_SENT, DOCUMENT_OPENED, DOCUMENT_FIELD_INSERTED, DOCUMENT_RECIPIENT_COMPLETED, DOCUMENT_RECIPIENT_REJECTED or DOCUMENT_COMPLETED. | [optional] 

## Methods

### NewEsignEvent

`func NewEsignEvent() *EsignEvent`

NewEsignEvent instantiates a new EsignEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEventWithDefaults

`func NewEsignEventWithDefaults() *EsignEvent`

NewEsignEventWithDefaults instantiates a new EsignEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EsignEvent) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsignEvent) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsignEvent) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsignEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *EsignEvent) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EsignEvent) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EsignEvent) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EsignEvent) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *EsignEvent) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *EsignEvent) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEmail

`func (o *EsignEvent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EsignEvent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EsignEvent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EsignEvent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *EsignEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EsignEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EsignEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EsignEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EsignEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EsignEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EsignEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EsignEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EsignEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


