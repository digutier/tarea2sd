package biblio

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"

	"github.com/golang/protobuf/proto"

	pb "github.com/digutier/tarea2sd/biblio"
)

//Server asdasd
type Server struct {
}

//UploadLibro es para blabla
func (s *Server) UploadLibro(stream BiblioService_UploadLibroServer) error{
	
}