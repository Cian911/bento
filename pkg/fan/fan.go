package fan

type Fan struct {
	IPAddress string
	ID        string
	Port      int
	Password  string
}

func NewFan() *Fan {
	return &Fan{}
}

func Connect() {}

func Send() {}

func Receive() {}
