21

Figure 2.9: The system depicted has input x ( t ) and output y ( t ) . Mathematically, systems operate on
function(s) to produce other function(s). In many ways, systems are like functions, rules that yield a
value for the dependent variable (our output signal) for each value of its independent variable (its input
signal). The notation y ( t ) = S [ x ( t )] corresponds to this block diagram. We term S [ ] the input-output
relation for the system.

This notation mimics the mathematical symbology of a function: A system's input is analogous to an
independent variable and its output the dependent variable. For the mathematically inclined, a system is a
functional : a function of a function (signals are functions).
Simple systems can be connected together one system's output becomes another's input to accomplish
some overall design. Interconnection topologies can be quite complicated, but usually consist of weaves of
three basic interconnection forms.

2.5.1 Cascade Interconnection

Figure 2.10: Interconnecting systems so that one system's output serves as the input to another is the
cascade con guration.

The simplest form is when one system's output is connected only to another's input. Mathematically,
w ( t ) = S [ x ( t )] , and y ( t ) = S [ w ( t )] , with the information contained in x ( t ) processed by the rst, then
the second system. In some cases, the ordering of the systems matter, in others it does not. For example, in
the fundamental model of communication (Figure 1.3) the ordering most certainly matters.

2.5.2 Parallel Interconnection

Figure 2.11: The parallel con guration.

A signal x ( t ) is routed to two (or more) systems, with this signal appearing as the input to all systems
simultaneously and with equal strength. Block diagrams have the convention that signals going to more
than one system are not split into pieces along the way. Two or more systems operate on x ( t ) and their
outputs are added together to create the output y ( t ) . Thus, y ( t ) = S ] x ( t )]+ S [ x ( t )] , and the information
in x ( t ) is processed separately by both systems.
