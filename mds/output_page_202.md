196 CHAPTER 6. INFORMATION COMMUNICATION

### 6.2 Types of Communication Channels

Electrical communications channels are either wireline or wireless channels. Wireline channels physically
connect transmitter to receiver with a wire which could be a twisted pair, coaxial cable or optic ber.
Consequently, wireline channels are more private and much less prone to interference. Simple wireline
channels connect a single transmitter to a single receiver: a point-to-point connection as with the telephone.
Listening in on a conversation requires that the wire be tapped and the voltage measured. Some wireline
channels operate in broadcast modes: one or more transmitter is connected to several receivers. One simple
example of this situation is cable television. Computer networks can be found that operate in point-to-point
or in broadcast modes. Wireless channels are much more public, with a transmitter's antenna radiating
a signal that can be received by any antenna su ciently close enough. In contrast to wireline channels
where the receiver takes in only the transmitter's signal, the receiver's antenna will react to electromagnetic
radiation coming from any source. This feature has two faces: The smiley face says that a receiver can take
in transmissions from any source, letting receiver electronics select wanted signals and disregarding others,
thereby allowing portable transmission and reception, while the frowny face says that interference and noise
are much more prevalent than in wireline situations. A noisier channel subject to interference compromises
the exibility of wireless communication.
Maxwell's equations neatly summarize the physics of all electromagnetic phenomena, including cir-
cuits, radio, and optic ber transmission.

r E =
@
( H ) div ( E ) =

r H = E +
@
( E ) div ( H ) = 0

(6.1)

where E is the electric eld, H the magnetic eld, dielectric permittivity, magnetic permeability,
electrical conductivity, and is the charge density. Kirchho 's Laws represent special cases of these equations
for circuits. We are not going to solve Maxwell's equations here; do bear in mind that a fundamental
understanding of communications channels ultimately depends on uency with Maxwell's equations. Perhaps
the most important aspect of them is that they are linear with respect to the electrical and magnetic elds.
Thus, the elds (and therefore the voltages and currents) resulting from two or more sources will add .

note: Nonlinear electromagnetic media do exist. The equations as written here are simpler
versions that apply to free-space propagation and conduction in metals. Nonlinear media are be-
coming increasingly important in optic ber communications, which are also governed by Maxwell's
equations.

### 6.3 Wireline Channels

Wireline channels were the rst used for electrical communications in the mid-nineteenth century for the
telegraph. Here, the channel is one of several wires connecting transmitter to receiver. The transmitter
simply creates a voltage related to the message signal and applies it to the wire(s). We must have a circuit
a closed path that supports current ow. In the case of single-wire communications, the earth is used as the
current's return path. In fact, the term ground for the reference node in circuits originated in single-wire
telegraphs. You can imagine that the earth's electrical characteristics are highly variable, and they are.
Single-wire metallic channels cannot support high-quality signal transmission having a bandwidth beyond a
few hundred Hertz over any appreciable distance.
Consequently, most wireline channels today essentially consist of pairs of conducting wires (see Figure 6.1),
and the transmitter applies a message-related voltage across the pair. How these pairs of wires are physically
con gured greatly a ects their transmission characteristics. One example is twisted pair , wherein the wires
This content is available online at http://cnx.org/content/m0099/2.13/ .
James Clerk Maxwell's biography can be found at http://www-groups.dcs.st-andrews.ac.uk/~history/Biographies/
Maxwell.html .
This content is available online at http://cnx.org/content/m0100/2.29/ .
