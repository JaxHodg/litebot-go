package main

func MemberHasPermission(env *CommandEnvironment, permission int) (bool, error) { //TODO: Check for admin
	guildID := env.message.GuildID
	userID := env.message.Author.ID
	member, err := env.session.State.Member(guildID, userID)
	if err != nil {
		if member, err = env.session.GuildMember(guildID, userID); err != nil {
			return false, err
		}
	}

	// Iterate through the role IDs stored in member.Roles
	// to check permissions
	for _, roleID := range member.Roles {
		role, err := env.session.State.Role(guildID, roleID)
		if err != nil {
			return false, err
		}
		if role.Permissions&permission != 0 {
			return true, nil
		}
	}

	return false, nil
}
