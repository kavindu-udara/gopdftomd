201

components at various frequencies propagate at the same speed. Transmitted signal amplitude does decay
exponentially along the transmission line. Note that in the high-frequency regime the space constant is small,
which means the signal attenuates little.

### 6.4 Wireless Channels

Wireless channels exploit the prediction made by Maxwell's equation that electromagnetic elds propagate
in free space like light. When a voltage is applied to an antenna, it creates an electromagnetic eld that
propagates in all directions (although antenna geometry a ects how much power ows in any given direction)
that induces electric currents in the receiver's antenna. Antenna geometry determines how energetic a eld
a voltage of a given frequency creates. In general terms, the dominant factor is the relation of the antenna's
size to the eld's wavelength. The fundamental equation relating frequency and wavelength for a propagating
wave is
f = c

Thus, wavelength and frequency are inversely related: High frequency corresponds to small wavelengths.
For example, a 1 MHz electromagnetic eld has a wavelength of 300 m. Antennas having a size or distance
from the ground comparable to the wavelength radiate elds most e ciently. Consequently, the lower the
frequency the bigger the antenna must be. Because most information signals are baseband signals, having
spectral energy at low frequencies, they must be modulated to higher frequencies to be transmitted over
wireless channels.
For most antenna-based wireless systems, how the signal diminishes as the receiver moves further from
the transmitter derives by considering how radiated power changes with distance from the transmitting
antenna. An antenna radiates a given amount of power into free space, and ideally this power propagates
without loss in all directions. Considering a sphere centered at the transmitter, the total power, which is
found by integrating the radiated power over the surface of the sphere, must be constant regardless of the
sphere's radius. This requirement results from the conservation of energy. Thus, if p ( d ) represents the power
integrated with respect to direction at a distance d from the antenna, the total power will be p ( d ) 4 d . For
this quantity to be a constant, we must have

p ( d ) /
1

which means that the received signal amplitude A must be proportional to the transmitter's amplitude A
and inversely related to distance from the transmitter.

A =
kA
(6.17)

for some value of the constant k . Thus, the further from the transmitter the receiver is located, the weaker
the received signal. Whereas the attenuation found in wireline channels can be controlled by physical
parameters and choice of transmission frequency, the inverse-distance attenuation found in wireless channels
persists across all frequencies.

Exercise 6.3 (Solution on p. 255.)
Why don't signals attenuate according to the inverse-square law in a conductor? What is the
di erence between the wireline and wireless cases?

The speed of propagation is governed by the dielectric constant and magnetic permeability of free
space.

c =
1

= 3 10 m/s

(6.18)
This content is available online at http://cnx.org/content/m0101/2.15/ .
