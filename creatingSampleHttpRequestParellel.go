package main
import "encoding/json"
import "net/http"
import "io/ioutil"
import "fmt"
import "strconv"
import "time"
import "runtime"

func creatingSampleHttpRequestParellel() {

	runtime.GOMAXPROCS(8) 

	for i := 1; i <= 10; i++ {
		go func(index int) {
			response, err := http.Get("http://5a530e1477e1d20012fa066a.mockapi.io/employeedata/" + strconv.Itoa(index))
			if err == nil {
				responseBody, readErr := ioutil.ReadAll(response.Body) 
				
				if readErr == nil {
					fmt.Println(responseBody)
					empList := new(Employee)
					json.Unmarshal(responseBody, &empList)

					fmt.Println(empList.Id)
				}
			}
		}(i);
	}

	time.Sleep(1000 * time.Millisecond)
}