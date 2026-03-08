244 CHAPTER 6. INFORMATION COMMUNICATION

Problem 6.9: A Novel Communication System
A clever system designer claims that the depicted transmitter (Figure 6.35) has, despite its complexity,
advantages over the usual amplitude modulation system. The message signal m ( t ) is bandlimited to W Hz,
and the carrier frequency ( f W ) . The channel attenuates the transmitted signal x ( t ) and adds white

noise of spectral height
N
.

Figure 6.35

The transfer function H ( f ) is given by H ( f ) =

(
j; if f < 0

j; if f > 0

(a) Find an expression for the spectrum of x ( t ) . Sketch your answer.
(b) Show that the usual coherent receiver demodulates this signal.
(c) Find the signal-to-noise ratio that results when this receiver is used.
(d) Find a superior receiver (one that yields a better signal-to-noise ratio), and analyze its performance.

Problem 6.10: Multi-Tone Digital Communication
In a so-called multi-tone system, several bits are gathered together and transmitted simultaneously on
di erent carrier frequencies during a T second interval. For example, B bits would be transmitted according
to

x ( t ) = A
X
b sin 2 ( k + 1) f t ; 0 t < T (6.64)

Here, f is the frequency o set for each bit and it is harmonically related to the bit interval T . The value of
b is either 1 or +1 .

(a) Find a receiver for this transmission scheme.
(b) An ELEC 241 alumni likes digital systems so much that he decides to produce a discrete-time version.
He samples the received signal (sampling interval T = T=N ). How should N be related to B , the
number of simultaneously transmitted bits?
(c) The alumni wants to nd a simple form for the receiver so that his software implementation runs as
e ciently as possible. How would you recommend he implement the receiver?

Problem 6.11: City Radio Channels
In addition to additive white noise, metropolitan cellular radio channels also contain multipath : the atten-
uated signal and a delayed, further attenuated signal are received superimposed. As shown in Figure 6.36,
multipath occurs because the buildings re ect the signal and the re ected path length between transmitter
and receiver is longer than the direct path.
