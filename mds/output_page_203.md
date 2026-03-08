197

Figure 6.1: Coaxial cable consists of one conductor wrapped around the central conductor. This type
of cable supports broader bandwidth signals than twisted pair, and nds use in cable television and
Ethernet.

Figure 6.2: The so-called distributed parameter model for two-wire cables has the depicted circuit
model structure. Element values depend on geometry and the properties of materials used to construct
the transmission line.
coaxial cable , where a concentric conductor surrounds a central wire with a dielectric material in between.
Coaxial cable, fondly called co-ax by engineers, is what Ethernet uses as its channel. In either case, wireline
channels form a dedicated circuit between transmitter and receiver. As we shall nd subsequently, several
transmissions can share the circuit by amplitude modulation techniques; commercial cable TV is an example.
These information-carrying circuits are designed so that interference from nearby electromagnetic sources is
minimized. Thus, by the time signals arrive at the receiver, they are relatively interference- and noise-free.
Both twisted pair and co-ax are examples of transmission lines , which all have the circuit model
shown in Figure 6.2 for an in nitesimally small length. This circuit model arises from solving Maxwell's
equations for the particular transmission line geometry. The series resistance comes from the conductor
used in the wires and from the conductor's geometry. The inductance and the capacitance derive from
transmission line geometry, and the parallel conductance from the medium between the wire pair. Note
that all the circuit elements have values expressed by the product of a constant times a length; this notation
represents that element values here have per-unit-length units. For example, the series resistance
e
R has units
of ohms/meter. For coaxial cable, the element values depend on the inner conductor's radius r , the outer
radius of the dielectric r , the conductivity of the conductors , and the conductivity , dielectric constant
, and magnetic permittivity of the dielectric as

e
R =
1

1
+
1
e
C =
2

e
G =
2

e
L =

ln
r

(6.2)

For twisted pair, having a separation d between the conductors that have conductivity and common radius
r and that are immersed in a medium having dielectric and magnetic properties, the element values are then
