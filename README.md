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

1. Build the control binary.

```sh
go build -o bcctl control/main.go
```

2. Generate some keys. You should see `node.key`, `node.pub`, `master.key` and `master.pub` in the directory after running
this:

```sh
./bcctl -generate
```

3. Now copy the relevant keys and build the agent:

```sh
mkdir -p agent/fs
cp master.pub node.key agent/fs/
go build -o bcagent agent/main.go
```

4. Install the agent (`bcagent`) somewhere, then run it.

5. Now using `bcctl` send a message. The agent should log it.

```
$ ./bcctl -send hi
+QpbZfgMkRsHXlhzNOhMvWZvrG6L+x/9UId5S7zBzasAMVexY0EycTu2I/kZwCpaW0Sv0tTF+D7oBUWMPUI+Ag==
Sent 68
```

```
Started listener
Reading packet from 127.0.0.1:48809
Message: hello
```

You're done!

## To-do

- Letting nodes respond to commands
- Shell command execution
- Basic machine stats

## To-maybe-do

- Hiding commands in mDNS announcements (if that's normal on your network anyway)
- Using nodes as traffic relays
