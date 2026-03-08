253

Before transmitting, ip a coin that has probability p of coming up heads
If only one of the N computer's coins comes up heads, its transmission occurs successfully, and the
others must wait until that transmission is complete and then resume the algorithm.
If none or more than one head comes up, the N computers will either remain silent (no heads) or a
collision will occur (more than one head). This unsuccessful transmission situation will be detected
by all computers once the signals have propagated the length of the cable, and the algorithm resumes
(return to the beginning).

(a) What is the optimal probability to use for ipping the coin? In other words, what should p be to
maximize the probability that exactly one computer transmits?
(b) What is the probability of one computer transmitting when this optimal value of p is used as the
number of computers grows to in nity?
(c) Using this optimal probability, what is the average number of coin ips that will be necessary to resolve
the access so that one computer successfully transmits?
(d) Evaluate this algorithm. Is it realistic? Is it e cient?

Problem 6.32: Repeaters
Because signals attenuate with distance from the transmitter, repeaters are frequently employed for both
analog and digital communication. For example, let's assume that the transmitter and receiver are D m
apart, and a repeater is positioned halfway between them (Figure 6.42). What the repeater does is amplify
its received signal to exactly cancel the attenuation encountered along the rst leg and to re-transmit the
signal to the ultimate receiver. However, the signal the repeater receives contains white noise as well as the
transmitted signal. The receiver experiences the same amount of white noise as the repeater.

Figure 6.42

(a) What is the block diagram for this system?
(b) For an amplitude-modulation communication system, what is the signal-to-noise ratio of the demodu-
lated signal at the receiver? Is this better or worse than the signal-to-noise ratio when no repeater is
present?
(c) For digital communication, we must consider the system's capacity. Is the capacity larger with the
repeater system than without it? If so, when; if not, why not?

Problem 6.33: Designing a Speech Communication System
We want to examine both analog and digital communication alternatives for a dedicated speech transmission
system. Assume the speech signal has a 5 kHz bandwidth. The wireless link between transmitter and receiver
is such that 200 watts of power can be received at a pre-assigned carrier frequency. We have some latitude in
choosing the transmission bandwidth, but the noise power added by the channel increases with bandwidth
with a proportionality constant of 0.1 watt/kHz.

(a) Design an analog system for sending speech under this scenario. What is the received signal-to-noise
ratio under these design constraints?
(b) How many bits must be used in the A/D converter to achieve the same signal-to-noise ratio?
(c) Is the bandwidth required by the digital channel to send the samples without error greater or smaller
than the analog bandwidth?
