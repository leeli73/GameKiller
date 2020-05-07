package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/xujiajun/nutsdb"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		t := r.FormValue("type")
		hostname := r.FormValue("hostname")
		if t == "add" {
			if setHostname(hostname) {
				w.Write([]byte(`添加成功`))
			} else {
				w.Write([]byte(`添加失败`))
			}
			return
		}
	} else {
		query := r.URL.Query()
		passwd := query["passwd"][0]
		if passwd != Config.WebPasswd {
			w.Write([]byte(`<script>alert("请输入正确的密码");window.location="/"</script>`))
			return
		}
		f, err := os.Open(path.Join("www", "index.html"))
		if err != nil {
			w.Write([]byte("出现未知错误！暂停访问"))
			return
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			w.Write([]byte("出现未知错误！暂停访问"))
			return
		}
		var list []string
		err = db.View(func(tx *nutsdb.Tx) error {
			entries, err := tx.GetAll("gamekiller")
			if err != nil {
				return err
			}
			for _, entry := range entries {
				//fmt.Println(string(entry.Key), string(entry.Value))
				list = append(list, string(entry.Key))
			}
			return nil
		})
		resStr := ""
		for i := 0; i < len(list); i++ {
			resStr += `<tr>
			<td>` + fmt.Sprintf("%d", i+1) + `</td>
			<td>` + list[i] + `</td>
			<td><button type="button" class="btn btn-danger" onclick="delete('` + list[i] + `')">删除</button></td>
		  </tr>`
		}
		if len(list) == 0 {
			resStr = "<tr>暂无数据</tr>"
		}
		strData := string(data)
		strData = strings.Replace(strData, "{{list}}", resStr, 1)
		w.Write([]byte(strData))
	}
}
