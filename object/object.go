package object

import "encoding/json"

// AccountNameRequest struct
type AccountNameRequest struct {
	FirstName string `json:"first_name"`
	ID        int    `json:"id"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
}

// AccountPushConversations struct
type AccountPushConversations struct {
	Count int                            `json:"count"`
	Items []AccountPushConversationsItem `json:"items"`
}

// AccountPushConversationsItem struct
type AccountPushConversationsItem struct {
	DisabledUntil int `json:"disabled_until"`
	PeerID        int `json:"peer_id"`
	Sound         int `json:"sound"`
}

// AccountPushParams struct
type AccountPushParams struct {
	AppRequest     []string `json:"app_request"`
	Birthday       []string `json:"birthday"`
	Chat           []string `json:"chat"`
	Comment        []string `json:"comment"`
	EventSoon      []string `json:"event_soon"`
	Friend         []string `json:"friend"`
	FriendAccepted []string `json:"friend_accepted"`
	FriendFound    []string `json:"friend_found"`
	GroupAccepted  []string `json:"group_accepted"`
	GroupInvite    []string `json:"group_invite"`
	Like           []string `json:"like"`
	Mention        []string `json:"mention"`
	Msg            []string `json:"msg"`
	NewPost        []string `json:"new_post"`
	PhotosTag      []string `json:"photos_tag"`
	Reply          []string `json:"reply"`
	Repost         []string `json:"repost"`
	SdkOpen        []string `json:"sdk_open"`
	WallPost       []string `json:"wall_post"`
	WallPublish    []string `json:"wall_publish"`
}

// AccountOffer struct
type AccountOffer struct {
	Description      string `json:"description"`
	ID               int    `json:"id"`
	Img              string `json:"img"`
	Instruction      string `json:"instruction"`
	InstructionHTML  string `json:"instruction_html"`
	Price            int    `json:"price"`
	ShortDescription string `json:"short_description"`
	Tag              string `json:"tag"`
	Title            string `json:"title"`
}

// AudioAudioFull struct
type AudioAudioFull struct {
	ID          int    `json:"id"`
	OwnerID     int    `json:"owner_id"`
	Artist      string `json:"artist"`
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	Date        int    `json:"date"`
	URL         string `json:"url"`
	IsHq        bool   `json:"is_hq"`
	LyricsID    int    `json:"lyrics_id"`
	AlbumID     int    `json:"album_id"`
	GenreID     int    `json:"genre_id"`
	TrackCode   string `json:"track_code"`
	IsExplicit  bool   `json:"is_explicit"`
	NoSearch    int    `json:"no_search"`
	MainArtists []struct {
		Name   string `json:"name"`
		ID     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"main_artists"`
}

// BaseCountry struct
type BaseCountry struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// BaseObject struct
type BaseObject struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// BaseObjectCount struct
type BaseObjectCount struct {
	Count int `json:"count"`
}

// BaseObjectWithName struct
type BaseObjectWithName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// BaseRequestParam struct
type BaseRequestParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GroupEvent struct
type GroupEvent struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
}

// GiftsGift struct
type GiftsGift struct {
	Date     int         `json:"date"`
	FromID   int         `json:"from_id"`
	Gift     GiftsLayout `json:"gift"`
	GiftHash string      `json:"gift_hash"`
	ID       int         `json:"id"`
	Message  string      `json:"message"`
	Privacy  int         `json:"privacy"`
}

// GiftsLayout struct
type GiftsLayout struct {
	ID       int    `json:"id"`
	Thumb256 string `json:"thumb_256"`
	Thumb48  string `json:"thumb_48"`
	Thumb96  string `json:"thumb_96"`
}

// GroupsGroup struct
type GroupsGroup struct {
	AdminLevel  int    `json:"admin_level"`
	Deactivated string `json:"deactivated"`
	ID          int    `json:"id"`
	IsAdmin     int    `json:"is_admin"`
	IsClosed    int    `json:"is_closed"`
	IsMember    int    `json:"is_member"`
	Name        string `json:"name"`
	Photo100    string `json:"photo_100"`
	Photo200    string `json:"photo_200"`
	Photo50     string `json:"photo_50"`
	ScreenName  string `json:"screen_name"`
	Type        string `json:"type"`
}

// MessagesMessage struct
type MessagesMessage struct {
	AdminAuthorID         int                         `json:"admin_author_id"`
	Action                messagesMessageAction       `json:"action"`
	Attachments           []messagesMessageAttachment `json:"attachments"`
	ConversationMessageID int                         `json:"conversation_message_id"`
	Date                  int                         `json:"date"`
	FromID                int                         `json:"from_id"`
	FwdMessages           []MessagesMessage           `json:"fwd_messages"`
	Geo                   baseGeo                     `json:"geo"`
	ID                    int                         `json:"id"`
	Important             bool                        `json:"important"`
	IsHidden              bool                        `json:"is_hidden"`
	Keyboard              MessagesKeyboard            `json:"keyboard"`
	Out                   int                         `json:"out"`
	Payload               string                      `json:"payload"`
	PeerID                int                         `json:"peer_id"`
	RandomID              int                         `json:"random_id"`
	Text                  string                      `json:"text"`
	UpdateTime            int                         `json:"update_time"`
}

// PhotosPhoto struct
type PhotosPhoto struct {
	AccessKey string        `json:"access_key"`
	AlbumID   int           `json:"album_id"`
	Date      int           `json:"date"`
	Height    int           `json:"height"`
	ID        int           `json:"id"`
	Images    []photosImage `json:"images"`
	Lat       float64       `json:"lat"`
	Long      float64       `json:"long"`
	OwnerID   int           `json:"owner_id"`
	PostID    int           `json:"post_id"`
	Text      string        `json:"text"`
	UserID    int           `json:"user_id"`
	Width     int           `json:"width"`
}

// UsersUser struct
type UsersUser struct {
	ID                     int               `json:"id"`
	FirstName              string            `json:"first_name"`
	LastName               string            `json:"last_name"`
	IsClosed               bool              `json:"is_closed"`
	CanAccessClosed        bool              `json:"can_access_closed"`
	Sex                    int               `json:"sex"`
	Nickname               string            `json:"nickname"`
	Domain                 string            `json:"domain"`
	ScreenName             string            `json:"screen_name"`
	Bdate                  string            `json:"bdate"`
	City                   BaseObject        `json:"city"`
	Country                BaseObject        `json:"country"`
	Photo50                string            `json:"photo_50"`
	Photo100               string            `json:"photo_100"`
	Photo200               string            `json:"photo_200"`
	PhotoMax               string            `json:"photo_max"`
	Photo200Orig           string            `json:"photo_200_orig"`
	Photo400Orig           string            `json:"photo_400_orig"`
	PhotoMaxOrig           string            `json:"photo_max_orig"`
	PhotoID                string            `json:"photo_id"`
	HasPhoto               int               `json:"has_photo"`
	HasMobile              int               `json:"has_mobile"`
	IsFriend               int               `json:"is_friend"`
	FriendStatus           int               `json:"friend_status"`
	Online                 int               `json:"online"`
	CanPost                int               `json:"can_post"`
	CanSeeAllPosts         int               `json:"can_see_all_posts"`
	CanSeeAudio            int               `json:"can_see_audio"`
	CanWritePrivateMessage int               `json:"can_write_private_message"`
	CanSendFriendRequest   int               `json:"can_send_friend_request"`
	Facebook               string            `json:"facebook"`
	FacebookName           string            `json:"facebook_name"`
	Twitter                string            `json:"twitter"`
	Instagram              string            `json:"instagram"`
	Site                   string            `json:"site"`
	Status                 string            `json:"status"`
	StatusAudio            AudioAudioFull    `json:"status_audio"`
	LastSeen               usersLastSeen     `json:"last_seen"`
	CropPhoto              usersCropPhoto    `json:"crop_photo"`
	Verified               int               `json:"verified"`
	FollowersCount         int               `json:"followers_count"`
	Blacklisted            int               `json:"blacklisted"`
	BlacklistedByMe        int               `json:"blacklisted_by_me"`
	IsFavorite             int               `json:"is_favorite"`
	IsHiddenFromFeed       int               `json:"is_hidden_from_feed"`
	CommonCount            int               `json:"common_count"`
	Occupation             usersOccupation   `json:"occupation"`
	Career                 usersCareer       `json:"career"`
	Military               usersMilitary     `json:"military"`
	University             int               `json:"university"`
	UniversityName         string            `json:"university_name"`
	Faculty                int               `json:"faculty"`
	FacultyName            string            `json:"faculty_name"`
	Graduation             int               `json:"graduation"`
	HomeTown               string            `json:"home_town"`
	Relation               int               `json:"relation"`
	Personal               usersPersonal     `json:"personal"`
	Interests              string            `json:"interests"`
	Music                  string            `json:"music"`
	Activities             string            `json:"activities"`
	Movies                 string            `json:"movies"`
	Tv                     string            `json:"tv"`
	Books                  string            `json:"books"`
	Games                  string            `json:"games"`
	Universities           []usersUniversity `json:"universities"`
	Schools                []usersSchool     `json:"schools"`
	About                  string            `json:"about"`
	Relatives              []usersRelative   `json:"relatives"`
	Quotes                 string            `json:"quotes"`
	Lists                  []int             `json:"lists"`
	Deactivated            string            `json:"deactivated"`
	WallDefault            string            `json:"wall_default"`
	Trending               int               `json:"trending"`
	Timezone               int               `json:"timezone"`
	MaidenName             string            `json:"maiden_name"`
	Exports                usersExports      `json:"exports"`
	Counters               usersUserCounters `json:"counters"`
	Contacts               usersContacts     `json:"contacts"`
	// TODO education
}

// UsersUserMin struct
type UsersUserMin struct {
	Deactivated string `json:"deactivated"`
	FirstName   string `json:"first_name"`
	Hidden      int    `json:"hidden"`
	ID          int    `json:"id"`
	LastName    string `json:"last_name"`
}

type usersContacts struct {
	MobilePhone string `json:"mobile_phone"`
	HomePhone   string `json:"home_phone"`
}

type usersCareer struct {
	CityID    int    `json:"city_id"`
	Company   string `json:"company"`
	CountryID int    `json:"country_id"`
	From      int    `json:"from"`
	GroupID   int    `json:"group_id"`
	Position  string `json:"position"`
	Until     int    `json:"until"`
}

type usersCropPhoto struct {
	Crop  usersCropPhotoCrop `json:"crop"`
	Photo PhotosPhoto        `json:"photo"`
	Rect  usersCropPhotoRect `json:"rect"`
}

