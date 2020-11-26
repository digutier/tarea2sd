package biblio

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"io"
	"bytes"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"


	"github.com/golang/protobuf/proto"
)

//Server asdasd
type Server struct {
}

//UploadLibro es para blabla
func (s *Server) UploadLibro(stream BiblioService_UploadLibroServer) error{
	for {
		UploadLibroRequest, err := stream.Recv()
		if err == io.EOF {
			log.Print("No hay mas")
			break
		}
		if err != nil {
			return err
		}
		chunk := UploadLibroRequest.GetData()
		print(chunk)
		log.Print("Recibi uno c:")

	}
	return nil 
}