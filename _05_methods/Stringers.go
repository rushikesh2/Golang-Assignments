ppackage main

 import "fmt"

 type IPAddr [4]byte


//convert IPAddr{1, 2, 3, 4} should print as "1.2.3.4"
 func (i IPAddr) String() string {
 	return fmt.Sprintf("%d.%d.%d.%d",i[0],i[1],i[2],i[3])
 }
 func main() {
 	hosts := map[string]IPAddr{
 		"loopback":  {127, 0, 0, 1},
 		"googleDNS": {8, 8, 8, 8},
 	}
 	for name, ip := range hosts {
 		fmt.Printf("%v: %v\n", name, ip)
 	}
 }