type usersCropPhotoCrop struct {
	X  float64 `json:"x"`
	X2 float64 `json:"x2"`
	Y  float64 `json:"y"`
	Y2 float64 `json:"y2"`
}

type usersCropPhotoRect struct {
	X  float64 `json:"x"`
	X2 float64 `json:"x2"`
	Y  float64 `json:"y"`
	Y2 float64 `json:"y2"`
}

type usersExports struct {
	Facebook    int `json:"facebook"`
	Livejournal int `json:"livejournal"`
	Twitter     int `json:"twitter"`
}

type usersLastSeen struct {
	Platform int `json:"platform"`
	Time     int `json:"time"`
}

type usersMilitary struct {
	CountryID int    `json:"country_id"`
	From      int    `json:"from"`
	Unit      string `json:"unit"`
	UnitID    int    `json:"unit_id"`
	Until     int    `json:"until"`
}

type usersOccupation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type usersPersonal struct {
	Alcohol    int      `json:"alcohol"`
	InspiredBy string   `json:"inspired_by"`
	Langs      []string `json:"langs"`
	LifeMain   int      `json:"life_main"`
	PeopleMain int      `json:"people_main"`
	Political  int      `json:"political"`
	Religion   string   `json:"religion"`
	Smoking    int      `json:"smoking"`
}

type usersRelative struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type usersSchool struct {
	City          int    `json:"city"`
	Class         string `json:"class"`
	Country       int    `json:"country"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	Type          int    `json:"type"`
	TypeStr       string `json:"type_str"`
	YearFrom      int    `json:"year_from"`
	YearGraduated int    `json:"year_graduated"`
	YearTo        int    `json:"year_to"`
}

type usersUniversity struct {
	Chair           int    `json:"chair"`
	ChairName       string `json:"chair_name"`
	City            int    `json:"city"`
	Country         int    `json:"country"`
	EducationForm   string `json:"education_form"`
	EducationStatus string `json:"education_status"`
	Faculty         int    `json:"faculty"`
	FacultyName     string `json:"faculty_name"`
	Graduation      int    `json:"graduation"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
}

type usersUserBroadcast struct {
}

type usersUserCounters struct {
	Albums        int `json:"albums"`
	Audios        int `json:"audios"`
	Followers     int `json:"followers"`
	Friends       int `json:"friends"`
	Gifts         int `json:"gifts"`
	Groups        int `json:"groups"`
	Notes         int `json:"notes"`
	OnlineFriends int `json:"online_friends"`
	Pages         int `json:"pages"`
	Photos        int `json:"photos"`
	Subscriptions int `json:"subscriptions"`
	UserPhotos    int `json:"user_photos"`
	UserVideos    int `json:"user_videos"`
	Videos        int `json:"videos"`
}

type usersUserFull struct {
}

type usersUserFullXtrType struct {
}

type usersUserLim struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	NameGen string `json:"name_gen"`
	Photo   string `json:"photo"`
}

// VideoVideo struct
type VideoVideo struct {
	AccessKey   string             `json:"access_key"`
	AddingDate  int                `json:"adding_date"`
	CanAdd      int                `json:"can_add"`
	CanEdit     int                `json:"can_edit"`
	Comments    int                `json:"comments"`
	Date        int                `json:"date"`
	Description string             `json:"description"`
	Duration    int                `json:"duration"`
	Files       videoVideoFiles    `json:"files"`
	ID          int                `json:"id"`
	Live        basePropertyExists `json:"live"`
	OwnerID     int                `json:"owner_id"`
	Photo130    string             `json:"photo_130"`
	Photo320    string             `json:"photo_320"`
	Photo800    string             `json:"photo_800"`
	Player      string             `json:"player"`
	Processing  basePropertyExists `json:"processing"`
	Title       string             `json:"title"`
	Views       int                `json:"views"`
}

// Button struct
type Button []MessagesKeyboardButton

// MessagesKeyboard struct
type MessagesKeyboard struct {
	Buttons []Button `json:"buttons"`
	OneTime bool     `json:"one_time"`
}

// MessagesKeyboardButton struct
type MessagesKeyboardButton struct {
	Action MessagesKeyboardButtonAction `json:"action"`
	Color  string                       `json:"color"`
}

// MessagesKeyboardButtonAction struct
type MessagesKeyboardButtonAction struct {
	Label   string `json:"label"`
	Payload string `json:"payload"`
	Type    string `json:"type"`
}

//-----------------------------
// Далее сгенерированные объекты на основе JSON schema
//-----------------------------

type adsAccessRole string

type adsAccesses struct {
	ClientID string        `json:"client_id"`
	Role     adsAccessRole `json:"role"`
}

type adsAccount struct {
	AccessRole    adsAccessRole  `json:"access_role"`
	AccountID     int            `json:"account_id"`
	AccountStatus int            `json:"account_status"`
	AccountType   adsAccountType `json:"account_type"`
}

type adsAccountType string

type adsAdApproved int

type adsAdCostType int

type adsAdLayout struct {
	AdFormat    int                 `json:"ad_format"`
	CampaignID  int                 `json:"campaign_id"`
	CostType    adsAdLayoutCostType `json:"cost_type"`
	Description string              `json:"description"`
	ID          int                 `json:"id"`
	ImageSrc    int                 `json:"image_src"`
	ImageSrc2x  int                 `json:"image_src_2x"`
	LinkDomain  string              `json:"link_domain"`
	LinkURL     string              `json:"link_url"`
	PreviewLink string              `json:"preview_link"`
	Title       string              `json:"title"`
	Video       int                 `json:"video"`
}

type adsAdLayoutCostType int

type adsAdStatus int

type adsCampaign struct {
	AllLimit  string            `json:"all_limit"`
	DayLimit  string            `json:"day_limit"`
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	StartTime int               `json:"start_time"`
	Status    adsCampaignStatus `json:"status"`
	StopTime  int               `json:"stop_time"`
	Type      adsCampaignType   `json:"type"`
}

type adsCampaignStatus int

type adsCampaignType string

type adsCategory struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Subcategories []BaseObjectWithName `json:"subcategories"`
}

