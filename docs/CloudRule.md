# CloudRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is the COA expense account a matching bill books to. An upsert normalizes a slug (\&quot;cloud\&quot;) to its account number. | [optional] 
**Pattern** | Pointer to **string** | Pattern is the merchant substring the rule matches on, case-insensitively. It is also the key an upsert writes by. | [optional] 
**Priority** | Pointer to **int32** | Priority breaks ties: when several patterns match, the highest wins. | [optional] 

## Methods

### NewCloudRule

`func NewCloudRule() *CloudRule`

NewCloudRule instantiates a new CloudRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRuleWithDefaults

`func NewCloudRuleWithDefaults() *CloudRule`

NewCloudRuleWithDefaults instantiates a new CloudRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CloudRule) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudRule) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudRule) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudRule) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPattern

`func (o *CloudRule) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CloudRule) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CloudRule) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CloudRule) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPriority

`func (o *CloudRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CloudRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CloudRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CloudRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


