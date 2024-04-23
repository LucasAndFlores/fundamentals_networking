# Fundamentals of networking - Hussein Nasser

## Section 2:

### Client-server architecture
- All started like: Separate application into two components. Expensive workload can be done on the server, such as CPU or memory intensive tasks. “Light” tasks maybe can live on the client (frontend) side. Client calls servers to perform expensive tasks.
- Clients need to communicate with the server, a communication model.

### OSI Model (Open systems interconnection model)
- 7 layers.
- Why a communication model?
-  Is necessary to build this model to be agnostic to the application. Imagine if you need to build a server that how you convert your data to works on wifi, and another one for fiber, another one to work with LTE and so on. We need to build this with a standard model to made easy and agnostic
- Network equipment management  would be hard without a standard.
- Decoupled inovation. Fiber vs WIFI (radio)
- The seven layers are:
- - Layer 7 = Application - Http/FTP/gRPC.  
- - Layer 6 = Presentation - Encoding/Serialization
- - Layer 5 = Session - Connection establishment, TLS
- - Layer 4 = Transport - UDP/TCP/QUICK (only care about ports) (SPORT/DPORT)
- - Layer 3 = Network - IP (only care about packet) (SIP/DIP)
- - Layer 2 = Data link - Frames, MAC addresses Ethernet (only care about frames) (SMAC/DMAC)
- - Layer 1 = Physical - Electric signals, fiber, or radio waves. (Only care about frames and convert them into bits to be transmitted) 
- Example sending a POST request to an HTTPS webpage (sender perspective):
- - Layer 7 = Application - Post request with a JSON data to HTTPs server.  
- - Layer 6 = Presentation - Serialise JSON to flat byte strings.
- - Layer 5 = Session - Request to establish a TCP connection/TLS
- - Layer 4 = Transport - Sends SYN request targeting the port 443 (HTTPS standard port)
- - Layer 3 = Network - SYN is placed an IP package and adds the source/dest IPs
- - Layer 2 = Data link - Each packet goes into a single frame and adds the source/dest MAC addresses.
- - Layer 1 = Physical - Each frame becomes string of bits which converted into either radio signal (wifi), electric signal (Ethernet) or light (fiber) 
- Example sending a POST request to an HTTPS webpage (receiver perspective):
- - Layer 1 = Physical - Radio, electric, light is received and coverted into digital bits
- - Layer 2 = Data link - The bits from layer 1 is assembled into frames.
- - Layer 3 = Network - The frames from layer 2 are assembled into IP packet.
- - Layer 4 = Transport - The IP packets from layer 3 are assembled into TCP segments. Deals with congestion control/flow control in case of TCP. If the Segment is SYN, we don’t need to go further into more layers as we are still processing the connection request (HTTPS case here).
- - Layer 5 = Session - The connection session is established or identified. We only arrived at this layer when necessary (three way handshake is done)
- - Layer 6 = Presentation - Deserialise flat byte strings back to JSON for the app to consume.
- - Layer 7 = Application - Application understands the JSON POST and all the process starts (receive a request, talk to service, then talk to database, store this, generate a log, send a response). 
- Necessarily we don’t need to communicate from layer 7 to 1 or vice versa. We can have intermediate layers (like session layer) communicate with each other, to establish a new session before send the request.
- About the two steps above (send/receiver a request), not almost all the time, have a direct connection, meaning that we can have it other steps over this communication. For example: We can have the sender, that communicates with a Switch (layer 2 only, resolving MAC addresses and addressing correctly), the switch communicates with a router (solve the communication between the IPs) and then start the process with the receiver.
- Firewall blocks certain applications from sending data or block unwanted packets to come trough your network. Some firewall is a layer 4 application, because need to look to IP address to identify where this come from. Some firewall’s can “transparent”. With this one, you can block some IPs addresses, for example
- Load balancing, for example, is on layer 4, where we have one port to receive but other to distribute, and we distribute this rewriting the packet (from layer 3) and sending to the current address.
- VPN’s is a layer 3 application, where it takes the IP packets and put them in another IP packets, in a simple words. We also have some VPNs in layer 2.
- If we try to bring this to “more real example” of today, we can have a receiver sending a request to a Proxy server, that send to a Load balancer/CDN on layer 7 and then reaches our backend server.  A reverse proxy server, for example, will only work until layer 4, where you map something like “I’m receiving traffic on port 80 and sending this to the 3000 (which is your application port, for example)”.
- Is simpler to deal with the layers 5-6-7 as just one layer, application. And this is what TCP/IP model does.
- TCP/IP model only have 4 of 7 layers: Application (group layer 5, 6 and 7), layer 4 transport, layer 3 Internet and layer 2 Data link. The physical layer is not officially covered in this model.

