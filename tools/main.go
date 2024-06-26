/*
   原始資料轉化與加工程式
   功能：
      1、將csv轉為格式化的JSON檔案
      2、為每一則FAQ增加一個UID，以便日後做相關識別（留言、喜歡、不喜歡...）
*/
package main

import(
   "os"
   "fmt"
   "io"
   "time"
   "strings"
   "strconv"
   "io/ioutil"
   "math/rand"
   "encoding/csv"
   "encoding/json"
   "encoding/base64"
   // "golang.org/x/text/transform"
   // "golang.org/x/text/encoding/traditionalchinese"
)

// category,title,question,group,url,contact,answer
type Items struct {
   Title		string		`json:"title"`
   Question		string		`json:"question"`
   Answer		string		`json:"answer"`
   Category		string		`json:"category"`
   Contact		string		`json:"contact"` 
   Url			string		`json:"url"`
   UUID			string		`json:"uuid"`
}

type CSv struct {
   ID		string			`json:"id"`
   Serial	string			`json:"serial"`
   Name		string			`json:"name"`
   ITEMs	map[string][]Items	`json:"items"`
}

func W(fileName string, content []byte)(error) {
   err := ioutil.WriteFile(fileName, []byte(content), 0666)
   return err
}

// 1)學術 10)資訊 20)人事 30)經費 40)生活
func main() {
   if len(os.Args) < 2 {
      fmt.Println("<usage> go run main.go filename")
      return
   }
   FilePath := strings.ToLower(os.Args[1])
   if FilePath == "" {
      fmt.Println("filename is empty")
      return
   }
   fname := strings.Split(FilePath, ".")
   categoryLists := map[string]CSv {
      "academia": { ID: "1", Name:"學術" },
      "its": { ID: "10", Name:"資訊" },
      "hr": { ID: "20", Name:"人事" },
      "budget": { ID: "30", Name:"經費" },
      "life": { ID: "40", Name:"生活" },
   }
   datas := categoryLists[fname[0]] 
   title := map[string]int{}
   rand.Seed(time.Now().UnixNano())
   datas.Serial = strconv.FormatInt(int64(rand.Intn(1000)), 10)

   file, err := os.OpenFile(FilePath, os.O_RDONLY, 0777)
   if err != nil {
      fmt.Println(err.Error())
      return
   }
   r := csv.NewReader(file)
   r.Comma = ','   // 以何種字元作分隔，預設為`,`。所以這裡可拿掉這行
   r.Comment = '#'   // 註解
   r.LazyQuotes = true
   r.ReuseRecord = true
   r.TrimLeadingSpace = true
   // r.Decode = traditionalchinese.Big5.NewEncoder()
   i := 0
   items := map[string][]Items{}
   for {
      record, err := r.Read()
      if err == io.EOF {
         break
      }
      if i == 0 {
         for j := 0; j < len(record); j++ {
            title[strings.ToLower(record[j])] = j
         }
         i += 1
         continue
      }
      if len(title) != len(record) {
         fmt.Printf("第%d行資料長度錯誤，應為%d，但偵測到：%d\n", i+1, len(title), len(record))
	 fmt.Printf("%v\n", record[0])
	 os.Exit(3)
      }
      uid := base64.StdEncoding.EncodeToString([]byte(record[title["question"]]))
      t := Items {
         Title:	   record[title["title"]],
         Question: record[title["question"]],
         Answer:   record[title["answer"]],
         Category: record[title["category"]],
         Contact:  record[title["contact"]],
         Url:      record[title["url"]],
         UUID:     uid,
      }
      items[t.Category] = append(items[t.Category], t)
      // datas.ITEMs = append(datas.ITEMs, t)
      i += 1
   }
   datas.ITEMs = items
   s, err := json.Marshal(datas)
   if err != nil {
      fmt.Println(err.Error())
      return
   }
   part := strings.Split(FilePath, ".")
   filename := part[0];
   if len(part) > 2 {
      filename = filename + "." + part[1]
   }
   if err := W(filename + ".json", s); err != nil {
      fmt.Println(err.Error())
      return
   }
/*
   subCategory := map[string][]Items{}
   for _, item := range datas.ITEMs {
      if len(subCategory[item.Category]) == 0 {
         subCategory[item.Category] = []Items{}
      }
      subCategory[item.Category] = append(subCategory[item.Category], item)
   }
   i = 10
   for k, t := range subCategory {
      rand.Seed(time.Now().UnixNano())
      id := strconv.FormatInt(int64(i), 10)
      ss := CSv {
         ID: id,
         Serial: rand.Intn(1000),
         Name: k,
         ITEMs: t,
      }
      i += 10
      sub, err := json.Marshal(ss)
      if err != nil {
         fmt.Println(err.Error())
         return
      }
      if err := W(part[0] + id + ".json", sub); err != nil {
         fmt.Println(err.Error())
         return
      }
      fmt.Println(k)
   }
*/
}
