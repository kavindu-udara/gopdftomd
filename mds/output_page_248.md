242 CHAPTER 6. INFORMATION COMMUNICATION

(b) Sketch a block diagram for a communication system (transmitter and receiver) that employs comple-
mentary signal transmission to send a message m ( t ) .
(c) What is the receiver's signal-to-noise ratio? How does it compare to the standard system that sends
the signal by simple amplitude modulation?

Problem 6.4: Phase Modulation
A message signal m ( t ) phase modulates a carrier if the transmitted signal equals

x ( t ) = A sin 2 f t + m ( t )

where is known as the phase deviation. In this problem, the phase deviation is small. As with all analog
modulation schemes, assume that j m ( t ) j < 1 , the message is bandlimited to W Hz, and the carrier frequency
f is much larger than W .

(a) What is the transmission bandwidth?
(b) Find a receiver for this modulation scheme.
(c) What is the signal-to-noise ratio of the received signal?

Hint: Use the facts that cos ( x ) 1 and sin ( x ) x for small x .

Problem 6.5: Digital Amplitude Modulation
Two ELEC 241 students disagree about a homework problem. The issue concerns the discrete-time signal
s ( n ) cos (2 f n ) , where the signal s ( n ) has no special characteristics and the modulation frequency f is
known. Sammy says that he can recover s ( n ) from its amplitude-modulated version by the same approach
used in analog communications. Samantha says that approach won't work.

(a) What is the spectrum of the modulated signal?
(b) Who is correct? Why?
(c) The teaching assistant does not want to take sides. He tells them that if s ( n ) cos (2 f n ) and
s ( n ) sin (2 f n ) were both available, s ( n ) can be recovered. What does he have in mind?

Problem 6.6: Anti-Jamming
One way for someone to keep people from receiving an AM transmission is to transmit noise at the same
carrier frequency. Thus, if the carrier frequency is f so that the transmitted signal is A 1+ m ( t ) sin (2 f t )
the jammer would transmit A n ( t ) sin (2 f t + ) . The noise n ( t ) has a constant power density spectrum

over the bandwidth of the message m ( t ) . The channel adds white noise of spectral height
N
.

(a) What would be the output of a traditional AM receiver tuned to the carrier frequency f ?
(b) RU Electronics proposes to counteract jamming by using a di erent modulation scheme. The scheme's
transmitted signal has the form A (1 + m ( t )) c ( t ) where c ( t ) is a periodic carrier signal (period
)

having the indicated waveform (Figure 6.33). What is the spectrum of the transmitted signal with
the proposed scheme? Assume the message bandwidth W is much less than the fundamental carrier
frequency f .
(c) The jammer, unaware of the change, is transmitting with a carrier frequency of f , while the receiver
tunes a standard AM receiver to a harmonic of the carrier frequency. What is the signal-to-noise ratio
of the receiver tuned to the harmonic having the largest power that does not contain the jammer?
