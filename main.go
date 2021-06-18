package main

import (
	"fmt"

	basictype "github.com/chandrasekar-om/go-basic/basicType"
	"github.com/chandrasekar-om/go-basic/commonFunctions"
	"github.com/chandrasekar-om/go-basic/concurrency"
	errorhandling "github.com/chandrasekar-om/go-basic/errorHandling"
	"github.com/chandrasekar-om/go-basic/interfaceType"
	"github.com/chandrasekar-om/go-basic/ios"
	"github.com/chandrasekar-om/go-basic/method"
	"github.com/chandrasekar-om/go-basic/sorting"
	"github.com/chandrasekar-om/go-basic/timers"
)

const message = "Hello Go Basic"

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	// 	rw.WriteHeader(http.StatusOK)
	// 	rw.Write([]byte(message))
	// })
	// err := http.ListenAndServe(":9090", mux)
	// if err != nil {
	// 	log.Fatalf("Searver fail to start : %v", err)
	// }
	basictype.DoStruct()
	interfaceType.DoDrive()
	//aggregateType.DoArray()
	//fmt.Println(closure.Count()())
	//aggregateType.AccessFieldsWithDot()
	//aggregateType.AccessFieldsWithPointer()
	//aggregateType.AccessFieldsWithLiterals()
	//aggregateType.EncapsulateNewStructCreation()
	//
	//referenceType.DoSlice()
	//referenceType.DoPointer()
	//referenceType.DoMap()
	//referenceType.DoFunction()
	//interfaceType.DoDrive()
	method.DoMethod()
	r, e := errorhandling.IsPositive(-10)
	fmt.Println(r, e)
	errorhandling.DoPanic()
	errorhandling.DoDefer()

	concurrency.DoGoroutine()
	concurrency.DoChannels()
	concurrency.DoChannelBuffering()
	concurrency.DoChannelSynchronization()
	concurrency.DoChannelDirections()
	concurrency.DoSelectWithChannels()
	concurrency.DoTimeoutsWithChannels()
	concurrency.DoNonBlockingChannel()
	concurrency.DoClosingChannels()
	concurrency.DoRangeoverChannels()
	concurrency.DoWorkerPools()
	concurrency.DoWeightGroup()
	concurrency.DoRateLimiting()
	concurrency.DoAtomicCounters()
	concurrency.DoMutexes()
	concurrency.DoStatefulGoroutines()

	//Timer
	timers.DoTimers()
	timers.DoTickers()
	//Sorting
	sorting.DoSorting()
	sorting.DoSortingByFunctions()
	//collectionFunctions
	commonFunctions.FMain()
	commonFunctions.DoStringFunctions()
	ios.DoRead()
	ios.DoWrite()
}
