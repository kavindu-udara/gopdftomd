235

e ciency. Furthermore, today's networks use heterogeneous links. Communication paths that form the
Internet use wireline, optical ber, and satellite communication links.
The rst electrical communications network was the telegraph. Here the network consisted of telegraph
operators who transmitted the message e ciently using Morse code and routed the message so that it took
the shortest possible path to its destination while taking into account internal network failures (downed lines,
drunken operators). From today's perspective, the fact that this nineteenth century system handled digital
communications is astounding. Morse code, which assigned a sequence of dots and dashes to each letter of
the alphabet, served as the source coding algorithm. The signal set consisted of a short and a long pulse.
Rather than a matched lter, the receiver was the operator's ear, and he wrote the message (translating
from received bits to symbols).

Note: Because of the need for a comma between dot-dash sequences to de ne letter (symbol)
boundaries, the average number of bits/symbol, as described in Subtleties of Coding (Example 6.4),
exceeded the Source Coding Theorem's upper bound.

Internally, communication networks do have point-to-point communication links between network nodes
well described by the Fundamental Model of Communications. However, many messages share the com-
munications channel between nodes using what we call time-domain multiplexing : Rather than the
continuous communications mode implied in the Model as presented, message sequences are sent, sharing
in time the channel's capacity. At a grander viewpoint, the network must route messages decide what
nodes and links to use based on destination information the address that is usually separate from the
message information. Routing in networks is necessarily dynamic: The complete route taken by messages is
formed as the network handles the message, with nodes relaying the message having some notion of the best
possible path at the time of transmission. Note that no omnipotent router views the network as a whole and
pre-determines every message's route. Certainly in the case of the postal system dynamic routing occurs,
and can consider issues like inoperative and overly busy links. In the telephone system, routing takes place
when you place the call; the route is xed once the phone starts ringing. Modern communication networks
strive to achieve the most e cient (timely) and most reliable information delivery system possible.

### 6.34 Message Routing

Focusing on electrical networks, most analog ones make ine cient use of communication links because truly
dynamic routing is di cult, if not impossible, to obtain. In radio networks, such as commercial television,
each station has a dedicated portion of the electromagnetic spectrum, and this spectrum cannot be shared
with other stations or used in any other than the regulated way. The telephone network is more dynamic,
but once it establishes a call the path through the network is xed. The users of that path control its
use, and may not make e cient use of it (long pauses while one person thinks, for example). Telephone
network customers would be quite upset if the telephone company momentarily disconnected the path so
that someone else could use it. This kind of connection through a network xed for the duration of the
communication session is known as a circuit-switched connection.
During the 1960s, it was becoming clear that not only was digital communication technically superior,
but also that the wide variety of communication modes computer login, le transfer, and electronic mail
needed a di erent approach than point-to-point. The notion of computer networks was born then, and what
was then called the ARPANET, now called the Internet, was born. Computer networks elaborate the basic
network model by subdividing messages into smaller chunks called packets (Figure 6.26). The rationale for
the network enforcing smaller transmissions was that large le transfers would consume network resources
all along the route, and, because of the long transmission time, a communication failure might require
retransmission of the entire le. By creating packets, each of which has its own address and is routed
independently of others, the network can better manage congestion. The analogy is that the postal service,
rather than sending a long letter in the envelope you provide, opens the envelope, places each page in a
separate envelope, and using the address on your envelope, addresses each page's envelope accordingly, and
This content is available online at http://cnx.org/content/m0076/2.9/ .
