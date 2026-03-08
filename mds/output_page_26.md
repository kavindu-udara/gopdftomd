20 CHAPTER 2. SIGNALS AND SYSTEMS

This derivation follows because the complex exponential evaluated at an integer multiple of 2 equals one.

2.4.4 Unit Sample

The second-most important discrete-time signal is the unit sample , which is de ned to be

( n ) =

(
1 if n = 0

0 otherwise
(2.28)

Figure 2.8: The unit sample.

Examination of a discrete-time signal's plot, like that of the cosine signal shown in Figure 2.7, reveals that
all discrete-time signals consist of a sequence of delayed and scaled unit samples. Because the value of
a sequence at each integer m is denoted by s ( m ) and the unit sample delayed to occur at m is written
( n m ) , we can decompose any signal as a sum of unit samples delayed to the appropriate location and
scaled by the signal value.

s ( n ) =
X
s ( m ) ( n m ) (2.29)

This kind of decomposition is unique to discrete-time signals, and will prove useful subsequently.

2.4.5 Symbolic-valued Signals

Another interesting aspect of discrete-time signals is that their values do not need to be real numbers. We
do have real-valued discrete-time signals like the sinusoid, but we also have signals that denote the sequence
of characters typed on the keyboard. Such characters certainly aren't real numbers, and as a collection of
possible signal values, they have little mathematical structure other than that they are members of a set.
More formally, each element of the symbolic-valued signal s ( n ) takes on one of the values f a ; : : : ; a g which
comprise the alphabet A . This technical terminology does not mean we restrict symbols to being mem-
bers of the English or Greek alphabet. They could represent keyboard characters, bytes (8-bit quantities),
integers that convey daily temperature. Whether controlled by software or not, discrete-time systems are
ultimately constructed from digital circuits, which consist entirely of analog circuit elements. Furthermore,
the transmission and reception of discrete-time signals, like e-mail, is accomplished with analog signals and
systems. Understanding how discrete-time and analog signals and systems intertwine is perhaps the main
goal of this course.

### 2.5 Introduction to Systems

Signals are manipulated by systems . Mathematically, we represent what a system does by the notation
y ( t ) = S [ x ( t )] , with x representing the input signal and y the output signal.
This content is available online at http://cnx.org/content/m0005/2.19/ .
