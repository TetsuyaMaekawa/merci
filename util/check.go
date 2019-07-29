package util

// IsArg 引数チェック
func IsArg(splitMessage []string) bool {
	if len(splitMessage) == 1 || len(splitMessage) == 3 {
		return true
	}
	return false
}
