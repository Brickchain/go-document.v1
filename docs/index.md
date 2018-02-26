# Documents

This document should contain all document types used in the Brickchain system, that are used by more than one component.

## Base

All other documents are based on this Base document type.

```go
type Base struct {
  Context          string    `json:"@context,omitempty"`
  Type             string    `json:"@type"`
  SubType          string    `json:"@subtype,omitempty"`
  Timestamp        time.Time `json:"@timestamp"`
  ID               string    `json:"@id,omitempty"`
  Links            []Link    `json:"@links,omitempty"`
  Owners           []string  `json:"@owners,omitempty"`
  Callbacks        []string  `json:"@callbacks,omitempty"`
  CertificateChain string    `json:"@certificateChain,omitempty"`
  mu        *sync.Mutex
}
```

| Label     | Description       | Example
|:----------|:------------------|:--------
| Context   | ?                 | ?
| Type      | Document type     | "mandate", "receipt", "action" etc.
| SubType   | Document sub type | ?
| Timestamp | The timestamp of when the document was created | This is a time.Time object in golang
| ID        | A unique identifier (uuid) of the document | "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
| Links     | ?                 | ?
| Owners    | The document owners (described as what?) | ?
| Callbacks | A list of callback URIs | ["https://example.com/api"]
| CertificateChain | A compact format JWS containing a CertificateChain type document | 
| mu        | Used internally for handling mutexes | N/A


## Mandate

A Mandate is the user's credentials to perform an Action. A Mandate is typically signed by a Realm, but could also be signed by a user.

```go
type Mandate struct {
	Base
	Role          string            `json:"role,omitempty"`
	Label         string            `json:"label,omitempty"`
	TTL           int               `json:"ttl,omitempty"`
	Recipient     string            `json:"recipient,omitempty"`
	RecipientName string            `json:"recipientName,omitempty"`
	RecipientPK   *jose.JsonWebKey  `json:"recipientPublicKey,omitempty"`
	RequestID     string            `json:"requestId,omitempty"`
	Sender        string            `json:"sender,omitempty"`
	Params        map[string]string `json:"params,omitempty"`
}
```

| Label         | Description     | Example
|:--------------|:----------------|:--------
| Base          | This includes fields from the base document type | N/A
| Role          | The ID of a Role | "admin@example.com"
| Label         | Label to use when listing mandates | "Open doors"
| TTL           | Seconds until expiration after the document was created (Timestamp) | 86400 (24hr)
| Recipient     | ?               | ?
| RecipientName | The name of the recipient | "Donald J. Trump"
| RecipientPK   | The Public Key of the recipient (or a hash of it?) | ?
| Params        | Parameters to send to an Action | [foo: "bar", bar: "gurka"]


## Action

An Action document is what is sent to an Action Endpoint. Parameters in the Action document can be retrieved from an Action Descriptor or a web form.

```go
type Action struct {
	Base
	Role    string            `json:"role"`
	Mandate string            `json:"mandate,omitempty"`
	Nonce   string            `json:"nonce,omitempty"`
	Params  map[string]string `json:"params,omitempty"`
	Facts   []Part            `json:"facts,omitempty"`
}
```

| Label         | Description     | Example
|:--------------|:----------------|:--------
| Base          | This includes fields from the base document type | N/A
| Role          | The ID of a Role | "admin@example.com"
| Mandate       | JSON copy of the Mandate performing the Action | (Encrypted and signed JSON)
| Nonce         | Nonce used to avoid replay attacks | "sdkcjsdocijwe09uidf34e32r"
| Params        | All parameters needed to perform the Action | [ resource: "s09dufdf", foo: "bar" ]
| Facts         | Any Facts needed to perform the action | See the Facts document


## ActionDescriptor

An ActionDescriptor is what you access to perform an Action. The ActionDescriptor object be accessed through an API, or a QR code or something completely different.

