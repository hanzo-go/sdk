# CloudProjectionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to **map[string]bool** | Apps says, per console app, whether the org may open it. The SAME six keys are always present (studio, bot, world, platform, team, admin), so a client maps over it unconditionally; a key is false both when the plan does not grant the app and when commerce could not be reached, because a read that decides what to SHOW fails to LOCKED rather than to an error. | [optional] 
**Tier** | Pointer to **string** | Tier is the plan slug commerce resolved for the org, or \&quot;\&quot; when the org has no active licensing subscription — which the console treats as its free default. | [optional] 

## Methods

### NewCloudProjectionView

`func NewCloudProjectionView() *CloudProjectionView`

NewCloudProjectionView instantiates a new CloudProjectionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProjectionViewWithDefaults

`func NewCloudProjectionViewWithDefaults() *CloudProjectionView`

NewCloudProjectionViewWithDefaults instantiates a new CloudProjectionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *CloudProjectionView) GetApps() map[string]bool`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *CloudProjectionView) GetAppsOk() (*map[string]bool, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *CloudProjectionView) SetApps(v map[string]bool)`

SetApps sets Apps field to given value.

### HasApps

`func (o *CloudProjectionView) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetTier

`func (o *CloudProjectionView) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CloudProjectionView) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CloudProjectionView) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *CloudProjectionView) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


