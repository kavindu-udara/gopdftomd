205

### 6.9 Channel Models

Both wireline and wireless channels share characteristics, allowing us to use a common model for how the
channel a ects transmitted signals.

The transmitted signal is usually not ltered by the channel.
The signal can be attenuated.
The signal propagates through the channel at a speed equal to or less than the speed of light, which
means that the channel delays the transmission.
The channel may introduce additive interference and/or noise.

Letting represent the attenuation introduced by the channel, the receiver's input signal is related to the
transmitted one by
r ( t ) = x ( t ) + i ( t ) + n ( t ) (6.24)

This expression corresponds to the system model for the channel shown in Figure 6.4. In this book, we shall
assume that the noise is white.

Figure 6.4: The channel component of the fundamental model of communication (Figure 1.3) has the
depicted form. The attenuation is due to propagation loss. Adding the interference and noise is justi ed
by the linearity property of Maxwell's equations.

Exercise 6.8 (Solution on p. 255.)
Is this model for the channel linear?

As expected, the signal that emerges from the channel is corrupted, but does contain the transmitted signal.
Communication system design begins with detailing the channel model, then developing the transmitter and
receiver that best compensate for the channel's corrupting behavior. We characterize the channel's quality
by the signal-to-interference ratio ( SIR ) and the signal-to-noise ratio ( SNR ). The ratios are computed
according to the relative power of each within the transmitted signal's bandwidth . Assuming the
signal x ( t ) 's spectrum spans the frequency interval [ f ; f ] , these ratios can be expressed in terms of power
spectra.

SIR =

2

Z

P ( f ) df

Z

P ( f ) df

(6.25)

SNR =

2

Z

P ( f ) df ( f f )
(6.26)

In most cases, the interference and noise powers do not vary for a given receiver. Variations in signal-to-
interference and signal-to-noise ratios arise from the attenuation because of transmitter-to-receiver distance
variations.
This content is available online at http://cnx.org/content/m0516/2.11/ .
