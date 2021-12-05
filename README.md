# litebot

Lite-bot is a simple and customizable bot for server moderation with features like kick, ban, purge, and blocking words. Nearly everything can be disabled or customized.

Lite-bot is a simple and customizable bot for server moderation with features like marking a message as a spoiler and blocking terms. Nearly everything can be disabled or customized so you only have to worry about the features you actually want. More features are always the way.

## Documentation

### Features

- Custom prefix
- Join/Leave messages
- Kick/ban users
- Purge messages
- Block custom word list
- Mark messages as spoiler

### Help

- Lists available commands
- Can display extra information about commands `!help kick`

### Join Message

- Sends a message when a user joins the Server
- Can be enabled/disabled with `!enable joinmessage` and `!disable joinmessage`
- Can be configured with `!joinmessage #General` and `!joinmessage #General Welcome {user}`

### Leave Message

- Sends a message when a user leaves the Server
- Can be enabled/disabled with `!enable leavemessage` and `!disable leavemessage`
- Can be configured with `!leavemessage #General` and `!leavemessage #General Goodbye {user}`

### Kick

- Kicks the mentioned user
- `!kick @user#0000`

### Ban

- Bans the mentioned user
- `!ban @user#0000`

### Purge

- Deletes the specified number (1 to 99) of messages
- `!purge 50`

### Enable/Disable

- Used to enable or disable a command
- `!enable kick` or `!disable kick`

### Prefix

- Used to set a custom prefix
- `!prefix +`

### Block/Unblock

- Used to block a specified term
- Admins are exempt from message filtering
- `!block discord.gg`

### Spoiler

- Used to mark messages as spoilers
- lite-bot will delete and resend the message as a spoiler
- `!spoiler` or `!spoiler https://discord.com/channels/123456789/123456789/123456789`

## TODO

- Move/copy messages
  -!move #other-channel to move the message in the new channel
- Intro Music
  - Users can set a music clip that will play when they join the voice channel
  - Probably really annoying so limit clip length and/or number of daily plays
- Queue system
  - Useful for games like Among us or Animal Crossing where limited lobby size
- React for role system
  - Useful for big servers
- Fix DMs
  - Not really necessary right now
  - Can't process DMs because it can't determine what Guild the message is from
