package bank_account

type storeInterface interface {
}

type Server struct {
	store storeInterface
}

func NewServer(store storeInterface) *Server {
	return &Server{store: store}
}
