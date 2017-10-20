package main

import (
	"fmt"
	"strconv"
	"sync"
)

type testData struct {
	i    int64
	Name string
	Data string
}

func consumer(ch chan testData) {
	i := 0
	for {
		dt := <-ch
		i++
		if i%100000 == 0 {
			fmt.Printf("Data %d\n", dt.i)
		}
	}
}

func consumer2(ch chan *testData) {
	i := 0
	for {
		dt := <-ch
		i++
		if i%100000 == 0 {
			fmt.Printf("Data %d\n", dt.i)
		}
	}
}

func consumer3(dt *testData) {

	if dt.i%100000 == 0 {
		fmt.Printf("Data %d\n", dt.i)
	}

}

func consumer4(dt testData) {

	if dt.i%100000 == 0 {
		fmt.Printf("Data %d\n", dt.i)
	}

}

func test1(num int64) {
	c := make(chan testData)

	go consumer(c)
	for i := int64(0); i < num; i++ {
		data := strconv.Itoa(int(i)) + `Bacon ipsum dolor amet salami doner pork, frankfurter kevin andouille sausage picanha tail bresaola rump ham venison. Alcatra turkey shankle corned beef landjaeger spare ribs rump prosciutto biltong pig tri-tip turducken pork belly ham. Pastrami turkey corned beef swine leberkas filet mignon, fatback short ribs picanha doner sirloin. Ribeye shoulder sirloin cupim. Tri-tip andouille shank, cow rump frankfurter turducken filet mignon.
		
		Sausage beef filet mignon, ball tip biltong meatball kielbasa jerky burgdoggen. Strip steak salami pig, pastrami beef kevin chuck ground round fatback capicola pork belly tongue venison pancetta. Shoulder pork belly turducken tri-tip meatball chuck filet mignon ground round. Tenderloin boudin short ribs ham, filet mignon t-bone flank bacon strip steak.
		
		Landjaeger pancetta ball tip bresaola spare ribs. Chuck venison beef ribs pig, burgdoggen brisket meatloaf rump pork belly sirloin kevin tail ham hock chicken drumstick. Capicola tri-tip meatloaf, prosciutto chuck short loin pig turducken flank ham hock t-bone kevin tongue fatback ribeye. Burgdoggen kielbasa pig t-bone porchetta meatball hamburger doner spare ribs short loin. Fatback shankle cupim ground round ribeye tail leberkas tri-tip. Tenderloin doner pork loin drumstick landjaeger rump. Frankfurter brisket capicola tongue sirloin shankle pork ribeye.`

		c <- testData{Name: data, i: i}
	}
}

func test2(num int64) {
	c2 := make(chan *testData)
	go consumer2(c2)
	for i := int64(0); i < num; i++ {
		data := strconv.Itoa(int(i)) + `Bacon ipsum dolor amet salami doner pork, frankfurter kevin andouille sausage picanha tail bresaola rump ham venison. Alcatra turkey shankle corned beef landjaeger spare ribs rump prosciutto biltong pig tri-tip turducken pork belly ham. Pastrami turkey corned beef swine leberkas filet mignon, fatback short ribs picanha doner sirloin. Ribeye shoulder sirloin cupim. Tri-tip andouille shank, cow rump frankfurter turducken filet mignon.
		
		Sausage beef filet mignon, ball tip biltong meatball kielbasa jerky burgdoggen. Strip steak salami pig, pastrami beef kevin chuck ground round fatback capicola pork belly tongue venison pancetta. Shoulder pork belly turducken tri-tip meatball chuck filet mignon ground round. Tenderloin boudin short ribs ham, filet mignon t-bone flank bacon strip steak.
		
		Landjaeger pancetta ball tip bresaola spare ribs. Chuck venison beef ribs pig, burgdoggen brisket meatloaf rump pork belly sirloin kevin tail ham hock chicken drumstick. Capicola tri-tip meatloaf, prosciutto chuck short loin pig turducken flank ham hock t-bone kevin tongue fatback ribeye. Burgdoggen kielbasa pig t-bone porchetta meatball hamburger doner spare ribs short loin. Fatback shankle cupim ground round ribeye tail leberkas tri-tip. Tenderloin doner pork loin drumstick landjaeger rump. Frankfurter brisket capicola tongue sirloin shankle pork ribeye.`

		c2 <- &testData{Name: data, i: i}
	}
}

