package main

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{concurrentConn: cc, bucket: make(chan int, cc)}
}
func (cl *ConnLimiter) GetConn() bool {

	//select {
	//case cl.bucket <- 1:
	//	log.Println("true")
	//	return true
	//default:
	//	log.Println("false")
	//	return false
	//}
	log.Println("bucket size:", len(cl.bucket))
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("reached the rate limitation:%d", len(cl.bucket))
		return false
	}
	cl.bucket <- 1
	return true
}
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("new connect comming: %d", c)
}
