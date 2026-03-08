151

converter, the device that converts integers to amplitudes, assigns an amplitude equal to the value lying
halfway in the quantization interval. The integer 6 would be assigned to the amplitude 0.625 in this scheme.
The error introduced by converting a signal from analog to digital form by sampling and amplitude quanti-
zation then back again would be half the quantization interval for each amplitude value. Thus, the so-called
A/D error equals half the width of a quantization interval:
. As we have xed the input-amplitude range,
the more bits available in the A/D converter, the smaller the quantization error.
To analyze the amplitude quantization error more deeply, we need to compute the signal-to-noise ratio,
which equals the ratio of the signal power and the quantization error power. Assuming the signal is a

sinusoid, the signal power is the square of the rms amplitude: power ( s ) =
=
. The illustration

(Figure 5.6) details a single quantization interval. Its width is and the quantization error is denoted by

Figure 5.6: A single quantization interval is shown, along with a typical signal's value before amplitude
quantization s ( nT ) and after Q ( s ( nT )) . denotes the error thus incurred.

. To nd the power in the quantization error, we note that no matter into which quantization interval the
signal's value falls, the error will have the same characteristics. To calculate the rms value, we must square
the error and average it over the interval.

rms ( ) =

s

Z

d

=

(5.10)

Since the quantization interval width for a B -bit converter equals
2
= 2 , we nd that the signal-to-

noise ratio for the analog-to-digital conversion process equals

SNR =

1

=
3
2 = 6 B + 10 log 1 : 5 dB (5.11)

Thus, every bit increase in the A/D converter yields a 6 dB increase in the signal-to-noise ratio. The constant
term 10 log 1 : 5 equals 1.76.

Exercise 5.8 (Solution on p. 191.)
This derivation assumed the signal's amplitude lay in the range [ 1 ; 1] . What would the amplitude
quantization signal-to-noise ratio be if it lay in the range [ A; A ] ?

Exercise 5.9 (Solution on p. 191.)
How many bits would be required in the A/D converter to ensure that the maximum amplitude
quantization error was less than 60 db smaller than the signal's peak value?

Exercise 5.10 (Solution on p. 191.)
Music on a CD is stored to 16-bit accuracy. To what signal-to-noise ratio does this correspond?
