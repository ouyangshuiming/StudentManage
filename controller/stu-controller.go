package controller

import (
	"StudentManage/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// 添加一个学生，所以只接收post请求（可以用postman模拟post请求）
// 浏览器中输入：localhost:8080/create?name="xiaoming"&class="2"&chinese="88"&math="90"&english="88"
func Create(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" { //只有GET请求才进行处理，POST请求就不处理
		//Http响应中包含三部分：状态行，响应头，响应体

		//（1）设置响应头中的"Content-Type"字段为json
		response.Header().Set("Content-Type", "application/json")
		//（2）设置状态行中的状态码为200
		response.WriteHeader(200)

		var student models.Student
		student.Name = request.URL.Query().Get("name")
		student.Class = request.URL.Query().Get("class")

		//读出来的是个字符串，需要转成float32,strconv.ParseFloat()函数返回的是float64，所以还要转一下
		Chinese, _ := strconv.ParseFloat(request.URL.Query().Get("chinese"), 32)
		student.Chinese = float32(Chinese)

		Math, _ := strconv.ParseFloat(request.URL.Query().Get("math"), 32) //读出来的是个字符串，需要转成float32
		student.Math = float32(Math)

		English, _ := strconv.ParseFloat(request.URL.Query().Get("english"), 32) //读出来的是个字符串，需要转成float32
		student.English = float32(English)

		models.CreatStudent(student)

		fmt.Println("创建成功")
		fmt.Println(student)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

// 删除一个学生
func Delete(response http.ResponseWriter, request *http.Request) {
	if request.Method == "DELETE" {
		response.Header().Set("Content-Type", "application/json")
		//（2）设置状态行中的状态码为200
		response.WriteHeader(200)

		url := request.URL
		// 此时url返回的是一个 *url.URL类型，即使你使用*url取出来的也不是字符串，而是一个url.URL结构体的对象
		// 所以需要使用(*url).Path，读取其中其中的Path属性

		idStr := strings.Split((*url).Path, "/")[2] // “/queryone/1“切分得到三个字符串——>"","queryone",”1“

		id, _ := strconv.Atoi(idStr) //将字符串转化成数字

		models.DeleteStudent(id)

	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

// 修改学生成绩，请求方式只能为post
// 前端请求中的请求体：
// {
// "id": 1,
// "english": 92
// }

func Updata(response http.ResponseWriter, request *http.Request) {
	var s models.Student

	// 解析请求体中的数据，并且将解析到的json数据解码到student这个结构体对象中
	json.NewDecoder(request.Body).Decode(&s)
	// 这里也可以：用另一种写法：
	// bodyBytes, _ := ioutil.ReadAll(request.Body) //读取请求中的请求体
	// json.Unmarshal(bodyBytes, &s)                //将请求体中的数据转成结构体对象，并且赋值给s

	models.UpdateStudent(s)

	// 返回成功消息
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Student updated successfully!"))
	return
}

// 获取单个学生信息（根据学生的id进行查询）
func QueryOne(response http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" { //只有GET请求才进行处理，POST请求就不处理
		response.Header().Set("Content-Type", "application/json")
		//（2）设置状态行中的状态码为200
		response.WriteHeader(200)

		url := request.URL
		// 此时url返回的是一个 *url.URL类型，即使你使用*url取出来的也不是字符串，而是一个url.URL结构体的对象
		// 所以需要使用(*url).Path，读取其中其中的Path属性

		id := strings.Split((*url).Path, "/")[2] // “/queryone/1“切分得到三个字符串——>"","queryone",”1“

		//id是字符串，将其转化为数字
		idNum, _ := strconv.Atoi(id)

		var student *models.Student
		student, _ = models.GetStudent(idNum)

		result, _ := json.Marshal(student)
		//（3）往响应体中写json数据
		response.Write(result)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

//QueryOne客户端浏览器输入   http://localhost:8080/queryone/2  得到：
//{"id":1002,"name":"Jason2","class":"三年2班","chinese":88,"math":99,"english":78}

// 获取所有学生信息
func QueryAll(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" { //只有GET请求才进行处理，POST请求就不处理
		//Http响应中包含三部分：状态行，响应头，响应体

		//（1）设置响应头中的"Content-Type"字段为json
		response.Header().Set("Content-Type", "application/json")
		//（2）设置状态行中的状态码为200
		response.WriteHeader(200)

		students := models.GetAllStudent()
		result, _ := json.Marshal(students)
		//（3）往响应体中写json数据
		response.Write(result)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

//QueryAll客户端浏览器得到：
//[  {"id":1000,"name":"Jason0","class":"三年0班","chinese":88,"math":99,"english":78},
//   {"id":1001,"name":"Jason1","class":"三年1班","chinese":88,"math":99,"english":78},
//   {"id":1002,"name":"Jason2","class":"三年2班","chinese":88,"math":99,"english":78},
//   {"id":1003,"name":"Jason3","class":"三年3班","chinese":88,"math":99,"english":78},
//   {"id":1004,"name":"Jason4","class":"三年4班","chinese":88,"math":99,"english":78}
//]
