# CommerceFulfillment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CommerceFulfillmentStatus**](CommerceFulfillmentStatus.md) |  | [optional] 
**Carrier** | Pointer to **string** |  | [optional] 
**TrackingNumber** | Pointer to **string** |  | [optional] 
**ShippedAt** | Pointer to **time.Time** |  | [optional] 
**DeliveredAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCommerceFulfillment

`func NewCommerceFulfillment() *CommerceFulfillment`

NewCommerceFulfillment instantiates a new CommerceFulfillment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceFulfillmentWithDefaults

`func NewCommerceFulfillmentWithDefaults() *CommerceFulfillment`

NewCommerceFulfillmentWithDefaults instantiates a new CommerceFulfillment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CommerceFulfillment) GetStatus() CommerceFulfillmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommerceFulfillment) GetStatusOk() (*CommerceFulfillmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommerceFulfillment) SetStatus(v CommerceFulfillmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommerceFulfillment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCarrier

`func (o *CommerceFulfillment) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *CommerceFulfillment) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *CommerceFulfillment) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *CommerceFulfillment) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *CommerceFulfillment) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *CommerceFulfillment) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *CommerceFulfillment) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *CommerceFulfillment) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetShippedAt

`func (o *CommerceFulfillment) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *CommerceFulfillment) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *CommerceFulfillment) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *CommerceFulfillment) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### GetDeliveredAt

`func (o *CommerceFulfillment) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *CommerceFulfillment) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *CommerceFulfillment) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *CommerceFulfillment) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


