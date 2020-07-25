# litebot-go
Litebot rewrite in Golang using DiscordGo

## TODO:
- More user friendly
    - Give users more information with !help
- More universal code
    - Merge commands, events and other features into modules
    - Modules can be enabled/disabled
    - All modules will show up in help
- Fix DMs
    - Won't respond to DMs because it can't determine what Guild the message is from
    - Getting rid of CommandEnviroment and having each command figure it out on its own


## Documentation

### Features
- Join/Leave messages 
- Kick/ban users
- Purge messages
- Block custom word list

### Join Message
- Sends a message when a user joins the Server
- Can be enabled/disabled with `!enable joinmessage` and `!disable joinmessage`
- Can be configured with `!set joinmessage Welcome {user}` and `!set joinchannel #General`

### Leave Message
- Sends a message when a user leaves the Server
- Can be enabled/disabled with `!enable leavemessage` and `!disable leavemessage`
- Can be configured with `!set leavemessage Welcome {user}` and `!set leavechannel #General`

### Kick
- Kicks the mentioned user
- `!kick @user#0000`

### Ban
- Bans the mentioned user
- `!ban @user#0000`

### Purge
- Deletes the specified number (<100) of messages
- `!purge 50`

### Enable/Disable
- Used to enable or disable a command
- `!enable kick` or `!disable kick`

### Set
- Used to set a custom value
- `!set prefix !`

### Block/Unblock
- Used to block a specified term
- `!block discord.gg`