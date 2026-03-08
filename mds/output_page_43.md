37

on resistor values. The KVL equation can be rewritten as v = v + v . Substituting into it the resistor's
v-i relation, we have v = R i + R i . Yes, we temporarily eliminate the quantity we seek. Though
not obvious, it is the simplest way to solve the equations. One of the KCL equations says i = i , which
means that v = R i + R i = ( R + R ) i . Solving for the current in the output resistor, we have

i =
v + R
. We have now solved the circuit : We have expressed one voltage or current in terms of

sources and circuit-element values. To nd any other circuit quantities, we can back substitute this answer
into our original equations or ones we developed along the way. Using the v-i relation for the output resistor,
we obtain the quantity we seek.

v =
R + R
v

Exercise 3.3 (Solution on p. 94.)
Referring back to Figure 3.6 (page 35), a circuit should serve some useful purpose. What kind of
system does our circuit realize and, in terms of element values, what are the system's parameter(s)?

### 3.5 Power Dissipation in Resistor Circuits

We can nd voltages and currents in simple circuits containing resistors and voltage or current sources. We
should examine whether these circuits variables obey the Conservation of Power principle: since a circuit is
a closed system, it should not dissipate or create energy. For the moment, our approach is to investigate
rst a resistor circuit's power consumption/creation. Later, we will prove that because of KVL and KCL
all circuits conserve power.
As de ned on p. 32, the instantaneous power consumed/created by every circuit element equals the
product of its voltage and current. The total power consumed/created by a circuit equals the sum of each
element's power.

P =
X
v i

Recall that each element's current and voltage must obey the convention that positive current is de ned to
enter the positive-voltage terminal. With this convention, a positive value of v i corresponds to consumed
power, a negative value to created power. Because the total power in a circuit must be zero ( P = 0 ), some
circuit elements must create power while others consume it.
Consider the simple series circuit appearing in Figure 3.6 (page 35). In performing our calculations,
we de ned the current i to ow through the positive-voltage terminals of both resistors and found it to

equal i =
v + R
. The voltage across the resistor R is the output voltage and we found it to equal

v =
R + R
v . Consequently, calculating the power for this resistor yields

P =
R R + R )
v

Consequently, this resistor dissipates power because P is positive. This result should not be surprising since
we showed (p. 33) that the power consumed by any resistor equals either of the following.

v
or i R (3.3)

Since resistors are positive-valued, resistors always dissipate power . But where does a resistor's power
go? By Conversation of Power, the dissipated power must be absorbed somewhere. The answer is not
directly predicted by circuit theory, but is by physics. Current owing through a resistor makes it hot; its
power is dissipated by heat.
This content is available online at http://cnx.org/content/m17305/1.5/ .
