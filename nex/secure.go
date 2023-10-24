package nex

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/ironfall-invasion/globals"

	nex "github.com/PretendoNetwork/nex-go"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewServer()
	globals.SecureServer.SetPRUDPVersion(1)
	globals.SecureServer.SetPRUDPProtocolMinorVersion(4)
	globals.SecureServer.SetDefaultNEXVersion(nex.NewNEXVersion(3, 7, 1))
	globals.SecureServer.SetKerberosPassword(globals.KerberosPassword)
	globals.SecureServer.SetAccessKey("feb81c7c")

	globals.SecureServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("== IRONFALL Invasion - Secure ==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("================================")
	})

	registerCommonSecureServerProtocols()

	globals.SecureServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_IRONFALL_SECURE_SERVER_PORT")))
}
