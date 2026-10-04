# TeamPlanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is whether that plan&#39;s entitlement is live. | [optional] 
**GuestLimit** | Pointer to **int64** | GuestLimit is the plan&#39;s team.guests cap, when the plan carries one. | [optional] 
**Guests** | Pointer to **int64** | Guests is how many of those seats are guests. | [optional] 
**Plan** | Pointer to **string** | Plan is the licensed plan id, empty when it cannot be resolved here — an honest dash on the page, never a fabricated tier. | [optional] 
**Seats** | Pointer to **int64** | Seats is the org&#39;s distinct active human members. | [optional] 
**UpgradeUrl** | Pointer to **string** | UpgradeURL is where the page sends a caller who wants a bigger plan. | [optional] 

## Methods

### NewTeamPlanInfo

`func NewTeamPlanInfo() *TeamPlanInfo`

NewTeamPlanInfo instantiates a new TeamPlanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPlanInfoWithDefaults

`func NewTeamPlanInfoWithDefaults() *TeamPlanInfo`

NewTeamPlanInfoWithDefaults instantiates a new TeamPlanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *TeamPlanInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TeamPlanInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TeamPlanInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TeamPlanInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetGuestLimit

`func (o *TeamPlanInfo) GetGuestLimit() int64`

GetGuestLimit returns the GuestLimit field if non-nil, zero value otherwise.

### GetGuestLimitOk

`func (o *TeamPlanInfo) GetGuestLimitOk() (*int64, bool)`

GetGuestLimitOk returns a tuple with the GuestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestLimit

`func (o *TeamPlanInfo) SetGuestLimit(v int64)`

SetGuestLimit sets GuestLimit field to given value.

### HasGuestLimit

`func (o *TeamPlanInfo) HasGuestLimit() bool`

HasGuestLimit returns a boolean if a field has been set.

### GetGuests

`func (o *TeamPlanInfo) GetGuests() int64`

GetGuests returns the Guests field if non-nil, zero value otherwise.

### GetGuestsOk

`func (o *TeamPlanInfo) GetGuestsOk() (*int64, bool)`

GetGuestsOk returns a tuple with the Guests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuests

`func (o *TeamPlanInfo) SetGuests(v int64)`

SetGuests sets Guests field to given value.

### HasGuests

`func (o *TeamPlanInfo) HasGuests() bool`

HasGuests returns a boolean if a field has been set.

### GetPlan

`func (o *TeamPlanInfo) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *TeamPlanInfo) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *TeamPlanInfo) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *TeamPlanInfo) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSeats

`func (o *TeamPlanInfo) GetSeats() int64`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *TeamPlanInfo) GetSeatsOk() (*int64, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *TeamPlanInfo) SetSeats(v int64)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *TeamPlanInfo) HasSeats() bool`

HasSeats returns a boolean if a field has been set.

### GetUpgradeUrl

`func (o *TeamPlanInfo) GetUpgradeUrl() string`

GetUpgradeUrl returns the UpgradeUrl field if non-nil, zero value otherwise.

### GetUpgradeUrlOk

`func (o *TeamPlanInfo) GetUpgradeUrlOk() (*string, bool)`

GetUpgradeUrlOk returns a tuple with the UpgradeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeUrl

`func (o *TeamPlanInfo) SetUpgradeUrl(v string)`

SetUpgradeUrl sets UpgradeUrl field to given value.

### HasUpgradeUrl

`func (o *TeamPlanInfo) HasUpgradeUrl() bool`

HasUpgradeUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


