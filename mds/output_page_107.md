101

coe cients depend on frequency is termed the signal's spectrum . When we plot the magnitude and phase
of the coe cients to display which harmonics are present in a periodic signal, we are plotting the spectrum.
Let's calculate the spectrum of the periodic pulse signal shown in Figure 4.1.

Figure 4.1: Periodic pulse signal.

The pulse width is , the period T , and the amplitude A . The spectrum of this signal is given by

c =
1

Z

Ae
dt =
A 2 k
e
1

At this point, simplifying this expression requires knowing an interesting property.

1 e = e e e = e 2 j sin

Armed with this result, we can simply express the Fourier series coe cients for our pulse sequence.

c = Ae

sin

(4.9)

Because this signal is real-valued, we nd that the coe cients do indeed have conjugate symmetry: c = c .
Because of this property, we do not need to plot the spectrum for negative k ; it can easily be found from the
spectrum for positive k . The magnitude of a real-valued signal's spectrum has even symmetry,
the angle has odd symmetry .
j c j = j c j \ c = \ c

The periodic pulse signal has neither even nor odd symmetry in the time domain; consequently, neither
properties 4.2 or 4.3 apply. To plot the spectrum, we need to calculate its magnitude and phase.

j c j = A
sin

(4.10)

\ c =
k
+ neg
sin

!

sign ( k )

The function neg ( ) equals 1 if its argument is negative and zero otherwise. The somewhat complicated
expression for the phase results because the sine term in (4.9) can be negative. Because magnitudes must
be positive, the occasional negative values of the sine function must be accounted for by a phase shift of .
