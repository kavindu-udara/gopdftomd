212 CHAPTER 6. INFORMATION COMMUNICATION

From our work in Fourier series, we know that this signal's spectrum contains odd-harmonics of the

fundamental, which here equals
1 T
. Thus, strictly speaking, the signal's bandwidth is in nite. In practical

terms, we use the 90%-power bandwidth to assess the e ective range of frequencies consumed by the signal.
The rst and third harmonics contain that fraction of the total power, meaning that the e ective bandwidth

of our baseband signal is
3 T
or, expressing this quantity in terms of the datarate,
3 R
. Thus, a digital

communications signal requires more bandwidth than the datarate: a 1 Mbps baseband system requires a
bandwidth of at least 1.5 MHz. Listen carefully when someone describes the transmission bandwidth of
digital communication systems: Did they say megabits or megahertz?

Exercise 6.15 (Solution on p. 256.)
Show that indeed the rst and third harmonics contain 90% of the transmitted power. If the receiver

uses a front-end lter of bandwidth
3 T
, what is the total harmonic distortion of the received signal?

Exercise 6.16 (Solution on p. 256.)
What is the 90% transmission bandwidth of the modulated signal set?

### 6.15 Frequency Shift Keying

In frequency-shift keying (FSK), the bit a ects the frequency of a carrier sinusoid.

s ( t ) = Ap ( t ) sin (2 f t )

s ( t ) = Ap ( t ) sin (2 f t )
(6.41)

Figure 6.11

The frequencies f , f are usually harmonically related to the bit interval. In the depicted example,

f =
3
and f =
4
. As can be seen from the transmitted signal for our example bit stream (Figure 6.12),

the transitions at bit interval boundaries are smoother than those of BPSK.
To determine the bandwidth required by this signal set, we again consider the alternating bit stream.
Think of it as two signals added together: The rst comprised of the signal s ( t ) , the zero signal, s ( t ) ,

Figure 6.12: This plot shows the FSK waveform for same bitstream used in the BPSK example (Fig-
ure 6.8).
