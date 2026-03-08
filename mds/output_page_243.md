237

Figure 6.27: The gateway serves as an interface between local area networks and the Internet. The two
shown here translate between LAN and WAN protocols; one of these also interfaces between two LANs,
presumably because together the two LANs would be geographically too dispersed.

Figure 6.28: The Ethernet architecture consists of a single coaxial cable terminated at either end by a
resistor having a value equal to the cable's characteristic impedance. Computers attach to the Ethernet
through an interface known as a transceiver because it sends as well as receives bit streams represented
as analog voltages.
that shown in BPSK Signal Sets, with one signal the negative of the other. Computers are attached in
parallel, resulting in the circuit model for Ethernet shown in Figure 6.29.

Exercise 6.36 (Solution on p. 259.)
From the viewpoint of a transceiver's sending op-amp, what is the load it sees and what is the
transfer function between this output voltage and some other transceiver's receiving circuit? Why
should the output resistor R be large?

No one computer has more authority than any other to control when and how messages are sent. Without
scheduling authority, you might well wonder how one computer sends to another without the (large) inter-
ference that the other computers would produce if they transmitted at the same time. The innovation of
Ethernet is that computers schedule themselves by a random-access method. This method relies on the
fact that all packets transmitted over the coaxial cable can be received by all transceivers, regardless of
which computer might actually be the intended recipient. In communications terminology, Ethernet directly
supports broadcast. Each computer goes through the following steps to send a packet.

1. The computer senses the voltage across the cable to determine if some other computer is transmitting.
2. If another computer is transmitting, wait until the transmissions nish and go back to the rst step.
If the cable has no transmissions, begin transmitting the packet.
3. If the receiver portion of the transceiver determines that no other computer is also sending a packet,
continue transmitting the packet until completion.