type adsClient struct {
	AllLimit string `json:"all_limit"`
	DayLimit string `json:"day_limit"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
}

type adsCriteria struct {
	AgeFrom              int                `json:"age_from"`
	AgeTo                int                `json:"age_to"`
	Apps                 string             `json:"apps"`
	AppsNot              string             `json:"apps_not"`
	Birthday             int                `json:"birthday"`
	Cities               string             `json:"cities"`
	CitiesNot            string             `json:"cities_not"`
	Country              int                `json:"country"`
	Districts            string             `json:"districts"`
	Groups               string             `json:"groups"`
	InterestCategories   string             `json:"interest_categories"`
	Interests            string             `json:"interests"`
	Paying               int                `json:"paying"`
	Positions            string             `json:"positions"`
	Religions            string             `json:"religions"`
	RetargetingGroups    string             `json:"retargeting_groups"`
	RetargetingGroupsNot string             `json:"retargeting_groups_not"`
	SchoolFrom           int                `json:"school_from"`
	SchoolTo             int                `json:"school_to"`
	Schools              string             `json:"schools"`
	Sex                  adsCriteriaSex     `json:"sex"`
	Stations             string             `json:"stations"`
	Statuses             string             `json:"statuses"`
	Streets              string             `json:"streets"`
	Travellers           basePropertyExists `json:"travellers"`
	UniFrom              int                `json:"uni_from"`
	UniTo                int                `json:"uni_to"`
	UserBrowsers         string             `json:"user_browsers"`
	UserDevices          string             `json:"user_devices"`
	UserOs               string             `json:"user_os"`
}

type adsCriteriaSex int

type adsDemoStats struct {
	ID    int                `json:"id"`
	Stats adsDemostatsFormat `json:"stats"`
	Type  adsObjectType      `json:"type"`
}

type adsDemostatsFormat struct {
	Age     []adsStatsAge    `json:"age"`
	Cities  []adsStatsCities `json:"cities"`
	Day     string           `json:"day"`
	Month   string           `json:"month"`
	Overall int              `json:"overall"`
	Sex     []adsStatsSex    `json:"sex"`
	SexAge  []adsStatsSexAge `json:"sex_age"`
}

type adsFloodStats struct {
	Left    int `json:"left"`
	Refresh int `json:"refresh"`
}

type adsLinkStatus struct {
	Description string `json:"description"`
	RedirectURL string `json:"redirect_url"`
	Status      string `json:"status"`
}

type adsObjectType string

type adsParagraphs struct {
	Paragraph string `json:"paragraph"`
}

type adsPostStats struct {
}

type adsRejectReason struct {
	Comment string     `json:"comment"`
	Rules   []adsRules `json:"rules"`
}

type adsRules struct {
	Paragraphs []adsParagraphs `json:"paragraphs"`
	Title      string          `json:"title"`
}

type adsStats struct {
	ID    int            `json:"id"`
	Stats adsStatsFormat `json:"stats"`
	Type  adsObjectType  `json:"type"`
}

type adsStatsAge struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Value           string  `json:"value"`
}

type adsStatsCities struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Name            string  `json:"name"`
	Value           int     `json:"value"`
}

type adsStatsFormat struct {
	Clicks          int    `json:"clicks"`
	Day             string `json:"day"`
	Impressions     int    `json:"impressions"`
	JoinRate        int    `json:"join_rate"`
	Month           string `json:"month"`
	Overall         int    `json:"overall"`
	Reach           int    `json:"reach"`
	Spent           int    `json:"spent"`
	VideoClicksSite int    `json:"video_clicks_site"`
	VideoViews      int    `json:"video_views"`
	VideoViewsFull  int    `json:"video_views_full"`
	VideoViewsHalf  int    `json:"video_views_half"`
}

type adsStatsSex struct {
	ClicksRate      float64          `json:"clicks_rate"`
	ImpressionsRate float64          `json:"impressions_rate"`
	Value           adsStatsSexValue `json:"value"`
}

type adsStatsSexAge struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Value           string  `json:"value"`
}

type adsStatsSexValue string

type adsTargSettings struct {
}

type adsTargStats struct {
	AudienceCount  int     `json:"audience_count"`
	RecommendedCpc float64 `json:"recommended_cpc"`
	RecommendedCpm float64 `json:"recommended_cpm"`
}

type adsTargSuggestions struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type adsTargSuggestionsCities struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

type adsTargSuggestionsRegions struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type adsTargSuggestionsSchools struct {
	Desc   string                        `json:"desc"`
	ID     int                           `json:"id"`
	Name   string                        `json:"name"`
	Parent string                        `json:"parent"`
	Type   adsTargSuggestionsSchoolsType `json:"type"`
}

type adsTargSuggestionsSchoolsType string

type adsTargetGroup struct {
	AudienceCount int    `json:"audience_count"`
	Domain        string `json:"domain"`
	ID            int    `json:"id"`
	Lifetime      int    `json:"lifetime"`
	Name          string `json:"name"`
	Pixel         string `json:"pixel"`
}

type adsUsers struct {
	Accesses []adsAccesses `json:"accesses"`
	UserID   int           `json:"user_id"`
}

type appsApp struct {
	AuthorGroup     int                    `json:"author_group"`
	AuthorID        int                    `json:"author_id"`
	AuthorURL       string                 `json:"author_url"`
	Banner1120      string                 `json:"banner_1120"`
	Banner560       string                 `json:"banner_560"`
	CatalogPosition int                    `json:"catalog_position"`
	Description     string                 `json:"description"`
	Genre           string                 `json:"genre"`
	GenreID         int                    `json:"genre_id"`
	ID              int                    `json:"id"`
	Icon139         string                 `json:"icon_139"`
	Icon150         string                 `json:"icon_150"`
	Icon278         string                 `json:"icon_278"`
	Icon75          string                 `json:"icon_75"`
	International   int                    `json:"international"`
	IsInCatalog     int                    `json:"is_in_catalog"`
	LeaderboardType appsAppLeaderboardType `json:"leaderboard_type"`
	MembersCount    int                    `json:"members_count"`
	PlatformID      int                    `json:"platform_id"`
	PublishedDate   int                    `json:"published_date"`
	ScreenName      string                 `json:"screen_name"`
	Screenshots     []PhotosPhoto          `json:"screenshots"`
	Section         string                 `json:"section"`
	Title           string                 `json:"title"`
	Type            appsAppType            `json:"type"`
}

type appsAppLeaderboardType int

type appsAppType string

type appsLeaderboard struct {
	Level  int `json:"level"`
	Points int `json:"points"`
	Score  int `json:"score"`
	UserID int `json:"user_id"`
}

type audioAudio struct {
	AccessKey string `json:"access_key"`
	Artist    string `json:"artist"`
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}

type audioAudioUploadResponse struct {
	Audio    string `json:"audio"`
	Hash     string `json:"hash"`
	Redirect string `json:"redirect"`
	Server   int    `json:"server"`
}

type audioLyrics struct {
	LyricsID int    `json:"lyrics_id"`
	Text     string `json:"text"`
}

type baseCity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type baseCommentsInfo struct {
	CanPost       int `json:"can_post"`
	Count         int `json:"count"`
	GroupsCanPost int `json:"groups_can_post"`
}

type baseError struct {
	ErrorCode     int                `json:"error_code"`
	ErrorMsg      string             `json:"error_msg"`
	RequestParams []BaseRequestParam `json:"request_params"`
}

type baseGeo struct {
	Coordinates baseGeoCoordinates `json:"coordinates"`
	Place       basePlace          `json:"place"`
	Showmap     int                `json:"showmap"`
	Type        string             `json:"type"`
}

type baseGeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type baseImage struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type baseLikes struct {
	Count     int `json:"count"`
	UserLikes int `json:"user_likes"`
}

type baseLikesInfo struct {
	CanLike    int `json:"can_like"`
	CanPublish int `json:"can_publish"`
	Count      int `json:"count"`
	UserLikes  int `json:"user_likes"`
}

type baseLink struct {
	Application baseLinkApplication `json:"application"`
	Button      baseLinkButton      `json:"button"`
	Caption     string              `json:"caption"`
	Description string              `json:"description"`
	Photo       PhotosPhoto         `json:"photo"`
	PreviewPage string              `json:"preview_page"`
	PreviewURL  string              `json:"preview_url"`
	Product     baseLinkProduct     `json:"product"`
	Rating      baseLinkRating      `json:"rating"`
	Title       string              `json:"title"`
	URL         string              `json:"url"`
}

type baseLinkApplication struct {
	AppID float64                  `json:"app_id"`
	Store baseLinkApplicationStore `json:"store"`
}

type baseLinkApplicationStore struct {
	ID   float64 `json:"id"`
	Name string  `json:"name"`
}

type baseLinkButton struct {
	Action baseLinkButtonAction `json:"action"`
	Title  string               `json:"title"`
}

type baseLinkButtonAction struct {
	Type baseLinkButtonActionType `json:"type"`
	URL  string                   `json:"url"`
}

type baseLinkButtonActionType string

type baseLinkProduct struct {
	Price marketPrice `json:"price"`
}

type baseLinkRating struct {
	ReviewsCount int     `json:"reviews_count"`
	Stars        float64 `json:"stars"`
}

type baseOkResponse int

type basePlace struct {
	Address   string  `json:"address"`
	Checkins  int     `json:"checkins"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Created   int     `json:"created"`
	ID        int     `json:"id"`
	Icon      string  `json:"icon"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
	Type      string  `json:"type"`
}

type basePropertyExists int

type baseRepostsInfo struct {
	Count        int `json:"count"`
	UserReposted int `json:"user_reposted"`
}

type baseSticker struct {
	Images               []baseImage `json:"images"`
	ImagesWithBackground []baseImage `json:"images_with_background"`
	ProductID            int         `json:"product_id"`
	StickerID            int         `json:"sticker_id"`
}

type baseUploadServer struct {
	UploadURL string `json:"upload_url"`
}

type baseUserID struct {
	UserID int `json:"user_id"`
}

type boardDefaultOrder int

type boardTopic struct {
	Comments  int    `json:"comments"`
	Created   int    `json:"created"`
	CreatedBy int    `json:"created_by"`
	ID        int    `json:"id"`
	IsClosed  int    `json:"is_closed"`
	IsFixed   int    `json:"is_fixed"`
	Title     string `json:"title"`
	Updated   int    `json:"updated"`
	UpdatedBy int    `json:"updated_by"`
}

type boardTopicComment struct {
	Attachments []WallCommentAttachment `json:"attachments"`
	Date        int                     `json:"date"`
	FromID      int                     `json:"from_id"`
	ID          int                     `json:"id"`
	RealOffset  int                     `json:"real_offset"`
	Text        string                  `json:"text"`
}

type boardTopicPoll struct {
	AnswerID int           `json:"answer_id"`
	Answers  []pollsAnswer `json:"answers"`
	Created  int           `json:"created"`
	IsClosed int           `json:"is_closed"`
	OwnerID  int           `json:"owner_id"`
	PollID   int           `json:"poll_id"`
	Question string        `json:"question"`
	Votes    string        `json:"votes"`
}

type commonFriend int

type databaseCity struct {
}

type databaseFaculty struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type databaseRegion struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type databaseSchool struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type databaseUniversity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type docsDoc struct {
	AccessKey string         `json:"access_key"`
	Date      int            `json:"date"`
	Ext       string         `json:"ext"`
	ID        int            `json:"id"`
	OwnerID   int            `json:"owner_id"`
	Preview   docsDocPreview `json:"preview"`
	Size      int            `json:"size"`
	Title     string         `json:"title"`
	Type      int            `json:"type"`
	URL       string         `json:"url"`
}

type docsDocPreview struct {
	Photo docsDocPreviewPhoto `json:"photo"`
	Video docsDocPreviewVideo `json:"video"`
}

type docsDocPreviewPhoto struct {
	Sizes []photosPhotoSizes `json:"sizes"`
}

type docsDocPreviewVideo struct {
	Filesize int    `json:"filesize"`
	Height   int    `json:"height"`
	Src      string `json:"src"`
	Width    int    `json:"width"`
}

type docsDocTypes struct {
	Count int    `json:"count"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type docsDocUploadResponse struct {
	File string `json:"file"`
}

type faveFavesLink struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
	Photo100    string `json:"photo_100"`
	Photo200    string `json:"photo_200"`
	Photo50     string `json:"photo_50"`
	Title       string `json:"title"`
	URL         string `json:"url"`
}

type friendsFriendStatus struct {
	FriendStatus   friendsFriendStatusStatus `json:"friend_status"`
	ReadState      int                       `json:"read_state"`
	RequestMessage string                    `json:"request_message"`
	Sign           string                    `json:"sign"`
	UserID         int                       `json:"user_id"`
}

type friendsFriendStatusStatus int

type friendsFriendsList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type friendsMutualFriend struct {
	CommonCount   int            `json:"common_count"`
	CommonFriends []commonFriend `json:"common_friends"`
	ID            int            `json:"id"`
}

type friendsRequests struct {
	From   string                `json:"from"`
	Mutual friendsRequestsMutual `json:"mutual"`
	UserID int                   `json:"user_id"`
}

type friendsRequestsMutual struct {
	Count int                         `json:"count"`
	Users []friendsRequestsMutualUser `json:"users"`
}

type friendsRequestsMutualUser int

type friendsRequestsXtrMessage struct {
	From    string                `json:"from"`
	Message string                `json:"message"`
	Mutual  friendsRequestsMutual `json:"mutual"`
	UserID  int                   `json:"user_id"`
}

type friendsUserXtrLists struct {
}

type friendsUserXtrPhone struct {
}

type groupsBanInfo struct {
	AdminID int                 `json:"admin_id"`
	Comment string              `json:"comment"`
	Date    int                 `json:"date"`
	EndDate int                 `json:"end_date"`
	Reason  groupsBanInfoReason `json:"reason"`
}

type groupsBanInfoReason int

type groupsCallbackSettings struct {
	APIVersion string               `json:"api_version"`
	Events     groupsLongPollEvents `json:"events"`
}