```go
type ActionDescriptor struct {
	Base
	Label     string            `json:"label"`
	Roles     []string          `json:"roles"`
	UIURI     string            `json:"uiURI,omitempty"`
	UIData    string            `json:"uiData,omitempty"`
	NonceType string            `json:"nonceType,omitempty"`
	Nonce     string            `json:"nonce,omitempty"`
	NonceURI  string            `json:"nonceURI,omitempty"`
	ActionURI string            `json:"actionURI"`
	Params    map[string]string `json:"params,omitempty"`
	Scopes    []string          `json:"scopes,omitempty"`
}
```

| Label         | Description     | Example
|:--------------|:----------------|:--------
| Base          | This includes fields from the base document type | N/A
| Roles         | A list of IDs of Roles | ["admin@example.com", "employee@example.com"]
| UIURI         | The URI of a User Interface | "https://example.com/ui"
| UIData        | Any extra data provided to the UI | ?
| NonceType     | ? | ?
| NonceURI      | The URI to get the Nonce from | "https://example.com/nonce"
| ActionURI     | Where to post the final Action object | "https://example.com/action"
| Params        | Parameters to send to an Action | [foo: "bar", bar: "gurka"]
| Scopes        | Scopes to send to an Action | See Scope



## Receipt

A Receipt is a document proving the result of an Action. The Base document has the timestamp of when it was issued. The Intervals field is used for showing upcoming events, such as when the Receipt is issued for an upcoming event or booking - they can also be recurring.

```go
type Receipt struct {
	Base
	Role      string     `json:"role,omitempty"`
	Action    string     `json:"action,omitempty"`
	URI       string     `json:"viewuri,omitempty"`
	JWT       string     `json:"jwt,omitempty"`
	Intervals []Interval `json:"intervals,omitempty"`
	Label     string     `json:"label,omitempty"`
}
```

| Label         | Description     | Example
|:--------------|:----------------|:--------
| Base          | This includes fields from the base document type | N/A
| Role          | The ID of a Role | "admin@example.com"
| Label         | The Label of the Receipt | "Booking of Meeting Room Z"
| Action        | JSON copy of the Action giving the Receipt | (Encrypted and signed JSON)
| URI           | A reference to a User Interface displaying the Receipt | (URL or inline HTML)
| JWT           | JWT containing the data needed to make any changes to resources in the Receipt | ?
| Intervals     | A list of intervals for upcoming events | [ {Start: "2017-02-15T07:00:00+01:00", End: "2017-02-15T17:30:00+01:00"} ]
| Label         | A descriptive Label readable by users | "Receipt for Coffee", "Event in room Z"


## Fact

A fact is a proof of id derived from using a certain service (i.e. e-mail address with SMTP, phone number using SMS).

```go
type Fact struct {
  Base
  TTL    time.Duration          `json:"ttl,omitempty"`
  Issuer string                 `json:"iss,omitempty"`
  Label  string                 `json:"label,omitempty"`
  Data   map[string]interface{} `json:"data"`
}
```

| Label         | Description     | Example
|:--------------|:----------------|:--------
| Base          | This includes fields from the base document type | N/A
| TTL           | Seconds until expiration after the document was created (Timestamp) | 86400 (24hr)
| Issuer        | The issuer in plaintext | "Facebook", "e-mail" ...
| Label         | The id          | "user@example.com"
| Data          | ?               | ?

## Message

A message with Title and a message body, in plain text.

```go
type Message struct {
	Base
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}
```

| Label         | Description     | Example
|:--------------|:----------------|:--------
| Base          | This includes fields from the base document type | N/A
| Title         | The title of the Message | "Hello"
| Message       | The body of the Message  | "Longer text"


## PushMessage

Used to send messages with push notifications.

```go
type PushMessage struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	URI     string `json:"uri,omitempty"`
	Data    string `json:"data,omitempty"`
}
```

| Label         | Description     | Example
|:--------------|:----------------|:--------
| Title         | The title of the PushMessage | "Hello"
| Message       | The body of the PushMessage | "Long text"
| URI           | A reference to a Brickchain document | "https://example.com/document"
| Data          | A Brickchain document | Inline JSON


## Todo...

Documents still to be described:

 * Multipart
 * Realm
 * RealmDescriptor
 * ScopeRequest
 * SignatureRequest
 * Membership
