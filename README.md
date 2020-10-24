# litebot
Lite-bot is a simple and customizable bot for server moderation with features like kick, ban, and purge, and blocking words. Nearly everything can be disabled or customized so you only have to worry about the features you actually want. More features are on the way.

## TODO:
- More user friendly
    - Give users more information with !help\
- Queue system
    - Useful for games like Among us or Animal Crossing
- Fix DMs
    - Won't respond to DMs because it can't determine what Guild the message is from

## Documentation

### Features
- Custom prefix
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
- Admins are exempt from message filtering
- `!block discord.gg`
