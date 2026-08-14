# AuthorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AuthorData**](AuthorData.md) | Data carries the author. | [optional] 
**Msg** | Pointer to **string** | Msg is the envelope&#39;s message slot, empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot;. | [optional] 

## Methods

### NewAuthorResult

`func NewAuthorResult() *AuthorResult`

NewAuthorResult instantiates a new AuthorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorResultWithDefaults

`func NewAuthorResultWithDefaults() *AuthorResult`

NewAuthorResultWithDefaults instantiates a new AuthorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AuthorResult) GetData() AuthorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthorResult) GetDataOk() (*AuthorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthorResult) SetData(v AuthorData)`

SetData sets Data field to given value.

### HasData

`func (o *AuthorResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *AuthorResult) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AuthorResult) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AuthorResult) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AuthorResult) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


