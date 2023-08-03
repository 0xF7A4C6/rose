package scanner

import (
	"runtime"
)

func InitScanner() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(i int) {
			s := GetScannerThread(i)
			go s.Scan()
		}(i)
	}
}
