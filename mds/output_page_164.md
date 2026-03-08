158 CHAPTER 5. DIGITAL SIGNAL PROCESSING

The properties of the discrete-time Fourier transform mirror those of the analog Fourier transform. The
DTFT properties table shows similarities and di erences. One important common property is Parseval's
Theorem.
X
j s ( n ) j =

Z
S e df (5.31)

To show this important property, we simply substitute the Fourier transform expression into the frequency-
domain expression for power.

Z
j S e j df =

Z
X
s ( n ) e

!
X
s ( n ) e

!

df

=
X
s ( n ) s ( n )

Z
e df

(5.32)

Using the orthogonality relation (5.28), the integral equals ( m n ) , where ( n ) is the unit sample (Fig-
ure 5.8). Thus, the double sum collapses into a single sum because nonzero values occur only when n = m ,
giving Parseval's Theorem as a result. We term
P
s ( n ) the energy in the discrete-time signal s ( n ) in spite
of the fact that discrete-time signals don't consume (or produce for that matter) energy. This terminology
is a carry-over from the analog world.

Exercise 5.13 (Solution on p. 192.)
Suppose we obtained our discrete-time signal from values of the product s ( t ) p ( t ) , where the
duration of the component pulses in p ( t ) is . How is the discrete-time signal energy related to
the total energy contained in s ( t ) ? Assume the signal is bandlimited and that the sampling rate
was chosen appropriate to the Sampling Theorem's conditions.

The following table (5.1) summarizes the properties of the discrete-time Fourier transform. Note the
similarities and di erences with the corresponding properties of the Fourier transform of analog signals
(page 117). The notable di erences are the scaling properties that, for analog signals, are expressed as s ( at ) .
For discrete-time signals, the scaling factor a must be an integer and the resulting spectrum must be periodic
with period 1. For example, the spectrum of x ( n ) = s ( n= 2) , n even, x ( n ) = 0 , n odd, is derived as follows.

X
s ( n= 2) e =
X
s ( m ) e [ n = 2 m ]

= S e

Exercise 5.14 (Solution on p. 192)
Note that the spectrum of this stretched-signal example has a period equal to 1 = 2 and that the
spectrum is a compressed version of s ( n ) 's spectrum. Suppose that this stretched signal were
passed through a lowpass lter having a cuto frequency of 1 = 4 . What e ect would this ltering
have on the stretched signal? How would you interpret this e ect?

### 5.7 Discrete Fourier Transforms (DFT)

The discrete-time Fourier transform (and the continuous-time transform as well) can be evaluated when we
have an analytic expression for the signal. Suppose we just have a signal, such as the speech signal used in the
previous chapter, for which there is no formula. How then would you compute the spectrum? For example,
how did we compute a spectrogram such as the one shown in the speech signal example (Figure 4.17)? The
Discrete Fourier Transform allows the computation of spectra from discrete-time data. While in discrete-
time we can exactly calculate spectra, for analog signals no similar exact spectrum computation exists.
For analog-signal spectra, use must build special devices, which turn out in most cases to consist of A/D
Discrete-Time Fourier Transform Properties http<://cnx.org/content/m0506/latest/
This content is available online at http://cnx.org/content/m10249/2.27/ .