### Host to host communication:
- Send a message from host A to B. Each host network has a unique Media Access Control address  (MAC). Layer 2 communication.
- Let’s say that we have a network and we have 4 others host connected. Everyone in the network will get the message, but only B will accept it.
- Now, imagine a scenario where we have a lot of machines (millions). This would be hard, we should scan and transmit our message to everyone in this network, so we need something more effective. We need the addresses get better. “Routability”. So we get the IP address.
- The IP address is built in two parts. One part to identify the network, other is the host. We use network portion to eliminate networks. The part is used to find the host, and we still need the MAC addresses
- But inside of the host, we can have it many things happening at same time, so how do we know where to send this packet? Ports! Ports are used to address exactly what process you want to communicate with it.

## Section 3: Internet Protocol (IP)

### The IP Building blocks:
- Layer 3 property (network). Can be automatically or statically. Network and host portion. 4 bytes in IPv4 - 32 bits.
- a.b.c.d/x (a.b.c.d are integers) x is the network bits and remains are host.
- Example: 192.168.254.0/25. The first 24 bits (3 bytes) are network there rest 8 are for host. This is also called subnet.
- Subnet has a subnet mask, where is use to determine whether an IP address is the same subnet.
- Most networks consist of hosts and a Default gateway. Host A can talk to host B directly if both are in the same subnet. Otherwise, A sends it to someone who might know it, the gateway. The gateway has an IP address and each host should know it is a gateway.
- A good example of performance bottleneck affected by this: let’s say that we have a server in one subnet and the database is in another subnet. And we depend on router to connect these two parts. If our router is congested, we can have some delay between the communication. We can add a switch between these parts to have a performance increase and include them in the same subnet

### IP Packet:
- IP packets has headers and data sections. IP Packet header is 20 bytes and can go until 60. Data section can go up to 65536 packets.
- Anatomy of the packet: Version, total length.
- Time to live - How many hops can this packets survive. TTL
- Protocol: What protocol is inside the data section?
- Source and destination IP.
- Explicit congestion notification (ECN). It is a way that router controller can inform that is receiving to many requests and is getting congestion.

