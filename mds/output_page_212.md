206 CHAPTER 6. INFORMATION COMMUNICATION

### 6.10 Baseband Communication

We use analog communication techniques for analog message signals, like music, speech, and television.
Transmission and reception of analog signals using analog results in an inherently noisy received signal
(assuming the channel adds noise, which it almost certainly does).
The simplest form of analog communication is baseband communication .

Point of Interest: We use analog communication techniques for analog message signals, like
music, speech, and television. Transmission and reception of analog signals using analog results
in an inherently noisy received signal (assuming the channel adds noise, which it almost certainly
does).

Here, the transmitted signal equals the message times a transmitter gain.

x ( t ) = Gm ( t ) (6.27)

An example, which is somewhat out of date, is the wireline telephone system. You don't use baseband
communication in wireless systems simply because low-frequency signals do not radiate well. The receiver in
a baseband system can't do much more than lter the received signal to remove out-of-band noise (interference
is small in wireline channels). Assuming the signal occupies a bandwidth of W Hz (the signal's spectrum
extends from zero to W ), the receiver applies a lowpass lter having the same bandwidth, as shown in
Figure 6.5.

Figure 6.5: The receiver for baseband communication systems is quite simple: a lowpass lter having
the same bandwidth as the signal.

We use the signal-to-noise ratio of the receiver's output b m ( t ) to evaluate any analog-message commu-
nication system. Assume that the channel introduces an attenuation and white noise of spectral height
N
. The lter does not a ect the signal component we assume its gain is unity but does lter the

noise, removing frequency components above W Hz. In the lter's output, the received signal power equals
G power ( m ) and the noise power N W , which gives a signal-to-noise ratio of

SNR =
G power ( m ) W
(6.28)

The signal term power ( m ) will be proportional to the bandwidth W ; thus, in baseband communication the
signal-to-noise ratio varies only with transmitter gain and channel attenuation and noise level.

### 6.11 Modulated Communication

Especially for wireless channels, like commercial radio and television, but also for wireline systems like cable
television, an analog message signal must be modulated : The transmitted signal's spectrum occurs at much
higher frequencies than those occupied by the signal.

Point of Interest: We use analog communication techniques for analog message signals, like
music, speech, and television. Transmission and reception of analog signals using analog results
in an inherently noisy received signal (assuming the channel adds noise, which it almost certainly
does).
This content is available online at http://cnx.org/content/m0517/2.19/ .
This content is available online at http://cnx.org/content/m0518/2.26/ .
