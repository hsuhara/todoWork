package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Task struct {
	TaskName    string    `json:"task_name"`    // タスク名称
	Deadline    time.Time `json:"deadline"`     // 期限
	Content     string    `json:"content"`      // 内容
	CreatedDate time.Time `json:"created_date"` // 作成日付
	UpdatedDate time.Time `json:"updated_date"` // 更新日付
}

func main() {
	// "/todoes"パスへのリクエストを処理するハンドラ関数を設定。
	// 下記のパスにアクセスすると定義したハンドラ関数が実行される。
	http.HandleFunc("/todoes", func(w http.ResponseWriter, r *http.Request) {
		// 一旦、送信するためのタスクデータを作成
		var tasks []Task
		tasks = append(tasks, Task{TaskName: "タスク1",
			Deadline:    time.Now().Add(48 * time.Hour), // 48時間後の日時
			Content:     "This is a sample task.My name is su",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		})

		// データをJSONにエンコードする。
		jsonData, err := json.Marshal(tasks)
		if err != nil {
			// エンコードに失敗した場合、サーバーエラーを返す。
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// エンコードされたJSONデータをクライアントに送信する。
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	// HTTPサーバーを起動
	http.ListenAndServe(":8080", nil)
}
