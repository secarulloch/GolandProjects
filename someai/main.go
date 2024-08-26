package main

import (
	"context"
	"fmt"
	"time"
)

// Read from Input
//func main() {
//	for true {
//		fmt.Print("What is your name? > ")
//		reader := bufio.NewReader(os.Stdin)
//		line, err := reader.ReadString('\n')
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("hello %s\n", line)
//	}
//}

// Reading from a File
//func main() {
//	file, _ := os.OpenFile("hello.txt", os.O_RDONLY, 0666)
//	defer file.Close()
//	reader := bufio.NewReader(file)
//	for {
//		line, err := reader.ReadString('\n')
//		fmt.Printf("> %s", line)
//		if err != nil {
//			return
//		}
//	}
//}

// Custom Data: Go Structs
type Message struct {
	Hello   string
	ignored string
	Test    string
}

// Struct with Marshalling and a Tag Line
type Hello struct {
	Message string `json:"hellooo"`
}

//func main() {
//	h := Message{Hello: "world", ignored: "hello", Test: "test"}
//	AsString, _ := json.Marshal(h)
//	fmt.Printf("%s\n", AsString)
//	// Struct with Marshalling and a Tag Line
//	h1 := Hello{Message: "world"}
//	b, _ := json.Marshal(h1)
//	fmt.Printf("%s\n", string(b))
//	//fmt.Printf("%+v\n", h)
//}

// Writing and Reading Structs from Files
type Salary struct {
	Basic float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

// Marshalling a Struct to File via JSON
//func main() {
//	data := Employee{
//		FirstName: "Nicolas",
//		LastName:  "Modrzyk",
//		Email:     "hellonico at gmail.com",
//		Age:       43,
//		MonthlySalary: []Salary{{Basic: 15000.00}, {Basic: 16000.00},
//			{Basic: 17000.00}},
//	}
//	file, _ := json.MarshalIndent(data, "", " ")
//	_ = ioutil.WriteFile("my_salary.json", file, 0644)
//}

// Reading a Struct from a File Containing JSON
//func main() {
//	jsonFile, _ := os.Open("my_salary.json")
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//	var employee Employee
//	_ = json.Unmarshal(byteValue, &employee)
//	// fmt.Printf("%+v", employee)
//	// pretty-print content by reverting to JSON
//	json, _ := json.MarshalIndent(employee, "", "  ")
//	fmt.Println(string(json))
//}

// Slicing Program Arguments
//func main() {
//	programName, questions := os.Args[0], os.Args[1:]
//	fmt.Println("Starting:%s", programName)
//
//	if len(questions) == 0 {
//		fmt.Printf("Usage:%s <question1> <question2> ...", programName)
//	} else {
//		for i, question := range questions {
//			fmt.Printf("Question [%d] > %s\n", i, question)
//		}
//	}
//}

// Loading Environment Variables Using the godotenv Library
//func main() {
//	godotenv.Load()
//
//	s3Bucket := os.Getenv("S3_BUCKET")
//	secretKey := os.Getenv("SECRET_KEY")
//
//	fmt.Printf("S3: %s and secret: %s", s3Bucket, secretKey)
//}

// Asynchronous Code: Go Routines
//func printNumbers() {
//	for i := 0; i < 10; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Printf("%d", i)
//	}
//}
//
//func main() {
//	go printNumbers()
//	printNumbers()
//}

// Asynchronous Code: Go Routines and Channels
//func printNumbers(c chan int) {
//	for i := 0; i < 10; i++ {
//		c <- i
//		time.Sleep(100 * time.Millisecond)
//	}
//	close(c)
//}
//
//func main() {
//	c := make(chan int)
//	go printNumbers(c)
//
//	for num := range c {
//		fmt.Println(num)
//	}
//}

// For/Switch to Retrieve Values from the Go Channel
func printNumbers(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(c)
}

//func main() {
//	c := make(chan int)
//	go printNumbers(c)
//
//	for value := range c {
//		switch value {
//		case 0:
//			fmt.Println("Received 0")
//		case 1:
//			fmt.Println("Received 1")
//		default:
//			fmt.Println("Received other value")
//		}
//	}
//}

// Channel and Values: Even or Odd
//func main() {
//	c := make(chan int)
//	go printNumbers(c)
//
//	for value := range c {
//		switch value % 2 {
//		case 0:
//			fmt.Printf("Value: %d is even\n", value)
//		case 1:
//			fmt.Printf("Value: %d is odd\n", value)
//		default:
//			fmt.Println("Received a weird value")
//		}
//	}
//}

// Select from Different Asynchronous Sources
//func main() {
//	ch := make(chan string)
//
//	go func() {
//		time.Sleep(1 * time.Second)
//		ch <- fmt.Sprintf("hello")
//	}()
//
//	go func() {
//		time.Sleep(2 * time.Second)
//		ch <- fmt.Sprintf("world")
//	}()
//
//	for {
//		select {
//		case v := <-ch:
//			fmt.Printf("%s\n", v)
//		case <-time.After(3 * time.Second):
//			fmt.Println("waited 3 seconds")
//			os.Exit(0)
//		}
//	}
//}

// Context with Timeout
// Context When Timeout Is Cancelled
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//
//	go func() {
//		time.Sleep(2 * time.Second)
//		fmt.Println("Task finished")
//		cancel()
//	}()
//
//	select {
//	case <-ctx.Done():
//		fmt.Println("Context Done")
//		err := ctx.Err()
//		if err != nil {
//			fmt.Printf("err: %s", err)
//		}
//	}
//}

// Parent Context, Data, and Go Routines
func Task(ctx context.Context) {
	var i = 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context done")
			return
		default:
			i++
			fmt.Printf("Running [%s]...%d\n", ctx.Value("hello"), i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	go Task(context.WithValue(ctx, "hello", "world"))
	go Task(context.WithValue(ctx, "hello", "nico"))

	<-ctx.Done()
}
