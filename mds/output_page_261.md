255

### Solutions to Exercises in Chapter 6

Solution to Exercise 6.1 (p. 200)
In both cases, the answer depends less on geometry than on material properties. For coaxial cable, c =

1
. For twisted pair, c =
1

s

+ cosh

.

Solution to Exercise 6.2 (p. 200)
You can nd these frequencies from the spectrum allocation chart (Section 7.3). Light in the middle of the
visible band has a wavelength of about 600 nm, which corresponds to a frequency of 5 10 Hz . Cable
television transmits within the same frequency band as broadcast television (about 200 MHz or 2 10 Hz ).
Thus, the visible electromagnetic frequencies are over six orders of magnitude higher!
Solution to Exercise 6.3 (p. 201)
As shown previously (6.11), voltages and currents in a wireline channel, which is modeled as a transmission
line having resistance, capacitance and inductance, decay exponentially with distance. The inverse-square law
governs free-space propagation because such propagation is lossless, with the inverse-square law a consequence
of the conservation of power. The exponential decay of wireline channels occurs because they have losses
and some ltering.
Solution to Exercise 6.4 (p. 202)

Figure 6.43

Use the Pythagorean Theorem, ( h + R ) = R + d , where h is the antenna height, d is the distance from
the top of the earth to a tangency point with the earth's surface, and R the earth's radius. The line-of-sight
distance between two earth-based antennae equals

d =

q h R + h +

q h R + h (6.66)

As the earth's radius is much larger than the antenna height, we have to a good approximation that d =
p h R +
p h R . If one antenna is at ground elevation, say h = 0 , the other antenna's range is
p h R .
Solution to Exercise 6.5 (p. 202)
As frequency decreases, wavelength increases and can approach the distance between the earth's surface and
the ionosphere. Assuming a distance between the two of 80 km, the relation f = c gives a corresponding
frequency of 3.75 kHz. Such low carrier frequencies would be limited to low bandwidth analog communication
and to low datarate digital communications. The US Navy did use such a communication scheme to reach
all of its submarines at once.
Solution to Exercise 6.6 (p. 203)
Transmission to the satellite, known as the uplink, encounters inverse-square law power losses. Re ecting
o the ionosphere not only encounters the same loss, but twice. Re ection is the same as transmitting
exactly what arrives, which means that the total loss is the product of the uplink and downlink losses. The
geosynchronous orbit lies at an altitude of 35700 km . The ionosphere begins at an altitude of about 50 km.
The amplitude loss in the satellite case is proportional to 2 : 8 10 ; for Marconi, it was proportional to
4 : 4 10 . Marconi was very lucky.
Solution to Exercise 6.7 (p. 204)
If the interferer's spectrum does not overlap that of our communications channel the interferer is out-of-
band we need only use a bandpass lter that selects our transmission band and removes other portions of
the spectrum.
