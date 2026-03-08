5

Figure 1.3: The Fundamental Model of Communication.

Figure 1.4: A system operates on its input signal x ( t ) to produce an output y ( t ) . source produces a signal that will be absorbed by the sink . Examples
of time-domain signals produced by a source are music, speech, and characters typed on a keyboard. Signals
can also be functions of two variables an image is a signal that depends on two spatial variables or more
television pictures (video signals) are functions of two spatial variables and time. Thus, information sources
produce signals. In physical systems, each signal corresponds to an electrical voltage or current .
To be able to design systems, we must understand electrical science and technology. However, we rst need
to understand the big picture to appreciate the context in which the electrical engineer works.
In communication systems, messages signals produced by sources must be recast for transmission.
The block diagram has the message s ( t ) passing through a block labeled transmitter that produces the
signal x ( t ) . In the case of a radio transmitter, it accepts an input audio signal and produces a signal that
physically is an electromagnetic wave radiated by an antenna and propagating as Maxwell's equations predict.
In the case of a computer network, typed characters are encapsulated in packets, attached with a destination
address, and launched into the Internet. From the communication systems big picture perspective, the
same block diagram applies although the systems can be very di erent. In any case, the transmitter should
not operate in such a way that the message s ( t ) cannot be recovered from x ( t ) . In the mathematical sense,
the inverse system must exist, else the communication system cannot be considered reliable. (It is ridiculous
to transmit a signal in such a way that no one can recover the original. However, clever systems exist that
transmit signals so that only the in crowd can recover them. Such cryptographic systems underlie secret
communications.)
Transmitted signals next pass through the next stage, the evil channel . Nothing good happens to a
signal in a channel: It can become corrupted by noise, distorted, and attenuated among many possibilities.
The channel cannot be escaped (the real world is cruel), and transmitter design and receiver design focus
on how best to jointly fend o the channel's e ects on signals. The channel is another system in our block
diagram, and produces r ( t ) , the signal received by the receiver. If the channel were benign (good luck
nding such a channel in the real world), the receiver would serve as the inverse system to the transmitter,
and yield the message with no distortion. However, because of the channel, the receiver must do its best to
produce a received message ^ s ( t ) that resembles s ( t ) as much as possible. Shannon showed in his 1948 paper
that reliable for the moment, take this word to mean error-free digital communication was possible over
arbitrarily noisy channels. It is this result that modern communications systems exploit, and why many
communications systems are going digital. The module on Chapter 6, titled Information Communication,
details Shannon's theory of information, and there we learn of Shannon's result and how to use it.
Finally, the received message is passed to the information sink that somehow makes use of the message.
http://www-gap.dcs.st-and.ac.uk/~history/Biographies/Shannon.html
