# TicketGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int64** | ExpiresIn is how long the ticket is good for, in seconds. | [optional] 
**Ticket** | Pointer to **string** | Ticket is the grant itself. It is single-purpose and short-lived, and it travels in a query string because a WebSocket handshake carries no Authorization header a browser can set. | [optional] 
**Url** | Pointer to **string** | URL is the PATH to open, ticket included — not an absolute URL. Which host this address wears in public is the edge&#39;s answer and not this process&#39;s, so an absolute URL would be a guess; the client already knows the host it is talking to. It names the PAGE, which is what a caller embeds — the page finds its own socket, and a caller that wants the raw socket adds &#x60;/ws&#x60;. | [optional] 

## Methods

### NewTicketGrant

`func NewTicketGrant() *TicketGrant`

NewTicketGrant instantiates a new TicketGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketGrantWithDefaults

`func NewTicketGrantWithDefaults() *TicketGrant`

NewTicketGrantWithDefaults instantiates a new TicketGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *TicketGrant) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *TicketGrant) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *TicketGrant) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *TicketGrant) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetTicket

`func (o *TicketGrant) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *TicketGrant) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *TicketGrant) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *TicketGrant) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetUrl

`func (o *TicketGrant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TicketGrant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TicketGrant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TicketGrant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