### ICMP, PING, Traceroute:
- ICMP (Internet Control Message Protocol): Lives in layer 3, information between host. Only have a destination and source IP, no ports. Designed for informational message (host, port unreachable). Uses IP directly. Ping and trace route use it. Doesn’t require listeners or ports to be opened.
- Some firewalls block ICMP for security reasons. That is why PING not work in those cases.
- PING information: When you try “ping IP”, you will see ICMP sequence, the ttl (time to live) of the message and the time taken to reach the IP. If you reach something local, maybe the TTL will be lower, but if you try something like google or twitter, the TTL and the time will be higher.
- PING: Lets says that we have the IP 192.168.1.3 and we try to PING the IP 192.168.10.3. The message would that would be send by my device should be something like this: Source IP: 192.168.1.3, Destination IP: 192.168.10.3. TTL 100 ICMP echo request. Once the destination IP received my PING, the TTL would be decremented by the time taken, for example: Source IP: 192.168.1.3, Destination IP: 192.168.10.3. TTL 96 ICMP echo request. Continuing the PING, the IP 192.168.1.3 send back the response like this: Source IP: 192.168.10.3, Destination IP: 192.168.1.3, TTL 100 ICMP echo request. Once the message arrived in my IP, the TTL will reduced too.
- PING, unreachable case: Lets says that we have the IP 192.168.1.3 and we try to PING the IP 192.168.10.3. The message would that would be send by my device should be something like this: Source IP: 192.168.1.3, Destination IP: 192.168.10.3. TTL 3 (reduced just to cause the error) ICMP echo request.  So my IP packet reaches my router, then other route and boom, the TTL was reached!. Once this happens, the router with the IP 192.168.5.100 will send a message to me like this:  Source IP: 192.168.5.100, Destination IP: 192.168.1.3. TTL 100 ICMP destination unreachable.
- Traceroute: Identify the “entire path” of your IP Packets. Increment TTL slowly. 
### ARP:
- ARP = address resolution protocol.
- Generally, when we want to send frames (on layer 2), we always will need the MAC address. But most of the time, we will not know the MAC address. This is where ARP enter the game. ARP is basically a key-value cache in the source IP, where you will reach your gateway, the gateway will broadcast this packet to all machines in the subnet masks and find the MAC address, send it back to the host and store it in  an ARP table, mapping the IP with the MAC address.
- Sometimes, this broadcast between gateway and the machines in that subnet mask can cause ARP poisoning.
- If you want to do a load balancing using virtual IP address, ARP becomes a critical layer. You want to have one IP address, but this IP is shared between seven machines, so the IP address doesn’t change.

### Practice: Capturing IP, ARP and ICMP Packets with TCPDUMP:
- TCPDUMP is a command-line tool where you can analyse ARP, IP packets and ICMP, apply filters or even save the packets and read with TCPDUMP

## Section 4: UDP

### What is UDP:
- User datagram protocol. Is a layer 4 protocol and have the ability to addresses processes in a host using ports. Ports here acts like an identifier to send and receive data. This protocol is stateless and should be simple. Layer 4 protocol.
- Use cases: Video streaming, VPN (TCP can made VPN operation hard, for example), DNS, webRTC.
- With the UDP, we have the concept of Multiplexing and Demultiplexing. Multiplexing is the way that we can address a lot of stuff in one place and demultiplexing is the opposite.
- Let’s say that we have a machine A, with the IP 10.0.0.1 and we have the app1 running on port 3333, app2 running on port 4454. And we have another machine B, with the IP 10.0.0.3 and the machine have the app1 running on port 6969, and app2 running on port 3333.
- Now, lets say the machine A, the app1 want to send a message to app1 in the machine B. Machine A will multiplex all this apps (ports) to UDP. The machine B will demultiplex UDP datagrams to each app.
- So the communication will have the source port and IP, and destination IP and port. When the second machines respond, it will have the same data.

### UDP Datagram:
- UDP header is 8 bytes only (IPV4). Datagram slides into an IP packet as data. Ports are 16 bit (0 to 65535).
- Anatomy: always contains source port, destination port, length, checksum and data.

### Pros and cons:
- Pros: Simple protocol and prioritise latency over consistency. Header size is small and this make consume less resources, like less bandwidth and memory. Stateless!!! And with low latency because of this, doesn’t have any handshake, order, retransmission or guaranteed delivery (TCP).
- Cons: No previous acknowledgment, so can be easily spoofed. No guarantee delivery. Connection-less (anyone can send data without prior knowledge). No flow-control, congestion control and ordered packets.

## Section 5: TCP

