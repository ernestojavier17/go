package maps

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic tests",func() {
	arr := [...][]int{
		{20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5},
		{1,1,2,-2,5,2,4,4,-1,-2,5},
		{20,1,1,2,2,3,3,5,5,4,20,4,5},
		{10},
		{1,1,1,1,1,1,10,1,1,1,1},
		{5,4,3,2,1,5,4,3,2,10,10},
	}
	sol := [...]int{5,-1,5,10,10,1}
	for i,v := range arr {
		It(fmt.Sprintf("Testing input %v",v),func() {Expect(FindOdd(v)).To(Equal(sol[i]))})}
})
