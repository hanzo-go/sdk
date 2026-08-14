# AffiliateOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AffiliateData**](AffiliateData.md) | Data carries the affiliate row the action just wrote. | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewAffiliateOut

`func NewAffiliateOut() *AffiliateOut`

NewAffiliateOut instantiates a new AffiliateOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateOutWithDefaults

`func NewAffiliateOutWithDefaults() *AffiliateOut`

NewAffiliateOutWithDefaults instantiates a new AffiliateOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AffiliateOut) GetData() AffiliateData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AffiliateOut) GetDataOk() (*AffiliateData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AffiliateOut) SetData(v AffiliateData)`

SetData sets Data field to given value.

### HasData

`func (o *AffiliateOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *AffiliateOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AffiliateOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AffiliateOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AffiliateOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *AffiliateOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AffiliateOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AffiliateOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AffiliateOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


