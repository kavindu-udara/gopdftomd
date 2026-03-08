245

Figure 6.36

(a) Assume that the length of the direct path is d meters and the re ected path is 1.5 times as long. What
is the model for the channel, including the multipath and the additive noise?
(b) Assume d is 1 km. Find and sketch the magnitude of the transfer function for the multipath component
of the channel. How would you characterize this transfer function?
(c) Would the multipath a ect AM radio? If not, why not; if so, how so? Would analog cellular telephone,
which operates at much higher carrier frequencies (800 MHz vs. 1 MHz for radio), be a ected or not?
Analog cellular telephone uses amplitude modulation to transmit voice.
(d) How would the usual AM receiver be modi ed to minimize multipath e ects? Express your modi ed
receiver as a block diagram.

Problem 6.12: Downlink Signal Sets
In digital cellular telephone systems, the base station (transmitter) needs to relay di erent voice signals to
several telephones at the same time. Rather than send signals at di erent frequencies, a clever Rice engineer
suggests using a di erent signal set for each data stream. For example, for two simultaneous data streams,
she suggests BPSK signal sets that have the depicted basic signals (Figure 6.37).

Figure 6.37

Thus, bits are represented in data stream 1 by s ( t ) and s ( t ) and in data stream 2 by s ( t ) and s ( t ) ,
each of which are modulated by 900 MHz carrier. The transmitter sends the two data streams so that their
bit intervals align. Each receiver uses a matched lter for its receiver. The requirement is that each receiver
not receive the other's bit stream.

(a) What is the block diagram describing the proposed system?
(b) What is the transmission bandwidth required by the proposed system?
(c) Will the proposal work? Does the fact that the two data streams are transmitted in the same bandwidth
at the same time mean that each receiver's performance is a ected? Can each bit stream be received
without interference from the other?
