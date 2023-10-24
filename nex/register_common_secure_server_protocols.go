package nex

import (
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-common-go/matchmake-extension"
	matchmaking "github.com/PretendoNetwork/nex-protocols-common-go/matchmaking"
	matchmaking_ext "github.com/PretendoNetwork/nex-protocols-common-go/matchmaking-ext"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-common-go/nat-traversal"
	secureconnection "github.com/PretendoNetwork/nex-protocols-common-go/secure-connection"

	"github.com/PretendoNetwork/ironfall-invasion/globals"
)

func registerCommonSecureServerProtocols() {
	secureconnection.NewCommonSecureConnectionProtocol(globals.SecureServer)

	matchmaking.NewCommonMatchMakingProtocol(globals.SecureServer)
	matchmaking_ext.NewCommonMatchMakingExtProtocol(globals.SecureServer)

	commonMatchmakeExtensionProtocol := matchmake_extension.NewCommonMatchmakeExtensionProtocol(globals.SecureServer)

	commonMatchmakeExtensionProtocol.GetUserFriendPIDs(globals.GetUserFriendPIDs)

	nat_traversal.NewCommonNATTraversalProtocol(globals.SecureServer)
}
