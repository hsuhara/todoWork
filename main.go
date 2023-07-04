package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	Id          string    `json:"id"`           // タスクID
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
		tasks = append(tasks, Task{
			Id:          "1",
			TaskName:    "task1",
			Deadline:    time.Now().Add(48 * time.Hour),
			Content:     "task coment1",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
			Task{
				Id:          "2",
				TaskName:    "task2",
				Deadline:    time.Now().Add(48 * time.Hour),
				Content:     "task coment1",
				CreatedDate: time.Now(),
				UpdatedDate: time.Now(),
			},
		)

		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Param id is missing", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		var foundTask *Task
		for _, task := range tasks {
			if task.Id == id {
				foundTask = &task
				break
			}
		}
		if foundTask == nil {
			http.Error(w, fmt.Sprintf("No task foun with taskId:%s", id), http.StatusNotFound)
			return
		}

		json, err := json.Marshal(foundTask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(json)
	})

	http.ListenAndServe(":8080", nil)
}
