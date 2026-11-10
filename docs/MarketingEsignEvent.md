# MarketingEsignEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | Event is what happened: DOCUMENT_SIGNED (a recipient signed), DOCUMENT_COMPLETED, DOCUMENT_REJECTED or DOCUMENT_CANCELLED. Any other event changes nothing. | [optional] 
**Payload** | Pointer to [**MarketingEsignDocument**](MarketingEsignDocument.md) | Payload is the document as it stands after the event. | [optional] 

## Methods

### NewMarketingEsignEvent

`func NewMarketingEsignEvent() *MarketingEsignEvent`

NewMarketingEsignEvent instantiates a new MarketingEsignEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingEsignEventWithDefaults

`func NewMarketingEsignEventWithDefaults() *MarketingEsignEvent`

NewMarketingEsignEventWithDefaults instantiates a new MarketingEsignEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *MarketingEsignEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *MarketingEsignEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *MarketingEsignEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *MarketingEsignEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetPayload

`func (o *MarketingEsignEvent) GetPayload() MarketingEsignDocument`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *MarketingEsignEvent) GetPayloadOk() (*MarketingEsignDocument, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *MarketingEsignEvent) SetPayload(v MarketingEsignDocument)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *MarketingEsignEvent) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


