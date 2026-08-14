# O11yPagerdutyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**Class** | Pointer to **string** |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 
**ClientUrl** | Pointer to **string** |  | [optional] 
**Component** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Images** | Pointer to [**[]O11yPagerdutyImage**](O11yPagerdutyImage.md) |  | [optional] 
**Links** | Pointer to [**[]O11yPagerdutyLink**](O11yPagerdutyLink.md) |  | [optional] 
**RoutingKey** | Pointer to **interface{}** |  | [optional] 
**RoutingKeyFile** | Pointer to **string** |  | [optional] 
**ServiceKey** | Pointer to **interface{}** |  | [optional] 
**ServiceKeyFile** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** | Timeout is the maximum time allowed to invoke the pagerduty. Setting this to 0 does not impose a timeout. | [optional] 
**Url** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewO11yPagerdutyConfig

`func NewO11yPagerdutyConfig() *O11yPagerdutyConfig`

NewO11yPagerdutyConfig instantiates a new O11yPagerdutyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPagerdutyConfigWithDefaults

`func NewO11yPagerdutyConfigWithDefaults() *O11yPagerdutyConfig`

NewO11yPagerdutyConfigWithDefaults instantiates a new O11yPagerdutyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yPagerdutyConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yPagerdutyConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yPagerdutyConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yPagerdutyConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetClass

`func (o *O11yPagerdutyConfig) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *O11yPagerdutyConfig) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *O11yPagerdutyConfig) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *O11yPagerdutyConfig) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetClient

`func (o *O11yPagerdutyConfig) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *O11yPagerdutyConfig) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *O11yPagerdutyConfig) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *O11yPagerdutyConfig) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetClientUrl

`func (o *O11yPagerdutyConfig) GetClientUrl() string`

GetClientUrl returns the ClientUrl field if non-nil, zero value otherwise.

### GetClientUrlOk

`func (o *O11yPagerdutyConfig) GetClientUrlOk() (*string, bool)`

GetClientUrlOk returns a tuple with the ClientUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUrl

`func (o *O11yPagerdutyConfig) SetClientUrl(v string)`

SetClientUrl sets ClientUrl field to given value.

### HasClientUrl

`func (o *O11yPagerdutyConfig) HasClientUrl() bool`

HasClientUrl returns a boolean if a field has been set.

### GetComponent

`func (o *O11yPagerdutyConfig) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *O11yPagerdutyConfig) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *O11yPagerdutyConfig) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *O11yPagerdutyConfig) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetDescription

`func (o *O11yPagerdutyConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yPagerdutyConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yPagerdutyConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yPagerdutyConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *O11yPagerdutyConfig) GetDetails() map[string]map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *O11yPagerdutyConfig) GetDetailsOk() (*map[string]map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *O11yPagerdutyConfig) SetDetails(v map[string]map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *O11yPagerdutyConfig) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetGroup

`func (o *O11yPagerdutyConfig) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *O11yPagerdutyConfig) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *O11yPagerdutyConfig) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *O11yPagerdutyConfig) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yPagerdutyConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yPagerdutyConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yPagerdutyConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yPagerdutyConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetImages

`func (o *O11yPagerdutyConfig) GetImages() []O11yPagerdutyImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *O11yPagerdutyConfig) GetImagesOk() (*[]O11yPagerdutyImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *O11yPagerdutyConfig) SetImages(v []O11yPagerdutyImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *O11yPagerdutyConfig) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetLinks

`func (o *O11yPagerdutyConfig) GetLinks() []O11yPagerdutyLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *O11yPagerdutyConfig) GetLinksOk() (*[]O11yPagerdutyLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *O11yPagerdutyConfig) SetLinks(v []O11yPagerdutyLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *O11yPagerdutyConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRoutingKey

`func (o *O11yPagerdutyConfig) GetRoutingKey() interface{}`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *O11yPagerdutyConfig) GetRoutingKeyOk() (*interface{}, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *O11yPagerdutyConfig) SetRoutingKey(v interface{})`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *O11yPagerdutyConfig) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.

### SetRoutingKeyNil

`func (o *O11yPagerdutyConfig) SetRoutingKeyNil(b bool)`

 SetRoutingKeyNil sets the value for RoutingKey to be an explicit nil

### UnsetRoutingKey
`func (o *O11yPagerdutyConfig) UnsetRoutingKey()`

UnsetRoutingKey ensures that no value is present for RoutingKey, not even an explicit nil
### GetRoutingKeyFile

`func (o *O11yPagerdutyConfig) GetRoutingKeyFile() string`

GetRoutingKeyFile returns the RoutingKeyFile field if non-nil, zero value otherwise.

### GetRoutingKeyFileOk

`func (o *O11yPagerdutyConfig) GetRoutingKeyFileOk() (*string, bool)`

GetRoutingKeyFileOk returns a tuple with the RoutingKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKeyFile

`func (o *O11yPagerdutyConfig) SetRoutingKeyFile(v string)`

SetRoutingKeyFile sets RoutingKeyFile field to given value.

### HasRoutingKeyFile

`func (o *O11yPagerdutyConfig) HasRoutingKeyFile() bool`

HasRoutingKeyFile returns a boolean if a field has been set.

### GetServiceKey

`func (o *O11yPagerdutyConfig) GetServiceKey() interface{}`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *O11yPagerdutyConfig) GetServiceKeyOk() (*interface{}, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *O11yPagerdutyConfig) SetServiceKey(v interface{})`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *O11yPagerdutyConfig) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### SetServiceKeyNil

`func (o *O11yPagerdutyConfig) SetServiceKeyNil(b bool)`

 SetServiceKeyNil sets the value for ServiceKey to be an explicit nil

### UnsetServiceKey
`func (o *O11yPagerdutyConfig) UnsetServiceKey()`

UnsetServiceKey ensures that no value is present for ServiceKey, not even an explicit nil
### GetServiceKeyFile

`func (o *O11yPagerdutyConfig) GetServiceKeyFile() string`

GetServiceKeyFile returns the ServiceKeyFile field if non-nil, zero value otherwise.

### GetServiceKeyFileOk

`func (o *O11yPagerdutyConfig) GetServiceKeyFileOk() (*string, bool)`

GetServiceKeyFileOk returns a tuple with the ServiceKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKeyFile

`func (o *O11yPagerdutyConfig) SetServiceKeyFile(v string)`

SetServiceKeyFile sets ServiceKeyFile field to given value.

### HasServiceKeyFile

`func (o *O11yPagerdutyConfig) HasServiceKeyFile() bool`

HasServiceKeyFile returns a boolean if a field has been set.

### GetSeverity

`func (o *O11yPagerdutyConfig) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *O11yPagerdutyConfig) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *O11yPagerdutyConfig) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *O11yPagerdutyConfig) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *O11yPagerdutyConfig) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *O11yPagerdutyConfig) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *O11yPagerdutyConfig) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *O11yPagerdutyConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTimeout

`func (o *O11yPagerdutyConfig) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *O11yPagerdutyConfig) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *O11yPagerdutyConfig) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *O11yPagerdutyConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *O11yPagerdutyConfig) GetUrl() interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yPagerdutyConfig) GetUrlOk() (*interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yPagerdutyConfig) SetUrl(v interface{})`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yPagerdutyConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *O11yPagerdutyConfig) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *O11yPagerdutyConfig) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


