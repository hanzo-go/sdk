# CloudPlanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is whether that plan&#39;s entitlement is live. | [optional] 
**GuestLimit** | Pointer to **int32** | GuestLimit is the plan&#39;s team.guests cap, when the plan carries one. | [optional] 
**Guests** | Pointer to **int32** | Guests is how many of those seats are guests. | [optional] 
**Plan** | Pointer to **string** | Plan is the licensed plan id, empty when it cannot be resolved here — an honest dash on the page, never a fabricated tier. | [optional] 
**Seats** | Pointer to **int32** | Seats is the org&#39;s distinct active human members. | [optional] 
**UpgradeUrl** | Pointer to **string** | UpgradeURL is where the page sends a caller who wants a bigger plan. | [optional] 

## Methods

### NewCloudPlanInfo

`func NewCloudPlanInfo() *CloudPlanInfo`

NewCloudPlanInfo instantiates a new CloudPlanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPlanInfoWithDefaults

`func NewCloudPlanInfoWithDefaults() *CloudPlanInfo`

NewCloudPlanInfoWithDefaults instantiates a new CloudPlanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CloudPlanInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CloudPlanInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CloudPlanInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CloudPlanInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetGuestLimit

`func (o *CloudPlanInfo) GetGuestLimit() int32`

GetGuestLimit returns the GuestLimit field if non-nil, zero value otherwise.

### GetGuestLimitOk

`func (o *CloudPlanInfo) GetGuestLimitOk() (*int32, bool)`

GetGuestLimitOk returns a tuple with the GuestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestLimit

`func (o *CloudPlanInfo) SetGuestLimit(v int32)`

SetGuestLimit sets GuestLimit field to given value.

### HasGuestLimit

`func (o *CloudPlanInfo) HasGuestLimit() bool`

HasGuestLimit returns a boolean if a field has been set.

### GetGuests

`func (o *CloudPlanInfo) GetGuests() int32`

GetGuests returns the Guests field if non-nil, zero value otherwise.

### GetGuestsOk

`func (o *CloudPlanInfo) GetGuestsOk() (*int32, bool)`

GetGuestsOk returns a tuple with the Guests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuests

`func (o *CloudPlanInfo) SetGuests(v int32)`

SetGuests sets Guests field to given value.

### HasGuests

`func (o *CloudPlanInfo) HasGuests() bool`

HasGuests returns a boolean if a field has been set.

### GetPlan

`func (o *CloudPlanInfo) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CloudPlanInfo) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CloudPlanInfo) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CloudPlanInfo) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSeats

`func (o *CloudPlanInfo) GetSeats() int32`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *CloudPlanInfo) GetSeatsOk() (*int32, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *CloudPlanInfo) SetSeats(v int32)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *CloudPlanInfo) HasSeats() bool`

HasSeats returns a boolean if a field has been set.

### GetUpgradeUrl

`func (o *CloudPlanInfo) GetUpgradeUrl() string`

GetUpgradeUrl returns the UpgradeUrl field if non-nil, zero value otherwise.

### GetUpgradeUrlOk

`func (o *CloudPlanInfo) GetUpgradeUrlOk() (*string, bool)`

GetUpgradeUrlOk returns a tuple with the UpgradeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeUrl

`func (o *CloudPlanInfo) SetUpgradeUrl(v string)`

SetUpgradeUrl sets UpgradeUrl field to given value.

### HasUpgradeUrl

`func (o *CloudPlanInfo) HasUpgradeUrl() bool`

HasUpgradeUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