type groupsContactsItem struct {
	Desc   string `json:"desc"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	UserID int    `json:"user_id"`
}

type groupsCountersGroup struct {
	Albums int `json:"albums"`
	Audios int `json:"audios"`
	Docs   int `json:"docs"`
	Market int `json:"market"`
	Photos int `json:"photos"`
	Topics int `json:"topics"`
	Videos int `json:"videos"`
}

type groupsCover struct {
	Enabled int         `json:"enabled"`
	Images  []baseImage `json:"images"`
}

type groupsGroupBanInfo struct {
	Comment string `json:"comment"`
	EndDate int    `json:"end_date"`
}

type groupsGroupCategory struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Subcategories []BaseObjectWithName `json:"subcategories"`
}

type groupsGroupCategoryFull struct {
	ID            int                   `json:"id"`
	Name          string                `json:"name"`
	PageCount     int                   `json:"page_count"`
	PagePreviews  []GroupsGroup         `json:"page_previews"`
	Subcategories []groupsGroupCategory `json:"subcategories"`
}

type groupsGroupCategoryType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type groupsGroupFull struct {
}

type groupsGroupLink struct {
	Desc            string `json:"desc"`
	EditTitle       int    `json:"edit_title"`
	ID              int    `json:"id"`
	ImageProcessing int    `json:"image_processing"`
	URL             string `json:"url"`
}

type groupsGroupPublicCategoryList struct {
	ID           int                       `json:"id"`
	Name         string                    `json:"name"`
	SubtypesList []groupsGroupCategoryType `json:"subtypes_list"`
}

type groupsGroupSettings struct {
	Access             int                             `json:"access"`
	Address            string                          `json:"address"`
	Audio              int                             `json:"audio"`
	Description        string                          `json:"description"`
	Docs               int                             `json:"docs"`
	ObsceneFilter      int                             `json:"obscene_filter"`
	ObsceneStopwords   int                             `json:"obscene_stopwords"`
	ObsceneWords       string                          `json:"obscene_words"`
	Photos             int                             `json:"photos"`
	Place              placesPlaceMin                  `json:"place"`
	PublicCategory     int                             `json:"public_category"`
	PublicCategoryList []groupsGroupPublicCategoryList `json:"public_category_list"`
	PublicSubcategory  int                             `json:"public_subcategory"`
	Rss                string                          `json:"rss"`
	Subject            int                             `json:"subject"`
	SubjectList        []groupsSubjectItem             `json:"subject_list"`
	Title              string                          `json:"title"`
	Topics             int                             `json:"topics"`
	Video              int                             `json:"video"`
	Wall               int                             `json:"wall"`
	Website            string                          `json:"website"`
	Wiki               int                             `json:"wiki"`
}

type groupsGroupXtrInvitedBy struct {
	AdminLevel groupsGroupXtrInvitedByAdminLevel `json:"admin_level"`
	ID         string                            `json:"id"`
	InvitedBy  int                               `json:"invited_by"`
	IsAdmin    int                               `json:"is_admin"`
	IsClosed   int                               `json:"is_closed"`
	IsMember   int                               `json:"is_member"`
	Name       string                            `json:"name"`
	Photo100   string                            `json:"photo_100"`
	Photo200   string                            `json:"photo_200"`
	Photo50    string                            `json:"photo_50"`
	ScreenName string                            `json:"screen_name"`
	Type       groupsGroupXtrInvitedByType       `json:"type"`
}

type groupsGroupXtrInvitedByAdminLevel int

type groupsGroupXtrInvitedByType string

type groupsGroupsArray struct {
	Count int                     `json:"count"`
	Items []groupsGroupsArrayItem `json:"items"`
}

type groupsGroupsArrayItem int

type groupsLinksItem struct {
	Desc      string `json:"desc"`
	EditTitle int    `json:"edit_title"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Photo100  string `json:"photo_100"`
	Photo50   string `json:"photo_50"`
	URL       string `json:"url"`
}

type groupsLongPollEvents struct {
	AudioNew             int `json:"audio_new"`
	BoardPostDelete      int `json:"board_post_delete"`
	BoardPostEdit        int `json:"board_post_edit"`
	BoardPostNew         int `json:"board_post_new"`
	BoardPostRestore     int `json:"board_post_restore"`
	GroupChangePhoto     int `json:"group_change_photo"`
	GroupChangeSettings  int `json:"group_change_settings"`
	GroupJoin            int `json:"group_join"`
	GroupLeave           int `json:"group_leave"`
	GroupOfficersEdit    int `json:"group_officers_edit"`
	LeadFormsNew         int `json:"lead_forms_new"`
	MarketCommentDelete  int `json:"market_comment_delete"`
	MarketCommentEdit    int `json:"market_comment_edit"`
	MarketCommentNew     int `json:"market_comment_new"`
	MarketCommentRestore int `json:"market_comment_restore"`
	MessageAllow         int `json:"message_allow"`
	MessageDeny          int `json:"message_deny"`
	MessageNew           int `json:"message_new"`
	MessageReply         int `json:"message_reply"`
	MessageTypingState   int `json:"message_typing_state"`
	MessagesEdit         int `json:"messages_edit"`
	PhotoCommentDelete   int `json:"photo_comment_delete"`
	PhotoCommentEdit     int `json:"photo_comment_edit"`
	PhotoCommentNew      int `json:"photo_comment_new"`
	PhotoCommentRestore  int `json:"photo_comment_restore"`
	PhotoNew             int `json:"photo_new"`
	PollVoteNew          int `json:"poll_vote_new"`
	UserBlock            int `json:"user_block"`
	UserUnblock          int `json:"user_unblock"`
	VideoCommentDelete   int `json:"video_comment_delete"`
	VideoCommentEdit     int `json:"video_comment_edit"`
	VideoCommentNew      int `json:"video_comment_new"`
	VideoCommentRestore  int `json:"video_comment_restore"`
	VideoNew             int `json:"video_new"`
	WallPostNew          int `json:"wall_post_new"`
	WallReplyDelete      int `json:"wall_reply_delete"`
	WallReplyEdit        int `json:"wall_reply_edit"`
	WallReplyNew         int `json:"wall_reply_new"`
	WallReplyRestore     int `json:"wall_reply_restore"`
	WallRepost           int `json:"wall_repost"`
}

type groupsLongPollServer struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     int    `json:"ts"`
}

type groupsLongPollSettings struct {
	APIVersion string               `json:"api_version"`
	Events     groupsLongPollEvents `json:"events"`
	IsEnabled  bool                 `json:"is_enabled"`
}

type groupsMarketInfo struct {
	ContactID    int            `json:"contact_id"`
	Currency     marketCurrency `json:"currency"`
	CurrencyText string         `json:"currency_text"`
	Enabled      int            `json:"enabled"`
	MainAlbumID  int            `json:"main_album_id"`
	PriceMax     int            `json:"price_max"`
	PriceMin     int            `json:"price_min"`
}

type groupsMemberRole struct {
	ID   int                    `json:"id"`
	Role groupsMemberRoleStatus `json:"role"`
}

type groupsMemberRoleStatus string

type groupsMemberStatus struct {
	Member int `json:"member"`
	UserID int `json:"user_id"`
}

type groupsMemberStatusFull struct {
	Invitation int `json:"invitation"`
	Member     int `json:"member"`
	Request    int `json:"request"`
	UserID     int `json:"user_id"`
}

type groupsOnlineStatus struct {
	Minutes int                    `json:"minutes"`
	Status  groupsOnlineStatusType `json:"status"`
}

type groupsOnlineStatusType string

type groupsOwnerXtrBanInfo struct {
	BanInfo groupsBanInfo             `json:"ban_info"`
	Group   GroupsGroup               `json:"group"`
	Profile UsersUser                 `json:"profile"`
	Type    groupsOwnerXtrBanInfoType `json:"type"`
}

type groupsOwnerXtrBanInfoType string

type groupsRoleOptions string

type groupsSubjectItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type groupsTokenPermissionSetting struct {
	Name    string `json:"name"`
	Setting int    `json:"setting"`
}

type groupsTokenPermissions struct {
	Mask        int                            `json:"mask"`
	Permissions []groupsTokenPermissionSetting `json:"permissions"`
}

type groupsUserXtrRole struct {
}

type lang string

type leadsChecked struct {
	Reason    string             `json:"reason"`
	Result    leadsCheckedResult `json:"result"`
	Sid       string             `json:"sid"`
	StartLink string             `json:"start_link"`
}

type leadsCheckedResult string

type leadsComplete struct {
	Cost     int            `json:"cost"`
	Limit    int            `json:"limit"`
	Spent    int            `json:"spent"`
	Success  baseOkResponse `json:"success"`
	TestMode int            `json:"test_mode"`
}

type leadsEntry struct {
	Aid       int    `json:"aid"`
	Comment   string `json:"comment"`
	Date      int    `json:"date"`
	Sid       string `json:"sid"`
	StartDate int    `json:"start_date"`
	Status    int    `json:"status"`
	TestMode  int    `json:"test_mode"`
	UID       int    `json:"uid"`
}

type leadsLead struct {
	Completed   int           `json:"completed"`
	Cost        int           `json:"cost"`
	Days        leadsLeadDays `json:"days"`
	Impressions int           `json:"impressions"`
	Limit       int           `json:"limit"`
	Spent       int           `json:"spent"`
	Started     int           `json:"started"`
}

type leadsLeadDays struct {
	Completed   int `json:"completed"`
	Impressions int `json:"impressions"`
	Spent       int `json:"spent"`
	Started     int `json:"started"`
}

type leadsStart struct {
	TestMode int    `json:"test_mode"`
	VkSid    string `json:"vk_sid"`
}

type marketCurrency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type marketMarketAlbum struct {
	Count       int         `json:"count"`
	ID          int         `json:"id"`
	OwnerID     int         `json:"owner_id"`
	Photo       PhotosPhoto `json:"photo"`
	Title       string      `json:"title"`
	UpdatedTime int         `json:"updated_time"`
}

type marketMarketCategory struct {
	ID      int           `json:"id"`
	Name    string        `json:"name"`
	Section marketSection `json:"section"`
}

type marketMarketItem struct {
	Availability marketMarketItemAvailability `json:"availability"`
	Category     marketMarketCategory         `json:"category"`
	Date         int                          `json:"date"`
	Description  string                       `json:"description"`
	ID           int                          `json:"id"`
	OwnerID      int                          `json:"owner_id"`
	Price        marketPrice                  `json:"price"`
	ThumbPhoto   string                       `json:"thumb_photo"`
	Title        string                       `json:"title"`
}

type marketMarketItemAvailability int

type marketMarketItemFull struct {
}

type marketPrice struct {
	Amount   string         `json:"amount"`
	Currency marketCurrency `json:"currency"`
	Text     string         `json:"text"`
}

type marketSection struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type messagesChat struct {
	AdminID      int                      `json:"admin_id"`
	ID           int                      `json:"id"`
	Kicked       int                      `json:"kicked"`
	Left         int                      `json:"left"`
	Photo100     string                   `json:"photo_100"`
	Photo200     string                   `json:"photo_200"`
	Photo50      string                   `json:"photo_50"`
	PushSettings messagesChatPushSettings `json:"push_settings"`
	Title        string                   `json:"title"`
	Type         string                   `json:"type"`
	Users        []messagesChatUser       `json:"users"`
}

