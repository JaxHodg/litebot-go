# litebot-go
Litebot rewrite in Golang using DiscordGo

## TODO:
- Save data using structs
- Message filtering (Custom word list)
- Fix DMs
- Clean up Events

## Documentation

### Features
- Join/Leave messages 
- Kick/ban users
- Purge messages

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