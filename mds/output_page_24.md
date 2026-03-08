18 CHAPTER 2. SIGNALS AND SYSTEMS

2.2.6 Square Wave

The square wave (Figure 2.6) sq ( t ) is a periodic signal like the sinusoid. It too has an amplitude and a
period, which must be speci ed to characterize the signal. We nd subsequently that the sine wave is a
simpler signal than the square wave.

Figure 2.6: The square wave.

### 2.3 Signal Decomposition

A signal's complexity is not related to how wiggly it is. Rather, a signal expert looks for ways of decomposing
a given signal into a sum of simpler signals , which we term the signal decomposition . Though we will
never compute a signal's complexity, it essentially equals the number of terms in its decomposition. In
writing a signal as a sum of component signals, we can change the component signal's gain by multiplying
it by a constant and by delaying it. More complicated decompositions could contain derivatives or integrals
of simple signals.

Example 2.2
As an example of signal complexity, we can express the pulse p ( t ) as a sum of delayed unit steps.

p ( t ) = u( t ) u( t ) (2.25)

Thus, the pulse is a more complex signal than the step. Be that as it may, the pulse is very useful
to us.

Exercise 2.4 (Solution on p. 30.)
Express a square wave having period T and amplitude A as a superposition of delayed and
amplitude-scaled pulses.

Because the sinusoid is a superposition of two complex exponentials, the sinusoid is more complex. We could
not prevent ourselves from the pun in this statement. Clearly, the word complex is used in two di erent
ways here. The complex exponential can also be written (using Euler's relation (2.16)) as a sum of a sine and
a cosine. We will discover that virtually every signal can be decomposed into a sum of complex exponentials,
and that this decomposition is very useful. Thus, the complex exponential is more fundamental, and Euler's
relation does not adequately reveal its complexity.

### 2.4 Discrete-Time Signals

So far, we have treated what are known as analog signals and systems. Mathematically, analog signals are
functions having continuous quantities as their independent variables, such as space and time. Discrete-time
signals are functions de ned on the integers; they are sequences. One of the fundamental results of signal
theory details the conditions under which an analog signal can be converted into a discrete-time one and
retrieved without error . This result is important because discrete-time signals can be manipulated by
This content is available online at http://cnx.org/content/m0008/2.12/ .
This content is available online at http://cnx.org/content/m0009/2.24/ .
