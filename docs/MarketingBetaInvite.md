# MarketingBetaInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **map[string]string** | Headers are the headers the message carries beside its own: the one-click unsubscribe (RFC 8058), List-Unsubscribe and List-Unsubscribe-Post. | [optional] 
**Html** | Pointer to **string** | HTML is the same message as HTML. | [optional] 
**Id** | Pointer to **string** | ID is the invite&#39;s id: the link&#39;s utm_content and what the sent report names. | [optional] 
**Link** | Pointer to **string** | Link is the member&#39;s referral link for this invite. | [optional] 
**Subject** | Pointer to **string** | Subject is the message&#39;s subject. | [optional] 
**Text** | Pointer to **string** | Text is the message&#39;s plain text, its unsubscribe line last. | [optional] 
**To** | Pointer to **string** | To is the invited address. | [optional] 

## Methods

### NewMarketingBetaInvite

`func NewMarketingBetaInvite() *MarketingBetaInvite`

NewMarketingBetaInvite instantiates a new MarketingBetaInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBetaInviteWithDefaults

`func NewMarketingBetaInviteWithDefaults() *MarketingBetaInvite`

NewMarketingBetaInviteWithDefaults instantiates a new MarketingBetaInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *MarketingBetaInvite) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MarketingBetaInvite) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MarketingBetaInvite) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MarketingBetaInvite) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHtml

`func (o *MarketingBetaInvite) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *MarketingBetaInvite) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *MarketingBetaInvite) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *MarketingBetaInvite) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetId

`func (o *MarketingBetaInvite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingBetaInvite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingBetaInvite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingBetaInvite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLink

`func (o *MarketingBetaInvite) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *MarketingBetaInvite) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *MarketingBetaInvite) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *MarketingBetaInvite) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetSubject

`func (o *MarketingBetaInvite) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MarketingBetaInvite) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MarketingBetaInvite) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MarketingBetaInvite) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *MarketingBetaInvite) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MarketingBetaInvite) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MarketingBetaInvite) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MarketingBetaInvite) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTo

`func (o *MarketingBetaInvite) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MarketingBetaInvite) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MarketingBetaInvite) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *MarketingBetaInvite) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


