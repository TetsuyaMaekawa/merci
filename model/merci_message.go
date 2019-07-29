package model

// MerciMessage メッセージテーブルマッピングオブジェクト
type MerciMessage struct {
}

// TableName テーブル名返却
func (m *MerciMessage) TableName() string {
	return "merci_message"
}
