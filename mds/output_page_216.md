210 CHAPTER 6. INFORMATION COMMUNICATION

Figure 6.7

(a)

(b)

Figure 6.8: The upper plot shows how a baseband signal set for transmitting the bit sequence 0110 .
The lower one shows an amplitude-modulated variant suitable for wireless channels. (Solution on p. 256.)
What is the expression for the signal arising from a digital transmitter sending the bit stream b ( n ) ,
n = f : : : ; 1 ; 0 ; 1 ; : : : g using the signal set s ( t ) , s ( t ) , each signal of which has duration T ?

### 6.14 Binary Phase Shift Keying

A commonly used example of a signal set consists of pulses that are negatives of each other (Figure 6.7).

s ( t ) = Ap ( t )

s ( t ) = Ap ( t )
(6.38)

Here, we have a baseband signal set suitable for wireline transmission. The entire bit stream b ( n ) is
represented by a sequence of these signals. Mathematically, the transmitted signal has the form

x ( t ) =
X
( 1) Ap ( t nT ) (6.39)

and graphically Figure 6.8 shows what a typical transmitted signal might be.
This way of representing a bit stream changing the bit changes the sign of the transmitted signal is
known as binary phase shift keying and abbreviated BPSK. The name comes from concisely expressing
this popular way of communicating digital information. The word binary is clear enough (one binary-valued
quantity is transmitted during a bit interval). Changing the sign of sinusoid amounts to changing shifting
the phase by (although we don't have a sinusoid yet). The word keying re ects back to the rst electrical
communication system, which happened to be digital as well: the telegraph.
This content is available online at http://cnx.org/content/m10280/2.14/ .
