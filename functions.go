package main

import (
	"github.com/bwmarrin/discordgo"
)

func MemberHasPermission(env *CommandEnvironment, permission int) (bool, bool, error) { // Perm, Admin, Error
	// Iterate through the role IDs stored in member.Roles
	// to check permissions
	for _, roleID := range env.Member.Roles {
		role, err := env.session.State.Role(env.Guild.ID, roleID)
		if err != nil {
			return false, false, err
		}
		if role.Permissions&permission != 0 {
			if role.Permissions&discordgo.PermissionAdministrator != 0 {
				return true, true, nil
			}
			return true, false, nil
		}
	}

	return false, false, nil
}
