package webserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/afaguilarr/go-example-webserver/app/src/cmd"
	"github.com/afaguilarr/go-example-webserver/app/src/crypto_client"
	"github.com/afaguilarr/go-example-webserver/app/src/http_helpers"
	"github.com/afaguilarr/go-example-webserver/app/src/users_client"
	"github.com/afaguilarr/go-example-webserver/proto"
)

type LogOutHandler struct {
	uc users_client.UsersClientHandlerInterface
}

func NewLogOutHandler(uc users_client.UsersClientHandlerInterface) LogOutHandler {
	return LogOutHandler{
		uc: uc,
	}
}

func (lh *LogOutHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	log.Println("LogOut endpoint called")
	switch r.Method {
	case http.MethodPost:
		cc := crypto_client.NewCryptoClientHandler(cmd.CryptoHost, cmd.DefaultPort)
		err := cc.CreateConnection()
		if err != nil {
			log.Fatalf("while creating Crypto Connection: %s", err)
		}
		defer func() {
			err := cc.CloseConnection()
			if err != nil {
				log.Fatalf("while closing Crypto Connection: %s", err)
			}
		}()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := lh.uc.LogOut(ctx, &proto.LogOutRequest{
			Username: "jiji",
		})
		if err != nil {
			http_helpers.ErrorHandler(w, r, http.StatusInternalServerError, "There was an error D:")
		}
		_, err = fmt.Fprintf(w, "Greeting: %s", resp.GetEncryptedValue())
		if err != nil {
			log.Fatalf("Something went wrong with the 'gRPC test endpoint': %s", err)
		}
	default:
		http_helpers.ErrorHandler(w, r, http.StatusMethodNotAllowed, "")
	}
}
