package nex

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PretendoNetwork/ironfall-invasion/globals"
	nex "github.com/PretendoNetwork/nex-go/v2"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewPRUDPServer()

	globals.AuthenticationEndpoint = nex.NewPRUDPEndPoint(1)
	globals.AuthenticationEndpoint.ServerAccount = globals.AuthenticationServerAccount
	globals.AuthenticationEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.AuthenticationEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	globals.AuthenticationServer.BindPRUDPEndPoint(globals.AuthenticationEndpoint)

	globals.AuthenticationServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(3, 7, 1))
	globals.AuthenticationServer.ByteStreamSettings.UseStructureHeader = true
	globals.AuthenticationServer.AccessKey = "feb81c7c"

	globals.AuthenticationEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("== IRONFALL Invasion - Auth ==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID)
		fmt.Printf("Method ID: %#v\n", request.MethodID)
		fmt.Println("==============================")
	})

	registerCommonAuthenticationServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_IRONFALL_AUTHENTICATION_SERVER_PORT"))

	globals.AuthenticationServer.Listen(port)
}
