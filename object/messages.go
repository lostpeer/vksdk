package object // import "github.com/severecloud/vksdk/object"

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
	Users        []int                    `json:"users"`
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
	State         string                    `json:"state"`
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
	Audio  AudioAudioFull `json:"audio"`
	Doc    docsDoc        `json:"doc"`
	Link   baseLink       `json:"link"`
	Market baseLink       `json:"market"`
	Photo  PhotosPhoto    `json:"photo"`
	Share  baseLink       `json:"share"`
	Type   string         `json:"type"`
	Video  VideoVideo     `json:"video"`
	Wall   baseLink       `json:"wall"`
}

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
	ConversationMessageID int                       `json:"conversation_message_id"`
	Email                 string                    `json:"email"`
	MemberID              int                       `json:"member_id"`
	Message               string                    `json:"message"`
	Photo                 messagesChatSettingsPhoto `json:"photo"`
	Text                  string                    `json:"text"`
	Type                  string                    `json:"type"`
}

type messagesMessageAttachment struct {
	Audio             AudioAudioFull       `json:"audio"`
	Doc               docsDoc              `json:"doc"`
	Gift              GiftsLayout          `json:"gift"`
	Link              baseLink             `json:"link"`
	Market            marketMarketItem     `json:"market"`
	MarketMarketAlbum marketMarketAlbum    `json:"market_market_album"`
	Photo             PhotosPhoto          `json:"photo"`
	Sticker           baseSticker          `json:"sticker"`
	Type              string               `json:"type"`
	Video             VideoVideo           `json:"video"`
	Wall              wallWallpostAttached `json:"wall"`
	WallReply         wallWallComment      `json:"wall_reply"`
}

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