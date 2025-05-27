package constants

type Link string

const (
	LinkVerifyCode Link = "https://id.kick.com/oauth/authorize"
	LinkToken      Link = "https://id.kick.com/oauth/token"
	LinkUsers      Link = "https://api.kick.com/public/v1/users"
	LinkChannels   Link = "https://api.kick.com/public/v1/channels"
)

type Scopes string

const (
	ScopesUserRead        Scopes = "user:read"
	ScopesChannelRead     Scopes = "channel:read"
	ScopesChannelWrite    Scopes = "channel:write"
	ScopesChatWrite       Scopes = "chat:write"
	ScopesStreamKeyRead   Scopes = "streamkey:read"
	ScopesEventsSubscribe Scopes = "events:subscribe"
	ScopesModerationBan   Scopes = "moderation:ban"
)

func (link Link) ToString() string {
	return string(link)
}

func (scopes Scopes) ToString() string {
	return string(scopes)
}
