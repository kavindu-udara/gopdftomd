# Chapter 7

# Appendix

### 7.1 Decibels

The decibel scale expresses amplitudes and power values logarithmically . The de nitions for these di er,
but are consistent with each other.

power ( s; in decibels ) = 10 log
power ( s ) s )

amplitude ( s; in decibels ) = 20 log
amplitude ( s ) s )

(7.1)

Here power ( s ) and amplitude ( s ) represent a reference power and amplitude, respectively. Quantifying
power or amplitude in decibels essentially means that we are comparing quantities to a standard or that we
want to express how they changed. You will hear statements like The signal went down by 3 dB and The
lter's gain in the stopband is 60 (Decibels is abbreviated dB.).

Exercise 7.1 (Solution on p. 265.)
The pre x deci implies a tenth; a decibel is a tenth of a Bel. Who is this measure named for?

The consistency of these two de nitions arises because power is proportional to the square of amplitude:

power ( s ) / amplitude ( s ) (7.2)

Plugging this expression into the de nition for decibels, we nd that

10 log
power ( s ) s )
= 10 log
amplitude ( s ) ( s )

= 20 log
amplitude ( s ) s )

(7.3)

Because of this consistency, stating relative change in terms of decibels is unambiguous . A factor
of 10 increase in amplitude corresponds to a 20 dB increase in both amplitude and power!
The accompanying table provides nice decibel values. Converting decibel values back and forth is fun,
and tests your ability to think of decibel values as sums and/or di erences of the well-known values and of
ratios as products and/or quotients. This conversion rests on the logarithmic nature of the decibel scale.
For example, to nd the decibel value for
p , we halve the decibel value for 2 ; 26 dB equals 10 + 10 + 6 dB
that corresponds to a ratio of 10 10 4 = 400 . Decibel quantities add; ratio values multiply.
One reason decibels are used so much is the frequency-domain input-output relation for linear systems:
Y ( f ) = X ( f ) H ( f ) . Because the transfer function multiplies the input signal's spectrum, to nd the output
This content is available online at http://cnx.org/content/m0082/2.16/ .

261