type messagesChatFull struct {
	AdminID      int                        `json:"admin_id"`
	ID           int                        `json:"id"`
	Kicked       int                        `json:"kicked"`
	Left         int                        `json:"left"`
	Photo100     string                     `json:"photo_100"`
	Photo200     string                     `json:"photo_200"`
	Photo50      string                     `json:"photo_50"`
	PushSettings messagesChatPushSettings   `json:"push_settings"`
	Title        string                     `json:"title"`
	Type         string                     `json:"type"`
	Users        []messagesUserXtrInvitedBy `json:"users"`
}

type messagesChatPushSettings struct {
	DisabledUntil int `json:"disabled_until"`
	Sound         int `json:"sound"`
}

type messagesChatSettingsPhoto struct {
	Photo100 string `json:"photo_100"`
	Photo200 string `json:"photo_200"`
	Photo50  string `json:"photo_50"`
}

type messagesChatSettingsState string

type messagesChatUser int

type messagesConversation struct {
	CanWrite     messagesConversationCanWrite     `json:"can_write"`
	ChatSettings messagesConversationChatSettings `json:"chat_settings"`
	Important    bool                             `json:"important"`
	InRead       int                              `json:"in_read"`
	OutRead      int                              `json:"out_read"`
	Peer         messagesConversationPeer         `json:"peer"`
	PushSettings messagesConversationPushSettings `json:"push_settings"`
	Unanswered   bool                             `json:"unanswered"`
	UnreadCount  int                              `json:"unread_count"`
}

type messagesConversationCanWrite struct {
	Allowed bool `json:"allowed"`
	Reason  int  `json:"reason"`
}

type messagesConversationChatSettings struct {
	MembersCount  int                       `json:"members_count"`
	Photo         messagesChatSettingsPhoto `json:"photo"`
	PinnedMessage messagesPinnedMessage     `json:"pinned_message"`
	State         messagesChatSettingsState `json:"state"`
	Title         string                    `json:"title"`
}

type messagesConversationPeer struct {
	ID      int    `json:"id"`
	LocalID int    `json:"local_id"`
	Type    string `json:"type"`
}

type messagesConversationPushSettings struct {
	DisabledForever bool `json:"disabled_forever"`
	DisabledUntil   int  `json:"disabled_until"`
	NoSound         bool `json:"no_sound"`
}

type messagesConversationWithMessage struct {
	Conversation messagesConversation `json:"conversation"`
	LastMessage  MessagesMessage      `json:"last_message"`
}

type messagesDialog struct {
	Important  int             `json:"important"`
	InRead     int             `json:"in_read"`
	Message    MessagesMessage `json:"message"`
	OutRead    int             `json:"out_read"`
	Unanswered int             `json:"unanswered"`
	Unread     int             `json:"unread"`
}

type messagesHistoryAttachment struct {
	Attachment messagesHistoryMessageAttachment `json:"attachment"`
	MessageID  int                              `json:"message_id"`
}

type messagesHistoryMessageAttachment struct {
	Audio  AudioAudioFull                       `json:"audio"`
	Doc    docsDoc                              `json:"doc"`
	Link   baseLink                             `json:"link"`
	Market baseLink                             `json:"market"`
	Photo  PhotosPhoto                          `json:"photo"`
	Share  baseLink                             `json:"share"`
	Type   messagesHistoryMessageAttachmentType `json:"type"`
	Video  VideoVideo                           `json:"video"`
	Wall   baseLink                             `json:"wall"`
}

type messagesHistoryMessageAttachmentType string

type messagesLastActivity struct {
	Online int `json:"online"`
	Time   int `json:"time"`
}

type messagesLongpollMessages struct {
	Count int               `json:"count"`
	Items []MessagesMessage `json:"items"`
}

type messagesLongpollParams struct {
	Key    string `json:"key"`
	Pts    int    `json:"pts"`
	Server string `json:"server"`
	Ts     int    `json:"ts"`
}

type messagesMessageAction struct {
	ConversationMessageID int                         `json:"conversation_message_id"`
	Email                 string                      `json:"email"`
	MemberID              int                         `json:"member_id"`
	Message               string                      `json:"message"`
	Photo                 messagesChatSettingsPhoto   `json:"photo"`
	Text                  string                      `json:"text"`
	Type                  messagesMessageActionStatus `json:"type"`
}

type messagesMessageActionStatus string

type messagesMessageAttachment struct {
	Audio             AudioAudioFull                `json:"audio"`
	Doc               docsDoc                       `json:"doc"`
	Gift              GiftsLayout                   `json:"gift"`
	Link              baseLink                      `json:"link"`
	Market            marketMarketItem              `json:"market"`
	MarketMarketAlbum marketMarketAlbum             `json:"market_market_album"`
	Photo             PhotosPhoto                   `json:"photo"`
	Sticker           baseSticker                   `json:"sticker"`
	Type              messagesMessageAttachmentType `json:"type"`
	Video             VideoVideo                    `json:"video"`
	Wall              WallWallpostAttached          `json:"wall"`
	WallReply         WallWallComment               `json:"wall_reply"`
}

type messagesMessageAttachmentType string

type messagesPinnedMessage struct {
	Attachments           []messagesMessageAttachment `json:"attachments"`
	ConversationMessageID int                         `json:"conversation_message_id"`
	Date                  int                         `json:"date"`
	FromID                int                         `json:"from_id"`
	FwdMessages           []MessagesMessage           `json:"fwd_messages"`
	Geo                   baseGeo                     `json:"geo"`
	ID                    int                         `json:"id"`
	PeerID                int                         `json:"peer_id"`
	Text                  string                      `json:"text"`
}

type messagesUserXtrInvitedBy struct {
}

type newsfeedItemAudio struct {
	Audio  newsfeedItemAudioAudio `json:"audio"`
	PostID int                    `json:"post_id"`
}

type newsfeedItemAudioAudio struct {
	Count int              `json:"count"`
	Items []AudioAudioFull `json:"items"`
}

type newsfeedItemFriend struct {
	Friends newsfeedItemFriendFriends `json:"friends"`
}

type newsfeedItemFriendFriends struct {
	Count int          `json:"count"`
	Items []baseUserID `json:"items"`
}

type newsfeedItemNote struct {
	Notes newsfeedItemNoteNotes `json:"notes"`
}

type newsfeedItemNoteNotes struct {
	Count int                    `json:"count"`
	Items []newsfeedNewsfeedNote `json:"items"`
}

type newsfeedItemPhoto struct {
	Photos newsfeedItemPhotoPhotos `json:"photos"`
	PostID int                     `json:"post_id"`
}

type newsfeedItemPhotoPhotos struct {
	Count int                     `json:"count"`
	Items []newsfeedNewsfeedPhoto `json:"items"`
}

type newsfeedItemPhotoTag struct {
	PhotoTags newsfeedItemPhotoTagPhotoTags `json:"photo_tags"`
	PostID    int                           `json:"post_id"`
}

type newsfeedItemPhotoTagPhotoTags struct {
	Count int                     `json:"count"`
	Items []newsfeedNewsfeedPhoto `json:"items"`
}

type newsfeedItemTopic struct {
	Comments baseCommentsInfo `json:"comments"`
	Likes    baseLikesInfo    `json:"likes"`
	PostID   int              `json:"post_id"`
	Text     string           `json:"text"`
}

type newsfeedItemVideo struct {
	Video newsfeedItemVideoVideo `json:"video"`
}

type newsfeedItemVideoVideo struct {
	Count int          `json:"count"`
	Items []VideoVideo `json:"items"`
}

type newsfeedItemWallpost struct {
	Attachments []WallWallpostAttachment `json:"attachments"`
	Comments    baseCommentsInfo         `json:"comments"`
	CopyHistory []WallWallpost           `json:"copy_history"`
	Geo         baseGeo                  `json:"geo"`
	Likes       baseLikesInfo            `json:"likes"`
	PostID      int                      `json:"post_id"`
	PostSource  WallPostSource           `json:"post_source"`
	PostType    newsfeedItemWallpostType `json:"post_type"`
	Reposts     baseRepostsInfo          `json:"reposts"`
	Text        string                   `json:"text"`
}

type newsfeedItemWallpostType string

type newsfeedList struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type newsfeedListFull struct {
}

type newsfeedNewsfeedItem struct {
}

type newsfeedNewsfeedItemType string

type newsfeedNewsfeedNote struct {
	Comments int    `json:"comments"`
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Title    string `json:"title"`
}

type newsfeedNewsfeedPhoto struct {
}

type notesNote struct {
	CanComment int    `json:"can_comment"`
	Comments   int    `json:"comments"`
	Date       int    `json:"date"`
	ID         int    `json:"id"`
	OwnerID    int    `json:"owner_id"`
	Text       string `json:"text"`
	TextWiki   string `json:"text_wiki"`
	Title      string `json:"title"`
	ViewURL    string `json:"view_url"`
}

type notesNoteComment struct {
	Date    int    `json:"date"`
	ID      int    `json:"id"`
	Message string `json:"message"`
	Nid     int    `json:"nid"`
	Oid     int    `json:"oid"`
	ReplyTo int    `json:"reply_to"`
	UID     int    `json:"uid"`
}

type notificationsFeedback struct {
	Attachments []WallWallpostAttachment `json:"attachments"`
	FromID      int                      `json:"from_id"`
	Geo         baseGeo                  `json:"geo"`
	ID          int                      `json:"id"`
	Likes       baseLikesInfo            `json:"likes"`
	Text        string                   `json:"text"`
	ToID        int                      `json:"to_id"`
}

type notificationsNotification struct {
	Date     int                             `json:"date"`
	Feedback notificationsFeedback           `json:"feedback"`
	Parent   notificationsNotificationParent `json:"parent"`
	Reply    notificationsReply              `json:"reply"`
	Type     string                          `json:"type"`
}

type notificationsNotificationParent struct {
}

type notificationsNotificationsComment struct {
	Date    int          `json:"date"`
	ID      int          `json:"id"`
	OwnerID int          `json:"owner_id"`
	Photo   PhotosPhoto  `json:"photo"`
	Post    WallWallpost `json:"post"`
	Text    string       `json:"text"`
	Topic   boardTopic   `json:"topic"`
	Video   VideoVideo   `json:"video"`
}

type notificationsReply struct {
	Date int `json:"date"`
	ID   int `json:"id"`
	Text int `json:"text"`
}

type oauthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	RedirectURI      string `json:"redirect_uri"`
}

type objects interface{}

type ordersAmount struct {
	Amounts  []ordersAmountItem `json:"amounts"`
	Currency string             `json:"currency"`
}

