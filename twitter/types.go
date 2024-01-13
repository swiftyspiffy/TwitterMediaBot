package twitter

type GetGuestTokenResponse struct {
	GuestToken string `json:"guest_token"`
}

type Tweet struct {
	Data Data `json:"data,omitempty"`
}
type AffiliatesHighlightedLabel struct {
}
type Description struct {
	Urls []any `json:"urls,omitempty"`
}
type LegacyEntities struct {
	Description Description `json:"description,omitempty"`
}
type ResultLegacy struct {
	CreatedAt               string         `json:"created_at,omitempty"`
	DefaultProfile          bool           `json:"default_profile,omitempty"`
	DefaultProfileImage     bool           `json:"default_profile_image,omitempty"`
	Description             string         `json:"description,omitempty"`
	Entities                LegacyEntities `json:"entities,omitempty"`
	FastFollowersCount      int            `json:"fast_followers_count,omitempty"`
	FavouritesCount         int            `json:"favourites_count,omitempty"`
	FollowersCount          int            `json:"followers_count,omitempty"`
	FriendsCount            int            `json:"friends_count,omitempty"`
	HasCustomTimelines      bool           `json:"has_custom_timelines,omitempty"`
	IsTranslator            bool           `json:"is_translator,omitempty"`
	ListedCount             int            `json:"listed_count,omitempty"`
	Location                string         `json:"location,omitempty"`
	MediaCount              int            `json:"media_count,omitempty"`
	Name                    string         `json:"name,omitempty"`
	NormalFollowersCount    int            `json:"normal_followers_count,omitempty"`
	PinnedTweetIdsStr       []any          `json:"pinned_tweet_ids_str,omitempty"`
	PossiblySensitive       bool           `json:"possibly_sensitive,omitempty"`
	ProfileBannerURL        string         `json:"profile_banner_url,omitempty"`
	ProfileImageURLHTTPS    string         `json:"profile_image_url_https,omitempty"`
	ProfileInterstitialType string         `json:"profile_interstitial_type,omitempty"`
	ScreenName              string         `json:"screen_name,omitempty"`
	StatusesCount           int            `json:"statuses_count,omitempty"`
	TranslatorType          string         `json:"translator_type,omitempty"`
	Verified                bool           `json:"verified,omitempty"`
	WithheldInCountries     []any          `json:"withheld_in_countries,omitempty"`
}
type Category struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	IconName string `json:"icon_name,omitempty"`
}
type Professional struct {
	RestID           string     `json:"rest_id,omitempty"`
	ProfessionalType string     `json:"professional_type,omitempty"`
	Category         []Category `json:"category,omitempty"`
}
type UserResult struct {
	Typename                   string                     `json:"__typename,omitempty"`
	ID                         string                     `json:"id,omitempty"`
	RestID                     string                     `json:"rest_id,omitempty"`
	AffiliatesHighlightedLabel AffiliatesHighlightedLabel `json:"affiliates_highlighted_label,omitempty"`
	IsBlueVerified             bool                       `json:"is_blue_verified,omitempty"`
	ProfileImageShape          string                     `json:"profile_image_shape,omitempty"`
	Legacy                     ResultLegacy               `json:"legacy,omitempty"`
	Professional               Professional               `json:"professional,omitempty"`
}
type UserResults struct {
	Result UserResult `json:"result,omitempty"`
}
type Core struct {
	UserResults UserResults `json:"user_results,omitempty"`
}
type UnmentionData struct {
}
type EditControl struct {
	EditTweetIds       []string `json:"edit_tweet_ids,omitempty"`
	EditableUntilMsecs string   `json:"editable_until_msecs,omitempty"`
	IsEditEligible     bool     `json:"is_edit_eligible,omitempty"`
	EditsRemaining     string   `json:"edits_remaining,omitempty"`
}
type Views struct {
	State string `json:"state,omitempty"`
}
type AdditionalMediaInfo struct {
	Monetizable bool `json:"monetizable,omitempty"`
}
type ExtMediaAvailability struct {
	Status string `json:"status,omitempty"`
}
type Large struct {
	H      int    `json:"h,omitempty"`
	W      int    `json:"w,omitempty"`
	Resize string `json:"resize,omitempty"`
}
type Medium struct {
	H      int    `json:"h,omitempty"`
	W      int    `json:"w,omitempty"`
	Resize string `json:"resize,omitempty"`
}
type Small struct {
	H      int    `json:"h,omitempty"`
	W      int    `json:"w,omitempty"`
	Resize string `json:"resize,omitempty"`
}
type Thumb struct {
	H      int    `json:"h,omitempty"`
	W      int    `json:"w,omitempty"`
	Resize string `json:"resize,omitempty"`
}
type Sizes struct {
	Large  Large  `json:"large,omitempty"`
	Medium Medium `json:"medium,omitempty"`
	Small  Small  `json:"small,omitempty"`
	Thumb  Thumb  `json:"thumb,omitempty"`
}
type OriginalInfo struct {
	Height     int   `json:"height,omitempty"`
	Width      int   `json:"width,omitempty"`
	FocusRects []any `json:"focus_rects,omitempty"`
}
type Variants struct {
	ContentType string `json:"content_type,omitempty"`
	URL         string `json:"url,omitempty"`
	Bitrate     int    `json:"bitrate,omitempty"`
}
type VideoInfo struct {
	AspectRatio    []int      `json:"aspect_ratio,omitempty"`
	DurationMillis int        `json:"duration_millis,omitempty"`
	Variants       []Variants `json:"variants,omitempty"`
}
type Media struct {
	DisplayURL           string               `json:"display_url,omitempty"`
	ExpandedURL          string               `json:"expanded_url,omitempty"`
	IDStr                string               `json:"id_str,omitempty"`
	Indices              []int                `json:"indices,omitempty"`
	MediaKey             string               `json:"media_key,omitempty"`
	MediaURLHTTPS        string               `json:"media_url_https,omitempty"`
	Type                 string               `json:"type,omitempty"`
	URL                  string               `json:"url,omitempty"`
	AdditionalMediaInfo  AdditionalMediaInfo  `json:"additional_media_info,omitempty"`
	ExtMediaAvailability ExtMediaAvailability `json:"ext_media_availability,omitempty"`
	Sizes                Sizes                `json:"sizes,omitempty"`
	OriginalInfo         OriginalInfo         `json:"original_info,omitempty"`
	VideoInfo            VideoInfo            `json:"video_info,omitempty"`
}
type Entities struct {
	Hashtags     []any   `json:"hashtags,omitempty"`
	Media        []Media `json:"media,omitempty"`
	Symbols      []any   `json:"symbols,omitempty"`
	Timestamps   []any   `json:"timestamps,omitempty"`
	Urls         []any   `json:"urls,omitempty"`
	UserMentions []any   `json:"user_mentions,omitempty"`
}
type ExtendedEntities struct {
	Media []Media `json:"media,omitempty"`
}
type Legacy struct {
	BookmarkCount             int              `json:"bookmark_count,omitempty"`
	Bookmarked                bool             `json:"bookmarked,omitempty"`
	CreatedAt                 string           `json:"created_at,omitempty"`
	ConversationIDStr         string           `json:"conversation_id_str,omitempty"`
	DisplayTextRange          []int            `json:"display_text_range,omitempty"`
	Entities                  Entities         `json:"entities,omitempty"`
	ExtendedEntities          ExtendedEntities `json:"extended_entities,omitempty"`
	FavoriteCount             int              `json:"favorite_count,omitempty"`
	Favorited                 bool             `json:"favorited,omitempty"`
	FullText                  string           `json:"full_text,omitempty"`
	IsQuoteStatus             bool             `json:"is_quote_status,omitempty"`
	Lang                      string           `json:"lang,omitempty"`
	PossiblySensitive         bool             `json:"possibly_sensitive,omitempty"`
	PossiblySensitiveEditable bool             `json:"possibly_sensitive_editable,omitempty"`
	QuoteCount                int              `json:"quote_count,omitempty"`
	ReplyCount                int              `json:"reply_count,omitempty"`
	RetweetCount              int              `json:"retweet_count,omitempty"`
	Retweeted                 bool             `json:"retweeted,omitempty"`
	UserIDStr                 string           `json:"user_id_str,omitempty"`
	IDStr                     string           `json:"id_str,omitempty"`
}
type Result struct {
	Typename       string        `json:"__typename,omitempty"`
	RestID         string        `json:"rest_id,omitempty"`
	Core           Core          `json:"core,omitempty"`
	UnmentionData  UnmentionData `json:"unmention_data,omitempty"`
	EditControl    EditControl   `json:"edit_control,omitempty"`
	IsTranslatable bool          `json:"is_translatable,omitempty"`
	Views          Views         `json:"views,omitempty"`
	Source         string        `json:"source,omitempty"`
	Legacy         Legacy        `json:"legacy,omitempty"`
}
type TweetResult struct {
	Result Result `json:"result,omitempty"`
}
type Data struct {
	TweetResult TweetResult `json:"tweetResult,omitempty"`
}
