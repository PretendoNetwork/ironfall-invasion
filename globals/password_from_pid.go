package globals

import (
	"context"

	"github.com/PretendoNetwork/nex-go/v2/types"
	pb_account "github.com/PretendoNetwork/grpc-go/account"
	"github.com/PretendoNetwork/nex-go/v2"
	"google.golang.org/grpc/metadata"
)

func PasswordFromPID(pid *types.PID) (string, uint32) {
	ctx := metadata.NewOutgoingContext(context.Background(), GRPCAccountCommonMetadata)

	response, err := GRPCAccountClient.GetNEXPassword(ctx, &pb_account.GetNEXPasswordRequest{Pid: pid.LegacyValue()	})
	if err != nil {
		globals.Logger.Error(err.Error())
		return "", nex.Errors.RendezVous.InvalidUsername
	}

	return response.Password, 0
}
