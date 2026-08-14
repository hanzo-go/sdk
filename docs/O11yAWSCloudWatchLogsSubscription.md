# O11yAWSCloudWatchLogsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterPattern** | Pointer to **string** | https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html \&quot;\&quot; implies no filtering is required | [optional] 
**LogGroupNamePrefix** | Pointer to **string** | subscribe to all logs groups with specified prefix. eg: &#x60;/aws/rds/&#x60; | [optional] 

## Methods

### NewO11yAWSCloudWatchLogsSubscription

`func NewO11yAWSCloudWatchLogsSubscription() *O11yAWSCloudWatchLogsSubscription`

NewO11yAWSCloudWatchLogsSubscription instantiates a new O11yAWSCloudWatchLogsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAWSCloudWatchLogsSubscriptionWithDefaults

`func NewO11yAWSCloudWatchLogsSubscriptionWithDefaults() *O11yAWSCloudWatchLogsSubscription`

NewO11yAWSCloudWatchLogsSubscriptionWithDefaults instantiates a new O11yAWSCloudWatchLogsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterPattern

`func (o *O11yAWSCloudWatchLogsSubscription) GetFilterPattern() string`

GetFilterPattern returns the FilterPattern field if non-nil, zero value otherwise.

### GetFilterPatternOk

`func (o *O11yAWSCloudWatchLogsSubscription) GetFilterPatternOk() (*string, bool)`

GetFilterPatternOk returns a tuple with the FilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPattern

`func (o *O11yAWSCloudWatchLogsSubscription) SetFilterPattern(v string)`

SetFilterPattern sets FilterPattern field to given value.

### HasFilterPattern

`func (o *O11yAWSCloudWatchLogsSubscription) HasFilterPattern() bool`

HasFilterPattern returns a boolean if a field has been set.

### GetLogGroupNamePrefix

`func (o *O11yAWSCloudWatchLogsSubscription) GetLogGroupNamePrefix() string`

GetLogGroupNamePrefix returns the LogGroupNamePrefix field if non-nil, zero value otherwise.

### GetLogGroupNamePrefixOk

`func (o *O11yAWSCloudWatchLogsSubscription) GetLogGroupNamePrefixOk() (*string, bool)`

GetLogGroupNamePrefixOk returns a tuple with the LogGroupNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogGroupNamePrefix

`func (o *O11yAWSCloudWatchLogsSubscription) SetLogGroupNamePrefix(v string)`

SetLogGroupNamePrefix sets LogGroupNamePrefix field to given value.

### HasLogGroupNamePrefix

`func (o *O11yAWSCloudWatchLogsSubscription) HasLogGroupNamePrefix() bool`

HasLogGroupNamePrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