### What is TCP (Transmission control protocol):
- Stands for transmission control protocol (focus on Transmission and protocol!!!). Layer 4 protocol. Ability to address processes in host using ports (similar to UDP on these points). It “controls” the transmission unlike UDP, which is a “firehouse”. It is stateful, mostly because you have to establish a connection and requires a handshake. 20 bytes headers Segment (can go to 60). You send segments.
- Use cases: Reliable communication, remote shell, database connections, Web communications and any bidirectional communication.
- TCP connection: connection is on Layer 5 (session). Connection should be established before sends data, otherwise, can be dropped or denied. Connection is an agreement between client and server. The connection is identified by 4 properties (2 pairs): SourceIp-SourcePort; DestinationIp-DestinationPort. From all this identifiers, the system hash all of this and store one hash in memory (or disk) and this is locked up in a lookup in the operation system, and this matches something called file descriptor. And this file descriptor contains the session and the state. So, you basically hash it, store. When a communication happen, you look up and know it if a connection is established or not.
- Can’t send data outside of a connection. Sometimes called socket or file descriptor. Requires a 3-way TCP handshake. Segments are sequenced and ordered, and this can only be possible in TCP, because at the end, all of this will turn into IP packets. And IP packets can’t be ordered or sequenced. Segments are acknowledged and because of this, lost segments are retransmitted.
- Now, simulating a connection establishment: Lets say the app1 on 10.0.01 want to send data to AppX on 10.0.0.2. App1 send SYN (with source IP, source port, SYN, destination IP and destination port) to AppX to synchronous sequence numbers. AppX sends SYNC/ACK to synchronous is sequence number. App1 ACKs AppX SYN. This is the 3-way TCP handshake
- Simulating sending data: App1 sends data to AppX. App1 encapsulates the data in a segment and send it. AppX acknowledges the segment. Can App1 send new segments before ACK of old segments arrives? Yes, they can, but it have a limit here. The limit will be established by the server and this should be set as flow control. But the router can handle it? This should be the congestion control.
- Parallel execution-acknowledgment: App1 sends segments 1,2,3 in parallel to AppX. AppX acknowledge all of them with a single ACK 3.
- Close connection: App1 wants to close the connection. App1 sends FIN, AppX ACK. AppX sends FIN, App1 ACK. Four way handshake.

### TCP Segment:
- TCP segment Header is 20 bytes and can go up to 60 bytes, this happens because we may have all TCP stuff (like sequences, acknowledgment, flow control and more) inside of the headers. TCP segments slides into IP packets as data. Port are 16 bit.
- Window size, inside of the headers, is responsible to tell to the client how much data can be handle it by the server.
- Maximum Segment size: Segment size depends the MTU of the network. Usually 512 bytes can go up to 1460. Default MTU in the internet is 1500 (results in MSS 1460). Jumbo frames MTU goes to 9000 or more, mss can be larger in jumbo frames cases.

### Flow control:
- How much the receiver can handle?
- When TCP segments arrive they are put in receiver’s buffer. If we kept sending data, the receiver will be overwhelmed and segments will be dropped. Solution? Let the send know how much you can handle.
- Window size (receiver window) RWND: 16 bit - up to 64kb. Updated with each acknowledgment. Tells the send how much  to send before waiting for ACK. Receiver can decide to decrease window size (out of memory) more important stuff.
- Sliding window. The sender can’t keep waiting for receiver to acknowledge all the segments. Whatever gets acknowledge moves. We “slide” the window between the segments, and once this segments are acknowledged, we include more. Sender maintains the sliding window for the receiver
- Window scaling: 64kb is too small. We can’t increase the bits on the segment. Meet window scaling factor (0-14). Window size can go up to 1GB. Only exchanged during the handshake.
- Summarising: Receiver host has a limit. We need to let the sender know how much it can send. Receiver window is in the segment. Sender maintains the sliding window to know how much it can send. Window scaling can increase that.

