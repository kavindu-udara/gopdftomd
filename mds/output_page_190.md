184 CHAPTER 5. DIGITAL SIGNAL PROCESSING

(c) In particular, assume our lter is FIR, with the unit-sample response having duration q + 1 . If the
input has duration N , what is the duration of the lter's output to this signal?

(d) Let the lter be a boxcar averager: h ( n ) =
1 + 1
for n = f 0 ; : : : ; q g and zero otherwise. Let the input

be a pulse of unit height and duration N . Find the lter's output when N =
q + 1
, q an odd integer.

Problem 5.12: A Digital Filter
A digital lter has the unit-sample response depicted in Figure 5.30.

Figure 5.30

(a) What is the di erence equation that de nes this lter's input-output relationship?
(b) What is this lter's transfer function?

(c) What is the lter's output when the input is sin
n
?

Problem 5.13: A Special Discrete-Time Filter
Consider a FIR lter governed by the di erence equation

y ( n ) =
1
x ( n + 2) +
2
x ( n + 1) + x ( n ) +
2
x ( n 1) +
1
x ( n 2)

(a) Find this lter's unit-sample response.
(b) Find this lter's transfer function. Characterize this transfer function ( i.e. , what classic lter category
does it fall into).
(c) Suppose we take a sequence and stretch it out by a factor of three.

x ( n ) =

(
s
n
n = 3 m; m = f : : : ; 1 ; 0 ; 1 ; : : : g

0 otherwise

Sketch the sequence x ( n ) for some example s ( n ) . What is the lter's output to this input? In
particular, what is the output at the indices where the input x ( n ) is intentionally zero? Now how
would you characterize this system?

Problem 5.14: Simulating the Real World
Much of physics is governed by di erential equations, and we want to use signal processing methods to
simulate physical problems. The idea is to replace the derivative with a discrete-time approximation and
solve the resulting di erential equation. For example, suppose we have the di erential equation

d
y ( t ) + ay ( t ) = x ( t )
