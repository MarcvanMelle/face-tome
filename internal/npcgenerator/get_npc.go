package npcgenerator

import (
	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"
)

// GetNPC returns the fully generated NPC response based on provided protocol buffer parameters
func GetNPC(request *api.GetNPCRequest) (*api.GetNPCResponse, error) {
	npcResponse := &api.GetNPCResponse{NpcData: &api.NPC{FirstName: "Dougie"}}

	return npcResponse, nil
}
