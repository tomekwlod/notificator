package teams

// https://learn.microsoft.com/en-us/microsoftteams/platform/task-modules-and-cards/cards/cards-format?tabs=adaptive-md%2Cdesktop%2Cconnector-html
// msgCard.Text = "Here are some examples of formatted stuff like " +
//
//	"<br> * this list itself  <br> * **bold** <br> * *italic* <br> * ***bolditalic***"
//
// msgCard.Text = "This is a **body** of this message (I will add something useful later)"
type MsgCard struct {
	// https://learn.microsoft.com/en-us/outlook/actionable-messages/message-card-reference
	Title      string `json:"title"`
	Text       string `json:"text"`
	ThemeColor string `json:"themeColor"` // deprecated but tey want to bring it back: https://github.com/MicrosoftDocs/msteams-docs/issues/10062
}
