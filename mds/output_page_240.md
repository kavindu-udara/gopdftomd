234 CHAPTER 6. INFORMATION COMMUNICATION

communication's ability to transmit a wider variety of signals than analog systems, point-to-point
digital systems can be organized into global (and beyond as well) systems that provide e cient and
exible information transmission. Computer networks , explored in the next section, are what we
call such systems today. Even analog-based networks, such as the telephone system, employ modern
computer networking ideas rather than the purely analog systems of the past.

Consequently, with the increased speed of digital computers, the development of increasingly e cient algo-
rithms, and the ability to interconnect computers to form a communications infrastructure, digital commu-
nication is now the best choice for many situations.

### 6.33 Communication Networks

Communication networks elaborate the Fundamental Model of Communications (Figure 1.3). The model
shown in Figure 6.25 describes point-to-point communications well, wherein the link between transmitter
and receiver is straightforward, and they have the channel to themselves. One modern example of this
communications mode is the modem that connects a personal computer with an information server via
a telephone line. The key aspect, some would say aw, of this model is that the channel is dedicated :
Only one communications link through the channel is allowed for all time. Regardless whether we have
a wireline or wireless channel, communication bandwidth is precious, and if it could be shared without
signi cant degradation in communications performance (measured by signal-to-noise ratio for analog signal
transmission and by bit-error probability for digital transmission) so much the better.

Figure 6.25: The prototypical communications network whether it be the postal service, cellular
telephone, or the Internet consists of nodes interconnected by links. Messages formed by the source are
transmitted within the network by dynamic routing. Two routes are shown. The longer one would be
used if the direct link were disabled or congested.

The idea of a network rst emerged with perhaps the oldest form of organized communication: the postal
service. Most communication networks, even modern ones, share many of its aspects.

A user writes a letter, serving in the communications context as the message source.
This message is sent to the network by delivery to one of the network's public entry points. Entry
points in the postal case are mailboxes, post o ces, or your friendly mailman or mailwoman picking
up the letter.
The communications network delivers the message in the most e cient (timely) way possible, trying
not to corrupt the message while doing so.
The message arrives at one of the network's exit points, and is delivered to the recipient (what we have
termed the message sink).

Exercise 6.35 (Solution on p. 259.)
Develop the network model for the telephone system, making it as analogous as possible with the
postal service-communications network metaphor.

What is most interesting about the network system is the ambivalence of the message source and sink about
how the communications link is made. What they do care about is message integrity and communications
This content is available online at http://cnx.org/content/m0075/2.11/ .