### Congestion control:
- How much data a network can handle it.
- The might handle the load but the middle boxes might not. The routers in the middle have a limit. We don’t want to congest the network with data. We need to avoid congestion. A new window > Congestion window (CWND), acknowledged by the sender.
- Two congestion control algorithms: TCP Slow start (the faster one): Start slow but goes fast. CWND and increase +1 MSS (Maximum segment size) after each ACK, it follows a path that is exponential. Congestion avoidance: once slow start reaches it threshold this kicks in. CWND and + 1 MSS after complete RTT.
- CWND must not exceeds RWND.
- Slow start: CWND starts with 1 MSS (or more), Send 1 segment and waits for ACK. With EACH ACK received, CWND is incremented by 1 MSS. Until we reach slow start threshold we switch to congestion avoidance.
- Congestion avoidance: Send CWND worth of segments and wait for ACK, Only when ALL segments are ACKed add UP to one MSS to CWND. Plus one.
- We don’t want routers dropping packets! Can routers let us know when congestion hit? Yes, with ECN (Explicit congestion notification). Routers and middle boxes can tag IP packets with ECN, which is an IP header bit, and the receiver will copy this bit back to the sender. So with all of this, routers don’t drop packets just let me know you are reaching your limit. Once this happens, you can reduce the CWND.
- While the receiver may handle large data middle boxes not. Middle routers buffers may fill up. Need to control the congestion in the network. Sender can send segments up to CWND or RWND without ACK. Isn’t normally a problem in hosts connected directly.  You want to send data as much as possible, but without waiting for acknowledgement.

### Congestion detection: Slow start vs Congestion avoidance:
- The moment we get timeouts, dup ACKs or packets drops. The slow start threshold reduced to half of whatever unacknowledged data is sent (roughly CWND/2 if all CWND worth of data is unacknowledged). The CWND is reset to 1 and we start over. Min slow start threshold is 2*MSS

### NAT(Network Address Translation):
- IPv4 is limited only 4 Billion. Private IP is different from Public IP address, and this maybe can be identified by the beginning of the number and submasks.  For example, IPs like 192.168.x.x or 10.0.0.x is private, they can’t be routable in the internet. Internal hosts can be assigned private addresses. Only your router needs public IP address. Router needs to translate requests.
- NAT Applications: Private to public -> So we don’t run out IPv4. Port forwarding -> Add a NAT entry in the router to forward packets to 80 to a machine in your LAN. No need to have root access to listen on port 80 on your device. Expose your local web server publically. Layer 4 Load balancing -> HAProxy NAT Mode - your load balancer is your gateway. Client send a request to a bogus IP. Router intercepts that packet and replaces the service IP with a destination server. Layer 4 reverse proxy.

### TCP Connection states:
- TCP is stateful protocol. Both client and server need to maintain all sorts of state. Window sizes, sequences and state of the connection. The connection goes trough many states.

### Pros and cons of TCP:
- Pros: Guarantee delivery. No one can send data without prior knowledge. Flow control and congestion control. Order packets no corruption or app level work. Secure and can’t be easily spoofed.
- Cons: Large header overhead compared to UDP. More bandwidth. Stateful - consumes memory on server and client. Considered high latency for certain workloads (slow start, congestion, acks). Does to much at a low level (This can cause TCP Head of line blocking, and this can also explain how to not do parallel request to a server, for example, is a bad idea). TCP meltdown (not a good candidate to VPN for example).

## Popular networking protocols:

