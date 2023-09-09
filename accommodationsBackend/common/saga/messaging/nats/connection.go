package nats

import (
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

var (
	natsComp *NATSComponent
	once     sync.Once
)

// NATSComponent contains reusable logic related to handling NATS connections.
type NATSComponent struct {
	cmu  sync.Mutex
	nc   *nats.Conn
	name string
}

func NewNATSComponent(name string) *NATSComponent {
	return &NATSComponent{
		name: name,
	}
}
func GetNATSComponent() *NATSComponent {
	once.Do(func() {
		log.Println("Pravi se nova prvi put!")
		natsComp = NewNATSComponent("accommodation-booking")
		err := natsComp.ConnectToServer("nats://nats:4222")
		if err != nil {
			log.Println("Greska pri konektovanju na server (GetNatsComp).")
			log.Fatal(err)
		} else {
			log.Println("Uspesno konektovan na server (GetNatsComp)")
		}
	})
	return natsComp
}
func (c *NATSComponent) ConnectToServer(url string, options ...nats.Option) error {
	c.cmu.Lock()
	defer c.cmu.Unlock()

	nc, err := nats.Connect(url, options...)
	if err != nil {
		return err
	}
	c.nc = nc
	return nil
}

func (c *NATSComponent) NATS() *nats.Conn {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	return c.nc
}

func (c *NATSComponent) JetStreamContext(opts ...nats.JSOpt) (nats.JetStreamContext, error) {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	jsContext, err := c.nc.JetStream(opts...)
	return jsContext, err
}

func (c *NATSComponent) Name() string {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	return c.name
}

func (c *NATSComponent) Shutdown() error {
	c.NATS().Close()
	return nil
}
