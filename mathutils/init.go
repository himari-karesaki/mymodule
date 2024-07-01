package mathutils

import "fmt"

// importされた時点で初期化が実行される
func init() {
	fmt.Println("mathutilsパッケージが初期化されました")
}
