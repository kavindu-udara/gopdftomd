259

For N = 7 , p < 0 : 31 .
Solution to Exercise 6.33 (p. 230)

In a length- N block, N single-bit and
double-bit errors can occur. The number of non-zero vectors
resulting from H b c must equal or exceed the sum of these two numbers.

2 1 N +
N ( N 1)
or 2
N + N + 2
(6.68)

The rst two solutions that attain equality are (5,1) and (90,78) codes. However, no perfect code exists
other than the single-bit error correcting Hamming code. (Perfect codes satisfy relations like (6.68) with
equality.)
Solution to Exercise 6.34 (p. 233)
To convert to bits/second, we divide the capacity stated in bits/transmission by the bit interval duration T .
Solution to Exercise 6.35 (p. 234)
The network entry point is the telephone handset, which connects you to the nearest station. Dialing the
telephone number informs the network of who will be the message recipient. The telephone system forms an
electrical circuit between your handset and your friend's handset. Your friend receives the message via the
same device the handset that served as the network entry point.
Solution to Exercise 6.36 (p. 237)

The transmitting op-amp sees a load given by R + Z k
R
, where N is the number of transceivers

other than this one attached to the coaxial cable. The transfer function to some other transceiver's receiver
circuit is R divided by this load.
Solution to Exercise 6.37 (p. 238)
The worst-case situation occurs when one computer begins to transmit just before the other's packet arrives.
Transmitters must sense a collision before packet transmission ends. The time taken for one computer's
packet to travel the Ethernet's length and for the other computer's transmission to arrive equals the round-
trip, not one-way, propagation time.
Solution to Exercise 6.38 (p. 239)
The cable must be a factor of ten shorter: It cannot exceed 100 m. Di erent minimum packet sizes means
di erent packet formats, making connecting old and new systems together more complex than need be.
Solution to Exercise 6.39 (p. 239)
When you pick up the telephone, you initiate a dialog with your network interface by dialing the number.
The network looks up where the destination corresponding to that number is located, and routes the call
accordingly. The route remains xed as long as the call persists. What you say amounts to high-level
protocol while establishing the connection and maintaining it corresponds to low-level protocol.