func test3(num int64) {
	mu := sync.RWMutex{}
	for i := int64(0); i < num; i++ {
		data := strconv.Itoa(int(i)) + `Bacon ipsum dolor amet salami doner pork, frankfurter kevin andouille sausage picanha tail bresaola rump ham venison. Alcatra turkey shankle corned beef landjaeger spare ribs rump prosciutto biltong pig tri-tip turducken pork belly ham. Pastrami turkey corned beef swine leberkas filet mignon, fatback short ribs picanha doner sirloin. Ribeye shoulder sirloin cupim. Tri-tip andouille shank, cow rump frankfurter turducken filet mignon.
		
		Sausage beef filet mignon, ball tip biltong meatball kielbasa jerky burgdoggen. Strip steak salami pig, pastrami beef kevin chuck ground round fatback capicola pork belly tongue venison pancetta. Shoulder pork belly turducken tri-tip meatball chuck filet mignon ground round. Tenderloin boudin short ribs ham, filet mignon t-bone flank bacon strip steak.
		
		Landjaeger pancetta ball tip bresaola spare ribs. Chuck venison beef ribs pig, burgdoggen brisket meatloaf rump pork belly sirloin kevin tail ham hock chicken drumstick. Capicola tri-tip meatloaf, prosciutto chuck short loin pig turducken flank ham hock t-bone kevin tongue fatback ribeye. Burgdoggen kielbasa pig t-bone porchetta meatball hamburger doner spare ribs short loin. Fatback shankle cupim ground round ribeye tail leberkas tri-tip. Tenderloin doner pork loin drumstick landjaeger rump. Frankfurter brisket capicola tongue sirloin shankle pork ribeye.`
		mu.Lock()
		consumer3(&testData{Name: data, i: i})
		mu.Unlock()
	}
}

func test4(num int64) {
	mu := sync.RWMutex{}
	for i := int64(0); i < num; i++ {
		data := strconv.Itoa(int(i)) + `Bacon ipsum dolor amet salami doner pork, frankfurter kevin andouille sausage picanha tail bresaola rump ham venison. Alcatra turkey shankle corned beef landjaeger spare ribs rump prosciutto biltong pig tri-tip turducken pork belly ham. Pastrami turkey corned beef swine leberkas filet mignon, fatback short ribs picanha doner sirloin. Ribeye shoulder sirloin cupim. Tri-tip andouille shank, cow rump frankfurter turducken filet mignon.
		
		Sausage beef filet mignon, ball tip biltong meatball kielbasa jerky burgdoggen. Strip steak salami pig, pastrami beef kevin chuck ground round fatback capicola pork belly tongue venison pancetta. Shoulder pork belly turducken tri-tip meatball chuck filet mignon ground round. Tenderloin boudin short ribs ham, filet mignon t-bone flank bacon strip steak.
		
		Landjaeger pancetta ball tip bresaola spare ribs. Chuck venison beef ribs pig, burgdoggen brisket meatloaf rump pork belly sirloin kevin tail ham hock chicken drumstick. Capicola tri-tip meatloaf, prosciutto chuck short loin pig turducken flank ham hock t-bone kevin tongue fatback ribeye. Burgdoggen kielbasa pig t-bone porchetta meatball hamburger doner spare ribs short loin. Fatback shankle cupim ground round ribeye tail leberkas tri-tip. Tenderloin doner pork loin drumstick landjaeger rump. Frankfurter brisket capicola tongue sirloin shankle pork ribeye.`
		mu.Lock()
		consumer4(testData{Name: data, i: i})
		mu.Unlock()
	}

}

func main() {

	test1(10000000)
	test2(10000000)

}
