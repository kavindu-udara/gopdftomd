134 CHAPTER 4. FREQUENCY DOMAIN

the system linear and time-invariant? If so, nd the transfer function; if not, what linearity or time-
invariance criterion does reverberation violate.
(b) A music group known as the ROwls is having trouble selling its recordings. The record company's
engineer gets the idea of applying di erent delay to the low and high frequencies and adding the result
to create a new musical e ect. Thus, the ROwls' audio would be separated into two parts (one less
than the frequency f , the other greater than f ), these would be delayed by and respectively,
and the resulting signals added. Draw a block diagram for this new audio processing system, showing
its various components.
(c) How does the magnitude of the system's transfer function depend on the two delays?

Problem 4.17: Echoes in Telephone Systems
A frequently encountered problem in telephones is echo. Here, because of acoustic coupling between the ear
piece and microphone in the handset, what you hear is also sent to the person talking. That person thus
not only hears you, but also hears her own speech delayed (because of propagation delay over the telephone
network) and attenuated (the acoustic coupling gain is less than one). Furthermore, the same problem
applies to you as well: The acoustic coupling occurs in her handset as well as yours.

(a) Develop a block diagram that describes this situation.
(b) Find the transfer function between your voice and what the listener hears.
(c) Each telephone contains a system for reducing echoes using electrical means. What simple system
could null the echoes?

Problem 4.18: E ective Drug Delivery
In most patients, it takes time for the concentration of an administered drug to achieve a constant level
in the blood stream. Typically, if the drug concentration in the patient's intravenous line is C u( t ) , the
concentration in the patient's blood stream is C (1 e ) u( t ) .

(a) Assuming the relationship between drug concentration in the patient's drug and the delivered concen-
tration can be described as a linear, time-invariant system, what is the transfer function?
(b) Sometimes, the drug delivery system goes awry and delivers drugs with little control. What would the
patient's drug concentration be if the delivered concentration were a ramp? More precisely, if it were
C t u( t ) ?
(c) A clever doctor wants to have the exibility to slow down or speed up the patient's drug concentration.
In other words, the concentration is to be C 1 e u( t ) , with b bigger or smaller than a . How
should the delivered drug concentration signal be changed to achieve this concentration pro le?

Problem 4.19: Catching Speeders with Radar
RU Electronics has been contracted to design a Doppler radar system. Radar transmitters emit a signal that
bounces o any conducting object. Signal di erences between what is sent and the radar return is processed
and features of interest extracted. In Doppler systems, the object's speed along the direction of the radar
beam is the feature the design must extract. The transmitted signal is a sinusoid: x ( t ) = A cos (2 f t ) . The
measured return signal equals B cos 2 (( f + f) t + ' ) , where the Doppler o set frequency f equals
10 v , where v is the car's velocity coming toward the transmitter.

(a) Design a system that uses the transmitted and return signals as inputs and produces f .
(b) One problem with designs based on overly simplistic design goals is that they are sensitive to unmodeled
assumptions. How would you change your design, if at all, so that whether the car is going away or
toward the transmitter could be determined?
(c) Suppose two objects traveling di erent speeds provide returns. How would you change your design, if
at all, to accommodate multiple returns?

Problem 4.20: Demodulating an AM Signal
Let m ( t ) denote the signal that has been amplitude modulated.

x ( t ) = A 1 + m ( t ) sin (2 f t )
