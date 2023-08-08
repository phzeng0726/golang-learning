package questions

import (
	"fmt"
	"net/http"
)

type Note struct {
	Id   int
	Name string
}

// 使用 Go 的 net/http 庫，建立一個簡單的 RESTful API，能夠處理用戶的 CRUD（新增、讀取、更新和刪除）請求。
func NoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

// 使用 Go 的 net/http 庫，建立一個簡單的 RESTful API，能夠處理用戶的 CRUD（新增、讀取、更新和刪除）請求。
func GetNotes(w http.ResponseWriter, r *http.Request) {
	// 返回用戶列表的 JSON
	// 例如，使用 json.Marshal() 將 Notes 切片轉換成 JSON 格式
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	// 解析請求的 JSON 數據，創建新用戶並添加到 Notes 切片中
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	// 解析請求的 JSON 數據，更新指定用戶的信息
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	// 從請求的路由參數中獲取用戶 ID，並從 Notes 切片中刪除該用戶
}
