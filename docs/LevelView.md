# LevelView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownlineCount** | Pointer to **int64** | DownlineCount is how many orgs sit exactly this many hops below the caller. It is 0 in the schedule quoted to a caller that has not applied, which has no downline to count. | [optional] 
**Level** | Pointer to **int64** | Level is the upline distance from the org whose spend is being shared: 1 is the direct referrer, 2 and 3 the referrers above it. Nothing accrues past 3. | [optional] 
**RateBps** | Pointer to **int64** | RateBps is the commission paid at this level, in basis points OF Hanzo&#39;s margin (2000 &#x3D; 20% of margin, never of the customer&#39;s bill). Level 1 is the affiliate&#39;s own negotiated rate; 2 and 3 are platform switches read live, so this is the schedule actually in force, not one compiled in. | [optional] 

## Methods

### NewLevelView

`func NewLevelView() *LevelView`

NewLevelView instantiates a new LevelView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLevelViewWithDefaults

`func NewLevelViewWithDefaults() *LevelView`

NewLevelViewWithDefaults instantiates a new LevelView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownlineCount

`func (o *LevelView) GetDownlineCount() int64`

GetDownlineCount returns the DownlineCount field if non-nil, zero value otherwise.

### GetDownlineCountOk

`func (o *LevelView) GetDownlineCountOk() (*int64, bool)`

GetDownlineCountOk returns a tuple with the DownlineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlineCount

`func (o *LevelView) SetDownlineCount(v int64)`

SetDownlineCount sets DownlineCount field to given value.

### HasDownlineCount

`func (o *LevelView) HasDownlineCount() bool`

HasDownlineCount returns a boolean if a field has been set.

### GetLevel

`func (o *LevelView) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LevelView) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LevelView) SetLevel(v int64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *LevelView) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetRateBps

`func (o *LevelView) GetRateBps() int64`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *LevelView) GetRateBpsOk() (*int64, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *LevelView) SetRateBps(v int64)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *LevelView) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