type ordersAmountItem struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Votes       string `json:"votes"`
}

type ordersOrder struct {
	Amount              int    `json:"amount"`
	AppOrderID          int    `json:"app_order_id"`
	CancelTransactionID int    `json:"cancel_transaction_id"`
	Date                int    `json:"date"`
	ID                  int    `json:"id"`
	Item                string `json:"item"`
	ReceiverID          int    `json:"receiver_id"`
	Status              string `json:"status"`
	TransactionID       int    `json:"transaction_id"`
	UserID              int    `json:"user_id"`
}

type pagesPrivacySettings int

type pagesWikipage struct {
	CreatorID   int                  `json:"creator_id"`
	CreatorName int                  `json:"creator_name"`
	EditorID    int                  `json:"editor_id"`
	EditorName  string               `json:"editor_name"`
	GroupID     int                  `json:"group_id"`
	ID          int                  `json:"id"`
	Title       string               `json:"title"`
	Views       int                  `json:"views"`
	WhoCanEdit  pagesPrivacySettings `json:"who_can_edit"`
	WhoCanView  pagesPrivacySettings `json:"who_can_view"`
}

type pagesWikipageFull struct {
	Created                  int                  `json:"created"`
	CreatorID                int                  `json:"creator_id"`
	CurrentUserCanEdit       int                  `json:"current_user_can_edit"`
	CurrentUserCanEditAccess int                  `json:"current_user_can_edit_access"`
	Edited                   int                  `json:"edited"`
	EditorID                 int                  `json:"editor_id"`
	GroupID                  int                  `json:"group_id"`
	HTML                     string               `json:"html"`
	ID                       int                  `json:"id"`
	Source                   string               `json:"source"`
	Title                    string               `json:"title"`
	ViewURL                  string               `json:"view_url"`
	Views                    int                  `json:"views"`
	WhoCanEdit               pagesPrivacySettings `json:"who_can_edit"`
	WhoCanView               pagesPrivacySettings `json:"who_can_view"`
}

type pagesWikipageVersion struct {
	Edited     int    `json:"edited"`
	EditorID   int    `json:"editor_id"`
	EditorName string `json:"editor_name"`
	ID         int    `json:"id"`
	Length     int    `json:"length"`
}

type photosCommentXtrPid struct {
	Attachments    []WallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"`
	FromID         int                     `json:"from_id"`
	ID             int                     `json:"id"`
	Likes          baseLikesInfo           `json:"likes"`
	Pid            int                     `json:"pid"`
	ReplyToComment int                     `json:"reply_to_comment"`
	ReplyToUser    int                     `json:"reply_to_user"`
	Text           string                  `json:"text"`
}

type photosImage struct {
	Height int             `json:"height"`
	Type   photosImageType `json:"type"`
	URL    string          `json:"url"`
	Width  int             `json:"width"`
}

type photosImageType string

type photosListItem string

type photosMarketAlbumUploadResponse struct {
	Gid    int    `json:"gid"`
	Hash   string `json:"hash"`
	Photo  string `json:"photo"`
	Server int    `json:"server"`
}

type photosMarketUploadResponse struct {
	CropData string `json:"crop_data"`
	CropHash string `json:"crop_hash"`
	GroupID  int    `json:"group_id"`
	Hash     string `json:"hash"`
	Photo    string `json:"photo"`
	Server   int    `json:"server"`
}

type photosMessageUploadResponse struct {
	Hash   string `json:"hash"`
	Photo  string `json:"photo"`
	Server int    `json:"server"`
}

type photosOwnerUploadResponse struct {
	Hash   string `json:"hash"`
	Photo  string `json:"photo"`
	Server int    `json:"server"`
}

type photosPhotoAlbum struct {
	Created     int         `json:"created"`
	Description string      `json:"description"`
	ID          int         `json:"id"`
	OwnerID     int         `json:"owner_id"`
	Size        int         `json:"size"`
	Thumb       PhotosPhoto `json:"thumb"`
	Title       string      `json:"title"`
	Updated     int         `json:"updated"`
}

type photosPhotoAlbumFull struct {
	CanUpload          int                                      `json:"can_upload"`
	CommentsDisabled   int                                      `json:"comments_disabled"`
	Created            int                                      `json:"created"`
	Description        string                                   `json:"description"`
	ID                 int                                      `json:"id"`
	OwnerID            int                                      `json:"owner_id"`
	PrivacyComment     []photosPhotoAlbumFullPrivacyCommentItem `json:"privacy_comment"`
	PrivacyView        []photosPhotoAlbumFullPrivacyViewItem    `json:"privacy_view"`
	Size               int                                      `json:"size"`
	Sizes              []photosPhotoSizes                       `json:"sizes"`
	ThumbID            int                                      `json:"thumb_id"`
	ThumbIsLast        int                                      `json:"thumb_is_last"`
	ThumbSrc           string                                   `json:"thumb_src"`
	Title              string                                   `json:"title"`
	Updated            int                                      `json:"updated"`
	UploadByAdminsOnly int                                      `json:"upload_by_admins_only"`
}

type photosPhotoAlbumFullPrivacyCommentItem string

type photosPhotoAlbumFullPrivacyViewItem string

type photosPhotoFull struct {
	AccessKey  string          `json:"access_key"`
	AlbumID    int             `json:"album_id"`
	CanComment int             `json:"can_comment"`
	Comments   BaseObjectCount `json:"comments"`
	Date       int             `json:"date"`
	Height     int             `json:"height"`
	ID         int             `json:"id"`
	Images     []photosImage   `json:"images"`
	Lat        float64         `json:"lat"`
	Likes      baseLikes       `json:"likes"`
	Long       float64         `json:"long"`
	OwnerID    int             `json:"owner_id"`
	PostID     int             `json:"post_id"`
	Reposts    BaseObjectCount `json:"reposts"`
	Tags       BaseObjectCount `json:"tags"`
	Text       string          `json:"text"`
	UserID     int             `json:"user_id"`
	Width      int             `json:"width"`
}

type photosPhotoFullXtrRealOffset struct {
	AccessKey  string             `json:"access_key"`
	AlbumID    int                `json:"album_id"`
	CanComment int                `json:"can_comment"`
	Comments   BaseObjectCount    `json:"comments"`
	Date       int                `json:"date"`
	Height     int                `json:"height"`
	Hidden     basePropertyExists `json:"hidden"`
	ID         int                `json:"id"`
	Lat        float64            `json:"lat"`
	Likes      baseLikes          `json:"likes"`
	Long       float64            `json:"long"`
	OwnerID    int                `json:"owner_id"`
	Photo1280  string             `json:"photo_1280"`
	Photo130   string             `json:"photo_130"`
	Photo2560  string             `json:"photo_2560"`
	Photo604   string             `json:"photo_604"`
	Photo75    string             `json:"photo_75"`
	Photo807   string             `json:"photo_807"`
	PostID     int                `json:"post_id"`
	RealOffset int                `json:"real_offset"`
	Reposts    BaseObjectCount    `json:"reposts"`
	Sizes      []photosPhotoSizes `json:"sizes"`
	Tags       BaseObjectCount    `json:"tags"`
	Text       string             `json:"text"`
	UserID     int                `json:"user_id"`
	Width      int                `json:"width"`
}

type photosPhotoSizes struct {
	Height int                  `json:"height"`
	Src    string               `json:"src"`
	Type   photosPhotoSizesType `json:"type"`
	Width  int                  `json:"width"`
}

type photosPhotoSizesType string

type photosPhotoTag struct {
	Date       int     `json:"date"`
	ID         int     `json:"id"`
	PlacerID   int     `json:"placer_id"`
	TaggedName string  `json:"tagged_name"`
	UserID     int     `json:"user_id"`
	Viewed     int     `json:"viewed"`
	X          float64 `json:"x"`
	X2         float64 `json:"x2"`
	Y          float64 `json:"y"`
	Y2         float64 `json:"y2"`
}

type photosPhotoUpload struct {
	AlbumID   int    `json:"album_id"`
	UploadURL string `json:"upload_url"`
	UserID    int    `json:"user_id"`
}

type photosPhotoUploadResponse struct {
	Aid        int    `json:"aid"`
	Hash       string `json:"hash"`
	PhotosList string `json:"photos_list"`
	Server     int    `json:"server"`
}

type photosPhotoXtrRealOffset struct {
	AccessKey  string             `json:"access_key"`
	AlbumID    int                `json:"album_id"`
	Date       int                `json:"date"`
	Height     int                `json:"height"`
	Hidden     basePropertyExists `json:"hidden"`
	ID         int                `json:"id"`
	Lat        float64            `json:"lat"`
	Long       float64            `json:"long"`
	OwnerID    int                `json:"owner_id"`
	Photo1280  string             `json:"photo_1280"`
	Photo130   string             `json:"photo_130"`
	Photo2560  string             `json:"photo_2560"`
	Photo604   string             `json:"photo_604"`
	Photo75    string             `json:"photo_75"`
	Photo807   string             `json:"photo_807"`
	PostID     int                `json:"post_id"`
	RealOffset int                `json:"real_offset"`
	Sizes      []photosPhotoSizes `json:"sizes"`
	Text       string             `json:"text"`
	UserID     int                `json:"user_id"`
	Width      int                `json:"width"`
}

type photosPhotoXtrTagInfo struct {
	AccessKey  string             `json:"access_key"`
	AlbumID    int                `json:"album_id"`
	Date       int                `json:"date"`
	Height     int                `json:"height"`
	ID         int                `json:"id"`
	Lat        float64            `json:"lat"`
	Long       float64            `json:"long"`
	OwnerID    int                `json:"owner_id"`
	Photo1280  string             `json:"photo_1280"`
	Photo130   string             `json:"photo_130"`
	Photo2560  string             `json:"photo_2560"`
	Photo604   string             `json:"photo_604"`
	Photo75    string             `json:"photo_75"`
	Photo807   string             `json:"photo_807"`
	PlacerID   int                `json:"placer_id"`
	PostID     int                `json:"post_id"`
	Sizes      []photosPhotoSizes `json:"sizes"`
	TagCreated int                `json:"tag_created"`
	TagID      int                `json:"tag_id"`
	Text       string             `json:"text"`
	UserID     int                `json:"user_id"`
	Width      int                `json:"width"`
}

type photosWallUploadResponse struct {
	Hash   string `json:"hash"`
	Photo  string `json:"photo"`
	Server int    `json:"server"`
}

