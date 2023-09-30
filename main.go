package main

// func main() {
// 	start := time.Now()
// 	defer func() {
// 		fmt.Println(time.Since(start))
// 	}()
// 	arr := []string{"ankit", "sharma", "shivam"}
// 	for _, name := range arr {
// 		go calling(name)
// 	}

// 	time.Sleep(1 * time.Second)
// }

// func calling(ip string) {
// 	fmt.Println(ip)
// 	// time.Sleep(2 * time.Second)
// }

// func main() {
// 	wg := &sync.WaitGroup{}
// 	mut := &sync.RWMutex{}
// 	arr := []int{1, 2, 3}
// 	wg.Add(4)
// 	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
// 		mut.Lock()
// 		arr = append(arr, 4)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)
// 	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
// 		mut.Lock()
// 		arr = append(arr, 5)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)
// 	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
// 		mut.Lock()
// 		arr = append(arr, 6)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)
// 	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
// 		mut.RLock()
// 		fmt.Printf("ankit %v\n", arr)
// 		mut.RUnlock()
// 		wg.Done()
// 	}(wg, mut)
// 	wg.Wait()
// 	fmt.Println(arr)
// }

// sync communication unbuffered communication
// func main() {
// 	chann := make(chan int)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)
// 	go func(myChan chan int, wg *sync.WaitGroup) {
// 		x := <-myChan
// 		fmt.Println(x)
// 		defer wg.Done()
// 	}(chann, wg)
// 	go func(myChan chan int, wg *sync.WaitGroup) {
// 		defer wg.Done()
// 		myChan <- 10
// 	}(chann, wg)
// 	wg.Wait()
// }

// channels
// func main() {
// 	chann := make(chan int, 4)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)
// 	go func(myChan chan int, wg *sync.WaitGroup) {
// 		fmt.Println(<-myChan)
// 		fmt.Println(<-myChan)
// 		fmt.Println(<-myChan)
// 		myChan <- 10
// 		defer wg.Done()
// 	}(chann, wg)
// 	go func(myChan chan int, wg *sync.WaitGroup) {
// 		defer wg.Done()
// 		myChan <- 10
// 		myChan <- 20
// 		myChan <- 40
// 		fmt.Println(<-myChan)
// 		close(myChan)
// 	}(chann, wg)
// 	wg.Wait()
// }

// 0 case in channels
// func main() {
// 	chann := make(chan int, 3)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)
// 	go func(myChan chan int, wg *sync.WaitGroup) {
// 		val, isClose := <-myChan
// 		fmt.Println(isClose, val)
// 		defer wg.Done()
// 	}(chann, wg)
// 	go func(myChan chan int, wg *sync.WaitGroup) {
// 		defer wg.Done()
// 		close(myChan)

// 	}(chann, wg)
// 	wg.Wait()
// }

// https://docs.google.com/document/d/1URAlMYnXUyTdhuXZjfII_sNhm5jHL3jRPmlYlPg_0Cw/edit
// Go interface is a group of method signatures that a type can implement.
// an interface is just a collection of method signatures.
// group together typres based on the their behaviour
// In Go methods are functions that are associated with a specific type

// closer in go
// func outerFunction() func() int {
// 	x := 10
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// func main() {
// 	innerFunc := outerFunction()
// 	fmt.Println(innerFunc())
// 	fmt.Println(innerFunc())
// }

// methods in golng
// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// func main() {
// 	myCircle := Circle{Radius: 5.0}
// 	area := myCircle.Area()
// 	fmt.Printf("The area of the circle is %.2f\n", area)
// }

// type Shape interface {
//   area() float32
// }

// type Rectangle struct {
//   length, breadth float32
// }

// func (r Rectangle) area() float32 {
//   return r.length * r.breadth
// }

// type Triangle struct {
//   base, height float32
// }

// func (t Triangle) area() float32 {
//     return 0.5 * t.base * t.height
// }

// func calculate(s Shape) float32 {
//   return s.area()
// }
// func main() {
//   r := Rectangle{7, 4}
//   t := Triangle{8, 12}

//   rectangleArea := calculate(r)
//   fmt.Println("Area of Rectangle:", rectangleArea)

//   triangleArea := calculate(t)
//   fmt.Println("Area of Triangle:", triangleArea)

// }

//converting a Go data structure into its JSON
// data := MyStruct{Field1: "value1", Field2: 42}
// jsonData, err := json.Marshal(data)

//converting JSON data into a Go data structure.
// jsonData := []byte(`{"Field1": "value1", "Field2": 42}`)
// var data MyStruct
// err := json.Unmarshal(jsonData, &data)

// bind JSON data from an HTTP request body to a Go struct and perform validation
// var data MyStruct
// err := c.ShouldBindJSON(&data)
