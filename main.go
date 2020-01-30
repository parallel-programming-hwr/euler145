package main

import (
	"fmt"
	"runtime"
)

var MAX uint64 = 1000000000

type Res struct{//this struct helps communicating with the channel eg. if channel will terminate, or what number it calculated
	n uint64
	nr uint64
	nnr uint64
	thread uint64
	terminate bool
}
func reverseNum(num uint64) uint64{
	var ret uint64 = 0
	for num!=0{
		ret = 10*ret + num%10
		num /=10
	}
	return ret
}
func isAllOdd(num uint64) bool {
	for num !=0 {
		if num%2==0{
			return false
		}
		num /=10
	}
	return true
}
func findRev(start uint64, end uint64, step uint64, ch chan Res){
	for i:=start ; i < end; i+=step{
		if i%10==0{
			continue
		}
		h := i+reverseNum(i)
		if isAllOdd(h){
			r :=new(Res)
			r.n =i
			r.nr = reverseNum(i)
			r.nnr = h
			r.thread = start // is equal to thread
			r.terminate = false
			ch <- *r
		}
	}
	r:= new(Res)
	r.terminate = true
	ch<-*r
}
func main() {
	numThreads := runtime.NumCPU()
	ch := make(chan Res)
	for i:=0 ;i< numThreads; i++{
		go findRev(uint64(i),MAX, uint64(numThreads),ch)
	}
	var n uint64 = 0
	for {
		h:=<-ch
		if h.terminate{
			numThreads--
			if numThreads <=0{
				fmt.Printf("\nTotal reversibles found:%d\n", n)
				break
			}
			continue
		}
		n++
		//fmt.Printf("zahl=%d rev(zahl)=%d zahl+rev(zahl)=%d n=%d thread=%d\n", h.n,h.nr,h.nnr,n,h.thread)


	}

}