type placesCheckin struct {
	Date         int     `json:"date"`
	Distance     int     `json:"distance"`
	ID           int     `json:"id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	PlaceCity    int     `json:"place_city"`
	PlaceCountry int     `json:"place_country"`
	PlaceID      int     `json:"place_id"`
	PlaceIcon    string  `json:"place_icon"`
	PlaceTitle   string  `json:"place_title"`
	PlaceType    string  `json:"place_type"`
	Text         string  `json:"text"`
	UserID       int     `json:"user_id"`
}

type placesPlaceFull struct {
	Address    string  `json:"address"`
	Checkins   int     `json:"checkins"`
	City       int     `json:"city"`
	Country    int     `json:"country"`
	Created    int     `json:"created"`
	Distance   int     `json:"distance"`
	GroupID    int     `json:"group_id"`
	GroupPhoto string  `json:"group_photo"`
	ID         int     `json:"id"`
	Icon       string  `json:"icon"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Title      string  `json:"title"`
	Type       string  `json:"type"`
}

type placesPlaceMin struct {
	Address   string  `json:"address"`
	Checkins  int     `json:"checkins"`
	City      int     `json:"city"`
	Country   int     `json:"country"`
	Created   int     `json:"created"`
	ID        int     `json:"id"`
	Icon      string  `json:"icon"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
	Type      string  `json:"type"`
}

type placesTypes struct {
	ID    int    `json:"id"`
	Icon  string `json:"icon"`
	Title string `json:"title"`
}

type pollsAnswer struct {
	ID    int     `json:"id"`
	Rate  float64 `json:"rate"`
	Text  string  `json:"text"`
	Votes int     `json:"votes"`
}

type pollsPoll struct {
	Anonymous int           `json:"anonymous"`
	AnswerID  int           `json:"answer_id"`
	Answers   []pollsAnswer `json:"answers"`
	Created   int           `json:"created"`
	ID        int           `json:"id"`
	OwnerID   int           `json:"owner_id"`
	Question  string        `json:"question"`
	Votes     string        `json:"votes"`
}

type pollsVoters struct {
	AnswerID int              `json:"answer_id"`
	Users    pollsVotersUsers `json:"users"`
}

type pollsVotersUsers struct {
	Count int                    `json:"count"`
	Items []pollsVotersUsersItem `json:"items"`
}

type pollsVotersUsersItem int

type searchHint struct {
	Description string            `json:"description"`
	Global      int               `json:"global"`
	Group       GroupsGroup       `json:"group"`
	Profile     UsersUserMin      `json:"profile"`
	Section     searchHintSection `json:"section"`
	Type        searchHintType    `json:"type"`
}

type searchHintSection string

type searchHintType string

type secureLevel struct {
	Level int `json:"level"`
	UID   int `json:"uid"`
}

type secureSmsNotification struct {
	AppID   int    `json:"app_id"`
	Date    int    `json:"date"`
	ID      int    `json:"id"`
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}

type secureTokenChecked struct {
	Date    int            `json:"date"`
	Expire  int            `json:"expire"`
	Success baseOkResponse `json:"success"`
	UserID  int            `json:"user_id"`
}

type secureTransaction struct {
	Date    int `json:"date"`
	ID      int `json:"id"`
	UIDFrom int `json:"uid_from"`
	UIDTo   int `json:"uid_to"`
	Votes   int `json:"votes"`
}

type statsActivity struct {
	Comments     int `json:"comments"`
	Copies       int `json:"copies"`
	Hidden       int `json:"hidden"`
	Likes        int `json:"likes"`
	Subscribed   int `json:"subscribed"`
	Unsubscribed int `json:"unsubscribed"`
}

type statsCity struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type statsCountry struct {
	Code  string `json:"code"`
	Count int    `json:"count"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type statsPeriod struct {
	Activity   statsActivity `json:"activity"`
	PeriodFrom string        `json:"period_from"`
	PeriodTo   string        `json:"period_to"`
	Reach      statsReach    `json:"reach"`
	Visitors   statsViews    `json:"visitors"`
}

type statsReach struct {
	Age              []statsSexAge  `json:"age"`
	Cities           []statsCity    `json:"cities"`
	Countries        []statsCountry `json:"countries"`
	MobileReach      int            `json:"mobile_reach"`
	Reach            int            `json:"reach"`
	ReachSubscribers int            `json:"reach_subscribers"`
	Sex              []statsSexAge  `json:"sex"`
	SexAge           []statsSexAge  `json:"sex_age"`
}

type statsSexAge struct {
	Count int    `json:"count"`
	Value string `json:"value"`
}

type statsViews struct {
	Age         []statsSexAge  `json:"age"`
	Cities      []statsCity    `json:"cities"`
	Countries   []statsCountry `json:"countries"`
	MobileViews int            `json:"mobile_views"`
	Sex         []statsSexAge  `json:"sex"`
	SexAge      []statsSexAge  `json:"sex_age"`
	Views       int            `json:"views"`
	Visitors    int            `json:"visitors"`
}

type statsWallpostStat struct {
	Hide             int `json:"hide"`
	JoinGroup        int `json:"join_group"`
	Links            int `json:"links"`
	ReachSubscribers int `json:"reach_subscribers"`
	ReachTotal       int `json:"reach_total"`
	Report           int `json:"report"`
	ToGroup          int `json:"to_group"`
	Unsubscribe      int `json:"unsubscribe"`
}

type storiesReplies struct {
	Count int `json:"count"`
	New   int `json:"new"`
}

type storiesStoryLink struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type storiesStoryStats struct {
	Answer      storiesStoryStatsStat `json:"answer"`
	Bans        storiesStoryStatsStat `json:"bans"`
	OpenLink    storiesStoryStatsStat `json:"open_link"`
	Replies     storiesStoryStatsStat `json:"replies"`
	Shares      storiesStoryStatsStat `json:"shares"`
	Subscribers storiesStoryStatsStat `json:"subscribers"`
	Views       storiesStoryStatsStat `json:"views"`
}

type storiesStoryStatsStat struct {
	Count int                    `json:"count"`
	State storiesStoryStatsState `json:"state"`
}

type storiesStoryStatsState string

type storiesStoryType string

type storiesStoryVideo struct {
}

type utilsDomainResolved struct {
	ObjectID int                     `json:"object_id"`
	Type     utilsDomainResolvedType `json:"type"`
}

type utilsDomainResolvedType string

type utilsLastShortenedLink struct {
	AccessKey string `json:"access_key"`
	Key       string `json:"key"`
	ShortURL  string `json:"short_url"`
	Timestamp int    `json:"timestamp"`
	URL       string `json:"url"`
	Views     int    `json:"views"`
}

type utilsLinkChecked struct {
	Link   string                 `json:"link"`
	Status utilsLinkCheckedStatus `json:"status"`
}

type utilsLinkCheckedStatus string

type utilsLinkStats struct {
	Key   string       `json:"key"`
	Stats []utilsStats `json:"stats"`
}

type utilsLinkStatsExtended struct {
	Key   string               `json:"key"`
	Stats []utilsStatsExtended `json:"stats"`
}

type utilsShortLink struct {
	AccessKey string `json:"access_key"`
	Key       string `json:"key"`
	ShortURL  string `json:"short_url"`
	URL       string `json:"url"`
}

type utilsStats struct {
	Timestamp int `json:"timestamp"`
	Views     int `json:"views"`
}

type utilsStatsCity struct {
	CityID int `json:"city_id"`
	Views  int `json:"views"`
}

type utilsStatsCountry struct {
	CountryID int `json:"country_id"`
	Views     int `json:"views"`
}

type utilsStatsExtended struct {
	Cities    []utilsStatsCity    `json:"cities"`
	Countries []utilsStatsCountry `json:"countries"`
	SexAge    []utilsStatsSexAge  `json:"sex_age"`
	Timestamp int                 `json:"timestamp"`
	Views     int                 `json:"views"`
}

type utilsStatsSexAge struct {
	AgeRange string `json:"age_range"`
	Female   int    `json:"female"`
	Male     int    `json:"male"`
}

type videoCatBlock struct {
	CanHide int               `json:"can_hide"`
	ID      int               `json:"id"`
	Items   []videoCatElement `json:"items"`
	Name    string            `json:"name"`
	Next    string            `json:"next"`
	Type    videoCatBlockView `json:"type"`
	View    videoCatBlockView `json:"view"`
}

type videoCatBlockView string

type videoCatElement struct {
	CanAdd      int                 `json:"can_add"`
	CanEdit     int                 `json:"can_edit"`
	Comments    int                 `json:"comments"`
	Count       int                 `json:"count"`
	Date        int                 `json:"date"`
	Description string              `json:"description"`
	Duration    int                 `json:"duration"`
	ID          int                 `json:"id"`
	IsPrivate   int                 `json:"is_private"`
	OwnerID     int                 `json:"owner_id"`
	Photo130    string              `json:"photo_130"`
	Photo160    string              `json:"photo_160"`
	Photo320    string              `json:"photo_320"`
	Photo640    string              `json:"photo_640"`
	Photo800    string              `json:"photo_800"`
	Title       string              `json:"title"`
	Type        videoCatElementType `json:"type"`
	UpdatedTime int                 `json:"updated_time"`
	Views       int                 `json:"views"`
}

type videoCatElementType string

type videoSaveResult struct {
	Description string `json:"description"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	UploadURL   string `json:"upload_url"`
	VideoID     int    `json:"video_id"`
}

type videoUploadResponse struct {
	Size    int `json:"size"`
	VideoID int `json:"video_id"`
}

type videoVideoAlbum struct {
	ID      int    `json:"id"`
	OwnerID int    `json:"owner_id"`
	Title   string `json:"title"`
}

type videoVideoAlbumFull struct {
	Count       int    `json:"count"`
	ID          int    `json:"id"`
	IsSystem    int    `json:"is_system"`
	OwnerID     int    `json:"owner_id"`
	Photo160    string `json:"photo_160"`
	Photo320    string `json:"photo_320"`
	Title       string `json:"title"`
	UpdatedTime int    `json:"updated_time"`
}

type videoVideoFiles struct {
	External string `json:"external"`
	Mp1080   string `json:"mp_1080"`
	Mp240    string `json:"mp_240"`
	Mp360    string `json:"mp_360"`
	Mp480    string `json:"mp_480"`
	Mp720    string `json:"mp_720"`
}

type videoVideoFull struct {
	AccessKey      string                             `json:"access_key"`
	AddingDate     int                                `json:"adding_date"`
	CanAdd         int                                `json:"can_add"`
	CanComment     int                                `json:"can_comment"`
	CanEdit        int                                `json:"can_edit"`
	CanRepost      int                                `json:"can_repost"`
	Comments       int                                `json:"comments"`
	Date           int                                `json:"date"`
	Description    string                             `json:"description"`
	Duration       int                                `json:"duration"`
	Files          videoVideoFiles                    `json:"files"`
	ID             int                                `json:"id"`
	Likes          baseLikes                          `json:"likes"`
	Live           basePropertyExists                 `json:"live"`
	OwnerID        int                                `json:"owner_id"`
	Photo130       string                             `json:"photo_130"`
	Photo320       string                             `json:"photo_320"`
	Photo800       string                             `json:"photo_800"`
	Player         string                             `json:"player"`
	PrivacyComment []videoVideoFullPrivacyCommentItem `json:"privacy_comment"`
	PrivacyView    []videoVideoFullPrivacyViewItem    `json:"privacy_view"`
	Processing     basePropertyExists                 `json:"processing"`
	Repeat         int                                `json:"repeat"`
	Title          string                             `json:"title"`
	Views          int                                `json:"views"`
}

type videoVideoFullPrivacyCommentItem string

type videoVideoFullPrivacyViewItem string

type videoVideoTag struct {
	Date       int    `json:"date"`
	ID         int    `json:"id"`
	PlacerID   int    `json:"placer_id"`
	TaggedName string `json:"tagged_name"`
	UserID     int    `json:"user_id"`
	Viewed     int    `json:"viewed"`
}

type videoVideoTagInfo struct {
	AccessKey   string             `json:"access_key"`
	AddingDate  int                `json:"adding_date"`
	CanAdd      int                `json:"can_add"`
	CanEdit     int                `json:"can_edit"`
	Comments    int                `json:"comments"`
	Date        int                `json:"date"`
	Description string             `json:"description"`
	Duration    int                `json:"duration"`
	Files       videoVideoFiles    `json:"files"`
	ID          int                `json:"id"`
	Live        basePropertyExists `json:"live"`
	OwnerID     int                `json:"owner_id"`
	Photo130    string             `json:"photo_130"`
	Photo320    string             `json:"photo_320"`
	Photo800    string             `json:"photo_800"`
	PlacerID    int                `json:"placer_id"`
	Player      string             `json:"player"`
	Processing  basePropertyExists `json:"processing"`
	TagCreated  int                `json:"tag_created"`
	TagID       int                `json:"tag_id"`
	Title       string             `json:"title"`
	Views       int                `json:"views"`
}

// WallAppPost struct
type WallAppPost struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
}

// WallAttachedNote struct
type WallAttachedNote struct {
	Comments     int    `json:"comments"`
	Date         int    `json:"date"`
	ID           int    `json:"id"`
	OwnerID      int    `json:"owner_id"`
	ReadComments int    `json:"read_comments"`
	Title        string `json:"title"`
	ViewURL      string `json:"view_url"`
}

// WallCommentAttachment struct
type WallCommentAttachment struct {
	Audio             AudioAudioFull            `json:"audio"`
	Doc               docsDoc                   `json:"doc"`
	Link              baseLink                  `json:"link"`
	Market            marketMarketItem          `json:"market"`
	MarketMarketAlbum marketMarketAlbum         `json:"market_market_album"`
	Note              WallAttachedNote          `json:"note"`
	Page              pagesWikipageFull         `json:"page"`
	Photo             PhotosPhoto               `json:"photo"`
	Sticker           baseSticker               `json:"sticker"`
	Type              WallCommentAttachmentType `json:"type"`
	Video             VideoVideo                `json:"video"`
}

// WallCommentAttachmentType struct
type WallCommentAttachmentType string

// WallGraffiti struct
type WallGraffiti struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Photo200 string `json:"photo_200"`
	Photo586 string `json:"photo_586"`
}

// WallPostSource struct
type WallPostSource struct {
	Data     string             `json:"data"`
	Platform string             `json:"platform"`
	Type     WallPostSourceType `json:"type"`
	URL      string             `json:"url"`
}

// WallPostSourceType struct
type WallPostSourceType string

// WallPostType struct
type WallPostType string

// WallPostedPhoto struct
type WallPostedPhoto struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
}

// WallViews struct
type WallViews struct {
	Count int `json:"count"`
}

// WallWallComment struct
type WallWallComment struct {
	Attachments    []WallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"`
	FromID         int                     `json:"from_id"`
	ID             int                     `json:"id"`
	Likes          baseLikesInfo           `json:"likes"`
	RealOffset     int                     `json:"real_offset"`
	ReplyToComment int                     `json:"reply_to_comment"`
	ReplyToUser    int                     `json:"reply_to_user"`
	Text           string                  `json:"text"`
}

// WallWallpost struct
type WallWallpost struct {
	AccessKey   string                   `json:"access_key"`
	Attachments []WallWallpostAttachment `json:"attachments"`
	Date        int                      `json:"date"`
	FromID      int                      `json:"from_id"`
	Geo         baseGeo                  `json:"geo"`
	ID          int                      `json:"id"`
	OwnerID     int                      `json:"owner_id"`
	PostSource  WallPostSource           `json:"post_source"`
	PostType    WallPostType             `json:"post_type"`
	SignerID    int                      `json:"signer_id"`
	Text        string                   `json:"text"`
	Views       WallViews                `json:"views"`
}

// WallWallpostAttached struct
type WallWallpostAttached struct {
	Attachments []WallWallpostAttachment `json:"attachments"`
	CanDelete   int                      `json:"can_delete"`
	Comments    baseCommentsInfo         `json:"comments"`
	CopyOwnerID int                      `json:"copy_owner_id"`
	CopyPostID  int                      `json:"copy_post_id"`
	CopyText    string                   `json:"copy_text"`
	Date        int                      `json:"date"`
	FromID      int                      `json:"from_id"`
	Geo         baseGeo                  `json:"geo"`
	ID          int                      `json:"id"`
	Likes       baseLikesInfo            `json:"likes"`
	PostSource  WallPostSource           `json:"post_source"`
	PostType    WallPostType             `json:"post_type"`
	Reposts     baseRepostsInfo          `json:"reposts"`
	SignerID    int                      `json:"signer_id"`
	Text        string                   `json:"text"`
	ToID        int                      `json:"to_id"`
}

// WallWallpostAttachment struct
type WallWallpostAttachment struct {
	Album             photosPhotoAlbum           `json:"album"`
	App               WallAppPost                `json:"app"`
	Audio             AudioAudioFull             `json:"audio"`
	Doc               docsDoc                    `json:"doc"`
	Graffiti          WallGraffiti               `json:"graffiti"`
	Link              baseLink                   `json:"link"`
	Market            marketMarketItem           `json:"market"`
	MarketMarketAlbum marketMarketAlbum          `json:"market_market_album"`
	Note              WallAttachedNote           `json:"note"`
	Page              pagesWikipageFull          `json:"page"`
	Photo             PhotosPhoto                `json:"photo"`
	PhotosList        []photosListItem           `json:"photos_list"`
	Poll              pollsPoll                  `json:"poll"`
	PostedPhoto       WallPostedPhoto            `json:"posted_photo"`
	Type              WallWallpostAttachmentType `json:"type"`
	Video             VideoVideo                 `json:"video"`
}

// WallWallpostAttachmentType struct
type WallWallpostAttachmentType string

// WallWallpostFull struct
type WallWallpostFull struct {
}

// WallWallpostToID struct
type WallWallpostToID struct {
	Attachments []WallWallpostAttachment `json:"attachments"`
	Comments    baseCommentsInfo         `json:"comments"`
	CopyOwnerID int                      `json:"copy_owner_id"`
	CopyPostID  int                      `json:"copy_post_id"`
	Date        int                      `json:"date"`
	FromID      int                      `json:"from_id"`
	Geo         baseGeo                  `json:"geo"`
	ID          int                      `json:"id"`
	Likes       baseLikesInfo            `json:"likes"`
	PostID      int                      `json:"post_id"`
	PostSource  WallPostSource           `json:"post_source"`
	PostType    WallPostType             `json:"post_type"`
	Reposts     baseRepostsInfo          `json:"reposts"`
	SignerID    int                      `json:"signer_id"`
	Text        string                   `json:"text"`
	ToID        int                      `json:"to_id"`
}

// WidgetsCommentMedia struct
type WidgetsCommentMedia struct {
	ItemID   int                     `json:"item_id"`
	OwnerID  int                     `json:"owner_id"`
	ThumbSrc string                  `json:"thumb_src"`
	Type     WidgetsCommentMediaType `json:"type"`
}

// WidgetsCommentMediaType struct
type WidgetsCommentMediaType string

// WidgetsCommentReplies struct
type WidgetsCommentReplies struct {
	CanPost int                         `json:"can_post"`
	Count   int                         `json:"count"`
	Replies []WidgetsCommentRepliesItem `json:"replies"`
}

// WidgetsCommentRepliesItem struct
type WidgetsCommentRepliesItem struct {
	Cid   int                `json:"cid"`
	Date  int                `json:"date"`
	Likes WidgetsWidgetLikes `json:"likes"`
	Text  string             `json:"text"`
	UID   int                `json:"uid"`
	User  usersUserFull      `json:"user"`
}

// WidgetsWidgetComment struct
type WidgetsWidgetComment struct {
	Attachments []WallCommentAttachment `json:"attachments"`
	CanDelete   int                     `json:"can_delete"`
	Comments    WidgetsCommentReplies   `json:"comments"`
	Date        int                     `json:"date"`
	FromID      int                     `json:"from_id"`
	ID          int                     `json:"id"`
	Likes       baseLikesInfo           `json:"likes"`
	Media       WidgetsCommentMedia     `json:"media"`
	PostSource  WallPostSource          `json:"post_source"`
	PostType    int                     `json:"post_type"`
	Reposts     baseRepostsInfo         `json:"reposts"`
	Text        string                  `json:"text"`
	ToID        int                     `json:"to_id"`
	User        usersUserFull           `json:"user"`
}

// WidgetsWidgetLikes struct
type WidgetsWidgetLikes struct {
	Count int `json:"count"`
}

// WidgetsWidgetPage struct
type WidgetsWidgetPage struct {
	Comments    BaseObjectCount `json:"comments"`
	Date        int             `json:"date"`
	Description string          `json:"description"`
	ID          int             `json:"id"`
	Likes       BaseObjectCount `json:"likes"`
	PageID      string          `json:"page_id"`
	Photo       string          `json:"photo"`
	Title       string          `json:"title"`
	URL         string          `json:"url"`
}