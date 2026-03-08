# Chapter 4

# Frequency Domain

### 4.1 Introduction to the Frequency Domain

In developing ways of analyzing linear circuits, we invented the impedance method because it made solving
circuits easier. Along the way, we developed the notion of a circuit's frequency response or transfer function.
This notion, which also applies to all linear, time-invariant systems, describes how the circuit responds to a
sinusoidal input when we express it in terms of a complex exponential. We also learned the Superposition
Principle for linear systems: The system's output to an input consisting of a sum of two signals is the sum
of the system's outputs to each individual component.
The study of the frequency domain combines these two notions a system's sinusoidal response is easy to
nd and a linear system's output to a sum of inputs is the sum of the individual outputs to develop the
crucial idea of a signal's spectrum . We begin by nding that those signals that can be represented as a
sum of sinusoids is very large. In fact, all signals can be expressed as a superposition of sinusoids .
As this story unfolds, we'll see that information systems rely heavily on spectral ideas. For example,
radio, television, and cellular telephones transmit over di erent portions of the spectrum. In fact, spectrum
is so important that communications systems are regulated as to which portions of the spectrum they can use
by the Federal Communications Commission in the United States and by International Treaty for the world
(see Frequency Allocation Chart in Section 7.3). Calculating the spectrum is easy: The Fourier transform
de nes how we can nd a signal's spectrum.

### 4.2 Fourier Series

In an earlier module (Exercise 2.4), we showed that a square wave could be expressed as a superposition of
pulses. As useful as this decomposition was in this example, it does not generalize well to other periodic
signals: How can a superposition of pulses equal a smooth signal like a sinusoid? Because of the importance
of sinusoids to linear systems, you might wonder whether they could be added together to represent a large
number of periodic signals. You would be right and in good company as well. Leonhard Euler and Carl
Friedrich Gauss in particular worried about this problem, and Jean Baptiste Fourier got the credit even
though tough mathematical issues were not settled until later. They worked on what is now known as the
Fourier series : representing any periodic signal as a superposition of sinusoids.
But the Fourier series goes well beyond being another signal decomposition method. Rather, the Fourier
series begins our journey to appreciate how a signal can be described in either the time-domain or the
frequency-domain with no compromise. Let s ( t ) be a periodic signal with period T . We want to show
that periodic signals, even those that have constant-valued segments like a square wave, can be expressed
This content is available online at http://cnx.org/content/m0038/2.10/ .
This content is available online at http://cnx.org/content/m0042/2.28/ .
http://www-groups.dcs.st-and.ac.uk/~history/Biographies/Euler.html
http://www-groups.dcs.st-and.ac.uk/~history/Biographies/Gauss.html
http://www-groups.dcs.st-and.ac.uk/~history/Biographies/Fourier.html

97
