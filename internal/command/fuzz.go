// +build gofuzz

package command

func FuzzRunAction(data []byte) int {
	ac := ActionConfig{
		Args: []string{"ls"},
		// "FOO: !var pizza/margherita"
		YamlInline: string(data),
	}
	// err := runAction(&ac)
	runAction(&ac)
	// if err != nil {
	// 	return 0
	// }
	// return 1
}
