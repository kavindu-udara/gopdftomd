236 CHAPTER 6. INFORMATION COMMUNICATION

Figure 6.26: Long messages, such as les, are broken into separate packets, then transmitted over com-
puter networks. A packet, like a letter, contains the destination address, the return address (transmitter
address), and the data. The data includes the message part and a sequence number identifying its order
in the transmitted message.
and the network exit point must reassemble the original message accordingly.
Communications networks are now categorized according to whether they use packets or not. A system
like the telephone network is said to be circuit switched : The network establishes a xed route that lasts
the entire duration of the message. Circuit switching has the advantage that once the route is determined,
the users can use the capacity provided them however they like. Its main disadvantage is that the users
may not use their capacity e ciently, clogging network links and nodes along the way. Packet-switched
networks continuously monitor network utilization, and route messages accordingly. Thus, messages can, on
the average, be delivered e ciently, but the network cannot guarantee a speci c amount of capacity to the
users.

### 6.35 Network architectures and interconnection

The network structure its architecture (Figure 6.25) typi es what are known as wide area net-
works (WANs). The nodes, and users for that matter, are spread geographically over long distances. Long
has no precise de nition, and is intended to suggest that the communication links vary widely. The Internet
is certainly the largest WAN, spanning the entire earth and beyond. Local area networks , LANs, employ
a single communication link and special routing. Perhaps the best known LAN is Ethernet . LANs connect
to other LANs and to wide area networks through special nodes known as gateways (Figure 6.27). In the
Internet, a computer's address consists of a four byte sequence, which is known as its IP address (Inter-
net Protocol address). An example address is 128.42.4.32 : each byte is separated by a period. The rst
two bytes specify the computer's domain (here Rice University). Computers are also addressed by a more
human-readable form: a sequence of alphabetic abbreviations representing institution, type of institution,
and computer name. A given computer has both names ( 128.42.4.32 is the same as soma.rice.edu ).
Data transmission on the Internet requires the numerical form. So-called name servers translate between
alphabetic and numerical forms, and the transmitting computer requests this translation before the message
is sent to the network.

### 6.36 Ethernet

Ethernet uses as its communication medium a single length of coaxial cable (Figure 6.28). This cable serves
as the ether, through which all digital data travel. Electrically, computers interface to the coaxial cable
(Figure 6.28) through a device known as a transceiver . This device is capable of monitoring the voltage
appearing between the core conductor and the shield as well as applying a voltage to it. Conceptually it
consists of two op-amps, one applying a voltage corresponding to a bit stream (transmitting data) and another
This content is available online at http://cnx.org/content/m0077/2.10/ .
Ethernet http://cnx.org/content/m0078/latest/
This content is available online at http://cnx.org/content/m10284/2.13/ .
