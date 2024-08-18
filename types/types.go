package types

type AuthRequest struct {
	ResponseType string `json:"response_type"`
	ClientID     string `json:"client_id"`
	Scope        string `json:"scope"`
	RedirectUri  string `json:"redirect_uri"`
	State        string `json:"state"`
}

type AuthTokenRequest struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
}

type ClientResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int64  `json:"expires_in"`
}

type CurrentPlayingResponse struct {
	Item struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Album struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Images []struct {
				Url    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"images"`
		} `json:"album"`
		Artists []struct {
			ID   string `json:"id"`
			Href string `json:"href"`
			Name string `json:"name"`
			Type string `json:"type"`
			Uri  string `json:"uri"`
		} `json:"artists"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		DurationMs int `json:"duration_ms"`
	} `json:"item"`
}

type SaveRequest struct {
	IDs     []string `json:"ids"`
	IsSaved bool     `json:"is_saved"`
}
