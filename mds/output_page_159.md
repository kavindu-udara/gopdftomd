153

5.5.3 Sinusoids

Discrete-time sinusoids have the obvious form s ( n ) = A cos (2 fn + ) . As opposed to analog complex
exponentials and sinusoids that can have their frequencies be any real value, frequencies of their discrete-
time counterparts yield unique waveforms only when f lies in the interval
;
. This choice of frequency
interval is arbitrary; we can also choose the frequency to lie in the interval [0 ; 1) . How to choose a unit-length
interval for a sinusoid's frequency will become evident later.

5.5.4 Unit Sample

The second-most important discrete-time signal is the unit sample , which is de ned to be

( n ) =

(
1 n = 0

0 otherwise
(5.14)

Figure 5.8: The unit sample.

Examination of a discrete-time signal's plot, like that of the cosine signal shown in Figure 5.7, reveals that
all signals consist of a sequence of delayed and scaled unit samples. Because the value of a sequence at
each integer m is denoted by s ( m ) and the unit sample delayed to occur at m is written ( n m ) , we can
decompose any signal as a sum of unit samples delayed to the appropriate location and scaled by the signal
value.

s ( n ) =
X
s ( m ) ( n m ) (5.15)

This kind of decomposition is unique to discrete-time signals, and will prove useful subsequently.

5.5.5 Unit Step

The unit step in discrete-time is well-de ned at the origin, as opposed to the situation with analog signals.

u( n ) =

(
1 n 0

0 n < 0
(5.16)

5.5.6 Symbolic Signals

An interesting aspect of discrete-time signals is that their values do not need to be real numbers. We do
have real-valued discrete-time signals like the sinusoid, but we also have signals that denote the sequence of
characters typed on the keyboard. Such characters certainly aren't real numbers, and as a collection of pos-
sible signal values, they have little mathematical structure other than that they are members of a set. More
formally, each element of the symbolic-valued signal s ( n ) takes on one of the values f a ; : : : ; a g which
comprise the alphabet A . This technical terminology does not mean we restrict symbols to being mem-
bers of the English or Greek alphabet. They could represent keyboard characters, bytes (8-bit quantities),
integers that convey daily temperature. Whether controlled by software or not, discrete-time systems are
ultimately constructed from digital circuits, which consist entirely of analog circuit elements. Furthermore,
the transmission and reception of discrete-time signals, like e-mail, is accomplished with analog signals and
systems. Understanding how discrete-time and analog signals and systems intertwine is perhaps the main
goal of this course.
