# broadc2

***This doesn't actually do anything yet. It's a glorified "read some data and verify with ed25519" system.***

---

A "C2 system" that gets its commands from LAN broadcasts.

## How

broadc2 works by listening for UDP broadcasts that are signed by a certain key*.

Instead of having to maintain a server to issue commands, I issue them from anywhere in the LAN.

What will the network admins point fingers at now!? The PC upstairs? The one below?

*I'll do encryption too, likely with a shared secret that's also baked in.

## Why

You might say, "This isn't practical". You are right. I'm making it for fun, based on a specific
idea I thought of for a specific LAN. So not really botnet of the year.

The software also doesn't try to hide itself, so even if I did use it it'd be easy to block.

## How to use

```sh
# Build the control binary
go build ./cmd/control
# Generate keys 
./control -generate
# Copy the keys over to be embed in the agent
mkdir -p cmd/agent/fs
cp master.pub node.key cmd/agent/fs/
# Build the agent
go build ./cmd/agent
```

Then install the agent (`agent`) somewhere and run it.

Now using `control` send a message. The agent should log it, and also run the message as a shell command.

```
$ ./control -send hi
+QpbZfgMkRsHXlhzNOhMvWZvrG6L+x/9UId5S7zBzasAMVexY0EycTu2I/kZwCpaW0Sv0tTF+D7oBUWMPUI+Ag==
Sent 68
```

```
Started listener
Reading packet from 127.0.0.1:48809
Message: hi
```

You're done!

## To-do

- Letting nodes respond to commands
- Shell command execution
- Basic machine stats

## To-maybe-do

- Hiding commands in mDNS announcements (if that's normal on your network anyway)
- Using nodes as traffic relays
