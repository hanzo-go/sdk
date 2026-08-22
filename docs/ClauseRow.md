# ClauseRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controls** | Pointer to **[]string** | Controls are the ids behind it, strongest first. Empty when nothing covers it — and an absent control is never listed here, however it maps. | [optional] 
**Group** | Pointer to **string** | Group is the clause&#39;s section within the standard, when it has one. | [optional] 
**Id** | Pointer to **string** | ID is the clause id as the standard publishes it — \&quot;CC6.1\&quot;, \&quot;A.5.15\&quot;, \&quot;AC\&quot;. | [optional] 
**Level** | Pointer to **string** | Level is what the strongest control pointed at this clause is worth: \&quot;automated\&quot;, \&quot;partial\&quot; or \&quot;none\&quot;. | [optional] 
**Title** | Pointer to **string** | Title is the standard&#39;s own words for that clause. | [optional] 

## Methods

### NewClauseRow

`func NewClauseRow() *ClauseRow`

NewClauseRow instantiates a new ClauseRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClauseRowWithDefaults

`func NewClauseRowWithDefaults() *ClauseRow`

NewClauseRowWithDefaults instantiates a new ClauseRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControls

`func (o *ClauseRow) GetControls() []string`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *ClauseRow) GetControlsOk() (*[]string, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *ClauseRow) SetControls(v []string)`

SetControls sets Controls field to given value.

### HasControls

`func (o *ClauseRow) HasControls() bool`

HasControls returns a boolean if a field has been set.

### GetGroup

`func (o *ClauseRow) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ClauseRow) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ClauseRow) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ClauseRow) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *ClauseRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClauseRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClauseRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClauseRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *ClauseRow) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ClauseRow) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ClauseRow) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ClauseRow) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetTitle

`func (o *ClauseRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClauseRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClauseRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ClauseRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


