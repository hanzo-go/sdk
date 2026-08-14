# PlanInfo

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

### NewPlanInfo

`func NewPlanInfo() *PlanInfo`

NewPlanInfo instantiates a new PlanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanInfoWithDefaults

`func NewPlanInfoWithDefaults() *PlanInfo`

NewPlanInfoWithDefaults instantiates a new PlanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *PlanInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PlanInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PlanInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PlanInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetGuestLimit

`func (o *PlanInfo) GetGuestLimit() int32`

GetGuestLimit returns the GuestLimit field if non-nil, zero value otherwise.

### GetGuestLimitOk

`func (o *PlanInfo) GetGuestLimitOk() (*int32, bool)`

GetGuestLimitOk returns a tuple with the GuestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestLimit

`func (o *PlanInfo) SetGuestLimit(v int32)`

SetGuestLimit sets GuestLimit field to given value.

### HasGuestLimit

`func (o *PlanInfo) HasGuestLimit() bool`

HasGuestLimit returns a boolean if a field has been set.

### GetGuests

`func (o *PlanInfo) GetGuests() int32`

GetGuests returns the Guests field if non-nil, zero value otherwise.

### GetGuestsOk

`func (o *PlanInfo) GetGuestsOk() (*int32, bool)`

GetGuestsOk returns a tuple with the Guests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuests

`func (o *PlanInfo) SetGuests(v int32)`

SetGuests sets Guests field to given value.

### HasGuests

`func (o *PlanInfo) HasGuests() bool`

HasGuests returns a boolean if a field has been set.

### GetPlan

`func (o *PlanInfo) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PlanInfo) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PlanInfo) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *PlanInfo) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSeats

`func (o *PlanInfo) GetSeats() int32`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *PlanInfo) GetSeatsOk() (*int32, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *PlanInfo) SetSeats(v int32)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *PlanInfo) HasSeats() bool`

HasSeats returns a boolean if a field has been set.

### GetUpgradeUrl

`func (o *PlanInfo) GetUpgradeUrl() string`

GetUpgradeUrl returns the UpgradeUrl field if non-nil, zero value otherwise.

### GetUpgradeUrlOk

`func (o *PlanInfo) GetUpgradeUrlOk() (*string, bool)`

GetUpgradeUrlOk returns a tuple with the UpgradeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeUrl

`func (o *PlanInfo) SetUpgradeUrl(v string)`

SetUpgradeUrl sets UpgradeUrl field to given value.

### HasUpgradeUrl

`func (o *PlanInfo) HasUpgradeUrl() bool`

HasUpgradeUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


