package customisation

import (
	"fmt"

	"github.com/rxdn/gdl/objects"
	"github.com/rxdn/gdl/objects/guild/emoji"
)

type CustomEmoji struct {
	Name     string
	Id       uint64
	Animated bool
}

func NewCustomEmoji(name string, id uint64, animated bool) CustomEmoji {
	return CustomEmoji{
		Name: name,
		Id:   id,
	}
}

func (e CustomEmoji) String() string {
	if e.Animated {
		return fmt.Sprintf("<a:%s:%d>", e.Name, e.Id)
	} else {
		return fmt.Sprintf("<:%s:%d>", e.Name, e.Id)
	}
}

func (e CustomEmoji) BuildEmoji() *emoji.Emoji {
	return &emoji.Emoji{
		Id:       objects.NewNullableSnowflake(e.Id),
		Name:     e.Name,
		Animated: e.Animated,
	}
}

var (
	EmojiId         = NewCustomEmoji("id", 1365385845984596099, false)
	EmojiOpen       = NewCustomEmoji("open", 1350163453998010378, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1350163384926076990, false)
	EmojiClose      = NewCustomEmoji("close", 1350163368425816185, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1365385786949894267, false)
	EmojiReason     = NewCustomEmoji("reason", 1350163485626990714, false)
	EmojiSubject    = NewCustomEmoji("subject", 1365386000083325070, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1365386027572789299, false)
	EmojiClaim      = NewCustomEmoji("claim", 1350163319503454208, false)
	EmojiPanel      = NewCustomEmoji("panel", 1327350268974600263, false)
	EmojiRating     = NewCustomEmoji("rating", 1365385923822489610, false)
	EmojiStaff      = NewCustomEmoji("staff", 1365385976901537943, false)
	EmojiThread     = NewCustomEmoji("thread", 1365386013018558634, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1365385697892110356, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1365385912372170853, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1365385828590813255, false)
	//EmojiTime       = NewCustomEmoji("time", 974006684622159952, false)
)

// PrefixWithEmoji Useful for whitelabel bots
func PrefixWithEmoji(s string, emoji CustomEmoji, includeEmoji bool) string {
	if includeEmoji {
		return fmt.Sprintf("%s %s", emoji, s)
	} else {
		return s
	}
}
