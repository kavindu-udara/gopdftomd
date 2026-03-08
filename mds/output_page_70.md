64 CHAPTER 3. ANALOG SIGNAL PROCESSING

is the transistor, with the turning achieved by the input signal. Just as in this analogy, a power supply is a
source of constant voltage as the water tower is supposed to provide a constant water pressure.
A device that is much more convenient for providing gain (and other useful features as well) than the
transistor is the operational ampli er , also known as the op-amp . An op-amp is an integrated circuit (a
complicated circuit involving several transistors constructed on a chip) that provides a large voltage gain if
you attach the power supply. We can model the op-amp with a new circuit element: the dependent source.

### 3.18 Dependent Sources

A dependent source is either a voltage or current source whose value is proportional to some other voltage
or current in the circuit. Thus, there are four di erent kinds of dependent sources; to describe an op-amp, we
need a voltage-dependent voltage source. However, the standard circuit-theoretical model for a transistor
contains a current-dependent current source. Dependent sources do not serve as inputs to a circuit like
independent sources. They are used to model active circuits : those containing electronic elements. The
RLC circuits we have been considering so far are known as passive circuits .

Figure 3.37: Of the four possible dependent sources, depicted is a voltage-dependent voltage source in
the context of a generic circuit.

Figure 3.38 shows the circuit symbol for the op-amp and its equivalent circuit in terms of a voltage-
dependent voltage source.

Figure 3.38: The op-amp has four terminals to which connections can be made. Inputs attach to nodes
a and b , and the output is node c . As the circuit model on the right shows, the op-amp serves as an
ampli er for the di erence of the input node voltages.

Here, the output voltage equals an ampli ed version of the di erence of node voltages appearing across
its inputs. The dependent source model portrays how the op-amp works quite well. As in most active circuit
schematics, the power supply is not shown, but must be present for the circuit model to be accurate. Most
operational ampli ers require both positive and negative supply voltages for proper operation.
This content is available online at http://cnx.org/content/m0053/2.14/ .
Small Signal Model for Bipolar Transistor http://cnx.org/content/m1019/latest/
