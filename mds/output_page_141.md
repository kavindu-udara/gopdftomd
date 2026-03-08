135

(a)

(b)

Figure 4.29 m ( t ) so that it is less than one in magnitude. The
frequency f is very large compared to the frequency content of the signal. What we are concerned about
here is not transmission, but reception.

(a) The so-called coherent demodulator simply multiplies the signal x ( t ) by a sinusoid having the same
frequency as the carrier and lowpass lters the result. Analyze this receiver and show that it works.
Assume the lowpass lter is ideal.
(b) One issue in coherent reception is the phase of the sinusoid used by the receiver relative to that used
by the transmitter. Assuming that the sinusoid of the receiver has a phase , how does the output
depend on ? What is the worst possible value for this phase?
(c) The incoherent receiver is more commonly used because of the phase sensitivity problem inherent in
coherent reception. Here, the receiver full-wave recti es the received signal and lowpass lters the
result (again ideally). Analyze this receiver. Does its output di er from that of the coherent receiver
in a signi cant way?

Problem 4.21: Unusual Amplitude Modulation
We want to send a band-limited signal having the depicted spectrum (Figure 4.29(a)) using amplitude
modulation. I.B. Di erent suggests using the square-wave carrier shown below (Figure 4.29(b)). Well, it is
di erent, but his friends wonder if any technique can demodulate it.

(a) Find an expression for X ( f ) , the Fourier transform of the modulated signal.
(b) Sketch the magnitude of X ( f ) , being careful to label important magnitudes and frequencies.
(c) What demodulation technique obviously works?
(d) I.B. challenges three of his friends to demodulate x ( t ) some other way. One friend suggests modulating
x ( t ) with cos
, another wants to try modulating with cos ( t ) and the third thinks cos
will
work. Sketch the magnitude of the Fourier transform of the signal each student's approach produces.
Which student comes closest to recovering the original signal? Why?

Problem 4.22: Sammy Falls Asleep...
While sitting in ELEC 241 class, he falls asleep during a critical time when an AM receiver is being described.
The received signal has the form r ( t ) = A 1 + m ( t ) cos (2 f t + ) where the phase is unknown. The
message signal is m ( t ) ; it has a bandwidth of W Hz and a magnitude less than 1 ( j m ( t ) j < 1 ). The phase
is unknown. The instructor drew a diagram (Figure 4.30) for a receiver on the board; Sammy slept through
the description of what the unknown systems where.

(a) What are the signals x ( t ) and x ( t ) ?
(b) What would you put in for the unknown systems that would guarantee that the nal output contained
the message regardless of the phase?

Hint: Think of a trigonometric identity that would prove useful.

(c) Sammy may have been asleep, but he can think of a far simpler receiver. What is it?
