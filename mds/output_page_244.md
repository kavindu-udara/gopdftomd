238 CHAPTER 6. INFORMATION COMMUNICATION

Figure 6.29: The top circuit expresses a simpli ed circuit model for a transceiver. The output resistance
R must be much larger than Z so that the sum of the various transmitter voltages add to create the
Ethernet conductor-to-shield voltage that serves as the received signal r ( t ) for all transceivers. In this
case, the equivalent circuit shown in the bottom circuit applies.
ately cease transmission, waiting a random amount of time to attempt the transmission again (go to
step 1) until only one computer transmits and the others defer. The condition wherein two (or more)
computers' transmissions interfere with others is known as a collision .

The reason two computers waiting to transmit may not sense the other's transmission immediately arises
because of the nite propagation speed of voltage signals through the coaxial cable. The longest time any
computer must wait to determine if its transmissions do not encounter interference is
, where L is the
coaxial cable's length. The maximum-length-speci cation for Ethernet is 1 km. Assuming a propagation
speed of 2/3 the speed of light, this time interval is more than 10 s. As analyzed in Problem 6.31, the
number of these time intervals required to resolve the collision is, on the average, less than two!

Exercise 6.37 (Solution on p. 259.)
Why does the factor of two enter into this equation? Consider the worst-case situation of two
transmitting computers located at the Ethernet's ends.

Thus, despite not having separate communication paths among the computers to coordinate their transmis-
sions, the Ethernet random access protocol allows computers to communicate without only a slight degra-
dation in e ciency, as measured by the time taken to resolve collisions relative to the time the Ethernet is
used to transmit information.
A subtle consideration in Ethernet is the minimum packet size P . The time required to transmit such
packets equals
, where C is the Ethernet's capacity in bps. Ethernet now comes in two di erent types,
each with individual speci cations, the most distinguishing of which is capacity: 10 Mbps and 100 Mbps. If
the minimum transmission time is such that the beginning of the packet has not propagated the full length
of the Ethernet before the end-of-transmission, it is possible that two computers will begin transmission at
the same time and, by the time their transmissions cease, the other's packet will not have propagated to
the other. In this case, computers in-between the two will sense a collision, which renders both computer's
transmissions senseless to them, without the two transmitting computers knowing a collision has occurred at
all! For Ethernet to succeed, we must have the minimum packet transmission time exceed twice the voltage
