239

propagation time:
P
>
2 L
or

P >
2 LC
(6.63)

Thus, for the 10 Mbps Ethernet having a 1 km maximum length speci cation, the minimum packet size is
200 bits.

Exercise 6.38 (Solution on p. 259.)
The 100 Mbps Ethernet was designed more recently than the 10 Mbps alternative. To maintain
the same minimum packet size as the earlier, slower version, what should its length speci cation
be? Why should the minimum packet size remain the same?

### 6.37 Communication Protocols

The complexity of information transmission in a computer network reliable transmission of bits across
a channel, routing, and directing information to the correct destination within the destination computers
operating system demands an overarching concept of how to organize information delivery. No unique set of
rules satis es the various constraints communication channels and network organization place on information
transmission. For example, random access issues in Ethernet are not present in wide-area networks such as
the Internet. A protocol is a set of rules that governs how information is delivered. For example, to use
the telephone network, the protocol is to pick up the phone, listen for a dial tone, dial a number having a
speci c number of digits, wait for the phone to ring, and say hello. In radio, the station uses amplitude or
frequency modulation with a speci c carrier frequency and transmission bandwidth, and you know to turn on
the radio and tune in the station. In technical terms, no one protocol or set of protocols can be used for any
communication situation. Be that as it may, communication engineers have found that a common thread
runs through the organization of the various protocols. This grand design of information transmission
organization runs through all modern networks today.
What has been de ned as a networking standard is a layered, hierarchical protocol organization. As
shown in Figure 6.30, protocols are organized by function and level of detail. Segregation of information
transmission, manipulation, and interpretation into these categories directly a ects how communication
systems are organized, and what role(s) software systems ful ll. Although not thought about in this way
in earlier times, this organizational structure governs the way communication engineers think about all
communication systems, from radio to the Internet.

Exercise 6.39 (Solution on p. 259.)
How do the various aspects of establishing and maintaining a telephone conversation t into this
layered protocol organization?

We now explicitly state whether we are working in the physical layer (signal set design, for example), the
data link layer (source and channel coding), or any other layer. IP abbreviates Internet protocol, and governs
gateways (how information is transmitted between networks having di erent internal organizations). TCP
(transmission control protocol) governs how packets are transmitted through a wide-area network such as
the Internet. Telnet is a protocol that concerns how a person at one computer logs on to another computer
across a network. A moderately high level protocol such as telnet, is not concerned with what data links
(wireline or wireless) might have been used by the network or how packets are routed. Rather, it establishes
connections between computers and directs each byte (presumed to represent a typed character) to the
appropriate operation system component at each end. It is not concerned with what the characters mean
or what programs the person is typing to. That aspect of information transmission is left to protocols at
higher layers.
Recently, an important set of protocols created the World Wide Web. These protocols exist independently
of the Internet. The Internet insures that messages are transmitted e ciently and intact; the Internet is not
concerned (to date) with what messages contain. HTTP (hypertext transfer protocol) frame what messages
This content is available online at http://cnx.org/content/m0080/2.19/ .