### DNS:
- Every time that we use internet, and type a URL and accessing a website, we are using this protocol. The first thing that we do is query the DNS, our DNS resolvers and find out what is the IP address that is matching that domain, so that we can turn around and use TCP IP in order to establish TCP IP handshake to that IP address.
- Example > “www.husseinnasser.com” (www is subdomain, husseinnasser is the domain and com is the top level domain)
- We can add more that to DNS queries, like mail IP address, website mail address, any information about the domain itself, text information, service record, what port to connect.
- People can’t remember IPs and an IP can change while the domain remains. A domain is a text point to an IP or collection of IPs. Adds an additional layer of abstraction. We can serve the closest IP to a client requesting the same domain (cloudflare do that exactly). Load balancing and reverse proxy.
- A new addressing system means we need a mapping, DNS!. If you have an IP and you need a MAC, we use ARP. If you have the name and you need the IP, we use DNS. Built on top of UDP. Port 53. Support many records (MX, TXT, A, CNAME).
- How DNS works > First layer is a DNS resolver, which can have a lot of cache (like popular ones 1111 or 8888). If doesn’t contains this data there, this will be forwarded to ROOT server, which host the IP of TLD (top level domain, like .com) and forward again. On the top level domain server, it will host the IP of the ANS and forward. On the ANS (Authoritative name server), host the IP of the target server that you want to reach.
- Let’s say that we want to access “google.com” from our machine. 1: The computer ask to DNS resolver what is the IP of google.com. The resolver will check the cache and see if is there. In case of not, 2: The resolver ask to ROOT server where is the .com server -> 3: The root answer with the TLD server -> 4: Then the resolver ask ask to TLD server where is the authoritative name server of google.com -> 5: TLD server answer with the ANS server address -> 6: The resolver then ask to ANS server what is the IP of google.com -> 7: The ANS server answered with the IP -> 8: The resolver sends the response to the machine -> 9: The machine establish the TCP connection with the google server.
- Notice in some steps, we are not directly asking for “google.com” or things like that, but part of the queries, like funnelling the query until the right place. All this queries can be tracked by the transaction ID, which lives on the DNS header.
-  Why DNS was so many layers? To have a better performance. Image if everything lives in one place and not being distributed… this would be slower than have all this distribution (like TLS, ANS, ROOT servers) and we can apply this to the day by day. If we want to search one record in a billion rows, don’t search on billion rows, but less possible. This is why we have partitions, distributed systems, file systems and things like that.
- DNS is not encrypted by default. DoT / DoH attempts to address this. Many attacks against DNS (DNS hijacking/poisoning).
- nslookup: you can use options like “-type”, and the third parameter is a custom resolver that you want to query. dig can be used too.

### TLS:
- Most common way to encrypt/decrypt communication between two parts.
- Let’s take a common case: we want to do a GET / on the server on port 80. How this happens in some scenarios:
- “Vanilla HTTP”: We do the request to the server (with all layers, headers, packets and so on) and got a response with HTML. Everything that is the middle of this, like ISP, router, where all the IP packets will go through, it will see this requests.
- HTTPS: In the same case, we first will do a handshake, responsible to share an encryption key between the client and the server. After this, all the request/response will be encrypted and the port should be 443. So the client will send an encrypted GET request, the server will decrypt, process and send an encrypted response back.
- We encrypt the data with symmetric key algorithms. Symmetric key algorithms generally share the same key to encrypt and decrypt the data. Since both parts needs the same key, we need to exchange the symmetric key between the parts. Key exchange uses an asymmetric key (PKI). This also helps to auth the server. You can have extensions, like SNI, reshared key, 0RTT.
- TLS 1.2: The client will send a hello request to server. Server will generate a RSA public key (never share the private key) and send this shared key to the client. The client will generate a premaster symmetric key (encrypted) with the public from the server and it will send to server again (change cipher and fin). The server will use the RSA private key to decrypt this premaster key and compare then. And then the encryption can happen and the flow can continue. This changes on every session.
- We need to be aware that everything that is the middle of this communication or even libs can be a potential threat to our server, because if our private key is leaked, all the communication can be decrypted.
- Diffie hellman algorithm:  To build symmetric key with this algorithm, you will need a private key from server, a private key from client and the public key (the only which is shared).  If you combine the three together, you will get the symmetric key. It only can be generate by all parts, you can not merge two parts and the next one to get the symmetric key.
- TLS 1.3: The client will generate the private key and the public key. The client will send the private key encrypted and the public to the server. Then the server will generate a private key, combine with the 2 other keys and generate the main symmetric key. The server will encrypt their key and send back the key and the public one. Then the client will combine the three keys and generate the main symmetric key. Then the flow can continue again.

