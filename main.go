package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	Id          string    `json:"id"`           // タスクID
	Name        string    `json:"name"`         // タスク名称
	Deadline    time.Time `json:"deadline"`     // 期限
	Content     string    `json:"content"`      // 内容
	CreatedDate time.Time `json:"created_date"` // 作成日付
	UpdatedDate time.Time `json:"updated_date"` // 更新日付
}

func main() {
	// "/todoes"パスへのリクエストを処理するハンドラ関数を設定。
	// 下記のパスにアクセスすると定義したハンドラ関数が実行される。
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		// 一旦、送信するためのタスクデータを作成
		var tasks []Task
		tasks = append(tasks, Task{
			Id:          "1",
			Name:        "task1",
			Deadline:    time.Now().Add(48 * time.Hour),
			Content:     "task coment1",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
			Task{
				Id:          "2",
				Name:        "task2",
				Deadline:    time.Now().Add(48 * time.Hour),
				Content:     "task coment1",
				CreatedDate: time.Now(),
				UpdatedDate: time.Now(),
			},
		)

		if r.URL.Path == "/todos" {
			w.Header().Set("Content-Type", "application/json")

			if r.URL.RawQuery == "" {
				// タスク一覧取得
				json, err := json.Marshal(tasks)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(json)

			} else {
				//単一タスク取得
				id := r.URL.Query().Get("id")

				if id == "" {
					http.Error(w, "task id is missing", http.StatusBadRequest)
					return
				}

				var foundTask *Task
				for _, task := range tasks {
					if task.Id == id {
						foundTask = &task
						break
					}
				}
				if foundTask == nil {
					http.Error(w, fmt.Sprintf("not found task with taskId:%s", id), http.StatusNotFound)
					return
				}

				json, err := json.Marshal(foundTask)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(json)
			}

		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}

	})

	http.ListenAndServe(":8080", nil)
}
